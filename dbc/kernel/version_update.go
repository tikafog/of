// Code generated by ent, DO NOT EDIT.

package kernel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/kernel/predicate"
	"github.com/tikafog/of/dbc/kernel/version"
)

// VersionUpdate is the builder for updating Version entities.
type VersionUpdate struct {
	config
	hooks    []Hook
	mutation *VersionMutation
}

// Where appends a list predicates to the VersionUpdate builder.
func (vu *VersionUpdate) Where(ps ...predicate.Version) *VersionUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetCurrent sets the "Current" field.
func (vu *VersionUpdate) SetCurrent(i int) *VersionUpdate {
	vu.mutation.ResetCurrent()
	vu.mutation.SetCurrent(i)
	return vu
}

// SetNillableCurrent sets the "Current" field if the given value is not nil.
func (vu *VersionUpdate) SetNillableCurrent(i *int) *VersionUpdate {
	if i != nil {
		vu.SetCurrent(*i)
	}
	return vu
}

// AddCurrent adds i to the "Current" field.
func (vu *VersionUpdate) AddCurrent(i int) *VersionUpdate {
	vu.mutation.AddCurrent(i)
	return vu
}

// SetLast sets the "Last" field.
func (vu *VersionUpdate) SetLast(i int) *VersionUpdate {
	vu.mutation.ResetLast()
	vu.mutation.SetLast(i)
	return vu
}

// SetNillableLast sets the "Last" field if the given value is not nil.
func (vu *VersionUpdate) SetNillableLast(i *int) *VersionUpdate {
	if i != nil {
		vu.SetLast(*i)
	}
	return vu
}

// AddLast adds i to the "Last" field.
func (vu *VersionUpdate) AddLast(i int) *VersionUpdate {
	vu.mutation.AddLast(i)
	return vu
}

// Mutation returns the VersionMutation object of the builder.
func (vu *VersionUpdate) Mutation() *VersionMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VersionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vu.hooks) == 0 {
		if err = vu.check(); err != nil {
			return 0, err
		}
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VersionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vu.check(); err != nil {
				return 0, err
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			if vu.hooks[i] == nil {
				return 0, fmt.Errorf("kernel: uninitialized hook (forgotten import kernel/runtime?)")
			}
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VersionUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VersionUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VersionUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VersionUpdate) check() error {
	if v, ok := vu.mutation.Current(); ok {
		if err := version.CurrentValidator(v); err != nil {
			return &ValidationError{Name: "Current", err: fmt.Errorf(`kernel: validator failed for field "Version.Current": %w`, err)}
		}
	}
	if v, ok := vu.mutation.Last(); ok {
		if err := version.LastValidator(v); err != nil {
			return &ValidationError{Name: "Last", err: fmt.Errorf(`kernel: validator failed for field "Version.Last": %w`, err)}
		}
	}
	return nil
}

func (vu *VersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   version.Table,
			Columns: version.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: version.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Current(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldCurrent,
		})
	}
	if value, ok := vu.mutation.AddedCurrent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldCurrent,
		})
	}
	if value, ok := vu.mutation.Last(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldLast,
		})
	}
	if value, ok := vu.mutation.AddedLast(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldLast,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{version.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// VersionUpdateOne is the builder for updating a single Version entity.
type VersionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VersionMutation
}

// SetCurrent sets the "Current" field.
func (vuo *VersionUpdateOne) SetCurrent(i int) *VersionUpdateOne {
	vuo.mutation.ResetCurrent()
	vuo.mutation.SetCurrent(i)
	return vuo
}

// SetNillableCurrent sets the "Current" field if the given value is not nil.
func (vuo *VersionUpdateOne) SetNillableCurrent(i *int) *VersionUpdateOne {
	if i != nil {
		vuo.SetCurrent(*i)
	}
	return vuo
}

// AddCurrent adds i to the "Current" field.
func (vuo *VersionUpdateOne) AddCurrent(i int) *VersionUpdateOne {
	vuo.mutation.AddCurrent(i)
	return vuo
}

// SetLast sets the "Last" field.
func (vuo *VersionUpdateOne) SetLast(i int) *VersionUpdateOne {
	vuo.mutation.ResetLast()
	vuo.mutation.SetLast(i)
	return vuo
}

// SetNillableLast sets the "Last" field if the given value is not nil.
func (vuo *VersionUpdateOne) SetNillableLast(i *int) *VersionUpdateOne {
	if i != nil {
		vuo.SetLast(*i)
	}
	return vuo
}

// AddLast adds i to the "Last" field.
func (vuo *VersionUpdateOne) AddLast(i int) *VersionUpdateOne {
	vuo.mutation.AddLast(i)
	return vuo
}

// Mutation returns the VersionMutation object of the builder.
func (vuo *VersionUpdateOne) Mutation() *VersionMutation {
	return vuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VersionUpdateOne) Select(field string, fields ...string) *VersionUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Version entity.
func (vuo *VersionUpdateOne) Save(ctx context.Context) (*Version, error) {
	var (
		err  error
		node *Version
	)
	if len(vuo.hooks) == 0 {
		if err = vuo.check(); err != nil {
			return nil, err
		}
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VersionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuo.check(); err != nil {
				return nil, err
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			if vuo.hooks[i] == nil {
				return nil, fmt.Errorf("kernel: uninitialized hook (forgotten import kernel/runtime?)")
			}
			mut = vuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, vuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Version)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from VersionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VersionUpdateOne) SaveX(ctx context.Context) *Version {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VersionUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VersionUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VersionUpdateOne) check() error {
	if v, ok := vuo.mutation.Current(); ok {
		if err := version.CurrentValidator(v); err != nil {
			return &ValidationError{Name: "Current", err: fmt.Errorf(`kernel: validator failed for field "Version.Current": %w`, err)}
		}
	}
	if v, ok := vuo.mutation.Last(); ok {
		if err := version.LastValidator(v); err != nil {
			return &ValidationError{Name: "Last", err: fmt.Errorf(`kernel: validator failed for field "Version.Last": %w`, err)}
		}
	}
	return nil
}

func (vuo *VersionUpdateOne) sqlSave(ctx context.Context) (_node *Version, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   version.Table,
			Columns: version.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: version.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`kernel: missing "Version.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, version.FieldID)
		for _, f := range fields {
			if !version.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("kernel: invalid field %q for query", f)}
			}
			if f != version.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Current(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldCurrent,
		})
	}
	if value, ok := vuo.mutation.AddedCurrent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldCurrent,
		})
	}
	if value, ok := vuo.mutation.Last(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldLast,
		})
	}
	if value, ok := vuo.mutation.AddedLast(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldLast,
		})
	}
	_node = &Version{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{version.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
