package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/JueViGrace/bakery-server/internal/database"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/google/uuid"
)

type UserStore interface {
	GetUsers() ([]*types.UserResponse, error)
	GetUserById(id *uuid.UUID) (*types.UserResponse, error)
	UpdateUser(r *types.UpdateUserRequest) (*types.UserResponse, error)
	DeleteUser(id *uuid.UUID) error
}

func (s *storage) UserStore() UserStore {
	return NewUserStore(s.ctx, s.queries)
}

type userStore struct {
	ctx context.Context
	db  *database.Queries
}

func NewUserStore(ctx context.Context, db *database.Queries) UserStore {
	return &userStore{
		ctx: ctx,
		db:  db,
	}
}

func (us *userStore) GetUsers() ([]*types.UserResponse, error) {
	users := make([]*types.UserResponse, 0)

	dbUsers, err := us.db.GetUsers(us.ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range dbUsers {
		user, err := types.DbUserToUser(&u)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (us *userStore) GetUserById(id *uuid.UUID) (*types.UserResponse, error) {
	user := new(types.UserResponse)

	dbUser, err := us.db.GetUserById(us.ctx, id.String())
	if err != nil {
		return nil, err
	}

	user, err = types.DbUserToUser(&dbUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userStore) UpdateUser(r *types.UpdateUserRequest) (*types.UserResponse, error) {
	user := new(types.UserResponse)

	params, err := types.NewUpdateUserParams(r)
	if err != nil {
		return nil, err
	}

	dbUser, err := us.db.UpdateUser(us.ctx, *params)
	if err != nil {
		return nil, err
	}

	user, err = types.DbUserToUser(&dbUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userStore) DeleteUser(id *uuid.UUID) error {
	err := us.db.DeleteUser(us.ctx, database.DeleteUserParams{
		DeletedAt: sql.NullString{
			String: time.Now().UTC().String(),
			Valid:  true,
		},
		ID: id.String(),
	})
	if err != nil {
		return err
	}

	return nil
}
