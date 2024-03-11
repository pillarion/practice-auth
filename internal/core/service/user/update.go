package user

import (
	"context"

	mjournal "github.com/pillarion/practice-auth/internal/core/model/journal"
	desc "github.com/pillarion/practice-auth/internal/core/model/user"
	"github.com/pillarion/practice-auth/internal/core/tools/password"
)

// Update updates the user information.
//
// ctx context.Context, user *desc.User
// error
func (s *service) Update(ctx context.Context, user *desc.Info) error {
	if user.Password != "" {
		todb, err := password.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = todb
	}

	err := s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			errTx := s.userRepo.Update(ctx, user)
			if errTx != nil {
				return errTx
			}

			_, errTx = s.journalRepo.Insert(ctx, &mjournal.Journal{
				Action: "User updated",
			})
			if errTx != nil {
				return errTx
			}

			return nil
		})
	if err != nil {
		return err
	}

	return nil
}
