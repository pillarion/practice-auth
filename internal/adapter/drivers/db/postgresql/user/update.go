package postgresql

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	dto "github.com/pillarion/practice-auth/internal/core/dto/postgresql"
	desc "github.com/pillarion/practice-auth/internal/core/model/user"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
)

func (p *pg) Update(ctx context.Context, user *desc.User) error {
	userDTO := dto.UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	builderUpdate := sq.Update(usersTable).
		PlaceholderFormat(sq.Dollar)
	if userDTO.Name != "" {
		builderUpdate = builderUpdate.Set(usersTableNameColumn, userDTO.Name)
	}
	if userDTO.Email != "" {
		builderUpdate = builderUpdate.Set(usersTableEmailColumn, userDTO.Email)
	}
	if userDTO.Password != "" {
		builderUpdate = builderUpdate.Set(usersTablePasswordColumn, userDTO.Password)
	}
	if userDTO.Role != desc.RoleUnknown {
		builderUpdate = builderUpdate.Set(usersTableRoleColumn, userDTO.Role)
	}
	builderUpdate = builderUpdate.Set(usersTableUpdatedAtColumn, time.Now()).
		Where(sq.Eq{usersTableIDColumn: userDTO.ID})
	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "User.Update",
		QueryRaw: query,
	}
	_, err = p.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
