package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pillarion/practice-auth/internal/core/entity/config"
	"github.com/pillarion/practice-auth/internal/core/port/repository/user"
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
	pgx *pgxpool.Pool
}

// New initializes a new user repository using the provided database configuration.
//
// ctx context.Context, cfg *config.Database
// repo.UserRepo, error
func New(ctx context.Context, cfg *config.Database) (user.Repo, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Db, cfg.Pass)

	pgx, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &pg{
		pgx: pgx,
	}, nil
}
