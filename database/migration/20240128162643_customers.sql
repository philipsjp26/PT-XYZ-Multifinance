-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `customers`(
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `identity_number` VARCHAR(150) NOT NULL UNIQUE,
    `full_name` VARCHAR(255) NOT NULL,
    `legal_name` VARCHAR(255) NOT NULL,
    `place_of_birth` VARCHAR(255) NULL DEFAULT '',
    `date_of_birth` DATE NOT NULL,
    `salary` DECIMAL(10,2),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `customers`;
-- +goose StatementEnd
