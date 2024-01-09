package domain

import "time"

func NewOrder(id, customerId, paymentId int, name, description, productType string, deadline time.Time) *Order {
	now := time.Now()

	o := &Order{
		ID:          id,
		Name:        name,
		Description: description,
		Deadline:    &deadline,
		ProductType: productType,
		PaymentID:   paymentId,
		CustomerID:  customerId,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
	}

	o.StatusPending()

	return o
}

type Order struct {
	ID          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	Deadline    *time.Time `json:"deadline" db:"deadline"`
	ProductType string     `json:"product_type" db:"product_type"`
	Status      string     `json:"status" db:"status"`

	PaymentID  int      `json:"payment_id" db:"payment_id"`
	CustomerID int      `json:"customer_id" db:"customer_id"`
	Threats    []Threat `json:"threats,omitempty" db:"threats"`
	Reports    []Report `json:"reports,omitempty" db:"reports"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

func (o *Order) StatusPending() *Order {
	status := OrderStatus()

	o.Status = status[order_pending]

	return o
}
