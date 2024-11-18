package repositories

import (
	"golang_template/internal/database"
)

type Repository interface {
	UserRepository() UserRepository
	VideoRepository() VideoRepository
}

// var (
// ErrGlobal = errors.New("some global error")
// )

type repository struct {
	userRepository  UserRepository
	videoRepository VideoRepository
}

func NewRepository(db database.Database) Repository {
	userRepository := NewUserRepository(db)
	videoRepository := NewVideoRepository(db)
	return &repository{userRepository: userRepository, videoRepository: videoRepository}
}

func (r *repository) UserRepository() UserRepository {
	return r.userRepository
}

func (r *repository) VideoRepository() VideoRepository { return r.videoRepository }
