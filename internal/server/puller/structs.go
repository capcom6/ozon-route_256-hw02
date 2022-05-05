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

import "time"

type Target struct {
	Server   string
	Login    string
	Password string
}

type Message struct {
	From  string
	To    string
	Date  time.Time
	Title string
}

type Messages []Message

func (a Messages) Len() int           { return len(a) }
func (a Messages) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Messages) Less(i, j int) bool { return a[i].Date.Unix() < a[j].Date.Unix() }

type result struct {
	messages Messages
	err      error
}
