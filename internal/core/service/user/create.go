package user

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

func (s service) Create(ctx context.Context, user *desc.User) (int64, error) {
	res, err := s.userRepo.InsertUser(ctx, user)

	if err != nil {
		return 0, err
	}

	return res, nil
}
