package access

import (
	"github.com/pillarion/practice-auth/internal/core/port/repository/access"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
)

type pg struct {
	db db.Client
}

const (
	accessMatrixTable               = "access_matrix"
	accessMatrixTableIDColumn       = "id"
	accessMatrixTableRoleColumn     = "role"
	accessMatrixTableEndpointColumn = "endpoint"
)

// New initializes a new user repository using the provided database configuration.
//
// db: the database client.
// repo.UserRepo, error
func New(db db.Client) (access.Repo, error) {
	return &pg{
		db: db,
	}, nil
}
