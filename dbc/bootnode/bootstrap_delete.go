// Code generated by ent, DO NOT EDIT.

package bootnode

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/bootnode/bootstrap"
	"github.com/tikafog/of/dbc/bootnode/predicate"
)

// BootstrapDelete is the builder for deleting a Bootstrap entity.
type BootstrapDelete struct {
	config
	hooks    []Hook
	mutation *BootstrapMutation
}

// Where appends a list predicates to the BootstrapDelete builder.
func (bd *BootstrapDelete) Where(ps ...predicate.Bootstrap) *BootstrapDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BootstrapDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bd.hooks) == 0 {
		affected, err = bd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BootstrapMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bd.mutation = mutation
			affected, err = bd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bd.hooks) - 1; i >= 0; i-- {
			if bd.hooks[i] == nil {
				return 0, fmt.Errorf("bootnode: uninitialized hook (forgotten import bootnode/runtime?)")
			}
			mut = bd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BootstrapDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BootstrapDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: bootstrap.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bootstrap.FieldID,
			},
		},
	}
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// BootstrapDeleteOne is the builder for deleting a single Bootstrap entity.
type BootstrapDeleteOne struct {
	bd *BootstrapDelete
}

// Exec executes the deletion query.
func (bdo *BootstrapDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bootstrap.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BootstrapDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
