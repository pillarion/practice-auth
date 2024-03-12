package user

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

// Service defines the user service.
//
//go:generate minimock -o mock/ -s "_minimock.go"
type Service interface {
	Create(ctx context.Context, user *desc.Info) (int64, error)
	Get(ctx context.Context, id int64) (*desc.User, error)
	Update(ctx context.Context, user *desc.Info) error
	Delete(ctx context.Context, id int64) error
}
