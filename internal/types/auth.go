package types

import (
	"time"

	"github.com/JueViGrace/bakery-server/internal/database"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthDataHandler = func(*fiber.Ctx, *AuthData) error

type AuthData struct {
	UserId   uuid.UUID
	Username string
	Role     string
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	FirstName   string    `json:"firstName" validate:"required"`
	LastName    string    `json:"lastName" validate:"required"`
	Username    string    `json:"username" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password" validate:"required"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
	BirthDate   time.Time `json:"birthDate"`
	Address1    string    `json:"address1"`
	Address2    string    `json:"address2"`
	Gender      string    `json:"gender"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RecoverPasswordRequest struct {
	Password string `json:"password"`
}

func SignUpRequestToDbUser(r *SignUpRequest) (*database.CreateUserParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	pass, err := util.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}

	return &database.CreateUserParams{
		ID:          id.String(),
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Username:    r.Username,
		Email:       r.Email,
		Password:    pass,
		PhoneNumber: r.PhoneNumber,
		BirthDate:   r.BirthDate.String(),
		Address1:    r.Address1,
		Address2:    r.Address2,
		Gender:      r.Gender,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	}, nil
}
