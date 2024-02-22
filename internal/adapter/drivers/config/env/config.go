package config

import (
	"fmt"
	"os"
	"strconv"

	ecfg "github.com/pillarion/practice-auth/internal/core/entity/config"
)

const (
	rpcPortEnv = "GRPC_PORT"
	pgDBEnv    = "POSTGRES_DB"
	pgUserEnv  = "POSTGRES_USER"
	pgPassEnv  = "POSTGRES_PASSWORD"
	pgHostEnv  = "POSTGRES_HOST"
	pgPortEnv  = "POSTGRES_PORT"
)

// Get retrieves the configuration for the application.
//
// Returns *ecfg.Config, error.
func Get() (*ecfg.Config, error) {

	grpcPort, err := getEnv(rpcPortEnv)
	if err != nil {
		return nil, err
	}

	grpcPortInt, err := strconv.Atoi(grpcPort)
	if err != nil {
		return nil, err
	}

	pgdb, err := getEnv(pgDBEnv)
	if err != nil {
		return nil, err
	}

	pguser, err := getEnv(pgUserEnv)
	if err != nil {
		return nil, err
	}

	pgpass, err := getEnv(pgPassEnv)
	if err != nil {
		return nil, err
	}

	pghost, err := getEnv(pgHostEnv)
	if err != nil {
		return nil, err
	}

	pgport, err := getEnv(pgPortEnv)
	if err != nil {
		return nil, err
	}

	return &ecfg.Config{
		GRPC: ecfg.GRPC{
			Port: grpcPortInt,
		},
		Database: ecfg.Database{
			Db:   pgdb,
			User: pguser,
			Pass: pgpass,
			Host: pghost,
			Port: pgport,
		},
	}, nil
}

func getEnv(env string) (string, error) {
	val := os.Getenv(env)
	if val == "" {
		return "", fmt.Errorf("env %s is not set", env)
	}
	return val, nil
}
