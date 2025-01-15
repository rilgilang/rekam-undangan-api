package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rilgilang/sticker-collection-api/config/dotenv"
	"github.com/rilgilang/sticker-collection-api/internal/api/dto"
	"github.com/rilgilang/sticker-collection-api/internal/api/presenter"
	"github.com/rilgilang/sticker-collection-api/internal/pkg/logger"
	"github.com/rilgilang/sticker-collection-api/internal/repositories"
)

type StickerService interface {
	GetOneRandomSticker(ctx context.Context, tag []string) *presenter.Response
}

type stickerService struct {
	stickerRepo repositories.StickerRepository
	cfg         *dotenv.Config
}

func NewStickerService(stickerRepo repositories.StickerRepository, cfg *dotenv.Config) StickerService {
	return &stickerService{
		stickerRepo: stickerRepo,
		cfg:         cfg,
	}
}

func (s *stickerService) GetOneRandomSticker(ctx context.Context, tag []string) *presenter.Response {
	var (
		log      = logger.NewLog("sticker_handler", s.cfg.LoggerEnable)
		response = presenter.Response{}
	)

	log.Info("fetching sticker data from db")

	sticker, err := s.stickerRepo.FindOneRandom(ctx, nil)

	if err != nil {
		log.Error(fmt.Sprintf(`error fetching sticker data to db got %s`, err))
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	if sticker == nil {
		log.Warn("sticker not found in db")
		return response.WithCode(404).WithError(errors.New("sticker not found"))
	}

	data := dto.OneSticker{
		ID:        sticker.ID,
		Url:       sticker.Url,
		Tag:       sticker.Tag,
		CreatedAt: sticker.CreatedAt,
		UpdatedAt: sticker.UpdatedAt,
	}

	return response.WithCode(200).WithData(data)
}
