-- +goose Up
-- +goose StatementBegin
CREATE TABLE events
(
    user_id INTEGER,
    event_name varchar NOT NULL,
    event_date timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events
-- +goose StatementEnd
