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

package repositories

import (
	"context"
	"database/sql"

	"gitlab.ozon.dev/capcom6/homework-2/internal/server/core/domain"
)

type mailboxes struct {
	db *sql.DB
}

const (
	SQL_INSERT     = `INSERT INTO public.mailboxes (user_id, "server", login, "password", "encrypted") VALUES($1, $2, $3, $4, $5) RETURNING id;`
	SQL_SELECT     = `SELECT user_id, id, "server", login, "password", "encrypted" FROM public.mailboxes WHERE user_id = $1;`
	SQL_DELETE     = `DELETE FROM public.mailboxes WHERE user_id = $1 AND id = $2;`
	SQL_DELETE_ALL = `DELETE FROM public.mailboxes WHERE user_id = $1;`
)

func NewMailboxes(db *sql.DB) *mailboxes {
	return &mailboxes{
		db: db,
	}
}

func (r *mailboxes) Create(ctx context.Context, m domain.Mailbox) (int, error) {
	row := r.db.QueryRowContext(ctx, SQL_INSERT, m.UserId, m.Server, m.Login, m.Password, m.Encrypted)
	if row.Err() != nil {
		return 0, row.Err()
	}

	var id int

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *mailboxes) Select(ctx context.Context, userId string) ([]domain.Mailbox, error) {
	rows, err := r.db.QueryContext(ctx, SQL_SELECT, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []domain.Mailbox{}
	for rows.Next() {
		mb := domain.Mailbox{}
		if err := rows.Scan(&mb.UserId, &mb.Id, &mb.Server, &mb.Login, &mb.Password, &mb.Encrypted); err != nil {
			return nil, err
		}
		result = append(result, mb)
	}
	return result, nil
}

func (r *mailboxes) Delete(ctx context.Context, userId string, id int) (err error) {
	if id == 0 {
		_, err = r.db.ExecContext(ctx, SQL_DELETE_ALL, userId)
	} else {
		_, err = r.db.ExecContext(ctx, SQL_DELETE, userId, id)
	}
	return
}
