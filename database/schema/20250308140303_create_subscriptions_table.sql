-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL
        REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    feed_id UUID NOT NULL
        REFERENCES feeds(id) ON UPDATE CASCADE ON DELETE CASCADE,

    UNIQUE(user_id, feed_id)
);

CREATE INDEX ON subscriptions(feed_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
