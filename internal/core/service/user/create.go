package user

import (
	"context"

	mjournal "github.com/pillarion/practice-auth/internal/core/model/journal"
	muser "github.com/pillarion/practice-auth/internal/core/model/user"
	"github.com/pillarion/practice-auth/internal/core/tools/password"
)

// Create creates a new user.
//
// ctx: the context
// user: the user information to be inserted
// int64: the ID of the newly created user
// error: an error, if any
func (s *service) Create(ctx context.Context, user *muser.Info) (int64, error) {
	todb, err := password.Hash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = todb

	var res int64
	err = s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			var errTx error
			res, errTx = s.userRepo.Insert(ctx, user)
			if errTx != nil {
				return errTx
			}

			_, errTx = s.journalRepo.Insert(ctx, &mjournal.Journal{
				Action: "User created",
			})
			if errTx != nil {
				return errTx
			}

			return nil
		})
	if err != nil {
		return 0, err
	}

	return res, nil
}
