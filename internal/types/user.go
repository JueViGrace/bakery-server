package types

import (
	"time"

	"github.com/JueViGrace/bakery-server/internal/database"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	PhoneNumber string    `json:"phone_number"`
	BirthDate   string    `json:"birth_date"`
	Address1    string    `json:"address1"`
	Address2    string    `json:"address2"`
	Gender      string    `json:"gender"`
	Role        string    `json:"-"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   string    `json:"-"`
}

// todo: make addresses an array
type UpdateUserRequest struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	BirthDate   string    `json:"birth_date" validate:"required"`
	Address1    string    `json:"address1" validate:"required"`
	Address2    string    `json:"address2" validate:"required"`
	Gender      string    `json:"gender" validate:"required"`
}

type ChangeEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

func DbUserToUser(db *database.BakeryUser) (user *UserResponse, err error) {
	id, err := uuid.Parse(db.ID)
	if err != nil {
		return nil, err
	}

	user = &UserResponse{
		ID:          id,
		FirstName:   db.FirstName,
		LastName:    db.LastName,
		Username:    db.Username,
		Email:       db.Email,
		Password:    db.Password,
		PhoneNumber: db.PhoneNumber,
		BirthDate:   db.BirthDate,
		Address1:    db.Address1,
		Address2:    db.Address2,
		Gender:      db.Gender,
		Role:        db.Role,
		CreatedAt:   db.CreatedAt,
		UpdatedAt:   db.UpdatedAt,
		DeletedAt:   db.DeletedAt.String,
	}

	return
}

func NewUpdateUserParams(r *UpdateUserRequest) (*database.UpdateUserParams, error) {
	birthDate, err := time.Parse(time.DateTime, r.BirthDate)
	if err != nil {
		return nil, err
	}

	return &database.UpdateUserParams{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		PhoneNumber: r.PhoneNumber,
		BirthDate:   util.FormatDateForResponse(birthDate),
		Address1:    r.Address1,
		Address2:    r.Address2,
		Gender:      r.Gender,
		UpdatedAt:   time.Now().UTC().String(),
		ID:          r.ID.String(),
	}, nil
}
