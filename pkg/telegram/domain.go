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

// type Response struct {
// 	Ok     bool     `json:"ok"`
// 	Result []Update `json:"result"`
// }
type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
type Entities struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}
type Message struct {
	MessageID int        `json:"message_id"`
	From      From       `json:"from"`
	Chat      Chat       `json:"chat"`
	Date      int        `json:"date"`
	Text      string     `json:"text"`
	Entities  []Entities `json:"entities"`
}
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type SendMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
type DeleteMessage struct {
	ChatID    int `json:"chat_id"`
	MessageID int `json:"message_id"`
}
