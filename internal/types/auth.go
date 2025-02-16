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
	UserId    uuid.UUID
	SessionId uuid.UUID
	Username  string
	Role      string
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Username    string `json:"username"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	BirthDate   string `json:"birth_date" validate:"required"`
	Address1    string `json:"address1" validate:"required"`
	Address2    string `json:"address2" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RecoverPasswordRequest struct {
	Password string `json:"password" validate:"required"`
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

	birthDate, err := time.Parse(time.DateOnly, r.BirthDate)
	if err != nil {
		return nil, err
	}

	var username string = r.Username
	if username == "" {
		username = r.Email
	}

	return &database.CreateUserParams{
		ID:          id.String(),
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Username:    username,
		Email:       r.Email,
		Password:    pass,
		PhoneNumber: r.PhoneNumber,
		BirthDate:   util.FormatDateForResponse(birthDate),
		Address1:    r.Address1,
		Address2:    r.Address2,
		Gender:      r.Gender,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	}, nil
}
