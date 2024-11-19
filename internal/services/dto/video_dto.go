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
	Title       string                `json:"title" form:"title"`
	Description string                `json:"description" form:"description"`
	ChunkNumber int                   `json:"chunk_number" form:"chunk_number"`
	TotalChunk  int                   `json:"total_chunk" form:"total_chunk"`
	File        *multipart.FileHeader `json:"file" form:"file"`
}
