package services

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang_template/internal/repositories"
	"golang_template/internal/services/dto"
	"reflect"
	"testing"
)

type MockVideoRepository struct {
	videos []dto.Video
}

func (mv *MockVideoRepository) GetAllVideos(ctx context.Context) ([]*dto.VideoResponse, error) {
	var videoResponses []*dto.VideoResponse
	for _, video := range mv.videos {
		videoResponses = append(videoResponses, &dto.VideoResponse{
			Title: video.Title,
			// others if necessary
		})
	}
	return videoResponses, nil
}

func (mv *MockVideoRepository) GetVideoByTitle(ctx context.Context, video dto.Video) (*dto.VideoResponse, error) {
	for _, v := range mv.videos {
		if v.Title == video.Title {
			return &dto.VideoResponse{
				Title: v.Title,
				// others if necessary
			}, nil
		}
	}
	return nil, fmt.Errorf("video not found")
}

func (mv *MockVideoRepository) CreateVideo(ctx context.Context, video dto.Video) (*dto.VideoResponse, error) {
	if video.Title == "" {
		return nil, fmt.Errorf("invalid video title")
	}
	mv.videos = append(mv.videos, video)
	return &dto.VideoResponse{
		Title: video.Title,
		// others if necessary
	}, nil
}

func TestNewVideoService(t *testing.T) {
	type args struct {
		repo repositories.VideoRepository
	}
	tests := []struct {
		name string
		args args
		want VideoService
	}{
		{
			name: "Test with valid repository",
			args: args{repo: &MockVideoRepository{}},
			want: &videoService{repo: &MockVideoRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVideoService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVideoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_videoService_CreateVideo(t *testing.T) {
	type fields struct {
		repo repositories.VideoRepository
	}
	type args struct {
		ctx   *fiber.Ctx
		video dto.Video
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.VideoResponse
		wantErr bool
	}{
		{
			name: "Create video successfully",
			fields: fields{
				repo: &MockVideoRepository{},
			},
			args: args{
				ctx:   &fiber.Ctx{},
				video: dto.Video{Title: "Test Video"},
			},
			want:    &dto.VideoResponse{Title: "Test Video"},
			wantErr: false,
		},
		{
			name: "Create video with error",
			fields: fields{
				repo: &MockVideoRepository{},
			},
			args: args{
				ctx:   &fiber.Ctx{},
				video: dto.Video{Title: ""},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoService{
				repo: tt.fields.repo,
			}
			got, err := v.CreateVideo(tt.args.ctx, tt.args.video)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateVideo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_videoService_GetAllVideos(t *testing.T) {
	video := dto.Video{
		Title: "Test Video",
	}
	type fields struct {
		repo repositories.VideoRepository
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*dto.VideoResponse
		wantErr bool
	}{
		{
			name: "Get all videos successfully",
			fields: fields{
				repo: &MockVideoRepository{videos: []dto.Video{video}},
			},
			args: args{
				ctx: &fiber.Ctx{},
			},
			want:    []*dto.VideoResponse{{Title: "Test Video"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoService{
				repo: tt.fields.repo,
			}
			got, err := v.GetAllVideos(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllVideos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_videoService_GetVideoByTitle(t *testing.T) {
	video := dto.Video{
		Title: "Test Video",
	}

	type fields struct {
		repo repositories.VideoRepository
	}
	type args struct {
		ctx   *fiber.Ctx
		video dto.Video
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.VideoResponse
		wantErr bool
	}{
		{
			name: "Get video by title successfully",
			fields: fields{
				repo: &MockVideoRepository{videos: []dto.Video{video}},
			},
			args: args{
				ctx:   &fiber.Ctx{},
				video: dto.Video{Title: "Test Video"},
			},
			want:    &dto.VideoResponse{Title: "Test Video"},
			wantErr: false,
		},
		{
			name: "Get video by title with error",
			fields: fields{
				repo: &MockVideoRepository{videos: []dto.Video{video}},
			},
			args: args{
				ctx:   &fiber.Ctx{},
				video: dto.Video{Title: ""},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoService{
				repo: tt.fields.repo,
			}
			got, err := v.GetVideoByTitle(tt.args.ctx, tt.args.video)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideoByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVideoByTitle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
