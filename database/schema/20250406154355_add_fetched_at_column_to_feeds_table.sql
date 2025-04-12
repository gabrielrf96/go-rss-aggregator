-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds ADD COLUMN fetched_at TIMESTAMP;
CREATE INDEX ON feeds(fetched_at ASC NULLS FIRST);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feeds DROP COLUMN fetched_at;
-- +goose StatementEnd
