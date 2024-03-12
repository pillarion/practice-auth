package config

import (
	"fmt"
	"os"
	"strconv"

	ecfg "github.com/pillarion/practice-auth/internal/core/entity/config"
)

const (
	rpcPortEnv    = "GRPC_PORT"
	pgDBEnv       = "POSTGRES_DB"
	pgUserEnv     = "POSTGRES_USER"
	pgPassEnv     = "POSTGRES_PASSWORD"
	pgHostEnv     = "POSTGRES_HOST"
	pgPortEnv     = "POSTGRES_PORT"
	httpPortEnv   = "HTTP_PORT"
	swagerPortEnv = "SWAGGER_PORT"
	tlsCertEnv    = "TLS_CERT"
	tlsKeyEnv     = "TLS_KEY"
	tlsCAEnv      = "TLS_CA"
	tlsPathEnv    = "TLS_PATH"
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

	httpPort, err := getEnv(httpPortEnv)
	if err != nil {
		return nil, err
	}

	httpPortInt, err := strconv.Atoi(httpPort)
	if err != nil {
		return nil, err
	}

	swagerPort, err := getEnv(swagerPortEnv)
	if err != nil {
		return nil, err
	}

	swagerPortInt, err := strconv.Atoi(swagerPort)
	if err != nil {
		return nil, err
	}

	tlsPath, err := getEnv(tlsPathEnv)
	if err != nil {
		return nil, err
	}

	tlsCAcert, err := getEnv(tlsCAEnv)
	if err != nil {
		return nil, err
	}

	tlsCert, err := getEnv(tlsCertEnv)
	if err != nil {
		return nil, err
	}

	tlsKey, err := getEnv(tlsKeyEnv)
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
		HTTP: ecfg.HTTP{
			Port: httpPortInt,
		},
		Swagger: ecfg.Swagger{
			Port: swagerPortInt,
		},
		TLS: ecfg.TLS{
			Path: tlsPath,
			CA:   tlsCAcert,
			Cert: tlsCert,
			Key:  tlsKey,
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
