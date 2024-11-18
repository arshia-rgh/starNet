package repositories

import (
	"context"
	"errors"
	"fmt"
	"golang_template/internal/database"
	"golang_template/internal/ent"
	"golang_template/internal/ent/user"
	"golang_template/internal/services/dto"
	"log"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExists      = errors.New("username already exists")
	ErrInvalidPassword = errors.New("invalid password")
	ErrDatabaseError   = errors.New("database error")
)

type UserRepository interface {
	Get(ctx context.Context, userDto dto.User) (*ent.User, error)
	CreateUser(ctx context.Context, userData dto.User) (*ent.User, error)
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

func (r userRepository) Get(ctx context.Context, userDto dto.User) (*ent.User, error) {
	userData, err := r.db.EntClient().User.
		Query().
		Where(user.UsernameEQ(userDto.Username)).
		First(ctx)

	if ent.IsNotFound(err) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("querying user: %w", err)
	}

	return userData, nil
}

func (r userRepository) CreateUser(ctx context.Context, userData dto.User) (*ent.User, error) {

	// Create new user
	user, err := r.db.EntClient().User.
		Create().
		SetUsername(userData.Username).
		SetPassword(userData.Password).
		Save(ctx)

	log.Println(err)

	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	return user, nil
}
