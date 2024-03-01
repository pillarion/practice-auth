package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	dto "github.com/pillarion/practice-auth/internal/core/dto/postgresql"
	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

// SelectUser selects a user from the database based on the given ID.
//
// ctx - the context
// id - the user ID
// *desc.User, error - returns a user and an error
func (p *pg) Select(ctx context.Context, id int64) (*desc.User, error) {
	builderSelect := sq.Select(
		usersTableIDColumn,
		usersTableNameColumn,
		usersTableEmailColumn,
		usersTablePasswordColumn,
		usersTableRoleColumn,
		usersTableCreatedAtColumn,
		usersTableUpdatedAtColumn,
	).
		From(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{usersTableIDColumn: id})
	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}
	var userDTO dto.UserDTO
	err = p.pgx.QueryRow(ctx, query, args...).
		Scan(
			&userDTO.ID,
			&userDTO.Name,
			&userDTO.Email,
			&userDTO.Password,
			&userDTO.Role,
			&userDTO.CreatedAt,
			&userDTO.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}

	user := desc.User{
		ID:        userDTO.ID,
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		Role:      userDTO.Role,
		CreatedAt: userDTO.CreatedAt,
		UpdatedAt: userDTO.UpdatedAt.Time,
	}

	return &user, nil
}
