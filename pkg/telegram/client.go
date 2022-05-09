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

package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitlab.ozon.dev/capcom6/homework-2/pkg/httpex"
)

const BASE_URL = "https://api.telegram.org/bot"

type Telegram struct {
	token string
	http  httpex.Requester
}

func New(cfg Config) *Telegram {

	if cfg.HttpClient == nil {
		cfg.HttpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	return &Telegram{
		token: cfg.Token,
		http:  cfg.HttpClient,
	}
}

func (t *Telegram) SendMessage(msg *SendMessage) error {
	return t.post("sendMessage", msg)
}

func (t *Telegram) DeleteMessage(msg *DeleteMessage) error {
	return t.post("deleteMessage", msg)
}

func (t *Telegram) post(method string, payload interface{}) error {
	url := fmt.Sprintf("%s%s/%s", BASE_URL, t.token, method)
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(body)
	resp, err := t.http.Post(url, "application/json", reader)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("api error: %s", resp.Status)
}
