package repositories

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"golang_template/internal/database"
	"golang_template/internal/services/dto"
	"log"
	"time"
)

const collectionNameVideo = "vidoes"

type VideoRepository interface {
	GetAllVideos(ctx context.Context) ([]*dto.Video, error)
	GetVideoByTitle(ctx context.Context, video dto.Video) (*dto.Video, error)
	CreateVideo(ctx context.Context, video dto.Video) (*dto.Video, error)
}

type videoRepository struct {
	db database.Database
}

func NewVideoRepository(db database.Database) VideoRepository {
	return &videoRepository{db: db}
}

func (v *videoRepository) GetAllVideos(ctx context.Context) ([]*dto.Video, error) {
	query := fmt.Sprintf("FOR v IN %s RETURN v", collectionNameVideo)
	cursor, err := v.db.DB().Query(ctx, query, nil)
	if err != nil {
		return nil, fmt.Errorf("querying video: %w", err)
	}
	defer cursor.Close()

	var responseVideos []*dto.Video
	for {
		var responseVideo dto.Video
		meta, err := cursor.ReadDocument(ctx, &responseVideos)
		if shared.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("reading video document: %w", err)
		}
		responseVideos = append(responseVideos, &responseVideo)
		log.Printf("Got document with key '%s' from query\n", meta.Key)

	}
	return responseVideos, nil
}

func (v *videoRepository) GetVideoByTitle(ctx context.Context, video dto.Video) (*dto.Video, error) {
	query := fmt.Sprintf("FOR v IN %v FILTER v.title == @title RETURN v", collectionNameVideo)
	bindVars := map[string]interface{}{
		"title": video.Title,
	}
	cursor, err := v.db.DB().Query(ctx, query, &arangodb.QueryOptions{
		BindVars: bindVars,
	})
	if err != nil {
		return nil, fmt.Errorf("querying video: %w", err)
	}
	defer cursor.Close()

	var dbVideo dto.Video
	for {
		meta, err := cursor.ReadDocument(ctx, &dbVideo)
		if shared.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("reading user document: %w", err)
		}
		log.Printf("Got document with key '%s' from query\n", meta.Key)
	}

	return &dbVideo, nil
}

func (v *videoRepository) CreateVideo(ctx context.Context, video dto.Video) (*dto.Video, error) {
	coll, err := v.db.DB().Collection(ctx, collectionNameVideo)
	if err != nil {
		return nil, fmt.Errorf("opening collection: %w", err)
	}

	video.UploadedAt = time.Now()

	meta, err := coll.CreateDocument(ctx, video)
	if err != nil {
		return nil, fmt.Errorf("creating video: %w", err)
	}
	log.Printf("Created document with key '%s'\n", meta.Key)
	video.Key = meta.Key

	return &video, nil

}
