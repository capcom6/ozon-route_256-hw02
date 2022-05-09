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
	"testing"

	"github.com/gojuno/minimock/v3"
	"gitlab.ozon.dev/capcom6/homework-2/pkg/httpex"
)

const (
	TOKEN = "TelegramToken"
)

type body struct {
}

func (b *body) Read(p []byte) (int, error) {
	return 0, nil
}

func (b *body) Close() error {
	return nil
}

func TestTelegram_SendMessage(t *testing.T) {
	type args struct {
		msg  *SendMessage
		resp http.Response
		err  error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Send message",
			args: args{
				msg: &SendMessage{ChatID: 123, Text: "Test"},
				resp: http.Response{
					StatusCode: 200,
					Body:       &body{},
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Send message HTTP error",
			args: args{
				msg: &SendMessage{ChatID: 123, Text: "Test"},
				resp: http.Response{
					StatusCode: 400,
					Body:       &body{},
				},
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "Send message infrastructure error",
			args: args{
				msg:  &SendMessage{ChatID: 123, Text: "Test"},
				resp: http.Response{},
				err:  fmt.Errorf("infrastructure error"),
			},
			wantErr: true,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt.args.msg)

			req := httpex.NewRequesterMock(mc)
			req.PostMock.Return(&tt.args.resp, tt.args.err)
			req.PostMock.Expect(fmt.Sprintf("https://api.telegram.org/bot%s/%s", TOKEN, "sendMessage"), "application/json", bytes.NewReader(payload))

			tr := &Telegram{
				token: TOKEN,
				http:  req,
			}
			if err := tr.SendMessage(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Telegram.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTelegram_DeleteMessage(t *testing.T) {
	type args struct {
		msg  *DeleteMessage
		resp http.Response
		err  error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Delete message",
			args: args{
				msg: &DeleteMessage{ChatID: 123, MessageID: 456},
				resp: http.Response{
					StatusCode: 200,
					Body:       &body{},
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Delete message HTTP error",
			args: args{
				msg: &DeleteMessage{ChatID: 123, MessageID: 456},
				resp: http.Response{
					StatusCode: 400,
					Body:       &body{},
				},
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "Delete message infrastructure error",
			args: args{
				msg:  &DeleteMessage{ChatID: 123, MessageID: 456},
				resp: http.Response{},
				err:  fmt.Errorf("infrastructure error"),
			},
			wantErr: true,
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt.args.msg)

			req := httpex.NewRequesterMock(mc)
			req.PostMock.Return(&tt.args.resp, tt.args.err)
			req.PostMock.Expect(fmt.Sprintf("https://api.telegram.org/bot%s/%s", TOKEN, "deleteMessage"), "application/json", bytes.NewReader(payload))

			tr := &Telegram{
				token: TOKEN,
				http:  req,
			}
			if err := tr.DeleteMessage(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Telegram.DeleteMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
