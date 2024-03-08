package journal

import (
	"context"

	model "github.com/pillarion/practice-auth/internal/core/model/journal"
)

// Repo defines the journal repository
//
//go:generate minimock -o mock/ -s "_minimock.go"
type Repo interface {
	Insert(ctx context.Context, j *model.Journal) (int64, error)
}
