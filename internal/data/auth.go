package data

import (
	"context"
	"errors"
	"time"

	"github.com/JueViGrace/bakery-go/internal/db"
	"github.com/JueViGrace/bakery-go/internal/util"
	"github.com/google/uuid"
)

type AuthStore interface {
	SignIn(r SignInRequest) (string, error)
	SignUp(r SignUpRequest) (string, error)
	RecoverPassword(r RecoverPasswordRequest) (string, error)
	ChangeEmail(r ChangeEmailRequest) (string, error)
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

func (s *authStore) SignIn(r SignInRequest) (string, error) {
	user, err := s.db.GetUserByEmail(s.ctx, r.Email)
	if err != nil {
		return "", err
	}

	if user.DeletedAt.Valid {
		return "", errors.New("this user was deleted")
	}

	if !util.ValidatePassword(r.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return "", nil
}

func (s *authStore) SignUp(r SignUpRequest) (string, error) {
	newUser, err := SignUpRequestToDbUser(r)
	if err != nil {
		return "", err
	}

	_, err = s.db.CreateUser(s.ctx, *newUser)
	if err != nil {
		return "", err
	}

	msg := "User created!"

	return msg, nil
}

func (s *authStore) RecoverPassword(r RecoverPasswordRequest) (string, error) {
	return "", nil
}

func (s *authStore) ChangeEmail(r ChangeEmailRequest) (string, error) {
	return "", nil
}

func SignUpRequestToDbUser(r SignUpRequest) (*db.CreateUserParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	pass, err := util.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}

	return &db.CreateUserParams{
		ID:        id,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  pass,
		BirthDate: r.BirthDate,
		Phone:     r.Phone,
	}, nil
}
