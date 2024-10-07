package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/JueViGrace/bakery-go/internal/types"
	"github.com/JueViGrace/bakery-go/internal/util"
)

type AuthStore interface {
	SignIn(r types.SignInRequest) (string, error)
	SignUp(r types.SignUpRequest) (string, error)
	RecoverPassword(r types.RecoverPasswordRequest) (string, error)
	ChangeEmail(r types.ChangeEmailRequest) (string, error)
}

func (s *storage) AuthStore() AuthStore {
	return NewAuthStore(s.ctx, s.queries)
}

type authStore struct {
	ctx context.Context
	db  *database.Queries
}

func NewAuthStore(ctx context.Context, db *database.Queries) AuthStore {
	return &authStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *authStore) SignIn(r types.SignInRequest) (string, error) {
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

	tokenString, err := util.CreateJWT(fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *authStore) SignUp(r types.SignUpRequest) (string, error) {
	newUser, err := types.SignUpRequestToDbUser(r)
	if err != nil {
		return "", err
	}

	user, err := s.db.CreateUser(s.ctx, *newUser)
	if err != nil {
		return "", err
	}

	tokenString, err := util.CreateJWT(fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// TODO: finish

func (s *authStore) RecoverPassword(r types.RecoverPasswordRequest) (string, error) {
	return "", nil
}

func (s *authStore) ChangeEmail(r types.ChangeEmailRequest) (string, error) {
	return "", nil
}
