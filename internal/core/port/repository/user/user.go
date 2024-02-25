package user

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

// Repo defines the database interface.
type Repo interface {
	Insert(ctx context.Context, user *desc.User) (int64, error)
	Select(ctx context.Context, id int64) (*desc.User, error)
	Update(ctx context.Context, user *desc.User) error
	Delete(ctx context.Context, id int64) error
}
