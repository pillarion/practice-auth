package user

import (
	"context"

	mjournal "github.com/pillarion/practice-auth/internal/core/model/journal"
)

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			errTx := s.userRepo.Delete(ctx, id)
			if errTx != nil {
				return errTx
			}

			_, errTx = s.journalRepo.Insert(ctx, &mjournal.Journal{
				Action: "User deleted",
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
