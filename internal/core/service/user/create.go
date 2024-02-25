package user

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
	"github.com/pillarion/practice-auth/internal/core/tools/password"
)

func (s service) Create(ctx context.Context, user *desc.User) (int64, error) {
	todb, err := password.Hash(user.Password)
	if err != nil {

		return 0, err
	}
	user.Password = todb

	res, err := s.userRepo.Insert(ctx, user)
	if err != nil {

		return 0, err
	}

	return res, nil
}
