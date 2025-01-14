package repositories

import (
	"context"
	"github.com/rilgilang/sticker-collection-api/internal/consts"
	"github.com/rilgilang/sticker-collection-api/internal/entities"
	"gorm.io/gorm"
)

type StickerRepository interface {
	FindOneRandom(ctx context.Context, tag []string) (*entities.Sticker, error)
}

type stickerRepository struct {
	db *gorm.DB
}

func NewStickerRepo(db *gorm.DB) StickerRepository {
	return &stickerRepository{
		db: db,
	}
}

func (r *stickerRepository) FindOneRandom(ctx context.Context, tag []string) (*entities.Sticker, error) {
	sticker := entities.Sticker{}

	err := r.db.WithContext(ctx).Raw(`select * from stickers s order by random() limit 1;`).First(&sticker).Error

	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	return &sticker, nil
}
