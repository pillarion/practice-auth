package access

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	model "github.com/pillarion/practice-auth/internal/core/model/access"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
)

// AccessMatrix retrieves a matrix from the database based on the provided endpoint.
//
// ctx context.Context, endpoint string
// []model.Access, error
func (p *pg) AccessMatrix(ctx context.Context, endpoint string) ([]model.Access, error) {
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
		fmt.Println(err)
		return nil, err
	}
	q := db.Query{
		Name:     "AccessMatrix.Select",
		QueryRaw: query,
	}
	var a []model.Access
	err = p.db.DB().ScanAllContext(ctx, &a, q, args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return a, nil
}
