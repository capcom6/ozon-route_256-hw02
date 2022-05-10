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

package interpreter

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"gitlab.ozon.dev/capcom6/homework-2/internal/bot/core/domain"
	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
)

func Test_processor_start(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
		chunks []string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Answer
		wantErr bool
	}{
		{
			name: "Simple start",
			args: args{
				ctx:    nil,
				userId: "id",
				chunks: []string{},
			},
			want: domain.Answer{
				Message:      MSG_WELCOME,
				DeleteSource: false,
			},
			wantErr: false,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpc := pb.NewMailAggregatorClientMock(mc)

			p := &processor{
				client: grpc,
			}
			got, err := p.start(tt.args.ctx, tt.args.userId, tt.args.chunks)
			if (err != nil) != tt.wantErr {
				t.Errorf("processor.start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processor.start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processor_add(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
		chunks []string
	}
	tests := []struct {
		name    string
		args    args
		mock    *pb.MailboxCreate
		want    domain.Answer
		wantErr bool
	}{
		{
			name: "Zero arg",
			args: args{
				ctx:    nil,
				userId: "123",
				chunks: []string{},
			},
			want: domain.Answer{
				Message:      MSG_ADD_PARAMETERS_COUNT,
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Single arg",
			args: args{
				ctx:    nil,
				userId: "123",
				chunks: []string{"arg"},
			},
			want: domain.Answer{
				Message:      MSG_ADD_PARAMETERS_COUNT,
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Two args",
			args: args{
				ctx:    nil,
				userId: "123",
				chunks: []string{"arg", "arg"},
			},
			want: domain.Answer{
				Message:      MSG_ADD_PARAMETERS_COUNT,
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Valid add",
			args: args{ctx: context.TODO(), userId: "123", chunks: []string{"arg1", "arg2", "arg3"}},
			mock: &pb.MailboxCreate{
				UserId: "123",
				Mailbox: &pb.MailboxIn{
					Server:   "arg1",
					Login:    "arg2",
					Password: "arg3",
				},
			},
			want:    domain.Answer{Message: fmt.Sprintf("Ящик %s @ %s добавлен", "arg2", "arg1"), DeleteSource: true},
			wantErr: false,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpc := pb.NewMailAggregatorClientMock(mc)
			if tt.mock != nil {
				grpc.CreateMock.Return(&pb.Empty{}, nil)
				grpc.CreateMock.Expect(context.TODO(), tt.mock)
			}

			p := &processor{
				client: grpc,
			}
			got, err := p.add(tt.args.ctx, tt.args.userId, tt.args.chunks)
			if (err != nil) != tt.wantErr {
				t.Errorf("processor.add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processor.add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processor_list(t *testing.T) {
	type mock struct {
		expect *pb.MailboxGet
		ret    *pb.Mailboxes
	}
	type args struct {
		ctx    context.Context
		userId string
		chunks []string
	}
	tests := []struct {
		name    string
		args    args
		mock    mock
		want    domain.Answer
		wantErr bool
	}{
		{
			name: "Empty list",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				chunks: []string{},
			},
			mock: mock{
				expect: &pb.MailboxGet{
					UserId: "123",
				},
				ret: &pb.Mailboxes{
					Mailboxes: []*pb.MailboxOut{},
				},
			},
			want: domain.Answer{
				Message:      "Список ящиков:\nящики отсутствуют",
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Some mailboxes",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				chunks: []string{},
			},
			mock: mock{
				expect: &pb.MailboxGet{
					UserId: "123",
				},
				ret: &pb.Mailboxes{
					Mailboxes: []*pb.MailboxOut{{
						Id:     1,
						Server: "srv1",
						Login:  "user1",
					}},
				},
			},
			want: domain.Answer{
				Message:      "Список ящиков:\n" + fmt.Sprintf("%d. %s @ %s\n", 1, "user1", "srv1"),
				DeleteSource: false,
			},
			wantErr: false,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpc := pb.NewMailAggregatorClientMock(mc)
			grpc.SelectMock.Expect(context.TODO(), tt.mock.expect)
			grpc.SelectMock.Return(tt.mock.ret, nil)

			p := &processor{
				client: grpc,
			}
			got, err := p.list(tt.args.ctx, tt.args.userId, tt.args.chunks)
			if (err != nil) != tt.wantErr {
				t.Errorf("processor.list() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processor.list() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processor_delete(t *testing.T) {
	type mock struct {
		expect *pb.MailboxDelete
		ret    *pb.Mailboxes
	}
	type args struct {
		ctx    context.Context
		userId string
		chunks []string
	}
	tests := []struct {
		name    string
		args    args
		mock    mock
		want    domain.Answer
		wantErr bool
	}{
		{
			name: "Without id",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				chunks: []string{},
			},
			mock: mock{},
			want: domain.Answer{
				Message:      MSG_DELETE_PARAMETERS_COUNT,
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Invalid id",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				chunks: []string{"invalid"},
			},
			mock: mock{},
			want: domain.Answer{
				Message:      MSG_INVALID_ID,
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Valid id",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				chunks: []string{"456"},
			},
			mock: mock{
				expect: &pb.MailboxDelete{
					UserId: "123",
					Mailbox: &pb.MailboxId{
						Id: 456,
					},
				},
				ret: &pb.Mailboxes{
					Mailboxes: []*pb.MailboxOut{},
				},
			},
			want: domain.Answer{
				Message:      "Список ящиков:\nящики отсутствуют",
				DeleteSource: false,
			},
			wantErr: false,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpc := pb.NewMailAggregatorClientMock(mc)
			if tt.mock.expect != nil {
				grpc.DeleteMock.Expect(context.TODO(), tt.mock.expect)
				grpc.DeleteMock.Return(tt.mock.ret, nil)
			}

			p := &processor{
				client: grpc,
			}
			got, err := p.delete(tt.args.ctx, tt.args.userId, tt.args.chunks)
			if (err != nil) != tt.wantErr {
				t.Errorf("processor.delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processor.delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processor_pull(t *testing.T) {
	type mock struct {
		expect *pb.MailboxGet
		ret    *pb.Messages
		err    error
	}
	type args struct {
		ctx    context.Context
		userId string
		chunks []string
	}

	demoTime := time.Now()
	demoTimeUnix := demoTime.Unix()

	tests := []struct {
		name    string
		args    args
		mock    mock
		want    domain.Answer
		wantErr bool
	}{
		{
			name: "Empty mails",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				chunks: []string{},
			},
			mock: mock{
				expect: &pb.MailboxGet{
					UserId: "123",
				},
				ret: &pb.Messages{
					Messages: []*pb.Message{},
				},
			},
			want: domain.Answer{
				Message:      "Новые письма:\nновых писем нет",
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Some mails",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				chunks: []string{},
			},
			mock: mock{
				expect: &pb.MailboxGet{
					UserId: "123",
				},
				ret: &pb.Messages{
					Messages: []*pb.Message{
						{
							Title:     "Title",
							From:      "From",
							To:        "To",
							Timestamp: demoTimeUnix,
						},
					},
				},
			},
			want: domain.Answer{
				Message:      "Новые письма:\n" + fmt.Sprintf("От %s на %s\n%s\n%s\n\n", "From", "To", demoTime.Format(time.RFC3339), "Title"),
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Communication error",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				chunks: []string{},
			},
			mock: mock{
				expect: &pb.MailboxGet{
					UserId: "123",
				},
				ret: &pb.Messages{},
				err: fmt.Errorf("test error"),
			},
			want:    domain.Answer{},
			wantErr: true,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpc := pb.NewMailAggregatorClientMock(mc)
			if tt.mock.expect != nil {
				grpc.PullMock.Expect(context.TODO(), tt.mock.expect)
				grpc.PullMock.Return(tt.mock.ret, tt.mock.err)
			}

			p := &processor{
				client: grpc,
			}
			got, err := p.pull(tt.args.ctx, tt.args.userId, tt.args.chunks)
			if (err != nil) != tt.wantErr {
				t.Errorf("processor.pull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processor.pull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processor_help(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
		chunks []string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Answer
		wantErr bool
	}{
		{
			name: "Simple help",
			args: args{
				ctx:    nil,
				userId: "id",
				chunks: []string{},
			},
			want: domain.Answer{
				Message:      MSG_HELP,
				DeleteSource: false,
			},
			wantErr: false,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpc := pb.NewMailAggregatorClientMock(mc)

			p := &processor{
				client: grpc,
			}
			got, err := p.help(tt.args.ctx, tt.args.userId, tt.args.chunks)
			if (err != nil) != tt.wantErr {
				t.Errorf("processor.help() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processor.help() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processor_Process(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
		msg    string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Answer
		wantErr bool
	}{
		{
			name: "Start",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				msg:    "/start",
			},
			want: domain.Answer{
				Message:      MSG_WELCOME,
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Start whitespace",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				msg:    "/start ",
			},
			want: domain.Answer{
				Message:      MSG_WELCOME,
				DeleteSource: false,
			},
			wantErr: false,
		},
		{
			name: "Add",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				msg:    "/add server login password",
			},
			want:    domain.Answer{Message: fmt.Sprintf("Ящик %s @ %s добавлен", "login", "server"), DeleteSource: true},
			wantErr: false,
		},
		{
			name: "List",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				msg:    "/list",
			},
			want:    domain.Answer{Message: "Список ящиков:\nящики отсутствуют", DeleteSource: false},
			wantErr: false,
		},
		{
			name: "Pull",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				msg:    "/pull",
			},
			want:    domain.Answer{Message: "Новые письма:\nновых писем нет", DeleteSource: false},
			wantErr: false,
		},
		{
			name: "Delete",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				msg:    "/delete 456",
			},
			want:    domain.Answer{Message: "Список ящиков:\nящики отсутствуют", DeleteSource: false},
			wantErr: false,
		},
		{
			name: "Unknown",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				msg:    "/unknown 456",
			},
			want:    domain.Answer{Message: fmt.Sprintf(MSG_UNKNOWN, "/unknown 456"), DeleteSource: false},
			wantErr: false,
		},
		{
			name: "Empty",
			args: args{
				ctx:    context.TODO(),
				userId: "123",
				msg:    "",
			},
			want:    domain.Answer{Message: fmt.Sprintf(MSG_UNKNOWN, ""), DeleteSource: false},
			wantErr: false,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	grpc := pb.NewMailAggregatorClientMock(mc)
	grpc.CreateMock.When(context.TODO(), &pb.MailboxCreate{
		UserId: "123",
		Mailbox: &pb.MailboxIn{
			Server:   "server",
			Login:    "login",
			Password: "password",
		},
	}).Then(&pb.Empty{}, nil)
	grpc.SelectMock.When(context.TODO(), &pb.MailboxGet{
		UserId: "123",
	}).Then(&pb.Mailboxes{
		Mailboxes: []*pb.MailboxOut{},
	}, nil)
	grpc.PullMock.When(context.TODO(), &pb.MailboxGet{
		UserId: "123",
	}).Then(&pb.Messages{}, nil)
	grpc.DeleteMock.When(context.TODO(), &pb.MailboxDelete{
		UserId: "123",
		Mailbox: &pb.MailboxId{
			Id: 456,
		},
	}).Then(&pb.Mailboxes{}, nil)

	p := &processor{
		client: grpc,
	}

	for _, tt := range tests {
		got, err := p.Process(tt.args.ctx, tt.args.userId, tt.args.msg)
		if (err != nil) != tt.wantErr {
			t.Errorf("processor.Process() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("processor.Process() = %v, want %v", got, tt.want)
		}
	}
}
