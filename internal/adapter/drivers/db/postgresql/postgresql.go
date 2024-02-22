package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pillarion/practice-auth/internal/core/entity/config"
	desc "github.com/pillarion/practice-auth/internal/core/model/user"
	repo "github.com/pillarion/practice-auth/internal/core/port/repository/user"

	sq "github.com/Masterminds/squirrel"
)

type pg struct {
	pgx *pgxpool.Pool
}

// New initializes a new user repository using the provided database configuration.
//
// ctx context.Context, cfg *config.Database
// repo.UserRepo, error
func New(ctx context.Context, cfg *config.Database) (repo.UserRepo, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Db, cfg.Pass)

	pgx, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &pg{
		pgx: pgx,
	}, nil
}

func (p *pg) InsertUser(ctx context.Context, user *desc.User) (int64, error) {

	builderInsert := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("user_id", "name", "email", "password", "role").
		Values(user.ID, user.Name, user.Email, user.Password, user.Role).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	var userID int64
	err = p.pgx.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (p *pg) SelectUser(ctx context.Context, id int64) (*desc.User, error) {

	builderSelect := sq.Select("user_id", "name", "email", "password", "role").
		From("users").
		Where(sq.Eq{"user_id": id})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	var user desc.User
	err = p.pgx.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (p *pg) UpdateUser(ctx context.Context, user *desc.User) error {

	builderUpdate := sq.Update("users").
		PlaceholderFormat(sq.Dollar).
		Set("name", user.Name).
		Set("email", user.Email).
		Set("password", user.Password).
		Set("role", user.Role).
		Where(sq.Eq{"user_id": user.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	_, err = p.pgx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *pg) DeleteUser(ctx context.Context, id int64) error {

	builderDelete := sq.Delete("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_id": id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	_, err = p.pgx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
