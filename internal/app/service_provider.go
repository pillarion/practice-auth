package app

import (
	"context"
	"fmt"
	"log"

	grpcUserController "github.com/pillarion/practice-auth/internal/adapter/controller/grpc"
	configDriver "github.com/pillarion/practice-auth/internal/adapter/drivers/config/env"
	pgJournalDriver "github.com/pillarion/practice-auth/internal/adapter/drivers/db/postgresql/journal"
	pgUserDriver "github.com/pillarion/practice-auth/internal/adapter/drivers/db/postgresql/user"
	config "github.com/pillarion/practice-auth/internal/core/entity/config"
	journalRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	userRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	userServicePort "github.com/pillarion/practice-auth/internal/core/port/service/user"
	userService "github.com/pillarion/practice-auth/internal/core/service/user"
	closer "github.com/pillarion/practice-platform/pkg/closer"
	pgClient "github.com/pillarion/practice-platform/pkg/dbclient"
	txManager "github.com/pillarion/practice-platform/pkg/pgtxmanager"

	// statik
	_ "github.com/pillarion/practice-auth/statik"
)

type serviceProvider struct {
	config *config.Config

	dbDriver          pgClient.DB
	dbClient          pgClient.Client
	txManager         txManager.TxManager
	userRepository    userRepoPort.Repo
	journalRepository journalRepoPort.Repo

	userService userServicePort.Service

	userServer *grpcUserController.Server
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		cfg, err := configDriver.Get()
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) DBDriver(ctx context.Context) pgClient.DB {
	if s.dbDriver == nil {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			s.Config().Database.Host,
			s.Config().Database.Port,
			s.Config().Database.User,
			s.Config().Database.Db,
			s.Config().Database.Pass,
		)
		db, err := pgClient.NewDB(ctx, dsn)
		if err != nil {
			log.Fatalf("failed to create db driver: %v", err)
		}

		s.dbDriver = db
	}

	return s.dbDriver
}

func (s *serviceProvider) DBClient(ctx context.Context) pgClient.Client {
	if s.dbClient == nil {
		cl, err := pgClient.New(s.DBDriver(ctx))
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) txManager.TxManager {
	if s.txManager == nil {
		s.txManager = txManager.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) userRepoPort.Repo {
	if s.userRepository == nil {
		repo, err := pgUserDriver.New(s.DBClient(ctx))
		if err != nil {
			log.Fatalf("failed to create user repository: %v", err)
		}

		s.userRepository = repo
	}

	return s.userRepository
}

func (s *serviceProvider) JournalRepository(ctx context.Context) journalRepoPort.Repo {
	if s.journalRepository == nil {
		repo, err := pgJournalDriver.New(s.DBClient(ctx))
		if err != nil {
			log.Fatalf("failed to create user repository: %v", err)
		}

		s.journalRepository = repo
	}

	return s.journalRepository
}

func (s *serviceProvider) UserService(ctx context.Context) userServicePort.Service {
	if s.userService == nil {
		service := userService.NewService(s.UserRepository(ctx), s.JournalRepository(ctx), s.TxManager(ctx))

		s.userService = service
	}

	return s.userService
}

func (s *serviceProvider) UserServer(ctx context.Context) *grpcUserController.Server {
	if s.userServer == nil {
		server := grpcUserController.NewServer(s.UserService(ctx))

		s.userServer = server
	}

	return s.userServer
}
