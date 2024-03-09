package config

// Config holds the configuration for the application.
type Config struct {
	GRPC     GRPC     `yaml:"grpc"`
	Database Database `yaml:"database"`
	HTTP     HTTP     `yaml:"http"`
	Swagger  Swagger  `yaml:"swagger"`
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
