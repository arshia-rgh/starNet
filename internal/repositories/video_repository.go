package repositories

import (
	"context"
	"golang_template/internal/database"
	"golang_template/internal/services/dto"
)

type VideoRepository interface {
	GetAllVideos(ctx context.Context, video dto.Video) ([]*dto.Video, error)
	GetVideoByTitle(ctx context.Context, video dto.Video) (*dto.Video, error)
	CreateVideo(ctx context.Context, video dto.Video) (*dto.Video, error)
}

type videoRepository struct {
	db database.Database
}

func NewVideoRepository(db database.Database) VideoRepository {
	return &videoRepository{db: db}
}

func (v *videoRepository) GetAllVideos(ctx context.Context, video dto.Video) ([]*dto.Video, error) {

}

func (v *videoRepository) GetVideoByTitle(ctx context.Context, video dto.Video) (*dto.Video, error) {

}

func (v *videoRepository) CreateVideo(ctx context.Context, video dto.Video) (*dto.Video, error) {

}
