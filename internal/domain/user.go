package domain

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var roles = []string{"admin", "customer"}
var (
	ErrPasswordRequired  = errors.New("password is required and can't be blank")
	ErrPasswordMinLength = errors.New("password must be at least 8 characters")
	ErrPasswordMaxLength = errors.New("password must be a maximum of 16 characters")
)

func NewUser(
	id int,
	birthdate time.Time,
	name,
	lastName,
	email,
	role,
	avatarUrl,
	password string,
) *User {
	now := time.Now()

	var roleToAssign string

	if role == "" {
		roleToAssign = roles[1]
	}

	u := &User{
		ID:        id,
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Birthdate: birthdate,
		AvatarUrl: avatarUrl,
		Role:      roleToAssign,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	u.SetPassword(u.Password)

	return u
}

type User struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	LastName  string     `json:"last_name" db:"last_name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	Role      string     `json:"role" db:"role"`
	AvatarUrl string     `json:"avatar_url" db:"avatar_url"`
	Birthdate time.Time  `json:"birthdate" db:"birthdate"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired

	}

	if len(password) < 8 {
		return ErrPasswordMinLength
	}

	if len(password) > 16 {
		return ErrPasswordMaxLength
	}

	u.Password = fmt.Sprintf("%x", md5.Sum([]byte(password)))

	return nil
}

func (u *User) Validate() error {
	// TODO: validation

	return nil
}
