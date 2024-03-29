// Code generated by ent, DO NOT EDIT.

package media

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/media/discovery"
	"github.com/tikafog/of/dbc/media/predicate"
)

// DiscoveryDelete is the builder for deleting a Discovery entity.
type DiscoveryDelete struct {
	config
	hooks    []Hook
	mutation *DiscoveryMutation
}

// Where appends a list predicates to the DiscoveryDelete builder.
func (dd *DiscoveryDelete) Where(ps ...predicate.Discovery) *DiscoveryDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DiscoveryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dd.hooks) == 0 {
		affected, err = dd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscoveryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dd.mutation = mutation
			affected, err = dd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dd.hooks) - 1; i >= 0; i-- {
			if dd.hooks[i] == nil {
				return 0, fmt.Errorf("media: uninitialized hook (forgotten import media/runtime?)")
			}
			mut = dd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DiscoveryDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DiscoveryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: discovery.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: discovery.FieldID,
			},
		},
	}
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DiscoveryDeleteOne is the builder for deleting a single Discovery entity.
type DiscoveryDeleteOne struct {
	dd *DiscoveryDelete
}

// Exec executes the deletion query.
func (ddo *DiscoveryDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{discovery.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DiscoveryDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
