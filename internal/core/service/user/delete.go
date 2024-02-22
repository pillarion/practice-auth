package user

import "context"

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.userRepo.DeleteUser(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
