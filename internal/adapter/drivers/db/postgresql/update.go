package postgresql

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	dto "github.com/pillarion/practice-auth/internal/core/dto/postgresql"
	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

func (p *pg) UpdateUser(ctx context.Context, user *desc.User) error {

	userDTO := dto.UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	builderUpdate := sq.Update("users").
		PlaceholderFormat(sq.Dollar)

	if userDTO.Name != "" {
		builderUpdate = builderUpdate.Set("name", userDTO.Name)
	}
	if userDTO.Email != "" {
		builderUpdate = builderUpdate.Set("email", userDTO.Email)
	}
	if userDTO.Password != "" {
		builderUpdate = builderUpdate.Set("password", userDTO.Password)
	}
	if userDTO.Role != desc.RoleUnknown {
		builderUpdate = builderUpdate.Set("role", userDTO.Role)
	}

	builderUpdate = builderUpdate.Set("updated_at", time.Now()).
		Where(sq.Eq{"id": userDTO.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	_, err = p.pgx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
