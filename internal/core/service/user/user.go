package user

import (
	jrepo "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	urepo "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	"github.com/pillarion/practice-auth/internal/core/port/service/user"
	pgclient "github.com/pillarion/practice-auth/internal/core/tools/pgclient/port"
)

type service struct {
	userRepo    urepo.Repo
	journalRepo jrepo.Repo
	txManager   pgclient.TxManager
}

// NewService initializes a new service with the given user repository.
//
// userRepo: the user repository for the service.
// returns: a UserService port.
func NewService(ur urepo.Repo, jr jrepo.Repo, txManager pgclient.TxManager) user.Service {
	return &service{
		userRepo:    ur,
		journalRepo: jr,
		txManager:   txManager,
	}
}
