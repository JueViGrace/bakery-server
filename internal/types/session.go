package types

import (
	"github.com/JueViGrace/bakery-server/internal/database"
	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID
	UserId       uuid.UUID
	Username     string
	RefreshToken string
	AccessToken  string
}

func DbSessionToSession(s *database.BakerySession) (*Session, error) {
	id, err := uuid.Parse(s.ID)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(s.UserID)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:           id,
		UserId:       userId,
		Username:     s.Username,
		RefreshToken: s.RefreshToken,
		AccessToken:  s.AccessToken,
	}, nil
}

func CreateSessionToDb(r *Session) database.CreateSessionParams {
	return database.CreateSessionParams{
		ID:           r.ID.String(),
		RefreshToken: r.RefreshToken,
		AccessToken:  r.AccessToken,
		Username:     r.Username,
		UserID:       r.UserId.String(),
	}
}

func UpdateSessionToDb(r *Session) database.UpdateSessionParams {
	return database.UpdateSessionParams{
		RefreshToken: r.RefreshToken,
		AccessToken:  r.AccessToken,
		Username:     r.Username,
		ID:           r.ID.String(),
	}
}
