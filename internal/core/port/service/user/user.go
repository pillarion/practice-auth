package user

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

// Service defines the user service.
type Service interface {
	Create(ctx context.Context, user *desc.User) (int64, error)
	Get(ctx context.Context, id int64) (*desc.User, error)
	Update(ctx context.Context, user *desc.User) error
	Delete(ctx context.Context, id int64) error
}
