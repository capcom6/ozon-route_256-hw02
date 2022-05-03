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

	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedMailAggregatorServer
}

func New() *service {
	return &service{}
}

func (s *service) Create(context.Context, *pb.MailboxCreate) (*pb.MailboxOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (s *service) Get(context.Context, *pb.MailboxGet) (*pb.Mailboxes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *service) Delete(context.Context, *pb.MailboxId) (*pb.Mailboxes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (s *service) Pull(context.Context, *pb.MailboxGet) (*pb.Messages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
