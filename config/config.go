package config

// Config represents the server configurations
type Config struct {
	AppEnv  string
	AppPort string
	AppHost string
}

var Configurations = Config{
	AppEnv:  "DEV",
	AppPort: "8080",
	AppHost: "127.0.0.1:8080",
}
