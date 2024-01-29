package entity

type CustomerLimit struct {
	CustomerID  int64   `db:"customer_id"`
	Tenor       int     `db:"tenor"`
	LimitAmount float32 `db:"limit_amount"`
}
