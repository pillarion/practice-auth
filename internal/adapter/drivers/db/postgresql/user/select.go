package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	dto "github.com/pillarion/practice-auth/internal/core/dto/postgresql"
	desc "github.com/pillarion/practice-auth/internal/core/model/user"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
)

// SelectByID selects a user from the database based on the given ID.
//
// ctx - the context
// id - the user ID
// *desc.User, error - returns a user and an error
func (p *pg) SelectByID(ctx context.Context, id int64) (*desc.User, error) {
	ctx, span := p.tracer.Start(ctx, "User.Select")
	defer span.End()
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
	q := db.Query{
		Name:     "User.Select",
		QueryRaw: query,
	}
	var userDTO dto.UserDTO
	err = p.db.DB().ScanOneContext(ctx, &userDTO, q, args...)
	if err != nil {
		return nil, err
	}

	user := desc.User{
		Info: desc.Info{
			ID:       userDTO.ID,
			Name:     userDTO.Name,
			Email:    userDTO.Email,
			Password: userDTO.Password,
			Role:     userDTO.Role,
		},
		CreatedAt: userDTO.CreatedAt,
	}
	if userDTO.UpdatedAt.Valid {
		user.UpdatedAt = userDTO.UpdatedAt.Time
	}

	return &user, nil
}

// SelectByName selects a user by username.
//
// ctx context.Context, username string.
// *desc.User, error.
func (p *pg) SelectByName(ctx context.Context, username string) (*desc.User, error) {
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
		Where(sq.Eq{usersTableNameColumn: username})
	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "User.Select",
		QueryRaw: query,
	}
	var userDTO dto.UserDTO
	err = p.db.DB().ScanOneContext(ctx, &userDTO, q, args...)
	if err != nil {
		return nil, err
	}

	user := desc.User{
		Info: desc.Info{
			ID:       userDTO.ID,
			Name:     userDTO.Name,
			Email:    userDTO.Email,
			Password: userDTO.Password,
			Role:     userDTO.Role,
		},
		CreatedAt: userDTO.CreatedAt,
	}
	if userDTO.UpdatedAt.Valid {
		user.UpdatedAt = userDTO.UpdatedAt.Time
	}

	return &user, nil
}

func (p *pg) SelectPassword(ctx context.Context, username string) (string, error) {
	builderSelect := sq.Select(
		usersTablePasswordColumn,
	).
		From(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{usersTableNameColumn: username})
	query, args, err := builderSelect.ToSql()
	if err != nil {
		return "", err
	}
	q := db.Query{
		Name:     "Password.Select",
		QueryRaw: query,
	}
	var pass string
	err = p.db.DB().ScanOneContext(ctx, &pass, q, args...)
	if err != nil {
		return "", err
	}

	return pass, nil
}
