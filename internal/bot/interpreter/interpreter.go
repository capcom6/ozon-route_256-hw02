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
	"strings"

	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
)

const (
	MSG_UNKNOWN = `Неизвестная команда %s. Список всех команд: /help`
	MSG_HELP    = `Бот поддерживает следующие команды:
/help - справка по командам
/add <server> <login> <password> - добавление почтового ящика
/list - список почтовых ящиков
/delete <id> - удаление выбранного почтового ящика
/delete * - удаление всех почтовых ящиков
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
	return "", nil
}

func (p *processor) list(ctx context.Context, userId string, chunks []string) (string, error) {
	resp, err := p.client.Select(ctx, &pb.MailboxGet{
		UserId: userId,
	})
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}
	builder.WriteString("Список ящиков:\n")

	for _, m := range resp.GetMailboxes() {
		builder.WriteString(fmt.Sprintf("%d. %s @ %s\n", m.GetId(), m.GetLogin(), m.GetServer()))
	}

	if len(resp.GetMailboxes()) == 0 {
		builder.WriteString("ящики отсутствуют")
	}

	return builder.String(), nil
}

func (p *processor) remove(ctx context.Context, userId string, chunks []string) (string, error) {
	return "", nil
}

func (p *processor) pull(ctx context.Context, userId string, chunks []string) (string, error) {
	return "", nil
}

func (p *processor) help(ctx context.Context, userId string, chunks []string) (string, error) {
	return MSG_HELP, nil
}
