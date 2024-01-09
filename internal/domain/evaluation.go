package domain

import "time"

func NewEvaluation(id, orderId, customerId, rate int, comment string) *Evaluation {
	now := time.Now()

	e := &Evaluation{
		ID:         id,
		OrderID:    orderId,
		CustomerID: customerId,
		Rate:       rate,
		Comment:    comment,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  nil,
	}

	return e
}

type Evaluation struct {
	ID         int        `json:"id" db:"id"`
	OrderID    int        `json:"order_id" db:"order_id"`
	CustomerID int        `json:"customer_id" db:"customer_id"`
	Rate       int        `json:"rate" db:"rate"`
	Comment    string     `json:"comment" db:"comment"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}

func (e *Evaluation) Validate() error {
	// TODO: validation

	return nil
}
