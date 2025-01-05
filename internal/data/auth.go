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
	SignIn(r *types.SignInRequest) (*types.AuthResponse, error)
	SignUp(r *types.SignUpRequest) (*types.AuthResponse, error)
	RecoverPassword(r *types.RecoverPasswordRequest) (string, error)
	Refresh(r *types.RefreshRequest) (*types.AuthResponse, error)
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

func (s *authStore) SignIn(r *types.SignInRequest) (*types.AuthResponse, error) {
	user, err := s.db.GetUserByEmail(s.ctx, r.Email)
	if err != nil {
		return nil, err
	}

	if user.DeletedAt.Valid {
		return nil, errors.New("this user was deleted")
	}

	if !util.ValidatePassword(r.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	tokenString, err := util.CreateJWT(user.ID, fmt.Sprintf("%s %s", user.FirstName, user.LastName))
	if err != nil {
		return nil, err
	}

	res := types.NewAuthResponse(tokenString, "")

	return res, nil
}

func (s *authStore) SignUp(r *types.SignUpRequest) (*types.AuthResponse, error) {
	newUser, err := types.SignUpRequestToDbUser(r)
	if err != nil {
		return nil, err
	}

	user, err := s.db.CreateUser(s.ctx, *newUser)
	if err != nil {
		return nil, err
	}

	tokenString, err := util.CreateJWT(user.ID, fmt.Sprintf("%s %s", user.FirstName, user.LastName))
	if err != nil {
		return nil, err
	}

	res := types.NewAuthResponse(tokenString, "")

	return res, nil
}

// TODO: finish

func (s *authStore) RecoverPassword(r *types.RecoverPasswordRequest) (string, error) {
	return "", nil
}

func (s *authStore) Refresh(r *types.RefreshRequest) (*types.AuthResponse, error) {
	return &types.AuthResponse{}, nil
}
