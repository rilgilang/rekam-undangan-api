package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rilgilang/kosan-api/config/dotenv"
	"github.com/rilgilang/kosan-api/internal/api/presenter"
	"github.com/rilgilang/kosan-api/internal/entities"
	"github.com/rilgilang/kosan-api/internal/pkg/logger"
	"github.com/rilgilang/kosan-api/internal/repositories"
)

type PaymentHistoryService interface {
	GetAllPaymentHistoriesByRoomId(ctx context.Context, roomId string) *presenter.Response
	CreatePaymentHistory(ctx context.Context, payload entities.CreatePaymentHistoryPayload) *presenter.Response
}

type paymentHistoryService struct {
	paymentHistoryRepo repositories.PaymentHistory
	roomRepo           repositories.RoomRepository
	cfg                *dotenv.Config
}

func NewPaymentHistoryService(paymentHistoryRepo repositories.PaymentHistory, roomRepo repositories.RoomRepository, cfg *dotenv.Config) PaymentHistoryService {
	return &paymentHistoryService{
		paymentHistoryRepo: paymentHistoryRepo,
		roomRepo:           roomRepo,
		cfg:                cfg,
	}
}

func (s *paymentHistoryService) GetAllPaymentHistoriesByRoomId(ctx context.Context, roomId string) *presenter.Response {
	var (
		log      = logger.NewLog("payment_histories_handler", s.cfg.LoggerEnable)
		response = presenter.Response{}
	)

	paymentHistory, err := s.paymentHistoryRepo.FetchAllByRoomID(ctx, roomId)

	if err != nil {
		log.Error(fmt.Sprintf(`error fetching payment history data to db got %s`, err))
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	return response.WithCode(200).WithData(paymentHistory)
}

func (s *paymentHistoryService) CreatePaymentHistory(ctx context.Context, payload entities.CreatePaymentHistoryPayload) *presenter.Response {
	var (
		log      = logger.NewLog("payment_histories_handler", s.cfg.LoggerEnable)
		response = presenter.Response{}
	)

	room, err := s.roomRepo.FetchOne(ctx, payload.RoomID)
	if err != nil {
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	if room == nil {
		return response.WithCode(404).WithError(errors.New("room not found!"))
	}

	_, err = s.paymentHistoryRepo.Create(ctx, &payload)

	if err != nil {
		log.Error(fmt.Sprintf(`error create payment history data to db got %s`, err))
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	stay, err := s.roomRepo.ExtendStay(ctx, room.ID, room.CheckIn.AddDate(0, 0, 30), room.CheckIn.AddDate(0, 0, 60))
	if err != nil {
		return nil
	}

	return response.WithCode(200).WithData(stay)
}
