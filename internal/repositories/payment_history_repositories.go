package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/rilgilang/kosan-api/internal/consts"
	"github.com/rilgilang/kosan-api/internal/entities"
	"gorm.io/gorm"
	"time"
)

type PaymentHistory interface {
	FetchAllByRoomID(ctx context.Context, roomId string) ([]entities.PaymentHistory, error)
	Create(ctx context.Context, payload *entities.CreatePaymentHistoryPayload) (*entities.PaymentHistory, error)
}

type paymentHistoryRepository struct {
	db *gorm.DB
}

func NewPaymentHistoryRepo(db *gorm.DB) PaymentHistory {
	return &paymentHistoryRepository{
		db: db,
	}
}

func (r *paymentHistoryRepository) FetchAllByRoomID(ctx context.Context, roomId string) ([]entities.PaymentHistory, error) {
	var histories []entities.PaymentHistory

	err := r.db.WithContext(ctx).Find(&histories, "room_id = ?", roomId).Order("created_at desc").Error
	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	return histories, nil
}

func (r *paymentHistoryRepository) Create(ctx context.Context, payload *entities.CreatePaymentHistoryPayload) (*entities.PaymentHistory, error) {
	paymentHistory := entities.PaymentHistory{}
	id := uuid.New().String()

	err := r.db.WithContext(ctx).Model(&paymentHistory).Create(
		map[string]interface{}{
			"id":              id,
			"total":           payload.Total,
			"payment_receipt": payload.PaymentReceipt,
			"room_id":         payload.RoomID,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		}).Error

	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).First(&paymentHistory, "id = ?", id).Error

	if err != nil {
		if err.Error() == consts.SqlNoRow {
			return nil, nil
		}
		return nil, err
	}

	return &paymentHistory, nil
}
