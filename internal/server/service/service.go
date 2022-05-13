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
	"log"

	"gitlab.ozon.dev/capcom6/homework-2/internal/server/config"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/core/domain"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/core/ports"
	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
	"gitlab.ozon.dev/capcom6/homework-2/pkg/crypto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedMailAggregatorServer
	cfg    config.Service
	repo   ports.MailboxesRepository
	puller ports.MessagesRepository
}

func New(cfg config.Service, repo ports.MailboxesRepository, puller ports.MessagesRepository) *service {
	return &service{
		cfg:    cfg,
		repo:   repo,
		puller: puller,
	}
}

func (s *service) Create(ctx context.Context, msg *pb.MailboxCreate) (*pb.Empty, error) {
	mb := domain.Mailbox{
		UserId:   msg.GetUserId(),
		Server:   msg.Mailbox.GetServer(),
		Login:    msg.Mailbox.GetLogin(),
		Password: msg.Mailbox.GetPassword(),
	}

	if s.cfg.SecretKey != "" {
		encrypted, err := crypto.Encrypt(mb.Password, s.cfg.SecretKey)
		if err != nil {
			log.Println(err)
		} else {
			mb.Password = encrypted
			mb.Encrypted = true
		}
	}

	if _, err := s.repo.Create(ctx, mb); err != nil {
		return nil, status.Errorf(codes.Internal, "could not create: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s *service) Select(ctx context.Context, msg *pb.MailboxGet) (*pb.Mailboxes, error) {
	mbx, err := s.selectMailboxes(ctx, msg.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not select: %v", err)
	}

	resp := pb.Mailboxes{
		Mailboxes: mbx,
	}

	return &resp, nil
}

func (s *service) Delete(ctx context.Context, msg *pb.MailboxDelete) (*pb.Mailboxes, error) {
	if err := s.repo.Delete(ctx, msg.GetUserId(), int(msg.Mailbox.GetId())); err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete: %v", err)
	}

	mbx, err := s.selectMailboxes(ctx, msg.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not select: %v", err)
	}

	resp := pb.Mailboxes{
		Mailboxes: mbx,
	}

	return &resp, nil
}

func (s *service) Pull(ctx context.Context, msg *pb.MailboxGet) (*pb.Messages, error) {
	mbx, err := s.repo.Select(ctx, msg.GetUserId())
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(mbx); i++ {
		if mbx[i].Encrypted {
			if mbx[i].Password, err = crypto.Decrypt(mbx[i].Password, s.cfg.SecretKey); err != nil {
				return nil, fmt.Errorf("can't decrypt password for mailbox id %d", mbx[i].Id)
			}
		}
	}

	msgs, err := s.puller.Pull(mbx)
	if err != nil {
		return nil, err
	}

	resp := &pb.Messages{
		Messages: []*pb.Message{},
	}

	for _, m := range msgs {
		resp.Messages = append(resp.Messages, &pb.Message{
			Title:     m.Title,
			From:      m.From,
			To:        m.To,
			Timestamp: m.Date.Unix(),
		})
	}

	return resp, nil
}

func (s *service) selectMailboxes(ctx context.Context, userId string) ([]*pb.MailboxOut, error) {
	mbx, err := s.repo.Select(ctx, userId)
	if err != nil {
		return nil, err
	}

	mailboxes := make([]*pb.MailboxOut, 0)

	for _, m := range mbx {
		mailboxes = append(mailboxes, &pb.MailboxOut{
			Id:     int32(m.Id),
			Server: m.Server,
			Login:  m.Login,
		})
	}

	return mailboxes, nil
}
