package dotenv

import (
	"errors"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	AppName              string `env:"APP_NAME"`
	AppEnvironment       string `env:"APP_ENVIRONMENT"`
	AppPort              int    `env:"APP_PORT"`
	DBDialect            string `env:"DB_DIALECT"`
	DBHost               string `env:"DB_HOST"`
	DBPort               int    `env:"DB_PORT"`
	DBName               string `env:"DB_NAME"`
	DBUsername           string `env:"DB_USERNAME"`
	DBPassword           string `env:"DB_PASSWORD"`
	MinioEndpoint        string `env:"MINIO_ENDPOINT"`
	MinioAccessKey       string `env:"MINIO_ACCESS_KEY"`
	MinioSecretAccessKey string `env:"MINIO_SECRET_ACCESS_KEY"`
	JWTKey               string `env:"JWT_KEY"`
	JWTExpiredMin        int    `env:"JWT_EXPIRED_MINUTE"`
	LoggerEnable         bool   `env:"LOGGER_ENABLE"`
}

func NewLoadConfig() (*Config, error) {

	envrionment := os.Getenv("APP_ENVIRONMENT")

	if envrionment == "production" {

		appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
		if err != nil {
			return nil, errors.New("error convert env app port")
		}

		dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, errors.New("error convert env db port")
		}

		jwtExpiredMinute, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_MINUTE"))
		if err != nil {
			return nil, errors.New("error convert env db port")
		}

		cfg := &Config{
			AppName:              os.Getenv("APP_NAME"),
			AppPort:              appPort,
			DBDialect:            os.Getenv("DB_DIALECT"),
			DBHost:               os.Getenv("DB_HOST"),
			DBPort:               dbPort,
			DBName:               os.Getenv("DB_NAME"),
			DBUsername:           os.Getenv("DB_USERNAME"),
			DBPassword:           os.Getenv("DB_PASSWORD"),
			MinioEndpoint:        os.Getenv("MINIO_ENDPOINT"),
			MinioAccessKey:       os.Getenv("MINIO_ACCESS_KEY"),
			MinioSecretAccessKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
			JWTKey:               os.Getenv("JWT_EXPIRED_MINUTE"),
			JWTExpiredMin:        jwtExpiredMinute,
			LoggerEnable:         false,
		}

		return cfg, nil
	}

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
