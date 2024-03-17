package config

import "time"

// Config holds the configuration for the application.
type Config struct {
	GRPC     GRPC     `yaml:"grpc"`
	Database Database `yaml:"database"`
	HTTP     HTTP     `yaml:"http"`
	Swagger  Swagger  `yaml:"swagger"`
	TLS      TLS      `yaml:"tls"`
	JWT      JWT      `yaml:"jwt"`
}

// GRPC holds the configuration for the gRPC server.
type GRPC struct {
	Port int `yaml:"port"`
}

// HTTP holds the configuration for the HTTP server.
type HTTP struct {
	Port int `yaml:"port"`
}

// Swagger holds the configuration for the Swagger server.
type Swagger struct {
	Port int `yaml:"port"`
}

// Database holds the configuration for the database.
type Database struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Db   string `yaml:"db"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

// TLS holds the configuration for the TLS server.
type TLS struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
	CA   string `yaml:"ca"`
	Path string `yaml:"path"`
}

type JWT struct {
	Secret          string        `yaml:"secret"`
	AccessDuration  time.Duration `yaml:"accessDuration"`
	RefreshDuration time.Duration `yaml:"refreshDuration"`
}
