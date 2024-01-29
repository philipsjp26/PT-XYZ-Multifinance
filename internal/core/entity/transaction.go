package entity

import (
	"time"
)

type Transaction struct {
	CustomerID        int64      `db:"customer_id"`
	ContractNumber    string     `db:"contract_number"`
	Otr               float32    `db:"otr"`
	AdminFee          float32    `db:"admin_fee"`
	InstallmentAmount float32    `db:"installment_amount"`
	InterestAmount    float32    `db:"interest_amount"`
	CreatedAt         *time.Time `db:"created_at,omitempty"`
	UpdatedAt         *time.Time `db:"updated_at,omitempty"`
}
