package repositories

import (
	"context"
	"fmt"
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

	if len(tag) > 0 {
		queryTags := ` tag::jsonb @> '[`
		for i, s := range tag {
			queryTags += fmt.Sprintf(`"%s",`, s)
			if i+1 == len(tag) {
				queryTags += fmt.Sprintf(`"%s"`, s)
			}
		}
		queryTags += `]'`

		err := r.db.WithContext(ctx).Raw(fmt.Sprintf(`select * from stickers s where %s order by random() limit 1;`, queryTags)).First(&sticker).Error

		if err != nil {
			if err.Error() == consts.SqlNoRow {
				return nil, nil
			}
			return nil, err
		}

		return &sticker, nil
	}

	err := r.db.WithContext(ctx).Raw(`select * from stickers s order by random() limit 1;`).First(&sticker).Error

	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	return &sticker, nil
}
