package repositories

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/connection"
	"golang_template/internal/config"
	"golang_template/internal/database"
	"golang_template/internal/services/dto"
	"log"
)

type MockDatabase struct {
	client arangodb.Client
	db     arangodb.Database
}

func (db *MockDatabase) Close() error {
	return nil
}

func (db *MockDatabase) Client() arangodb.Client {
	return db.client
}

func (db *MockDatabase) DB() arangodb.Database {
	return db.db
}

func setUpNewVideo(db arangodb.Database) {
	ctx := context.Background()
	col, err := db.Collection(ctx, "videos")
	if err != nil {
		log.Fatalf("Failed to get collection: %v", err)
	}

	videos := []dto.Video{
		{Title: "Test Video 1", Description: "Test Description 1", FilePath: "/path/to/video1"},
		{Title: "Test Video 2", Description: "Test Description 2", FilePath: "/path/to/video2"},
	}

	for _, video := range videos {
		_, err := col.CreateDocument(ctx, video)
		if err != nil {
			log.Fatalf("Failed to create document: %v", err)
		}
	}
}

func setupTestVideoDatabase() database.Database {
	cfg, err := config.LoadConfig("../../config/config.yml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	dbConfig := &cfg.DB
	if dbConfig == nil {
		log.Fatalf("dbconfig is nil")
	}

	conn := connection.NewHttp2Connection(connection.DefaultHTTP2ConfigurationWrapper(
		connection.NewRoundRobinEndpoints([]string{fmt.Sprintf("http://%v:%v", dbConfig.Host, dbConfig.Port)}),
		true,
	))
	auth := connection.NewBasicAuth(dbConfig.User, dbConfig.Password)
	err = conn.SetAuthentication(auth)
	if err != nil {
		log.Fatalf("failed to set auth %v", err)
	}
	client := arangodb.NewClient(conn)
	db, err := client.GetDatabase(context.Background(), "starnettest", &arangodb.GetDatabaseOptions{SkipExistCheck: false})
	if err != nil {
		log.Fatalf("failed to open database, %v", err.Error())
	}

	return &MockDatabase{
		db:     db,
		client: client,
	}
}

func tearDown(db database.Database, collectionName string) {
	coll, err := db.DB().Collection(context.Background(), collectionName)
	if err != nil {
		log.Fatalf("failed to open collection in teearDown function, err: %v", err)
	}
	err = coll.Truncate(context.Background())
	if err != nil {
		log.Fatalf("failed to truncate the collection in the tearDown function, err: %v", err)
	}
}
