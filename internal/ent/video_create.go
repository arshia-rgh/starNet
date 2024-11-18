// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golang_template/internal/ent/video"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoCreate is the builder for creating a Video entity.
type VideoCreate struct {
	config
	mutation *VideoMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (vc *VideoCreate) SetTitle(s string) *VideoCreate {
	vc.mutation.SetTitle(s)
	return vc
}

// SetDescription sets the "description" field.
func (vc *VideoCreate) SetDescription(s string) *VideoCreate {
	vc.mutation.SetDescription(s)
	return vc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vc *VideoCreate) SetNillableDescription(s *string) *VideoCreate {
	if s != nil {
		vc.SetDescription(*s)
	}
	return vc
}

// SetFilePath sets the "file_path" field.
func (vc *VideoCreate) SetFilePath(s string) *VideoCreate {
	vc.mutation.SetFilePath(s)
	return vc
}

// SetUploadedAt sets the "uploaded_at" field.
func (vc *VideoCreate) SetUploadedAt(t time.Time) *VideoCreate {
	vc.mutation.SetUploadedAt(t)
	return vc
}

// SetNillableUploadedAt sets the "uploaded_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableUploadedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetUploadedAt(*t)
	}
	return vc
}

// Mutation returns the VideoMutation object of the builder.
func (vc *VideoCreate) Mutation() *VideoMutation {
	return vc.mutation
}

// Save creates the Video in the database.
func (vc *VideoCreate) Save(ctx context.Context) (*Video, error) {
	vc.defaults()
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideoCreate) SaveX(ctx context.Context) *Video {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VideoCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VideoCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VideoCreate) defaults() {
	if _, ok := vc.mutation.UploadedAt(); !ok {
		v := video.DefaultUploadedAt
		vc.mutation.SetUploadedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VideoCreate) check() error {
	if _, ok := vc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Video.title"`)}
	}
	if v, ok := vc.mutation.Title(); ok {
		if err := video.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Video.title": %w`, err)}
		}
	}
	if _, ok := vc.mutation.FilePath(); !ok {
		return &ValidationError{Name: "file_path", err: errors.New(`ent: missing required field "Video.file_path"`)}
	}
	if v, ok := vc.mutation.FilePath(); ok {
		if err := video.FilePathValidator(v); err != nil {
			return &ValidationError{Name: "file_path", err: fmt.Errorf(`ent: validator failed for field "Video.file_path": %w`, err)}
		}
	}
	if _, ok := vc.mutation.UploadedAt(); !ok {
		return &ValidationError{Name: "uploaded_at", err: errors.New(`ent: missing required field "Video.uploaded_at"`)}
	}
	return nil
}

func (vc *VideoCreate) sqlSave(ctx context.Context) (*Video, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VideoCreate) createSpec() (*Video, *sqlgraph.CreateSpec) {
	var (
		_node = &Video{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(video.Table, sqlgraph.NewFieldSpec(video.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.Title(); ok {
		_spec.SetField(video.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := vc.mutation.Description(); ok {
		_spec.SetField(video.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := vc.mutation.FilePath(); ok {
		_spec.SetField(video.FieldFilePath, field.TypeString, value)
		_node.FilePath = value
	}
	if value, ok := vc.mutation.UploadedAt(); ok {
		_spec.SetField(video.FieldUploadedAt, field.TypeTime, value)
		_node.UploadedAt = value
	}
	return _node, _spec
}

// VideoCreateBulk is the builder for creating many Video entities in bulk.
type VideoCreateBulk struct {
	config
	err      error
	builders []*VideoCreate
}

// Save creates the Video entities in the database.
func (vcb *VideoCreateBulk) Save(ctx context.Context) ([]*Video, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Video, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VideoCreateBulk) SaveX(ctx context.Context) []*Video {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VideoCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VideoCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}