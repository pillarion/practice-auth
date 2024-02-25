package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (p *pg) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{usersTableIDColumn: id})
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
