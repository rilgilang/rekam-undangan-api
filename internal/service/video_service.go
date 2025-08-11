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
	"strings"
)

type VideoService interface {
	GetAllVideo(ctx context.Context) *presenter.Response
	DownloadVideo(ctx context.Context, uniqueId string) *presenter.Response
	ProcessVideo(ctx context.Context, uniqueId, url string) *presenter.Response
}

type videoService struct {
	videoRepo repositories.VideoRepository
	storage   pkg.Storage
	cache     pkg.Cache
	cfg       *dotenv.Config
}

func NewVideoService(videoRepo repositories.VideoRepository, storage pkg.Storage, cache pkg.Cache, cfg *dotenv.Config) VideoService {
	return &videoService{
		videoRepo: videoRepo,
		storage:   storage,
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
		log.Info(fmt.Sprintf(`error get processing the url: %s`, err))
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	return response.WithCode(200).WithData(videos)
}

func (s *videoService) DownloadVideo(ctx context.Context, uniqueId string) *presenter.Response {
	var (
		log      = logger.NewLog("download_video_service", s.cfg.LoggerEnable)
		response = presenter.Response{}
	)

	video, err := s.videoRepo.FetchOneByUniqueId(ctx, uniqueId)
	if err != nil {
		log.Info(fmt.Sprintf(`error get processing the url: %s`, err))
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	if video == nil {
		return response.WithCode(404).WithError(errors.New("video not found!"))
	}

	fileByte, _, contentType, err := s.storage.GetFile(ctx, video.URL)
	if err != nil {
		return response.WithCode(500).WithError(errors.New("something went wrong!"))
	}

	fileName := strings.Split(video.URL, "/")[len(strings.Split(video.URL, "/"))-1]

	return response.WithCode(200).WithStream(&fileByte, fileName, *contentType)
}

func (s *videoService) ProcessVideo(ctx context.Context, uniqueId, url string) *presenter.Response {
	var (
		log      = logger.NewLog("process_video_service", s.cfg.LoggerEnable)
		response = presenter.Response{}
	)

	payload := map[string]interface{}{
		"id":        "",
		"user_id":   "",
		"url":       url,
		"unique_id": uniqueId,
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
