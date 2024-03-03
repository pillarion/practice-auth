package app

import (
	"context"
	"fmt"
	"log"

	guser "github.com/pillarion/practice-auth/internal/adapter/controller/grpc"
	config "github.com/pillarion/practice-auth/internal/adapter/drivers/config/env"
	pgjd "github.com/pillarion/practice-auth/internal/adapter/drivers/db/postgresql/journal"
	pgud "github.com/pillarion/practice-auth/internal/adapter/drivers/db/postgresql/user"
	cfg "github.com/pillarion/practice-auth/internal/core/entity/config"
	rpj "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	rpuser "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	spuser "github.com/pillarion/practice-auth/internal/core/port/service/user"
	suser "github.com/pillarion/practice-auth/internal/core/service/user"
	pgcd "github.com/pillarion/practice-auth/internal/core/tools/pgclient/adapter"
	db "github.com/pillarion/practice-auth/internal/core/tools/pgclient/port"
	pgcs "github.com/pillarion/practice-auth/internal/core/tools/pgclient/service"
)

type serviceProvider struct {
	config *cfg.Config

	dbDriver          db.DB
	dbClient          db.Client
	txManager         db.TxManager
	userRepository    rpuser.Repo
	journalRepository rpj.Repo

	userService spuser.Service

	userServer *guser.Server
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *cfg.Config {
	if s.config == nil {
		cfg, err := config.Get()
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) DBDriver(ctx context.Context) db.DB {
	if s.dbDriver == nil {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			s.Config().Database.Host,
			s.Config().Database.Port,
			s.Config().Database.User,
			s.Config().Database.Db,
			s.Config().Database.Pass,
		)
		db, err := pgcd.NewDB(ctx, dsn)
		if err != nil {
			log.Fatalf("failed to create db driver: %v", err)
		}

		s.dbDriver = db
	}

	return s.dbDriver
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pgcs.New(s.DBDriver(ctx))
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		//closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = pgcd.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) rpuser.Repo {
	if s.userRepository == nil {
		repo, err := pgud.New(s.DBClient(ctx))
		if err != nil {
			log.Fatalf("failed to create user repository: %v", err)
		}

		s.userRepository = repo
	}

	return s.userRepository
}

func (s *serviceProvider) JournalRepository(ctx context.Context) rpj.Repo {
	if s.journalRepository == nil {
		repo, err := pgjd.New(s.DBClient(ctx))
		if err != nil {
			log.Fatalf("failed to create user repository: %v", err)
		}

		s.journalRepository = repo
	}

	return s.journalRepository
}

func (s *serviceProvider) UserService(ctx context.Context) spuser.Service {
	if s.userService == nil {
		service := suser.NewService(s.UserRepository(ctx), s.JournalRepository(ctx), s.TxManager(ctx))

		s.userService = service
	}

	return s.userService
}

func (s *serviceProvider) UserServer(ctx context.Context) *guser.Server {
	if s.userServer == nil {
		server := guser.NewServer(s.UserService(ctx))

		s.userServer = server
	}

	return s.userServer
}
