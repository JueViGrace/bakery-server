package data

import (
	"context"
	"time"

	"github.com/JueViGrace/bakery-go/internal/db"
	"github.com/google/uuid"
)

type UserStore interface {
	GetUsers() ([]User, error)
	GetUser(id uuid.UUID) (*User, error)
	UpdateUser(ur UpdateUserRequest) (*User, error)
	DeleteUser(id uuid.UUID) error
}

type userStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewUserStore(ctx context.Context, db *db.Queries) UserStore {
	return &userStore{
		ctx: ctx,
		db:  db,
	}
}

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

func (us *userStore) GetUsers() ([]User, error) {
	users := make([]User, 0)

	dbUsers, err := us.db.GetUsers(us.ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range dbUsers {
		users = append(users, *dbUserToUser(u))
	}

	return users, nil
}

func (us *userStore) GetUser(id uuid.UUID) (*User, error) {
	user := new(User)

	dbUser, err := us.db.GetUserById(us.ctx, id)
	if err != nil {
		return nil, err
	}

	user = dbUserToUser(dbUser)

	return user, nil
}

func (us *userStore) UpdateUser(ur UpdateUserRequest) (*User, error) {
	user := new(User)

	dbUser, err := us.db.UpdateUser(us.ctx, *URToDbUser(ur))
	if err != nil {
		return nil, err
	}

	user = dbUserToUser(dbUser)

	return user, nil
}

func (us *userStore) DeleteUser(id uuid.UUID) error {
	err := us.db.DeleteUser(us.ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func dbUserToUser(du db.BakeryUser) *User {
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

func URToDbUser(ur UpdateUserRequest) *db.UpdateUserParams {
	return &db.UpdateUserParams{
		ID:        ur.ID,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		BirthDate: ur.BirthDate,
		Phone:     ur.Phone,
	}
}
