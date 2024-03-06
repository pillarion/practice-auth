package user

import (
	"context"

	mjournal "github.com/pillarion/practice-auth/internal/core/model/journal"
	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

func (s service) Get(ctx context.Context, id int64) (*desc.User, error) {
	var res *desc.User
	err := s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			var errTx error
			res, errTx = s.userRepo.Select(ctx, id)
			if errTx != nil {
				return errTx
			}

			_, errTx = s.journalRepo.Insert(ctx, &mjournal.Journal{
				Action: "User get",
			})
			if errTx != nil {
				return errTx
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}
