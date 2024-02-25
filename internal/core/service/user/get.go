package user

import (
	"context"

	desc "github.com/pillarion/practice-auth/internal/core/model/user"
)

func (s service) Get(ctx context.Context, id int64) (*desc.User, error) {
	res, err := s.userRepo.Select(ctx, id)
	if err != nil {

		return nil, err
	}

	return res, nil
}
