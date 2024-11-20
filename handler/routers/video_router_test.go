package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"golang_template/handler/controllers"
	"reflect"
	"testing"
)

type MockVideoController struct{}

func (mvc *MockVideoController) UploadVideo(c *fiber.Ctx) error {
	return nil
}

func (mvc *MockVideoController) PlayVideo(c *fiber.Ctx) error {
	return nil
}

func (mvc *MockVideoController) ShowAllVideos(c *fiber.Ctx) error {
	return nil
}

func TestNewVideoRouter(t *testing.T) {
	type args struct {
		controller controllers.VideoController
	}
	tests := []struct {
		name string
		args args
		want VideoRouter
	}{
		{
			name: "Test with valid controller",
			args: args{controller: &MockVideoController{}},
			want: &videoRouter{Controller: &MockVideoController{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVideoRouter(tt.args.controller); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVideoRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_videoRouter_AddProtectedRoutes(t *testing.T) {
	app := fiber.New()
	type fields struct {
		Controller controllers.VideoController
		app        *fiber.App
	}
	type args struct {
		router fiber.Router
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add protected routes",
			fields: fields{
				Controller: &MockVideoController{},
			},
			args: args{
				router: app.Group("/"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoRouter{
				Controller: tt.fields.Controller,
			}
			v.AddProtectedRoutes(tt.args.router)
			assert.Equal(t, 3, int(app.HandlersCount()))
		})
	}
}

func Test_videoRouter_AddPublicRoutes(t *testing.T) {
	app := fiber.New()
	type fields struct {
		Controller controllers.VideoController
	}
	type args struct {
		router fiber.Router
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add public routes",
			fields: fields{
				Controller: &MockVideoController{},
			},
			args: args{
				router: app.Group("/"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &videoRouter{
				Controller: tt.fields.Controller,
			}
			v.AddPublicRoutes(tt.args.router)
			assert.Equal(t, 2, int(app.HandlersCount()))
		})
	}
}
