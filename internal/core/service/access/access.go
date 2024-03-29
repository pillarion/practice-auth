package access

import (
	config "github.com/pillarion/practice-auth/internal/core/entity/config"
	accessRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/access"
	journalRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	userRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	accessServicePort "github.com/pillarion/practice-auth/internal/core/port/service/access"
	txManager "github.com/pillarion/practice-platform/pkg/pgtxmanager"
)

// Service defines the access service.
type service struct {
	accessRepo  accessRepoPort.Repo
	userRepo    userRepoPort.Repo
	txManager   txManager.TxManager
	journalRepo journalRepoPort.Repo
	jwtConfig   config.JWT
}

// NewService initializes a new service with the given access repository.
func NewService(
	ar accessRepoPort.Repo,
	ur userRepoPort.Repo,
	txManager txManager.TxManager,
	jr journalRepoPort.Repo,
	cfg config.JWT,
) accessServicePort.Service {
	return &service{
		accessRepo:  ar,
		userRepo:    ur,
		txManager:   txManager,
		journalRepo: jr,
		jwtConfig:   cfg,
	}
}
