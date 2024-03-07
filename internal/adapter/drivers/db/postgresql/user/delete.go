package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
)

func (p *pg) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{usersTableIDColumn: id})
	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "User.Delete",
		QueryRaw: query,
	}
	_, err = p.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
