package config

// Config holds the configuration for the application.
type Config struct {
	PortGRPC int    `json:"port_grpc"`
	DBuser   string `json:"db_user"`
	DBpass   string `json:"db_pass"`
	DBhost   string `json:"db_host"`
	DBport   string `json:"db_port"`
	DBname   string `json:"db_name"`
}
