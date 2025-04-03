package repositories

import (
	"context"
	"github.com/rilgilang/kosan-api/internal/consts"
	"github.com/rilgilang/kosan-api/internal/entities"
	"gorm.io/gorm"
	"time"
)

type RoomRepository interface {
	FetchAll(ctx context.Context) ([]entities.Room, error)
	FetchOne(ctx context.Context, roomId string) (*entities.Room, error)
	UpdateRenter(ctx context.Context, values map[string]string) (*entities.Room, error)
	ExtendStay(ctx context.Context, roomId string, checkin time.Time, checkout time.Time) (*entities.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepo(db *gorm.DB) RoomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) FetchAll(ctx context.Context) ([]entities.Room, error) {
	rooms := []entities.Room{}

	err := r.db.WithContext(ctx).Find(&rooms).Order("room_number ASC").Error

	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	return rooms, nil
}

func (r *roomRepository) FetchOne(ctx context.Context, roomId string) (*entities.Room, error) {
	room := entities.Room{}

	err := r.db.WithContext(ctx).First(&room, "id = ?", roomId).Error

	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	return &room, nil
}

func (r *roomRepository) UpdateRenter(ctx context.Context, values map[string]string) (*entities.Room, error) {
	room := entities.Room{}

	err := r.db.WithContext(ctx).Raw("UPDATE rooms SET renter = $1, id_card = $2 WHERE id = $3;", values["renter"], values["id_card"], values["id"]).Error

	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).First(&room, "id = ?", values["id"]).Error

	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	return &room, nil
}

func (r *roomRepository) ExtendStay(ctx context.Context, roomId string, checkin time.Time, checkout time.Time) (*entities.Room, error) {
	room := entities.Room{}

	// Use Model(&room) and Where before calling Updates
	err := r.db.WithContext(ctx).
		Model(&room).
		Where("id = ?", roomId).
		Updates(map[string]interface{}{
			"already_paid_this_month": true,
			"check_in":                checkin,
			"check_out":               checkout,
		}).Error

	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	err = r.db.WithContext(ctx).First(&room, "id = ?", roomId).Error

	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	return &room, nil
}
