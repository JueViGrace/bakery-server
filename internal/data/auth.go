package data

import (
	"context"
	"errors"

	"github.com/JueViGrace/bakery-server/internal/database"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/google/uuid"
)

type AuthStore interface {
	SignIn(r *types.SignInRequest) (*types.AuthResponse, error)
	SignUp(r *types.SignUpRequest) (*types.AuthResponse, error)
	Refresh(r *types.RefreshRequest, a *types.AuthData) (*types.AuthResponse, error)
	RecoverPassword(r *types.RecoverPasswordRequest) (string, error)
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

// TODO: limit sessions to 5?
func (s *authStore) SignIn(r *types.SignInRequest) (*types.AuthResponse, error) {
	user, err := s.db.GetUser(s.ctx, database.GetUserParams{
		Email:    r.Email,
		Username: r.Email,
	})
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.DeletedAt.Valid {
		return nil, errors.New("this user was deleted")
	}

	if !util.ValidatePassword(r.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	userId, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, err
	}

	sessionId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	newTokens, err := createTokens(userId, sessionId, user.Username)
	if err != nil {
		return nil, err
	}

	err = s.db.CreateSession(s.ctx, types.CreateSessionToDb(&types.Session{
		ID:           sessionId,
		UserId:       userId,
		Username:     user.Username,
		RefreshToken: newTokens.RefreshToken,
		AccessToken:  newTokens.AccessToken,
	}))
	if err != nil {
		return nil, err
	}

	return newTokens, nil
}

// TODO: check for conflicts
func (s *authStore) SignUp(r *types.SignUpRequest) (*types.AuthResponse, error) {
	newUser, err := types.SignUpRequestToDbUser(r)
	if err != nil {
		return nil, err
	}

	user, err := s.db.CreateUser(s.ctx, *newUser)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, err
	}

	sessionId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	newTokens, err := createTokens(userId, sessionId, user.Username)
	if err != nil {
		return nil, err
	}

	err = s.db.CreateSession(s.ctx, types.CreateSessionToDb(&types.Session{
		ID:           sessionId,
		UserId:       userId,
		Username:     user.Username,
		RefreshToken: newTokens.RefreshToken,
		AccessToken:  newTokens.AccessToken,
	}))
	if err != nil {
		return nil, err
	}

	return newTokens, nil
}

func (s *authStore) Refresh(r *types.RefreshRequest, a *types.AuthData) (*types.AuthResponse, error) {
	user, err := s.db.GetUserById(s.ctx, a.UserId.String())
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, err
	}

	newTokens, err := createTokens(userId, a.SessionId, user.Username)
	if err != nil {
		return nil, err
	}

	err = s.db.UpdateSession(s.ctx, types.UpdateSessionToDb(
		&types.Session{
			UserId:       userId,
			Username:     user.Username,
			RefreshToken: newTokens.RefreshToken,
			AccessToken:  newTokens.AccessToken,
		},
	))
	if err != nil {
		return nil, err
	}

	return newTokens, nil
}

// TODO: implement this
func (s *authStore) RecoverPassword(r *types.RecoverPasswordRequest) (string, error) {
	return "", nil
}

func createTokens(userId, sessionId uuid.UUID, username string) (*types.AuthResponse, error) {
	accessToken, err := util.CreateAccessToken(userId, sessionId, username)
	if err != nil {
		return nil, err
	}

	refreshToken, err := util.CreateRefreshToken(userId, sessionId, username)
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
