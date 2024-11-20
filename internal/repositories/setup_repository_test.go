package repositories

import (
	"context"
	"database/sql"
	"golang_template/internal/ent"
	"golang_template/internal/services/dto"
)

type MockDatabase struct {
	client *ent.Client
	db     *sql.DB
}

func (db *MockDatabase) Close() error {
	return nil
}

func (db *MockDatabase) EntClient() *ent.Client {
	return db.client
}

func (db *MockDatabase) DB() *sql.DB {
	return db.db
}

func setUpNewVideo(client *ent.Client) {
	video1 := dto.Video{
		Title:       "Test Video 1",
		Description: "Test Description 1",
		FilePath:    "/path/to/video1",
	}
	video2 := dto.Video{
		Title:       "Test Video 2",
		Description: "Test Description 2",
		FilePath:    "/path/to/video2",
	}
	videos := []dto.Video{video1, video2}
	videoCreateBulk := make([]*ent.VideoCreate, len(videos))
	for i, video := range videos {
		videoCreateBulk[i] = client.Video.Create().
			SetTitle(video.Title).
			SetDescription(video.Description).
			SetFilePath(video.FilePath)
	}
	client.Video.CreateBulk(videoCreateBulk...).Save(context.Background())
}
