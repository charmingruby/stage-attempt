package domain

import (
	"errors"
	"time"

	"github.com/charmingruby/stage-attempt/helpers"
	"github.com/charmingruby/stage-attempt/pkg"
)

var roles = []string{"developer", "customer"}
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
	phone,
	email,
	role,
	password string,
) *User {
	now := time.Now()

	u := &User{
		ID:              id,
		Name:            name,
		LastName:        lastName,
		Email:           email,
		Birthdate:       birthdate,
		Phone:           phone,
		Role:            helpers.If[string](role == "", roles[1], role),
		EmailVerifiedAt: nil,
		CreatedAt:       now,
		UpdatedAt:       now,
		DeletedAt:       nil,
	}

	u.SetPassword(password)

	println(u.Password)

	return u
}

type User struct {
	ID              int        `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	LastName        string     `json:"last_name" db:"last_name"`
	Email           string     `json:"email" db:"email"`
	Phone           string     `json:"phone" db:"phone"`
	Password        string     `json:"password" db:"password"`
	Role            string     `json:"role" db:"role"`
	Birthdate       time.Time  `json:"birthdate" db:"birthdate"`
	EmailVerifiedAt *time.Time `json:"email_verified_at" db:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
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

	hashedPassword, err := pkg.GenerateHash(password)
	if err != nil {
		println("error")
		return err
	}

	u.Password = hashedPassword

	return nil
}

func (u *User) Validate() error {
	// TODO: validation

	return nil
}
