package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rilgilang/rekam-undangan-api/config/dotenv"
	"github.com/rilgilang/rekam-undangan-api/internal/api/presenter"
	"github.com/rilgilang/rekam-undangan-api/internal/pkg"
	"github.com/rilgilang/rekam-undangan-api/internal/pkg/logger"
	"github.com/rilgilang/rekam-undangan-api/internal/repositories"
)

type VideoService interface {
	GetAllVideo(ctx context.Context) *presenter.Response
	ProcessVideo(ctx context.Context, url string) *presenter.Response
}

type videoService struct {
	videoRepo repositories.VideoRepository
	cache     pkg.Cache
	cfg       *dotenv.Config
}

func NewVideoService(videoRepo repositories.VideoRepository, cache pkg.Cache, cfg *dotenv.Config) VideoService {
	return &videoService{
		videoRepo: videoRepo,
		cache:     cache,
		cfg:       cfg,
	}
}

func (s *videoService) GetAllVideo(ctx context.Context) *presenter.Response {
	var (
		log      = logger.NewLog("process_video_service", s.cfg.LoggerEnable)
		response = presenter.Response{}
	)

	videos, err := s.videoRepo.FetchAll(ctx)
	if err != nil {
		fmt.Println("error wir --> ", err)
		log.Info(fmt.Sprintf(`error get processing the url: %s`, err))
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	return response.WithCode(200).WithData(videos)
}

func (s *videoService) ProcessVideo(ctx context.Context, url string) *presenter.Response {
	var (
		log      = logger.NewLog("process_video_service", s.cfg.LoggerEnable)
		response = presenter.Response{}
	)

	payload := map[string]interface{}{
		"id":      "",
		"user_id": "",
		"url":     url,
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("error wir --> ", err)
		log.Info(fmt.Sprintf(`error processing the url: %s`, err))
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	err = s.cache.Lpush(ctx, "video_queue", string(payloadBytes))

	if err != nil {
		fmt.Println("error wir --> ", err)
		log.Info(fmt.Sprintf(`error processing the url: %s`, err))
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	return response.WithCode(200).WithData("ok")
}
