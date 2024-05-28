-- +goose Up
-- +goose StatementBegin
CREATE TABLE projects if not exists (
    id uuid NOT NULL DEFAULT gen_random_uuid (), created_at timestamp with time zone DEFAULT now(), updated_at timestamp with time zone, deleted_at timestamp with time zone, name text, description text, project_type text, woocommerce_name text, woocommerce_description text, woocommerce_domain text, woocommerce_consumer_key text, woocommerce_consumer_secret text, shopify_consumer_name text, shopify_description text, shopify_domain text, shopify_consumer_key text, shopify_consumer_secret text, user_id uuid, validated_at jsonb, is_active boolean, PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE projects
-- +goose StatementEnd