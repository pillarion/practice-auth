package config

// Config holds the configuration for the application.
type Config struct {
	GRPC     GRPC     `yaml:"grpc"`
	Database Database `yaml:"database"`
}

// GRPC holds the configuration for the gRPC server.
type GRPC struct {
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
