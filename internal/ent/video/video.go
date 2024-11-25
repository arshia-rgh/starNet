// Code generated by ent, DO NOT EDIT.

package video

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the video type in the database.
	Label = "video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFilePath holds the string denoting the file_path field in the database.
	FieldFilePath = "file_path"
	// FieldUploadedAt holds the string denoting the uploaded_at field in the database.
	FieldUploadedAt = "uploaded_at"
	// Table holds the table name of the video in the database.
	Table = "videos"
)

// Columns holds all SQL columns for video fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldFilePath,
	FieldUploadedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// FilePathValidator is a validator for the "file_path" field. It is called by the builders before save.
	FilePathValidator func(string) error
	// DefaultUploadedAt holds the default value on creation for the "uploaded_at" field.
	DefaultUploadedAt time.Time
)

// OrderOption defines the ordering options for the Video queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByFilePath orders the results by the file_path field.
func ByFilePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFilePath, opts...).ToFunc()
}

// ByUploadedAt orders the results by the uploaded_at field.
func ByUploadedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUploadedAt, opts...).ToFunc()
}
