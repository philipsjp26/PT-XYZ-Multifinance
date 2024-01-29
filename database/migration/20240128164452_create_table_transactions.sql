-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `transactions`(
    `customer_id` BIGINT NOT NULL,
    `contract_number` VARCHAR(255) NOT NULL UNIQUE,
    `otr` DECIMAL(10,2) NOT NULL,
    `admin_fee` DECIMAL(10,2) NOT NULL,
    `installment_amount` DECIMAL(10,2) NOT NULL,
    `interest_amount` DECIMAL(10,2) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fk_customer_transaction` FOREIGN KEY (`customer_id`) REFERENCES `customers`(`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `transactions`;
-- +goose StatementEnd
