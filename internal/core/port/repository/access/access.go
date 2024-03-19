package access

import (
	"context"

	model "github.com/pillarion/practice-auth/internal/core/model/access"
)

// Repo defines the access repository.
type Repo interface {
	AccessMatrix(ctx context.Context, endpoint string) ([]model.Access, error)
}
