package data

import (
	"context"

	"github.com/JueViGrace/bakery-server/internal/database"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/google/uuid"
)

type SessionStore interface {
	GetSessionById(id uuid.UUID) (session *types.Session, err error)
	GetSessionByUsername(username string) (session *types.Session, err error)
	CreateSession(r *types.Session) (err error)
	UpdateSession(r *types.Session) (err error)
	DeleteSessionById(id uuid.UUID) (err error)
	DeleteSessionByToken(token string) (err error)
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

func (s *sessionStore) GetSessionById(id uuid.UUID) (*types.Session, error) {
	session := new(types.Session)

	dbSession, err := s.db.GetSessionById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	session, err = types.DbSessionToSession(&dbSession)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionStore) GetSessionByUsername(username string) (*types.Session, error) {
	session := new(types.Session)

	dbSession, err := s.db.GetSessionByUsername(s.ctx, username)
	if err != nil {
		return nil, err
	}

	session, err = types.DbSessionToSession(&dbSession)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionStore) CreateSession(r *types.Session) error {
	err := s.db.CreateSession(s.ctx, *types.CreateSessionToDb(r))
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionStore) UpdateSession(r *types.Session) error {
	err := s.db.UpdateSession(s.ctx, *types.UpdateSessionToDb(r))
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionStore) DeleteSessionById(id uuid.UUID) error {
	err := s.db.DeleteSessionById(s.ctx, id.String())
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionStore) DeleteSessionByToken(token string) error {
	err := s.db.DeleteSessionByToken(s.ctx, database.DeleteSessionByTokenParams{
		RefreshToken: token,
		AccessToken:  token,
	})
	if err != nil {
		return err
	}
	return nil
}
