-- +goose Up
ALTER TABLE urls
ADD COLUMN access_count INTEGER DEFAULT 0;

-- +goose Down
ALTER TABLE urls
DROP COLUMN access_count;
