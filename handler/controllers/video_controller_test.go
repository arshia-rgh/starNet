package controllers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"golang_template/internal/ent"
	"golang_template/internal/services"
	"golang_template/internal/services/dto"
	"log"
	"mime/multipart"
	"net/http/httptest"
	"reflect"
	"testing"
)

type MockVideoService struct {
	videos []dto.Video
}

func (m *MockVideoService) CreateVideo(ctx *fiber.Ctx, video dto.Video) (*ent.Video, error) {
	m.videos = append(m.videos, video)

	return &ent.Video{
		Title:       video.Title,
		Description: video.Description,
		FilePath:    video.Description,
	}, nil
}
func (m *MockVideoService) GetAllVideos(ctx *fiber.Ctx) ([]*ent.Video, error) {
	var v []*ent.Video
	for _, video := range m.videos {
		entVideo := &ent.Video{
			Title:       video.Title,
			Description: video.Description,
			FilePath:    video.FilePath,
		}
		v = append(v, entVideo)
	}
	if len(v) == 0 {
		return nil, &ent.NotFoundError{}
	}
	return v, nil
}
func (m *MockVideoService) GetVideoByTitle(ctx *fiber.Ctx, video dto.Video) (*ent.Video, error) {
	for _, v := range m.videos {
		if v.Title == video.Title {
			return &ent.Video{
				Title:       v.Title,
				Description: v.Description,
				FilePath:    v.FilePath,
			}, nil
		} else {
			return nil, &ent.NotFoundError{}
		}
	}
	return nil, &ent.NotFoundError{}
}

func TestNewVideoController(t *testing.T) {
	type args struct {
		videoService services.VideoService
	}
	tests := []struct {
		name string
		args args
		want VideoController
	}{
		{
			name: "test create new controller",
			args: args{videoService: &MockVideoService{}},
			want: NewVideoController(&MockVideoService{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVideoController(tt.args.videoService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVideoController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_videoController_PlayVideo(t *testing.T) {
	testVideo := dto.Video{
		Title:    "test",
		FilePath: "test path",
	}
	type fields struct {
		videoService services.VideoService
	}

	tests := []struct {
		name           string
		fields         fields
		forcedLoggedIn bool
		wantStatus     int
	}{
		{
			name:           "play video successfully",
			fields:         fields{videoService: &MockVideoService{videos: []dto.Video{testVideo}}},
			forcedLoggedIn: true,
			wantStatus:     200,
		},
		{
			name:           "play video with no video found",
			fields:         fields{videoService: &MockVideoService{}},
			forcedLoggedIn: true,
			wantStatus:     404,
		},
		{
			name:           "play video with not logged in",
			fields:         fields{videoService: &MockVideoService{}},
			forcedLoggedIn: false,
			wantStatus:     401,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			v := &videoController{
				videoService: tt.fields.videoService,
			}
			middleware := MockMiddleware{mockAuthMiddleware: MockAuthMiddleware{forceLoggedIn: tt.forcedLoggedIn}}
			authMiddleware := middleware.Auth()
			app.Use(authMiddleware.Handle())
			app.Get("/videos/:title/play", v.PlayVideo)

			req := httptest.NewRequest("GET", "/videos/test/play", nil)
			req.Header.Set("Content-Type", "application/json")

			res, err := app.Test(req, -1)
			if err != nil {
				log.Fatalln(err)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)

		})
	}
}

func Test_videoController_ShowAllVideos(t *testing.T) {
	testVideo1 := dto.Video{
		Title:    "test",
		FilePath: "test path",
	}
	testVideo2 := dto.Video{
		Title:    "test2",
		FilePath: "test path2",
	}
	testVideos := []dto.Video{testVideo1, testVideo2}
	type fields struct {
		videoService services.VideoService
	}

	tests := []struct {
		name       string
		fields     fields
		wantStatus int
	}{
		{
			name:       "get all videos successfully",
			fields:     fields{videoService: &MockVideoService{testVideos}},
			wantStatus: 200,
		},
		{
			name:       "get all videos with no videos found",
			fields:     fields{videoService: &MockVideoService{}},
			wantStatus: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			v := &videoController{
				videoService: tt.fields.videoService,
			}
			app.Get("/videos", v.ShowAllVideos)
			req := httptest.NewRequest("GET", "/videos", nil)
			req.Header.Set("Content-Type", "application/json")

			res, err := app.Test(req, -1)
			if err != nil {
				log.Fatalln(err)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}

func Test_videoController_UploadVideo(t *testing.T) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	_ = w.WriteField("title", "testvideo")
	_ = w.WriteField("chunk_number", "1")
	_ = w.WriteField("total_chunks", "1")
	part, err := w.CreateFormFile("file", "test_video.mp4")
	if err != nil {
		t.Fatalf("Error creating form file: %v", err)
	}
	part.Write([]byte("test video content"))
	w.Close()

	type fields struct {
		videoService services.VideoService
	}

	tests := []struct {
		name           string
		fields         fields
		body           bytes.Buffer
		forcedLoggedIn bool
		role           string
		wantStatus     int
	}{
		{
			name:           "upload video successfully",
			fields:         fields{videoService: &MockVideoService{}},
			body:           b,
			forcedLoggedIn: true,
			role:           "admin",
			wantStatus:     200,
		},
		{
			name:           "upload video with no admin role and logged in",
			fields:         fields{videoService: &MockVideoService{}},
			body:           b,
			forcedLoggedIn: true,
			role:           "anything but not admin",
			wantStatus:     401,
		},
		{
			name:           "upload video with wrong data",
			fields:         fields{videoService: &MockVideoService{}},
			body:           bytes.Buffer{},
			forcedLoggedIn: true,
			role:           "admin",
			wantStatus:     400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoController{
				videoService: tt.fields.videoService,
			}
			app := fiber.New()
			middleware := MockMiddleware{mockAuthMiddleware: MockAuthMiddleware{forceLoggedIn: tt.forcedLoggedIn, role: tt.role}}
			authMiddleware := middleware.Auth()
			app.Use(authMiddleware.Handle())
			app.Post("upload-video", v.UploadVideo)

			req := httptest.NewRequest("POST", "/upload-video", &tt.body)
			req.Header.Set("Content-Type", w.FormDataContentType())

			res, err := app.Test(req, -1)
			if err != nil {
				log.Fatalln(err)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)

		})
	}
}
