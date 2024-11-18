package dto

import "time"

type Video struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FilePath    string    `json:"file_path"`
	UploadedAt  time.Time `json:"uploaded_at"`
}
