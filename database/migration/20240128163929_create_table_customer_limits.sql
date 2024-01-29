-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `customer_limits`(
    `customer_id` BIGINT NOT NULL,
    `tenor` INT NOT NULL,
    `limit_amount` DECIMAL(10,2) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fk_customer_limit` FOREIGN KEY (`customer_id`) REFERENCES `customers`(`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `customer_limits`;
-- +goose StatementEnd
