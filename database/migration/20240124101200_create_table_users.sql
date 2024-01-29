-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- CREATE TABLE IF NOT EXISTS `users` (
--     `id` BIGINT NOT NULL auto_increment,
--     `name` VARCHAR(150) NOT NULL,
--     PRIMARY KEY(id)
-- );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
