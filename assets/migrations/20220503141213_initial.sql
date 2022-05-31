-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.mailboxes (
    user_id varchar(36) NOT NULL,
    id serial4 NOT NULL,
    "server" varchar(64) NOT NULL,
    login varchar(64) NOT NULL,
    "password" varchar(128) NOT NULL,
    CONSTRAINT mailboxes_pkey PRIMARY KEY (user_id, id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE public.mailboxes;
-- +goose StatementEnd