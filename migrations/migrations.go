package migrations

import (
	"github.com/rilgilang/kosan-api/internal/entities"
	"gorm.io/gorm"
)

var models = []interface{}{
	&entities.Room{},
	&entities.PaymentHistory{},
}

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "").AutoMigrate(models...)
}
