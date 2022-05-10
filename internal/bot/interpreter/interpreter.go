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
	"strconv"
	"strings"
	"time"

	"gitlab.ozon.dev/capcom6/homework-2/internal/bot/core/domain"
	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
)

const (
	MSG_UNKNOWN = `Неизвестная команда %s. Список всех команд: /help`
	MSG_HELP    = `Бот поддерживает следующие команды:
/help - справка по командам
/add <server> <login> <password> - добавление почтового ящика
/list - список почтовых ящиков
/delete <id> - удаление выбранного почтового ящика
/pull - получить список новых писем`
	MSG_WELCOME = `Добро пожаловать в почтовый бот Route256
Бот позволяет проверять наличие новой почты сразу в нескольких почтовых ящиках и отображать заголовки новых писем в чате.
Для получения справки по командам введите /help`
	MSG_ADD_PARAMETERS_COUNT    = `Недостаточно параметров. Формат команды: /add <server> <login> <password>`
	MSG_DELETE_PARAMETERS_COUNT = `Недостаточно параметров. Формат команды: /delete <id>`
	MSG_INVALID_ID              = `Некорректный идентификатор ящика`
)

type processor struct {
	client pb.MailAggregatorClient
}

func New(client pb.MailAggregatorClient) *processor {
	return &processor{
		client: client,
	}
}

func (p *processor) Process(ctx context.Context, userId string, msg string) (domain.Answer, error) {
	chunks := strings.Split(msg, " ")
	if len(chunks) == 0 {
		return domain.Answer{Message: fmt.Sprintf(MSG_UNKNOWN, msg)}, nil
	}

	cmd := chunks[0]
	switch cmd {
	case "/start":
		return p.start(ctx, userId, chunks[1:])
	case "/add":
		return p.add(ctx, userId, chunks[1:])
	case "/list":
		return p.list(ctx, userId, chunks[1:])
	case "/delete":
		return p.delete(ctx, userId, chunks[1:])
	case "/pull":
		return p.pull(ctx, userId, chunks[1:])
	case "/help":
		return p.help(ctx, userId, chunks[1:])
	}

	return domain.Answer{Message: fmt.Sprintf(MSG_UNKNOWN, msg)}, nil
}

func (p *processor) start(ctx context.Context, userId string, chunks []string) (domain.Answer, error) {
	return domain.Answer{Message: MSG_WELCOME}, nil
}

func (p *processor) add(ctx context.Context, userId string, chunks []string) (domain.Answer, error) {
	if len(chunks) != 3 {
		return domain.Answer{Message: MSG_ADD_PARAMETERS_COUNT}, nil
	}

	mb := pb.MailboxIn{
		Server:   chunks[0],
		Login:    chunks[1],
		Password: chunks[2],
	}

	req := pb.MailboxCreate{
		UserId:  userId,
		Mailbox: &mb,
	}

	if _, err := p.client.Create(ctx, &req); err != nil {
		return domain.Answer{}, err
	}

	return domain.Answer{
		Message:      fmt.Sprintf("Ящик %s @ %s добавлен", mb.Login, mb.Server),
		DeleteSource: true,
	}, nil
}

func (p *processor) list(ctx context.Context, userId string, chunks []string) (domain.Answer, error) {
	resp, err := p.client.Select(ctx, &pb.MailboxGet{
		UserId: userId,
	})
	if err != nil {
		return domain.Answer{}, err
	}

	ans := formatMailboxesList(resp.GetMailboxes())

	return domain.Answer{Message: ans}, nil
}

func (p *processor) delete(ctx context.Context, userId string, chunks []string) (domain.Answer, error) {
	if len(chunks) != 1 {
		return domain.Answer{Message: MSG_DELETE_PARAMETERS_COUNT}, nil
	}

	id, err := strconv.Atoi(chunks[0])
	if err != nil {
		return domain.Answer{Message: MSG_INVALID_ID}, nil
	}

	req := pb.MailboxDelete{
		UserId: userId,
		Mailbox: &pb.MailboxId{
			Id: int32(id),
		},
	}

	resp, err := p.client.Delete(ctx, &req)
	if err != nil {
		return domain.Answer{}, err
	}

	ans := formatMailboxesList(resp.GetMailboxes())

	return domain.Answer{Message: ans}, nil
}

func (p *processor) pull(ctx context.Context, userId string, chunks []string) (domain.Answer, error) {

	req := pb.MailboxGet{
		UserId: userId,
	}

	resp, err := p.client.Pull(ctx, &req)
	if err != nil {
		return domain.Answer{}, err
	}

	builder := strings.Builder{}
	builder.WriteString("Новые письма:\n")

	for _, m := range resp.GetMessages() {
		tm := time.Unix(m.GetTimestamp(), 0)
		builder.WriteString(fmt.Sprintf("От %s на %s\n%s\n%s\n\n", m.GetFrom(), m.GetTo(), tm.Format(time.RFC3339), m.GetTitle()))
	}

	if len(resp.GetMessages()) == 0 {
		builder.WriteString("новых писем нет")
	}

	return domain.Answer{Message: builder.String()}, nil
}

func (p *processor) help(ctx context.Context, userId string, chunks []string) (domain.Answer, error) {
	return domain.Answer{Message: MSG_HELP}, nil
}

func formatMailboxesList(lst []*pb.MailboxOut) string {
	builder := strings.Builder{}
	builder.WriteString("Список ящиков:\n")

	for _, m := range lst {
		builder.WriteString(fmt.Sprintf("%d. %s @ %s\n", m.GetId(), m.GetLogin(), m.GetServer()))
	}

	if len(lst) == 0 {
		builder.WriteString("ящики отсутствуют")
	}
	return builder.String()
}
