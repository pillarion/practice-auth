package user

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
	"github.com/pillarion/practice-auth/internal/core/tools/password"
)

// Update updates the user information.
//
// ctx context.Context, user *desc.User
// error
func (s service) Update(ctx context.Context, user *desc.User) error {

	if user.Password != "" {
		todb, _, _, err := password.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = todb
	}

	err := s.userRepo.UpdateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
