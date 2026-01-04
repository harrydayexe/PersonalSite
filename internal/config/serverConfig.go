package config

// Environment defines which environment the application is running in
type Environment string

const (
	Local      Environment = "local"
	Test       Environment = "test"
	Production Environment = "production"
)

// ServerConfig is a struct that holds the configuration for the server itself.
type ServerConfig struct {
	Environment  Environment `env:"ENVIRONMENT" envDefault:"local"`
	VerboseMode  bool        `env:"VERBOSE" envDefault:"false"`
	Port         int         `env:"PORT" envDefault:"8080"`
	ReadTimeout  int         `env:"READ_TIMEOUT" envDefault:"15"`
	WriteTimeout int         `env:"WRITE_TIMEOUT" envDefault:"15"`
	IdleTimeout  int         `env:"IDLE_TIMEOUT" envDefault:"60"`
}
