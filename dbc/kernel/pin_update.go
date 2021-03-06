// Code generated by entc, DO NOT EDIT.

package kernel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/kernel/pin"
	"github.com/tikafog/of/dbc/kernel/predicate"
)

// PinUpdate is the builder for updating Pin entities.
type PinUpdate struct {
	config
	hooks    []Hook
	mutation *PinMutation
}

// Where appends a list predicates to the PinUpdate builder.
func (pu *PinUpdate) Where(ps ...predicate.Pin) *PinUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetRid sets the "rid" field.
func (pu *PinUpdate) SetRid(s string) *PinUpdate {
	pu.mutation.SetRid(s)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PinUpdate) SetStatus(s string) *PinUpdate {
	pu.mutation.SetStatus(s)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PinUpdate) SetNillableStatus(s *string) *PinUpdate {
	if s != nil {
		pu.SetStatus(*s)
	}
	return pu
}

// SetStep sets the "step" field.
func (pu *PinUpdate) SetStep(s string) *PinUpdate {
	pu.mutation.SetStep(s)
	return pu
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (pu *PinUpdate) SetNillableStep(s *string) *PinUpdate {
	if s != nil {
		pu.SetStep(*s)
	}
	return pu
}

// SetPriority sets the "priority" field.
func (pu *PinUpdate) SetPriority(i int) *PinUpdate {
	pu.mutation.ResetPriority()
	pu.mutation.SetPriority(i)
	return pu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (pu *PinUpdate) SetNillablePriority(i *int) *PinUpdate {
	if i != nil {
		pu.SetPriority(*i)
	}
	return pu
}

// AddPriority adds i to the "priority" field.
func (pu *PinUpdate) AddPriority(i int) *PinUpdate {
	pu.mutation.AddPriority(i)
	return pu
}

// SetRelate sets the "relate" field.
func (pu *PinUpdate) SetRelate(s string) *PinUpdate {
	pu.mutation.SetRelate(s)
	return pu
}

// SetNillableRelate sets the "relate" field if the given value is not nil.
func (pu *PinUpdate) SetNillableRelate(s *string) *PinUpdate {
	if s != nil {
		pu.SetRelate(*s)
	}
	return pu
}

// SetUpdatedUnix sets the "updated_unix" field.
func (pu *PinUpdate) SetUpdatedUnix(i int64) *PinUpdate {
	pu.mutation.ResetUpdatedUnix()
	pu.mutation.SetUpdatedUnix(i)
	return pu
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (pu *PinUpdate) SetNillableUpdatedUnix(i *int64) *PinUpdate {
	if i != nil {
		pu.SetUpdatedUnix(*i)
	}
	return pu
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (pu *PinUpdate) AddUpdatedUnix(i int64) *PinUpdate {
	pu.mutation.AddUpdatedUnix(i)
	return pu
}

// SetComment sets the "comment" field.
func (pu *PinUpdate) SetComment(s string) *PinUpdate {
	pu.mutation.SetComment(s)
	return pu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (pu *PinUpdate) SetNillableComment(s *string) *PinUpdate {
	if s != nil {
		pu.SetComment(*s)
	}
	return pu
}

// ClearComment clears the value of the "comment" field.
func (pu *PinUpdate) ClearComment() *PinUpdate {
	pu.mutation.ClearComment()
	return pu
}

// Mutation returns the PinMutation object of the builder.
func (pu *PinUpdate) Mutation() *PinMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PinUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("kernel: uninitialized hook (forgotten import kernel/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PinUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PinUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PinUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pin.Table,
			Columns: pin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pin.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Rid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldRid,
		})
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldStatus,
		})
	}
	if value, ok := pu.mutation.Step(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldStep,
		})
	}
	if value, ok := pu.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pin.FieldPriority,
		})
	}
	if value, ok := pu.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pin.FieldPriority,
		})
	}
	if value, ok := pu.mutation.Relate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldRelate,
		})
	}
	if value, ok := pu.mutation.UpdatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: pin.FieldUpdatedUnix,
		})
	}
	if value, ok := pu.mutation.AddedUpdatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: pin.FieldUpdatedUnix,
		})
	}
	if value, ok := pu.mutation.Comment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldComment,
		})
	}
	if pu.mutation.CommentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pin.FieldComment,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PinUpdateOne is the builder for updating a single Pin entity.
type PinUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PinMutation
}

// SetRid sets the "rid" field.
func (puo *PinUpdateOne) SetRid(s string) *PinUpdateOne {
	puo.mutation.SetRid(s)
	return puo
}

// SetStatus sets the "status" field.
func (puo *PinUpdateOne) SetStatus(s string) *PinUpdateOne {
	puo.mutation.SetStatus(s)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PinUpdateOne) SetNillableStatus(s *string) *PinUpdateOne {
	if s != nil {
		puo.SetStatus(*s)
	}
	return puo
}

// SetStep sets the "step" field.
func (puo *PinUpdateOne) SetStep(s string) *PinUpdateOne {
	puo.mutation.SetStep(s)
	return puo
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (puo *PinUpdateOne) SetNillableStep(s *string) *PinUpdateOne {
	if s != nil {
		puo.SetStep(*s)
	}
	return puo
}

// SetPriority sets the "priority" field.
func (puo *PinUpdateOne) SetPriority(i int) *PinUpdateOne {
	puo.mutation.ResetPriority()
	puo.mutation.SetPriority(i)
	return puo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (puo *PinUpdateOne) SetNillablePriority(i *int) *PinUpdateOne {
	if i != nil {
		puo.SetPriority(*i)
	}
	return puo
}

// AddPriority adds i to the "priority" field.
func (puo *PinUpdateOne) AddPriority(i int) *PinUpdateOne {
	puo.mutation.AddPriority(i)
	return puo
}

// SetRelate sets the "relate" field.
func (puo *PinUpdateOne) SetRelate(s string) *PinUpdateOne {
	puo.mutation.SetRelate(s)
	return puo
}

// SetNillableRelate sets the "relate" field if the given value is not nil.
func (puo *PinUpdateOne) SetNillableRelate(s *string) *PinUpdateOne {
	if s != nil {
		puo.SetRelate(*s)
	}
	return puo
}

// SetUpdatedUnix sets the "updated_unix" field.
func (puo *PinUpdateOne) SetUpdatedUnix(i int64) *PinUpdateOne {
	puo.mutation.ResetUpdatedUnix()
	puo.mutation.SetUpdatedUnix(i)
	return puo
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (puo *PinUpdateOne) SetNillableUpdatedUnix(i *int64) *PinUpdateOne {
	if i != nil {
		puo.SetUpdatedUnix(*i)
	}
	return puo
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (puo *PinUpdateOne) AddUpdatedUnix(i int64) *PinUpdateOne {
	puo.mutation.AddUpdatedUnix(i)
	return puo
}

// SetComment sets the "comment" field.
func (puo *PinUpdateOne) SetComment(s string) *PinUpdateOne {
	puo.mutation.SetComment(s)
	return puo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (puo *PinUpdateOne) SetNillableComment(s *string) *PinUpdateOne {
	if s != nil {
		puo.SetComment(*s)
	}
	return puo
}

// ClearComment clears the value of the "comment" field.
func (puo *PinUpdateOne) ClearComment() *PinUpdateOne {
	puo.mutation.ClearComment()
	return puo
}

// Mutation returns the PinMutation object of the builder.
func (puo *PinUpdateOne) Mutation() *PinMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PinUpdateOne) Select(field string, fields ...string) *PinUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pin entity.
func (puo *PinUpdateOne) Save(ctx context.Context) (*Pin, error) {
	var (
		err  error
		node *Pin
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("kernel: uninitialized hook (forgotten import kernel/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PinUpdateOne) SaveX(ctx context.Context) *Pin {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PinUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PinUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PinUpdateOne) sqlSave(ctx context.Context) (_node *Pin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pin.Table,
			Columns: pin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pin.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`kernel: missing "Pin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pin.FieldID)
		for _, f := range fields {
			if !pin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("kernel: invalid field %q for query", f)}
			}
			if f != pin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Rid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldRid,
		})
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldStatus,
		})
	}
	if value, ok := puo.mutation.Step(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldStep,
		})
	}
	if value, ok := puo.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pin.FieldPriority,
		})
	}
	if value, ok := puo.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pin.FieldPriority,
		})
	}
	if value, ok := puo.mutation.Relate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldRelate,
		})
	}
	if value, ok := puo.mutation.UpdatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: pin.FieldUpdatedUnix,
		})
	}
	if value, ok := puo.mutation.AddedUpdatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: pin.FieldUpdatedUnix,
		})
	}
	if value, ok := puo.mutation.Comment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pin.FieldComment,
		})
	}
	if puo.mutation.CommentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pin.FieldComment,
		})
	}
	_node = &Pin{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
