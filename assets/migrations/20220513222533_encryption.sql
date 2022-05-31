-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.mailboxes
ADD COLUMN encrypted boolean NOT NULL DEFAULT FALSE;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.mailboxes DROP COLUMN "encrypted";
-- +goose StatementEnd