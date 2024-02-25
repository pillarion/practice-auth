package user

import (
	repo "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	"github.com/pillarion/practice-auth/internal/core/port/service/user"
)

type service struct {
	userRepo repo.Repo
}

// NewService initializes a new service with the given user repository.
//
// userRepo: the user repository for the service.
// returns: a UserService port.
func NewService(userRepo repo.Repo) user.Service {
	return &service{
		userRepo: userRepo,
	}
}
