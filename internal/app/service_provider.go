package app

import (
	"context"
	"fmt"
	"log"

	grpcAccessController "github.com/pillarion/practice-auth/internal/adapter/controller/access_grpc"
	grpcAuthController "github.com/pillarion/practice-auth/internal/adapter/controller/auth_grpc"
	grpcUserController "github.com/pillarion/practice-auth/internal/adapter/controller/user_grpc"
	configDriver "github.com/pillarion/practice-auth/internal/adapter/drivers/config/env"
	pgAccessDriver "github.com/pillarion/practice-auth/internal/adapter/drivers/db/postgresql/access"
	pgJournalDriver "github.com/pillarion/practice-auth/internal/adapter/drivers/db/postgresql/journal"
	pgUserDriver "github.com/pillarion/practice-auth/internal/adapter/drivers/db/postgresql/user"
	config "github.com/pillarion/practice-auth/internal/core/entity/config"
	accessRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/access"
	journalRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	userRepoPort "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	accessServicePort "github.com/pillarion/practice-auth/internal/core/port/service/access"
	authServicePort "github.com/pillarion/practice-auth/internal/core/port/service/auth"
	userServicePort "github.com/pillarion/practice-auth/internal/core/port/service/user"
	accessService "github.com/pillarion/practice-auth/internal/core/service/access"
	authService "github.com/pillarion/practice-auth/internal/core/service/auth"
	userService "github.com/pillarion/practice-auth/internal/core/service/user"
	grpcAccessPort "github.com/pillarion/practice-auth/pkg/access_v1"
	grpcAuthPort "github.com/pillarion/practice-auth/pkg/auth_v1"
	grpcUserPort "github.com/pillarion/practice-auth/pkg/user_v1"
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
	accessRepository  accessRepoPort.Repo

	userService   userServicePort.Service
	accessService accessServicePort.Service
	authService   authServicePort.Service

	userServer   grpcUserPort.UserV1Server
	authServer   grpcAuthPort.AuthV1Server
	accessServer grpcAccessPort.AccessV1Server
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

func (s *serviceProvider) AccessRepository(ctx context.Context) accessRepoPort.Repo {
	if s.accessRepository == nil {
		repo, err := pgAccessDriver.New(s.DBClient(ctx))
		if err != nil {
			log.Fatalf("failed to create access repository: %v", err)
		}

		s.accessRepository = repo
	}

	return nil
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

func (s *serviceProvider) AccessService(ctx context.Context) accessServicePort.Service {
	if s.accessService == nil {
		service := accessService.NewService(
			s.AccessRepository(ctx),
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.JournalRepository(ctx),
		)

		s.accessService = service
	}

	return s.accessService
}

func (s *serviceProvider) AuhtService(ctx context.Context) authServicePort.Service {
	if s.authService == nil {
		service := authService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.JournalRepository(ctx),
		)

		s.authService = service
	}

	return s.authService
}

func (s *serviceProvider) UserServer(ctx context.Context) grpcUserPort.UserV1Server {
	if s.userServer == nil {
		server := grpcUserController.NewServer(s.UserService(ctx))

		s.userServer = server
	}

	return s.userServer
}

func (s *serviceProvider) AccessServer(ctx context.Context) grpcAccessPort.AccessV1Server {
	if s.accessServer == nil {
		server := grpcAccessController.NewServer(s.AccessService(ctx))

		s.accessServer = server
	}

	return s.accessServer
}

func (s *serviceProvider) AuthServer(ctx context.Context) grpcAuthPort.AuthV1Server {
	if s.authServer == nil {
		server := grpcAuthController.NewServer(s.AuhtService(ctx))

		s.authServer = server
	}

	return s.authServer
}
