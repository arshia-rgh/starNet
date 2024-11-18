package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique().NotEmpty(),
		field.String("description").Optional(),
		field.String("file_path").Unique().NotEmpty(),
		field.Time("uploaded_at").Default(time.Now()),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return nil
}
