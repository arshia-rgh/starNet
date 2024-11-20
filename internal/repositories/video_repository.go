package repositories

import (
	"context"
	"fmt"
	"golang_template/internal/database"
	"golang_template/internal/ent"
	VideoClient "golang_template/internal/ent/video"
	"golang_template/internal/services/dto"
)

type VideoRepository interface {
	GetAllVideos(ctx context.Context) ([]*ent.Video, error)
	GetVideoByTitle(ctx context.Context, video dto.Video) (*ent.Video, error)
	CreateVideo(ctx context.Context, video dto.Video) (*ent.Video, error)
}

type videoRepository struct {
	db database.Database
}

func NewVideoRepository(db database.Database) VideoRepository {
	return &videoRepository{db: db}
}

func (v *videoRepository) GetAllVideos(ctx context.Context) ([]*ent.Video, error) {
	return v.db.EntClient().Video.Query().All(ctx)
}

func (v *videoRepository) GetVideoByTitle(ctx context.Context, video dto.Video) (*ent.Video, error) {
	return v.db.EntClient().Video.Query().Where(VideoClient.TitleEQ(video.Title)).Only(ctx)

}

func (v *videoRepository) CreateVideo(ctx context.Context, video dto.Video) (*ent.Video, error) {
	videoCreate := v.db.EntClient().Video.Create()
	if video.Title == "" {
		return nil, fmt.Errorf("title can not be empty\n")
	}
	videoCreate.SetTitle(video.Title)
	if video.Description != "" {
		videoCreate.SetDescription(video.Description)
	}
	videoCreate.SetFilePath(video.FilePath)

	return videoCreate.Save(ctx)

}
