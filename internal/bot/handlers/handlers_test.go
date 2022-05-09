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
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/capcom6/homework-2/internal/bot/core/ports"
)

type responseWriter struct {
	Headers    http.Header
	Body       []byte
	StatusCode int
}

func newResponseWriter() *responseWriter {
	return &responseWriter{
		Headers:    map[string][]string{},
		Body:       []byte{},
		StatusCode: 200,
	}
}

func (w *responseWriter) Header() http.Header {
	return w.Headers
}

func (w *responseWriter) Write(buf []byte) (int, error) {
	w.Body = append(w.Body, buf...)
	return len(buf), nil
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
}

type body struct {
	Body []byte
}

func (b *body) Read(p []byte) (int, error) {
	if len(b.Body) == 0 {
		return 0, io.EOF
	}

	var i int
	for i = 0; i < len(b.Body) && i < len(p); i++ {
		p[i] = b.Body[i]
	}

	b.Body = b.Body[i:]

	return i, nil
}

func (b *body) Close() error {
	return nil
}

func Test_handler_ServeHTTP(t *testing.T) {
	type fields struct {
		uri string
	}
	type args struct {
		r *http.Request
	}

	urlWoPath, _ := url.Parse("/")
	urlWPath, _ := url.Parse("/path")
	urlWInvalidPath, _ := url.Parse("/invalid/path")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *responseWriter
	}{
		{
			name:   "GET",
			fields: fields{uri: "path"},
			args: args{
				r: &http.Request{
					Method: "GET",
					URL:    urlWPath,
					Body: &body{
						Body: []byte{},
					},
				},
			},
			want: &responseWriter{
				StatusCode: http.StatusMethodNotAllowed,
			},
		},
		{
			name:   "PUT",
			fields: fields{uri: "path"},
			args: args{
				r: &http.Request{
					Method: "PUT",
					URL:    urlWPath,
					Body: &body{
						Body: []byte{},
					},
				},
			},
			want: &responseWriter{
				StatusCode: http.StatusMethodNotAllowed,
			},
		},
		{
			name:   "POST w/o path",
			fields: fields{uri: "path"},
			args: args{
				r: &http.Request{
					Method: "POST",
					URL:    urlWoPath,
					Body: &body{
						Body: []byte{},
					},
				},
			},
			want: &responseWriter{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:   "POST with invalid path",
			fields: fields{uri: "path"},
			args: args{
				r: &http.Request{
					Method: "POST",
					URL:    urlWInvalidPath,
					Body: &body{
						Body: []byte{},
					},
				},
			},
			want: &responseWriter{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:   "POST with invalid payload",
			fields: fields{uri: "path"},
			args: args{
				r: &http.Request{
					Method: "POST",
					URL:    urlWPath,
					Body: &body{
						Body: []byte("not a json"),
					},
				},
			},
			want: &responseWriter{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			name:   "POST with payload",
			fields: fields{uri: "path"},
			args: args{
				r: &http.Request{
					Method: "POST",
					URL:    urlWPath,
					Body: &body{
						Body: []byte(`{"update_id": 123}`),
					},
				},
			},
			want: &responseWriter{
				StatusCode: http.StatusOK,
			},
		},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tt := range tests {
		tg := ports.NewTelegramServiceMock(mc)
		ip := ports.NewInterpreterServiceMock(mc)

		rw := newResponseWriter()

		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				uri: tt.fields.uri,
				tg:  tg,
				ip:  ip,
			}
			h.ServeHTTP(rw, tt.args.r)

			assert.Equal(t, tt.want.StatusCode, rw.StatusCode)
		})
	}
}
