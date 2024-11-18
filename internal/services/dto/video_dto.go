package dto

import (
	"mime/multipart"
	"time"
)

type Video struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FilePath    string    `json:"file_path"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

type VideoUpload struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	File        *multipart.FileHeader `json:"file"`
}
