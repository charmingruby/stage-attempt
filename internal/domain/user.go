package domain

import "time"

type User struct {
	ID        int        `json:"id" db:"id"`
	CreateAt  time.Time  `json:"create_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
