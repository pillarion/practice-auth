package access

import (
	"context"

	model "github.com/pillarion/practice-auth/internal/core/model/access"
)

type Repo interface {
	AccessMatrix(ctx context.Context, endpoint string) ([]model.Access, error)
}
