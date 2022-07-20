// Code generated by entc, DO NOT EDIT.

package media

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/media/announce"
	"github.com/tikafog/of/dbc/media/predicate"
)

// AnnounceDelete is the builder for deleting a Announce entity.
type AnnounceDelete struct {
	config
	hooks    []Hook
	mutation *AnnounceMutation
}

// Where appends a list predicates to the AnnounceDelete builder.
func (ad *AnnounceDelete) Where(ps ...predicate.Announce) *AnnounceDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AnnounceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ad.hooks) == 0 {
		affected, err = ad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnounceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ad.mutation = mutation
			affected, err = ad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ad.hooks) - 1; i >= 0; i-- {
			if ad.hooks[i] == nil {
				return 0, fmt.Errorf("media: uninitialized hook (forgotten import media/runtime?)")
			}
			mut = ad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AnnounceDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AnnounceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: announce.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: announce.FieldID,
			},
		},
	}
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
}

// AnnounceDeleteOne is the builder for deleting a single Announce entity.
type AnnounceDeleteOne struct {
	ad *AnnounceDelete
}

// Exec executes the deletion query.
func (ado *AnnounceDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{announce.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AnnounceDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}