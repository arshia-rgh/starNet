package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"golang_template/internal/database"
	"golang_template/internal/services/dto"
	"log"
)

const collectionNameUser = "users"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExists      = errors.New("username already exists")
	ErrInvalidPassword = errors.New("invalid password")
	ErrDatabaseError   = errors.New("database error")
)

type UserRepository interface {
	Get(ctx context.Context, userDto dto.User) (*dto.User, error)
	CreateUser(ctx context.Context, userData dto.User) (*dto.User, error)
}

type userRepository struct {
	db database.Database
}

func test(err error) {
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("user not found")
	}
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Get(ctx context.Context, userDto dto.User) (*dto.User, error) {
	query := fmt.Sprintf("FOR u IN %v FILTER u.username == @username RETURN u", collectionNameUser)
	bindVars := map[string]interface{}{
		"username": userDto.Username,
	}

	cursor, err := r.db.DB().Query(ctx, query, &arangodb.QueryOptions{
		BindVars: bindVars,
	})
	if err != nil {
		return nil, fmt.Errorf("querying user: %w", err)
	}
	defer cursor.Close()

	var user dto.User
	for {
		meta, err := cursor.ReadDocument(ctx, &user)
		if shared.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("reading user document: %w", err)
		}
		log.Printf("Got document with key '%s' from query\n", meta.Key)
	}

	if user.Username == "" {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, userData dto.User) (*dto.User, error) {
	exists, err := r.db.DB().CollectionExists(ctx, collectionNameUser)
	if err != nil {
		return nil, fmt.Errorf("failed checking collection existense, %w", err)
	}
	if !exists {
		// Computed values for role field be auto being set to the normal
		computedValues := []arangodb.ComputedValue{
			{
				Name:       "role",
				Expression: "RETURN 'normal'",
				Overwrite:  true,
				ComputeOn:  []arangodb.ComputeOn{"insert"},
			},
		}

		properties := arangodb.CreateCollectionProperties{
			Schema:         nil,
			ComputedValues: computedValues,
		}

		if _, err = r.db.DB().CreateCollection(ctx, collectionNameUser, &properties); err != nil {
			return nil, fmt.Errorf("creating collection: %w", err)
		}

		// Unique index on username
		coll, err := r.db.DB().Collection(ctx, collectionNameUser)
		if err != nil {
			return nil, fmt.Errorf("opening collection: %w", err)
		}
		unique := true
		sparse := true

		if _, _, err := coll.EnsurePersistentIndex(ctx, []string{"username"}, &arangodb.CreatePersistentIndexOptions{
			Sparse: &sparse,
			Unique: &unique,
		}); err != nil {
			return nil, fmt.Errorf("creating index on username: %w", err)
		}

	}
	col, err := r.db.DB().Collection(ctx, collectionNameUser)
	if err != nil {
		return nil, fmt.Errorf("opening collection: %w", err)
	}

	meta, err := col.CreateDocument(ctx, userData)
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}
	log.Printf("Created document with key '%s'\n", meta.Key)
	userData.Key = meta.Key
	userData.ID = string(meta.ID)

	return &userData, nil
}
