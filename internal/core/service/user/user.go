package user

import (
	journalRepo "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	userRepo "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	"github.com/pillarion/practice-auth/internal/core/port/service/user"
	txmanager "github.com/pillarion/practice-auth/internal/core/tools/dbclient/port/pgtxmanager"
)

type service struct {
	userRepo    userRepo.Repo
	journalRepo journalRepo.Repo
	txManager   txmanager.TxManager
}

// NewService initializes a new service with the given user repository.
//
// userRepo: the user repository for the service.
// returns: a UserService port.
func NewService(ur userRepo.Repo, jr journalRepo.Repo, txManager txmanager.TxManager) user.Service {
	return &service{
		userRepo:    ur,
		journalRepo: jr,
		txManager:   txManager,
	}
}
