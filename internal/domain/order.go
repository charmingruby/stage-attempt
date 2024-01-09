package domain

import "time"

type Order struct {
	Id          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	Deadline    *time.Time `json:"deadline" db:"deadline"`
	ProductType string     `json:"product_type" db:"product_type"`
	Status      string     `json:"status" db:"status"`

	Threats    []Threat `json:"threats" db:"threats"`
	Reports    []Report `json:"reports" db:"reports"`
	PaymentID  int      `json:"payment_id" db:"payment_id"`
	CustomerID int      `json:"customer_id" db:"customer_id"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
