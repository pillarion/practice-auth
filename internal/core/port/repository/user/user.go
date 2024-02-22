package db

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

// UserRepo defines the database interface.
type UserRepo interface {
	InsertUser(ctx context.Context, user *desc.User) (int64, error)
	SelectUser(ctx context.Context, id int64) (*desc.User, error)
	UpdateUser(ctx context.Context, user *desc.User) error
	DeleteUser(ctx context.Context, id int64) error
}
