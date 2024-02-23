package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (p *pg) DeleteUser(ctx context.Context, id int64) error {

	builderDelete := sq.Delete("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

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
