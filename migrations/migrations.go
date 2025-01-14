package migrations

import (
	"github.com/rilgilang/sticker-collection-api/internal/entities"
	"gorm.io/gorm"
)

var models = []interface{}{
	&entities.User{},
	&entities.Sticker{},
}

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "").AutoMigrate(models...)
}
