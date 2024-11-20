package repositories

import (
	"context"
	"fmt"
	"golang_template/internal/database"
	VideoClient "golang_template/internal/ent/video"
	"golang_template/internal/services/dto"
)

type VideoRepository interface {
	GetAllVideos(ctx context.Context) ([]*dto.VideoResponse, error)
	GetVideoByTitle(ctx context.Context, video dto.Video) (*dto.VideoResponse, error)
	CreateVideo(ctx context.Context, video dto.Video) (*dto.VideoResponse, error)
}

type videoRepository struct {
	db database.Database
}

func NewVideoRepository(db database.Database) VideoRepository {
	return &videoRepository{db: db}
}

func (v *videoRepository) GetAllVideos(ctx context.Context) ([]*dto.VideoResponse, error) {
	videos, err := v.db.EntClient().Video.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var responseVideos []*dto.VideoResponse
	for _, v := range videos {
		responseVideo := dto.VideoResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			FilePath:    v.FilePath,
			UploadedAt:  v.UploadedAt,
		}

		responseVideos = append(responseVideos, &responseVideo)
	}
	return responseVideos, nil
}

func (v *videoRepository) GetVideoByTitle(ctx context.Context, video dto.Video) (*dto.VideoResponse, error) {
	dbVideo, err := v.db.EntClient().Video.Query().Where(VideoClient.TitleEQ(video.Title)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.VideoResponse{
		ID:          dbVideo.ID,
		Title:       dbVideo.Title,
		Description: dbVideo.Description,
		FilePath:    dbVideo.FilePath,
		UploadedAt:  dbVideo.UploadedAt,
	}, nil
}

func (v *videoRepository) CreateVideo(ctx context.Context, video dto.Video) (*dto.VideoResponse, error) {
	videoCreate := v.db.EntClient().Video.Create()
	if video.Title == "" {
		return nil, fmt.Errorf("title can not be empty\n")
	}
	videoCreate.SetTitle(video.Title)
	if video.Description != "" {
		videoCreate.SetDescription(video.Description)
	}
	videoCreate.SetFilePath(video.FilePath)

	dbVideo, err := videoCreate.Save(ctx)

	if err != nil {
		return nil, err
	}
	return &dto.VideoResponse{
		ID:          dbVideo.ID,
		Title:       dbVideo.Title,
		Description: dbVideo.Description,
		FilePath:    dbVideo.FilePath,
		UploadedAt:  dbVideo.UploadedAt,
	}, nil

}
