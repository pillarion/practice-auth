package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	dto "github.com/pillarion/practice-auth/internal/core/dto/postgresql"
	desc "github.com/pillarion/practice-auth/internal/core/model/user"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
)

// InsertUser inserts a new user into the database.
//
// ctx - the context for the database operation.
// user - the user object to be inserted.
// (int64, error) - returns the user_id of the inserted user and any error encountered.
func (p *pg) Insert(ctx context.Context, user *desc.User) (int64, error) {
	userDTO := dto.UserDTO{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	builderInsert := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(
			usersTableNameColumn,
			usersTableEmailColumn,
			usersTablePasswordColumn,
			usersTableRoleColumn,
		).
		Values(userDTO.Name, userDTO.Email, userDTO.Password, userDTO.Role).
		Suffix("RETURNING " + usersTableIDColumn)
	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "User.Insert",
		QueryRaw: query,
	}
	var userID int64
	err = p.db.DB().ScanOneContext(ctx, &userID, q, args...)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
