package data

import (
	"context"
	"time"

	"github.com/JueViGrace/bakery-go/internal/db"
)

type AuthStore interface {
	SignIn(r SignInRequest) (*string, error)
	SignUp(r SignUpRequest) (*string, error)
	RecoverPassword(r RecoverPasswordRequest) (*string, error)
	ChangeEmail(r ChangeEmailRequest) (*string, error)
}

type authStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewAuthStore(ctx context.Context, db *db.Queries) AuthStore {
	return &authStore{
		ctx: ctx,
		db:  db,
	}
}

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

// TODO: JWT

func (s *authStore) SignIn(r SignInRequest) (*string, error) {
	_, err := s.db.GetUserByEmail(s.ctx, r.Email)
	if err != nil {
		return nil, err
	}

	// TODO: check if is not deleted
	// TODO: check password

	return nil, nil
}

func (s *authStore) SignUp(r SignUpRequest) (*string, error) {
	_, err := s.db.CreateUser(s.ctx, *SignUpRequestToDbUser(r))
	if err != nil {
		return nil, err
	}

	msg := "User created!"

	return &msg, nil
}

func (s *authStore) RecoverPassword(r RecoverPasswordRequest) (*string, error) {
	return nil, nil
}

func (s *authStore) ChangeEmail(r ChangeEmailRequest) (*string, error) {
	return nil, nil
}

func SignUpRequestToDbUser(r SignUpRequest) *db.CreateUserParams {
	// TODO: BCRYPT

	return &db.CreateUserParams{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  r.Password,
		BirthDate: r.BirthDate,
		Phone:     r.Phone,
	}
}
