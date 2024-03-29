package postgresql

import (
	"github.com/pillarion/practice-auth/internal/core/port/repository/user"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	usersTable                = "users"
	usersTableIDColumn        = "id"
	usersTableNameColumn      = "name"
	usersTableEmailColumn     = "email"
	usersTablePasswordColumn  = "password"
	usersTableRoleColumn      = "role"
	usersTableCreatedAtColumn = "created_at"
	usersTableUpdatedAtColumn = "updated_at"
)

type pg struct {
	db     db.Client
	tracer trace.Tracer
}

// New initializes a new user repository using the provided database configuration.
//
// db: the database client.
// repo.UserRepo, error
func New(db db.Client) (user.Repo, error) {
	return &pg{
		db:     db,
		tracer: otel.Tracer("postgresql"),
	}, nil
}
