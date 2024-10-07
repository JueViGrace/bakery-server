package data

import (
	"context"

	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/JueViGrace/bakery-go/internal/types"
	"github.com/google/uuid"
)

type UserStore interface {
	GetUsers() ([]types.User, error)
	GetUserById(id uuid.UUID) (*types.User, error)
	UpdateUser(ur types.UpdateUserRequest) (*types.User, error)
	DeleteUser(id uuid.UUID) error
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

func (us *userStore) GetUsers() ([]types.User, error) {
	users := make([]types.User, 0)

	dbUsers, err := us.db.GetUsers(us.ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range dbUsers {
		users = append(users, *types.DbUserToUser(u))
	}

	return users, nil
}

func (us *userStore) GetUserById(id uuid.UUID) (*types.User, error) {
	user := new(types.User)

	dbUser, err := us.db.GetUserById(us.ctx, id)
	if err != nil {
		return nil, err
	}

	user = types.DbUserToUser(dbUser)

	return user, nil
}

func (us *userStore) UpdateUser(ur types.UpdateUserRequest) (*types.User, error) {
	user := new(types.User)

	dbUser, err := us.db.UpdateUser(us.ctx, *types.NewUpdateUserParams(ur))
	if err != nil {
		return nil, err
	}

	user = types.DbUserToUser(dbUser)

	return user, nil
}

func (us *userStore) DeleteUser(id uuid.UUID) error {
	err := us.db.DeleteUser(us.ctx, id)
	if err != nil {
		return err
	}

	return nil
}
