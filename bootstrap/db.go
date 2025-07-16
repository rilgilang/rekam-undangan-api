package bootstrap

import (
	"fmt"
	"github.com/rilgilang/rekam-undangan-api/config/dotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// simple db connection
func DatabaseConnection(config *dotenv.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Jakarta`,
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
