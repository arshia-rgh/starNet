package repositories

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"golang_template/internal/database"
	"golang_template/internal/ent"
	"golang_template/internal/ent/enttest"
	"golang_template/internal/services/dto"
	"reflect"
	"testing"
	"time"
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

func TestNewVideoRepository(t *testing.T) {
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
			args: args{db: &MockDatabase{}},
			want: &videoRepository{db: &MockDatabase{}},
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
	testClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	mockDB := &MockDatabase{client: testClient}

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
		want    *ent.Video
		wantErr bool
	}{
		{
			name: "create video successfully",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.Background(),
				video: dto.Video{
					Title:       "Test Video",
					Description: "Test Description",
					FilePath:    "/path/to/video",
				},
			},
			want: &ent.Video{
				ID:          1,
				Title:       "Test Video",
				Description: "Test Description",
				FilePath:    "/path/to/video",
				UploadedAt:  time.Now(),
			},
			wantErr: false,
		},
		{
			name:   "create video with no title (required)",
			fields: fields{db: mockDB},
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
			fields: fields{db: mockDB},
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
			if tt.want != nil {
				assert.Equal(t, tt.want.ID, got.ID)
				assert.Equal(t, tt.want.Title, got.Title)
				assert.Equal(t, tt.want.Description, got.Description)
				assert.Equal(t, tt.want.FilePath, got.FilePath)
			} else {
				assert.Equal(t, tt.want, got)
			}

		})
	}
}

func Test_videoRepository_GetAllVideos(t *testing.T) {
	mockClient := &ent.Client{}
	mockDB := &MockDatabase{client: mockClient}

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
		want    []*ent.Video
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
			want: []*ent.Video{
				{
					Title:       "Test Video 1",
					Description: "Test Description 1",
					FilePath:    "/path/to/video1",
				},
				{
					Title:       "Test Video 2",
					Description: "Test Description 2",
					FilePath:    "/path/to/video2",
				},
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoRepository{
				db: tt.fields.db,
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

func Test_videoRepository_GetVideoByTitle(t *testing.T) {
	mockClient := &ent.Client{}
	mockDB := &MockDatabase{client: mockClient}

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
		want    *ent.Video
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
					Title: "Test Video",
				},
			},
			want: &ent.Video{
				Title:       "Test Video",
				Description: "Test Description",
				FilePath:    "/path/to/video",
			},
			wantErr: false,
		},
		// Add more test cases as needed
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVideoByTitle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
