package user

import (
	"context"

	model "github.com/pillarion/practice-auth/internal/core/model/user"
)

// Repo defines the database interface.
//
//go:generate minimock -o mock/ -s "_minimock.go"
type Repo interface {
	Insert(ctx context.Context, user *model.Info) (int64, error)
	SelectByID(ctx context.Context, id int64) (*model.User, error)
	SelectByName(ctx context.Context, username string) (*model.User, error)
	SelectPassword(ctx context.Context, username string) (string, error)
	Update(ctx context.Context, user *model.Info) error
	Delete(ctx context.Context, id int64) error
}
