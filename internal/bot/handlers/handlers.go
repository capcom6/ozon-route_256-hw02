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

package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"gitlab.ozon.dev/capcom6/homework-2/internal/bot/core/ports"
	"gitlab.ozon.dev/capcom6/homework-2/pkg/telegram"
)

type handler struct {
	uri string
	tg  ports.TelegramService
	ip  ports.InterpreterService
}

func New(cfg Config) http.Handler {
	return &handler{
		uri: cfg.URI,
		tg:  cfg.TG,
		ip:  cfg.Processor,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.String()[1:] != h.uri {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payload, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	update := telegram.Update{}
	if err := json.Unmarshal(payload, &update); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if update.Message.From.IsBot {
		return
	}

	log.Printf("%+v\n", update)

	go func(u telegram.Update) {
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()

		ans, err := h.ip.Process(ctx, strconv.Itoa(u.Message.From.ID), u.Message.Text)
		if err != nil {
			log.Printf("error processing message %v\n", err)
			if ans.Message == "" {
				ans.Message = "К сожалению, произошла ошибка. Попробуйте позже."
			}
		}

		if ans.Message != "" {
			err = h.tg.SendMessage(&telegram.SendMessage{
				ChatID: u.Message.Chat.ID,
				Text:   ans.Message,
			})
			if err != nil {
				log.Printf("error sending message %v\n", err)
			}
		}
		if ans.DeleteSource {
			err = h.tg.DeleteMessage(&telegram.DeleteMessage{
				ChatID:    u.Message.Chat.ID,
				MessageID: u.Message.MessageID,
			})
			if err != nil {
				log.Printf("error deleting message %v\n", err)
			}
		}

	}(update)
}
