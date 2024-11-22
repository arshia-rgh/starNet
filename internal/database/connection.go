package database

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver/v2/connection"
	"golang_template/internal/config"
	"time"

	"github.com/arangodb/go-driver/v2/arangodb"
)

type Database interface {
	Close() error
	Client() arangodb.Client
	DB() arangodb.Database
}

type database struct {
	database arangodb.Database
	client   arangodb.Client
}

func NewDatabase(ctx context.Context, dbConfig *config.DatabaseConfig) (Database, error) {

	if dbConfig == nil {
		return nil, fmt.Errorf("database config cannot be nil")
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn := connection.NewHttp2Connection(connection.Http2Configuration{
		Endpoint: connection.NewRoundRobinEndpoints([]string{fmt.Sprintf("http://%v:%v", dbConfig.Host, dbConfig.Port)}),
	})

	auth := connection.NewBasicAuth(dbConfig.User, dbConfig.Password)
	err := conn.SetAuthentication(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to set authentication: %v", err)
	}

	client := arangodb.NewClient(conn)
	db, err := client.Database(timeoutCtx, dbConfig.DBName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	return &database{
		database: db,
		client:   client,
	}, nil
}

func (db *database) Close() error {
	return nil
}

func (db *database) Client() arangodb.Client {
	return db.client
}

func (db *database) DB() arangodb.Database {
	return db.database
}
