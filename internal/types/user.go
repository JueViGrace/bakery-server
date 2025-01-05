package types

import (
	"time"

	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	PhoneNumber string    `json:"phoneNumber"`
	BirthDate   time.Time `json:"birthDate"`
	Address1    string    `json:"address1"`
	Address2    string    `json:"address2"`
	Gender      string    `json:"gender"`
	Role        string    `json:"-"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   string    `json:"-"`
}

type UpdateUserRequest struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	PhoneNumber string    `json:"phoneNumber"`
	BirthDate   time.Time `json:"birthDate"`
	Address1    string    `json:"address1"`
	Address2    string    `json:"address2"`
	Gender      string    `json:"gender"`
}

type ChangeEmailRequest struct {
	Email string `json:"email"`
}

func DbUserToUser(db *database.BakeryUser) (user *UserResponse, err error) {
	id, err := uuid.Parse(db.ID)
	if err != nil {
		return nil, err
	}

	birthDate, err := time.Parse(time.DateTime, db.BirthDate)
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
		BirthDate:   birthDate,
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

func NewUpdateUserParams(r *UpdateUserRequest) *database.UpdateUserParams {
	return &database.UpdateUserParams{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		PhoneNumber: r.PhoneNumber,
		BirthDate:   r.BirthDate.String(),
		Address1:    r.Address1,
		Address2:    r.Address2,
		Gender:      r.Gender,
		UpdatedAt:   time.Now().UTC().String(),
		ID:          r.ID.String(),
	}
}
