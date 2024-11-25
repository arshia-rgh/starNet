package repositories

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"golang_template/internal/database"
	"golang_template/internal/services/dto"
	"reflect"
	"testing"
	"time"
)

func TestNewVideoRepository(t *testing.T) {
	db := setupTestVideoDatabase()
	type args struct {
		db database.Database
	}
	tests := []struct {
		name string
		args args
		want VideoRepository
	}{
		{
			name: "test new repository creation",
			args: args{db: db},
			want: &videoRepository{db: db},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVideoRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVideoRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_videoRepository_CreateVideo(t *testing.T) {
	db := setupTestVideoDatabase()
	defer tearDown(db, "videos")

	type fields struct {
		db database.Database
	}
	type args struct {
		ctx   context.Context
		video dto.Video
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Video
		wantErr bool
	}{
		{
			name: "create video successfully",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				video: dto.Video{
					Title:       "Test Video",
					Description: "Test Description",
					FilePath:    "/path/to/video",
				},
			},
			want: &dto.Video{
				Title:       "Test Video",
				Description: "Test Description",
				FilePath:    "/path/to/video",
				UploadedAt:  time.Now(),
			},
			wantErr: false,
		},
		{
			name:   "create video with no title (required)",
			fields: fields{db: db},
			args: args{
				ctx: context.Background(),
				video: dto.Video{
					// no title
					Description: "Test Description",
					FilePath:    "/path/to/video",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "create video with duplicated title",
			fields: fields{db: db},
			args: args{
				ctx: context.Background(),
				video: dto.Video{
					Title:       "Test Video",
					Description: "Test Description",
					FilePath:    "/path/to/video",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoRepository{
				db: tt.fields.db,
			}
			got, err := v.CreateVideo(tt.args.ctx, tt.args.video)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				assert.Equal(t, tt.want, got)
				return
			}
			assert.Equal(t, tt.want.Title, got.Title)
			assert.Equal(t, tt.want.Description, got.Description)
			assert.Equal(t, tt.want.FilePath, got.FilePath)

		})
	}
}

func Test_videoRepository_GetAllVideos(t *testing.T) {
	mockDB := setupTestVideoDatabase()
	defer tearDown(mockDB, "videos")

	type fields struct {
		db database.Database
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*dto.Video
		wantErr bool
	}{
		{
			name: "get all videos successfully",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.Background(),
			},
			want: []*dto.Video{
				{
					Title:       "Test Video 1",
					Description: "Test Description 1",
					FilePath:    "/path/to/video1",
					UploadedAt:  time.Now(),
				},
				{
					Title:       "Test Video 2",
					Description: "Test Description 2",
					FilePath:    "/path/to/video2",
					UploadedAt:  time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name:    "get all videos with no videos found",
			fields:  fields{db: mockDB},
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoRepository{
				db: tt.fields.db,
			}
			if tt.name == "get_all_videos_successfully" {
				// create test videos
				setUpNewVideo(mockDB.DB())
			}
			got, err := v.GetAllVideos(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				assert.Equal(t, tt.want, got)
				return
			}
			for i, video := range got {
				assert.Equal(t, tt.want[i].Title, video.Title)
				assert.Equal(t, tt.want[i].Description, video.Description)
				assert.Equal(t, tt.want[i].FilePath, video.FilePath)
				assert.WithinDuration(t, tt.want[i].UploadedAt, video.UploadedAt, time.Second)
			}
		})
	}
}

func Test_videoRepository_GetVideoByTitle(t *testing.T) {
	mockDB := setupTestVideoDatabase()
	defer tearDown(mockDB, "videos")
	// create test videos
	setUpNewVideo(mockDB.DB())

	type fields struct {
		db database.Database
	}
	type args struct {
		ctx   context.Context
		video dto.Video
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Video
		wantErr bool
	}{
		{
			name: "get video by title successfully",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.Background(),
				video: dto.Video{
					Title: "Test Video 1",
				},
			},
			want: &dto.Video{
				Title:       "Test Video 1",
				Description: "Test Description 1",
				FilePath:    "/path/to/video1",
				UploadedAt:  time.Now(),
			},
			wantErr: false,
		},
		{
			name:   "get video by title with wrong title",
			fields: fields{db: mockDB},
			args: args{
				ctx: context.Background(),
				video: dto.Video{
					Title: "some random wrong title",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoRepository{
				db: tt.fields.db,
			}

			got, err := v.GetVideoByTitle(tt.args.ctx, tt.args.video)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideoByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				assert.Equal(t, tt.want, got)
				return
			}
			assert.Equal(t, tt.want.Title, got.Title)
			assert.Equal(t, tt.want.Description, got.Description)
			assert.Equal(t, tt.want.FilePath, got.FilePath)
			assert.WithinDuration(t, tt.want.UploadedAt, got.UploadedAt, time.Second)
		})
	}
}
