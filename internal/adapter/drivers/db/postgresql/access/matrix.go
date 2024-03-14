package access

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	model "github.com/pillarion/practice-auth/internal/core/model/access"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
)

// AccessMatrix retrieves a matrix from the database based on the provided endpoint.
//
// ctx context.Context, endpoint string
// *model.Matrix, error
func (p *pg) AccessMatrix(ctx context.Context, endpoint string) (*model.Matrix, error) {
	builderSelect := sq.Select(
		accessMatrixTableIDColumn,
		accessMatrixTableEndpointColumn,
		accessMatrixTableRoleColumn,
	).
		From(accessMatrixTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{accessMatrixTableEndpointColumn: endpoint})
	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "AccessMatrix.Select",
		QueryRaw: query,
	}

	var am model.Matrix
	err = p.db.DB().ScanAllContext(ctx, &am, q, args...)
	if err != nil {
		return nil, err
	}

	return &am, nil
}
