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

package puller

import (
	"log"
	"sort"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/core/domain"
)

type Puller struct {
}

func New() *Puller {
	return &Puller{}
}

func (p *Puller) Pull(targets []domain.Mailbox) ([]domain.Message, error) {
	if len(targets) == 0 {
		return []domain.Message{}, nil
	}

	messages := domain.Messages{}

	ch := make(chan result, len(targets))
	defer close(ch)

	for _, t := range targets {
		go func(t domain.Mailbox, ch chan result) {
			res, err := p.pullSingle(t)

			ch <- result{
				messages: res,
				err:      err,
			}
		}(t, ch)

		// messages = append(messages, res...)
	}

	for range targets {
		res := <-ch
		if res.err != nil {
			log.Printf("pull error: %v\n", res.err)
			continue
		}
		messages = append(messages, res.messages...)
	}

	sort.Sort(sort.Reverse(messages))

	return messages, nil
}

func (p *Puller) pullSingle(t domain.Mailbox) ([]domain.Message, error) {
	c, err := client.DialTLS(t.Server+":993", nil)
	if err != nil {
		return nil, err
	}
	// log.Printf("Connected to %s\n", t.Server)
	// Don't forget to logout
	defer c.Logout()

	c.Timeout = 5 * time.Second

	// Login
	if err := c.Login(t.Login, t.Password); err != nil {
		return nil, err
	}
	// log.Printf("Logged in as %s\n", t.Login)

	// Select INBOX
	mbox, err := c.Select("INBOX", true)
	if err != nil {
		return nil, err
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	from := mbox.UnseenSeqNum
	to := mbox.Messages

	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message)

	done := make(chan error, 1)

	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, imap.FetchFlags}, messages)
	}()

	// log.Println("Unseen messages:")
	result := []domain.Message{}
	for msg := range messages {
		if contains(imap.SeenFlag, msg.Flags) {
			continue
		}

		// log.Printf("%+v\n", msg.Flags)
		// log.Println("* " + msg.Envelope.Subject)
		// log.Printf(" - %+v\n", msg)
		result = append(result, domain.Message{
			From:  msg.Envelope.From[0].Address(),
			To:    msg.Envelope.To[0].Address(),
			Date:  msg.Envelope.Date,
			Title: msg.Envelope.Subject,
		})
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return result, nil
}

func contains(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
