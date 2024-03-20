package auth

import (
	config "github.com/pillarion/practice-auth/internal/core/entity/config"
	journalRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	userRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	authServicePort "github.com/pillarion/practice-auth/internal/core/port/service/auth"
	txManager "github.com/pillarion/practice-platform/pkg/pgtxmanager"
)

type service struct {
	userRepo    userRepoPort.Repo
	txManager   txManager.TxManager
	journalRepo journalRepoPort.Repo
	jwtConfig   config.JWT
}

// NewService initializes a new service with the given access repository.
func NewService(
	ur userRepoPort.Repo,
	txManager txManager.TxManager,
	jr journalRepoPort.Repo,
	cfg config.JWT,
) authServicePort.Service {
	return &service{
		userRepo:    ur,
		txManager:   txManager,
		journalRepo: jr,
		jwtConfig:   cfg,
	}
}
