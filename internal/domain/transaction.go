package domain

import "time"

func NewTransaction(paymentType PaymentType, id, payerID int) *Transaction {
	now := time.Now()

	t := &Transaction{
		ID:        id,
		PayerID:   payerID,
		Payment:   paymentType,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	return t

}

type Transaction struct {
	ID        int         `json:"id" db:"id"`
	PayerID   int         `json:"payer_id" db:"payer_id"`
	Payment   PaymentType `json:"payment_type" db:"payment_type"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty" db:"deleted_at"`
}

type PaymentType struct {
	CreditPayment *CreditOpt `json:"credit_payment,omitempty" db:"credit_payment"`
	CashPayment   *CashOpt   `json:"cash_payment,omitempty" db:"cash_payment"`
}

type CreditOpt struct {
	Amount           float64 `json:"amount" db:"amount"`
	AmountOfPartials int64   `json:"amount_of_partials" db:"amount_of_partials"`
}

type CashOpt struct {
	Amount float64 `json:"amount" db:"amount"`
}

func (t *Transaction) Validate() error {
	// TODO: validation

	return nil
}
