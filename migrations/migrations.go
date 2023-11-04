package migrations

import (
	"digital_sekuriti_indonesia/internal/entities"
	"gorm.io/gorm"
)

var models = []interface{}{
	&entities.User{},
}

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
}
