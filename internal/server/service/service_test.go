// Copyright 2022 Aleksandr Soloshenko
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/core/domain"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/core/ports"
	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
)

var (
	mailboxIn = &pb.MailboxIn{
		Server:   "server",
		Login:    "login",
		Password: "password",
	}
	mailboxInInvalid = &pb.MailboxIn{
		Server:   "invalid",
		Login:    "login",
		Password: "password",
	}

	mailboxOut = &pb.MailboxOut{
		Id:     0,
		Server: "server",
		Login:  "login",
	}
	mailboxOutWithId = &pb.MailboxOut{
		Id:     111,
		Server: "server",
		Login:  "login",
	}
	mailboxOutInvalid = &pb.MailboxOut{
		Id:     0,
		Server: "invalid",
		Login:  "login",
	}

	mailbox = domain.Mailbox{
		UserId:   "123",
		Server:   "server",
		Login:    "login",
		Password: "password",
	}
	mailboxWithId = domain.Mailbox{
		Id:       111,
		UserId:   "123",
		Server:   "server",
		Login:    "login",
		Password: "password",
	}
	mailboxInvalid = domain.Mailbox{
		UserId:   "123",
		Server:   "invalid",
		Login:    "login",
		Password: "password",
	}

	errSome = fmt.Errorf("some error")
)

func Test_service_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		msg *pb.MailboxCreate
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Empty
		wantErr bool
	}{
		{
			name: "Simple create",
			args: args{
				ctx: context.TODO(),
				msg: &pb.MailboxCreate{
					UserId:  "123",
					Mailbox: mailboxIn,
				},
			},
			want:    &pb.Empty{},
			wantErr: false,
		},
		{
			name: "Invalid create",
			args: args{
				ctx: context.TODO(),
				msg: &pb.MailboxCreate{
					UserId:  "123",
					Mailbox: mailboxInInvalid,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	repo := ports.NewMailboxesRepositoryMock(mc)
	puller := ports.NewMessagesRepositoryMock(mc)

	repo.CreateMock.When(context.TODO(), mailbox).Then(789, nil)
	repo.CreateMock.When(context.TODO(), mailboxInvalid).Then(0, errSome)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:   repo,
				puller: puller,
			}
			got, err := s.Create(tt.args.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Select(t *testing.T) {
	type args struct {
		ctx context.Context
		msg *pb.MailboxGet
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Mailboxes
		wantErr bool
	}{
		{
			name: "Empty",
			args: args{
				ctx: context.TODO(),
				msg: &pb.MailboxGet{
					UserId: "000",
				},
			},
			want: &pb.Mailboxes{
				Mailboxes: []*pb.MailboxOut{},
			},
			wantErr: false,
		},
		{
			name: "Single",
			args: args{
				ctx: context.TODO(),
				msg: &pb.MailboxGet{
					UserId: "123",
				},
			},
			want: &pb.Mailboxes{
				Mailboxes: []*pb.MailboxOut{mailboxOut},
			},
			wantErr: false,
		},
		{
			name: "Multiple",
			args: args{
				ctx: context.TODO(),
				msg: &pb.MailboxGet{
					UserId: "456",
				},
			},
			want: &pb.Mailboxes{
				Mailboxes: []*pb.MailboxOut{mailboxOut, mailboxOutWithId},
			},
			wantErr: false,
		},
		{
			name: "Error",
			args: args{
				ctx: context.TODO(),
				msg: &pb.MailboxGet{
					UserId: "invalid",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	repo := ports.NewMailboxesRepositoryMock(mc)
	puller := ports.NewMessagesRepositoryMock(mc)

	repo.SelectMock.When(context.TODO(), "000").Then([]domain.Mailbox{}, nil)
	repo.SelectMock.When(context.TODO(), "123").Then([]domain.Mailbox{mailbox}, nil)
	repo.SelectMock.When(context.TODO(), "456").Then([]domain.Mailbox{mailbox, mailboxWithId}, nil)
	repo.SelectMock.When(context.TODO(), "invalid").Then(nil, errSome)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:   repo,
				puller: puller,
			}
			got, err := s.Select(tt.args.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
