package types

import (
	"github.com/JueViGrace/bakery-server/internal/database"
	"github.com/google/uuid"
)

type Session struct {
	UserId       uuid.UUID
	Username     string
	RefreshToken string
	AccessToken  string
}

func DbSessionToSession(s *database.BakerySession) (*Session, error) {
	userId, err := uuid.Parse(s.UserID)
	if err != nil {
		return nil, err
	}

	return &Session{
		UserId:       userId,
		Username:     s.Username,
		RefreshToken: s.RefreshToken,
		AccessToken:  s.AccessToken,
	}, nil
}

func CreateSessionToDb(r *Session) *database.CreateSessionParams {
	return &database.CreateSessionParams{
		RefreshToken: r.RefreshToken,
		AccessToken:  r.AccessToken,
		Username:     r.Username,
		UserID:       r.UserId.String(),
	}
}

func UpdateSessionToDb(r *Session) *database.UpdateSessionParams {
	return &database.UpdateSessionParams{
		RefreshToken: r.RefreshToken,
		AccessToken:  r.AccessToken,
		Username:     r.Username,
		UserID:       r.UserId.String(),
	}
}
