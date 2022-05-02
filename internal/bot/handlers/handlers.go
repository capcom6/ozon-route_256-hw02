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
	"encoding/json"
	"io"
	"log"
	"net/http"

	"gitlab.ozon.dev/capcom6/homework-2/pkg/telegram"
)

type handler struct {
	uri string
	tg  *telegram.Telegram
}

func New(cfg Config) http.Handler {
	return &handler{
		uri: cfg.URI,
		tg:  cfg.TG,
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

	if !update.Message.From.IsBot {
		err := h.tg.SendMessage(&telegram.SendMessage{
			ChatID: update.Message.Chat.ID,
			Text:   "Pong: " + update.Message.Text,
		})
		if err != nil {
			log.Println(err)
		}
	}

	// log.Println(string(payload))
	log.Printf("%+v\n", update)

	w.Write([]byte(r.URL.String()))
}
