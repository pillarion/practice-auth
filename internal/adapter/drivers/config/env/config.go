package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	ecfg "github.com/pillarion/practice-auth/internal/core/entity/config"
)

const (
	rpcPortEnv               = "GRPC_PORT"
	pgDBEnv                  = "POSTGRES_DB"
	pgUserEnv                = "POSTGRES_USER"
	pgPassEnv                = "POSTGRES_PASSWORD"
	pgHostEnv                = "POSTGRES_HOST"
	pgPortEnv                = "POSTGRES_PORT"
	httpPortEnv              = "HTTP_PORT"
	swagerPortEnv            = "SWAGGER_PORT"
	tlsCertEnv               = "TLS_CERT"
	tlsKeyEnv                = "TLS_KEY"
	tlsCAEnv                 = "TLS_CA"
	tlsPathEnv               = "TLS_PATH"
	jwtAccessDurationEnv     = "JWT_ACCESS_DURATION"
	jwtRefreshDurationEnv    = "JWT_REFRESH_DURATION"
	jwtSecretEnv             = "JWT_SECRET"
	metricsPortEnv           = "METRICS_PORT"
	metricsNamespaceEnv      = "METRICS_NAMESPACE"
	serviceNameEnv           = "SERVICE_NAME"
	traceCollectorAddressEnv = "TRACE_COLLECTOR_ADDRESS"
	logLevelEnv              = "LOG_LEVEL"
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

	jwtAccessDuration, err := getEnv(jwtAccessDurationEnv)
	if err != nil {
		return nil, err
	}

	jwtAccessDur, err := time.ParseDuration(jwtAccessDuration)
	if err != nil {
		return nil, err
	}

	jwtRefreshDuration, err := getEnv(jwtRefreshDurationEnv)
	if err != nil {
		return nil, err
	}

	jwtRefreshDur, err := time.ParseDuration(jwtRefreshDuration)
	if err != nil {
		return nil, err
	}

	jwtSecret, err := getEnv(jwtSecretEnv)
	if err != nil {
		return nil, err
	}

	metricsPort, err := getEnv(metricsPortEnv)
	if err != nil {
		return nil, err
	}

	metricsPortInt, err := strconv.Atoi(metricsPort)
	if err != nil {
		return nil, err
	}

	metricsNamespace, err := getEnv(metricsNamespaceEnv)
	if err != nil {
		return nil, err
	}

	serviceName, err := getEnv(serviceNameEnv)
	if err != nil {
		return nil, err
	}

	traceCollectorAddress, err := getEnv(traceCollectorAddressEnv)
	if err != nil {
		return nil, err
	}

	logLevel, err := getEnv(logLevelEnv)
	if err != nil {
		return nil, err
	}

	return &ecfg.Config{
		ServiceName: serviceName,
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
		JWT: ecfg.JWT{
			Secret:          jwtSecret,
			AccessDuration:  jwtAccessDur,
			RefreshDuration: jwtRefreshDur,
		},
		Metrics: ecfg.Metrics{
			Port:      metricsPortInt,
			Namespace: metricsNamespace,
		},
		Trace: ecfg.Trace{
			CollectorAddress: traceCollectorAddress,
		},
		Log: ecfg.Log{
			Level: logLevel,
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
