package data

import (
	"context"

	"github.com/JueViGrace/bakery-server/internal/database"
)

type SessionStore interface {
	GetTokenById(id string) (*database.BakerySession, error)
	GetTokenByToken(token string) (*database.BakerySession, error)
	DeleteTokenById(id string) error
	DeleteTokenByToken(token string) error
}

func (s *storage) SessionStore() SessionStore {
	return NewSessionStore(s.ctx, s.queries)
}

type sessionStore struct {
	ctx context.Context
	db  *database.Queries
}

func NewSessionStore(ctx context.Context, db *database.Queries) SessionStore {
	return &sessionStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *sessionStore) GetTokenById(id string) (*database.BakerySession, error) {
	session, err := s.db.GetTokenById(s.ctx, id)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *sessionStore) GetTokenByToken(token string) (*database.BakerySession, error) {
	session, err := s.db.GetTokenById(s.ctx, token)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *sessionStore) DeleteTokenById(id string) error {
	err := s.db.DeleteTokenById(s.ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionStore) DeleteTokenByToken(token string) error {
	err := s.db.DeleteTokenByToken(s.ctx, token)
	if err != nil {
		return err
	}
	return nil
}
