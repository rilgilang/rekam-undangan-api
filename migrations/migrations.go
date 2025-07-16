package migrations

import (
	"github.com/rilgilang/rekam-undangan-api/internal/entities"
	"gorm.io/gorm"
)

var models = []interface{}{
	&entities.User{},
	&entities.Video{},
}

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "").AutoMigrate(models...)
}
