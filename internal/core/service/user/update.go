package user

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

func (s service) Update(ctx context.Context, user *desc.User) error {
	err := s.userRepo.UpdateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
