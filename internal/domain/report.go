package domain

import (
	"time"
)

func NewReport(id, userId int, title, description string) *Report {
	now := time.Now()

	r := &Report{
		ID:          id,
		Title:       title,
		Description: description,
		UserID:      userId,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
	}

	return r
}

type Report struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	UserID      int        `json:"user_id" db:"user_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (r *Report) Validate() error {
	// TODO: validation

	return nil
}
