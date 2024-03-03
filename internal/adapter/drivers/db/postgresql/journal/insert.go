package journal

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/pillarion/practice-auth/internal/core/model/journal"
	db "github.com/pillarion/practice-auth/internal/core/tools/pgclient/port"
)

func (p *pg) Insert(ctx context.Context, j *desc.Journal) (int64, error) {
	builderInsert := sq.Insert(journalTable).
		PlaceholderFormat(sq.Dollar).
		Columns(
			journalTableActionColumn,
		).
		Values(j.Action).
		Suffix("RETURNING " + journalTableIDColumn)
	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "Journal.Insert",
		QueryRaw: query,
	}
	var journalID int64
	err = p.db.DB().ScanOneContext(ctx, &journalID, q, args...)
	if err != nil {
		return 0, err
	}

	return journalID, nil
}
