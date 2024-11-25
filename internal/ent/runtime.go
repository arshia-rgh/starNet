// Code generated by ent, DO NOT EDIT.

package ent

import (
	"golang_template/internal/ent/schema"
	"golang_template/internal/ent/user"
	"golang_template/internal/ent/video"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	videoFields := schema.Video{}.Fields()
	_ = videoFields
	// videoDescTitle is the schema descriptor for title field.
	videoDescTitle := videoFields[0].Descriptor()
	// video.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	video.TitleValidator = videoDescTitle.Validators[0].(func(string) error)
	// videoDescFilePath is the schema descriptor for file_path field.
	videoDescFilePath := videoFields[2].Descriptor()
	// video.FilePathValidator is a validator for the "file_path" field. It is called by the builders before save.
	video.FilePathValidator = videoDescFilePath.Validators[0].(func(string) error)
	// videoDescUploadedAt is the schema descriptor for uploaded_at field.
	videoDescUploadedAt := videoFields[3].Descriptor()
	// video.DefaultUploadedAt holds the default value on creation for the uploaded_at field.
	video.DefaultUploadedAt = videoDescUploadedAt.Default.(time.Time)
}
