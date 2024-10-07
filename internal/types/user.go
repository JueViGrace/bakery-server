package types

import (
	"time"

	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	BirthDate time.Time `json:"birth_date"`
	Phone     string    `json:"phone"`
	Role      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

type UpdateUserRequest struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Phone     string    `json:"phone"`
}

func DbUserToUser(du database.BakeryUser) *User {
	return &User{
		ID:        du.ID,
		FirstName: du.FirstName,
		LastName:  du.LastName,
		Email:     du.Email,
		Password:  du.Password,
		BirthDate: du.BirthDate,
		Phone:     du.Phone,
		Role:      du.Role,
		CreatedAt: du.CreatedAt.Time,
		UpdatedAt: du.UpdatedAt.Time,
		DeletedAt: du.DeletedAt.Time,
	}
}

func NewUpdateUserParams(ur UpdateUserRequest) *database.UpdateUserParams {
	return &database.UpdateUserParams{
		ID:        ur.ID,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		BirthDate: ur.BirthDate,
		Phone:     ur.Phone,
	}
}
