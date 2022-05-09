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

package ports

import (
	"context"

	"gitlab.ozon.dev/capcom6/homework-2/internal/bot/core/domain"
	"gitlab.ozon.dev/capcom6/homework-2/pkg/telegram"
)

type InterpreterService interface {
	Process(ctx context.Context, userId string, msg string) (domain.Answer, error)
}

type TelegramService interface {
	SendMessage(msg *telegram.SendMessage) error
	DeleteMessage(msg *telegram.DeleteMessage) error
}
