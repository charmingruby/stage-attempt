package domain

import "time"

type Thread struct {
	ID        int        `json:"id" db:"id"`
	ProductID int        `json:"product_id" db:"product_id"`
	Context   string     `json:"context" db:"context"`
	ErrorLog  string     `json:"error_log" db:"error_log"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	SolvedAt  *time.Time `json:"solved_at" db:"solved_at"`
}

func NewThread(id, productId int, context, errorLog string) *Thread {
	now := time.Now()

	t := &Thread{
		ID:        id,
		ProductID: productId,
		Context:   context,
		ErrorLog:  errorLog,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
		SolvedAt:  nil,
	}

	return t
}

func (t *Thread) Validate() error {
	// TODO: validation

	return nil
}
