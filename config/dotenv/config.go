package dotenv

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName       string `env:"APP_NAME"`
	AppPort       int    `env:"APP_PORT"`
	DBDialect     string `env:"DB_DIALECT"`
	DBHost        string `env:"DB_HOST"`
	DBPort        int    `env:"DB_PORT"`
	DBName        string `env:"DB_NAME"`
	DBUsername    string `env:"DB_USERNAME"`
	DBPassword    string `env:"DB_PASSWORD"`
	JWTKey        string `env:"JWT_KEY"`
	JWTExpiredMin int    `env:"JWT_EXPIRED_MINUTE"`
	LoggerEnable  bool   `env:"LOGGER_ENABLE"`
}

func NewLoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// parse
	var cfg Config
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
