package repositories

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"golang_template/internal/database"
	"golang_template/internal/services/dto"
	"log"
)

const collectionNameVideo = "videos"

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
	exists, err := v.db.DB().CollectionExists(ctx, collectionNameVideo)
	if err != nil {
		return nil, fmt.Errorf("failed checking collection existense, %w", err)
	}
	if !exists {
		// Computed values for uploaded_at field be auto
		computedValues := []arangodb.ComputedValue{
			{
				Name:       "uploaded_at",
				Expression: "RETURN DATE_NOW()",
				Overwrite:  true,
				ComputeOn:  []arangodb.ComputeOn{"insert"},
			},
		}

		properties := arangodb.CreateCollectionProperties{
			Schema:         nil,
			ComputedValues: computedValues,
		}

		if _, err = v.db.DB().CreateCollection(ctx, collectionNameVideo, &properties); err != nil {
			return nil, fmt.Errorf("creating collection: %w", err)
		}

		// Unique indexes on title and file_path
		coll, err := v.db.DB().Collection(ctx, collectionNameVideo)
		if err != nil {
			return nil, fmt.Errorf("opening collection: %w", err)
		}
		unique := true
		sparse := true

		if _, _, err := coll.EnsurePersistentIndex(ctx, []string{"title"}, &arangodb.CreatePersistentIndexOptions{
			Sparse: &sparse,
			Unique: &unique,
		}); err != nil {
			return nil, fmt.Errorf("creating index on title: %w", err)
		}

		if _, _, err := coll.EnsurePersistentIndex(ctx, []string{"file_path"}, &arangodb.CreatePersistentIndexOptions{
			Sparse: &sparse,
			Unique: &unique,
		}); err != nil {
			return nil, fmt.Errorf("creating index on file_path: %w", err)
		}
	}

	coll, err := v.db.DB().Collection(ctx, collectionNameVideo)
	if err != nil {
		return nil, fmt.Errorf("opening collection: %w", err)
	}

	meta, err := coll.CreateDocument(ctx, video)
	if err != nil {
		return nil, fmt.Errorf("creating video: %w", err)
	}
	log.Printf("Created document with key '%s'\n", meta.Key)
	video.Key = meta.Key

	return &video, nil

}
