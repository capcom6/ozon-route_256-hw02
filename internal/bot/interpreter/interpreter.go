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
)

type processor struct {
	client pb.MailAggregatorClient
}

func New(client pb.MailAggregatorClient) *processor {
	return &processor{
		client: client,
	}
}

func (p *processor) Process(ctx context.Context, userId string, msg string) (string, error) {
	chunks := strings.Split(msg, " ")
	if len(chunks) == 0 {
		return fmt.Sprintf(MSG_UNKNOWN, msg), nil
	}

	cmd := chunks[0]
	switch cmd {
	case "/start":
		return p.start(ctx, userId, chunks[1:])
	case "/add":
		return p.add(ctx, userId, chunks[1:])
	case "/list":
		return p.list(ctx, userId, chunks[1:])
	case "/remove":
		return p.remove(ctx, userId, chunks[1:])
	case "/pull":
		return p.pull(ctx, userId, chunks[1:])
	case "/help":
		return p.help(ctx, userId, chunks[1:])
	}

	return fmt.Sprintf(MSG_UNKNOWN, msg), nil
}

func (p *processor) start(ctx context.Context, userId string, chunks []string) (string, error) {
	return MSG_WELCOME, nil
}

func (p *processor) add(ctx context.Context, userId string, chunks []string) (string, error) {
	if len(chunks) != 3 {
		return "Недостаточно параметров. Формат команды: /add <server> <login> <password>", nil
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
		return "", err
	}

	return "Ящик добавлен", nil
}

func (p *processor) list(ctx context.Context, userId string, chunks []string) (string, error) {
	resp, err := p.client.Select(ctx, &pb.MailboxGet{
		UserId: userId,
	})
	if err != nil {
		return "", err
	}

	ans := formatMailboxesList(resp.GetMailboxes())

	return ans, nil
}

func (p *processor) remove(ctx context.Context, userId string, chunks []string) (string, error) {
	if len(chunks) != 1 {
		return "Недостаточно параметров. Формат команды: /delete <id>", nil
	}

	id, err := strconv.Atoi(chunks[0])
	if err != nil {
		return "Некорректный идентификатор ящика", err
	}

	req := pb.MailboxDelete{
		UserId: userId,
		Mailbox: &pb.MailboxId{
			Id: int32(id),
		},
	}

	resp, err := p.client.Delete(ctx, &req)
	if err != nil {
		return "", err
	}

	ans := formatMailboxesList(resp.GetMailboxes())

	return ans, nil
}

func (p *processor) pull(ctx context.Context, userId string, chunks []string) (string, error) {

	req := pb.MailboxGet{
		UserId: userId,
	}

	resp, err := p.client.Pull(ctx, &req)
	if err != nil {
		return "", err
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

	return builder.String(), nil
}

func (p *processor) help(ctx context.Context, userId string, chunks []string) (string, error) {
	return MSG_HELP, nil
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
