package types

import (
	"time"

	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/JueViGrace/bakery-go/internal/util"
	"github.com/google/uuid"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	Phone     string    `json:"phone"`
}

type RecoverPasswordRequest struct {
	Password string `json:"password"`
}

type ChangeEmailRequest struct {
	Email string `json:"email"`
}

func SignUpRequestToDbUser(r SignUpRequest) (*database.CreateUserParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	pass, err := util.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}

	return &database.CreateUserParams{
		ID:        id,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  pass,
		BirthDate: r.BirthDate,
		Phone:     r.Phone,
	}, nil
}
