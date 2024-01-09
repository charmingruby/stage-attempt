package domain

import "time"

const (
	payment_pending    = "pending"
	payment_canceled   = "canceled"
	payment_processing = "processing"
	payment_delivering = "delivering"
	payment_delivered  = "delivered"
)

func NewPayment(paymentType PaymentType, id, payerID int) *Payment {
	now := time.Now()

	t := &Payment{
		ID:        id,
		PayerID:   payerID,
		Payment:   paymentType,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	t.StatusPending()

	return t
}

type Payment struct {
	ID        int         `json:"id" db:"id"`
	PayerID   int         `json:"payer_id" db:"payer_id"`
	ReportID  int         `json:"report_id" db:"report_id"`
	Payment   PaymentType `json:"payment_type" db:"payment_type"`
	Status    string      `json:"status" db:"status"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty" db:"deleted_at"`
}

type PaymentType struct {
	CreditPayment *CreditOpt `json:"credit_payment,omitempty" db:"credit_payment"`
	CashPayment   *CashOpt   `json:"cash_payment,omitempty" db:"cash_payment"`
}

type CreditOpt struct {
	AmountInCents    int64 `json:"amount_in_cents" db:"amount_in_cents"`
	AmountOfPartials int64 `json:"amount_of_partials" db:"amount_of_partials"`
}

type CashOpt struct {
	AmountInCents int64 `json:"amount_in_cents" db:"amount_in_cents"`
}

func paymentStatus() map[string]string {
	return map[string]string{
		payment_pending:    "pending",
		payment_canceled:   "canceled",
		payment_processing: "processing",
		payment_delivering: "delivering",
		payment_delivered:  "delivered",
	}
}

func (p *Payment) StatusPending() *Payment {
	status := paymentStatus()

	p.Status = status[payment_pending]

	return p
}

func (p *Payment) Validate() error {
	// TODO: validation

	return nil
}
