-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id uuid NOT NULL DEFAULT gen_random_uuid (), created_at timestamp with time zone, updated_at timestamp with time zone, deleted_at timestamp with time zone, name text, email text NOT NULL, password text, role text, validated_at jsonb, last_login jsonb, is_active boolean, PRIMARY KEY (id)
);

CREATE UNIQUE INDEX users_email_key ON users USING btree ("email");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop TABLE users
-- +goose StatementEnd