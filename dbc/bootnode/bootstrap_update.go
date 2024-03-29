// Code generated by ent, DO NOT EDIT.

package bootnode

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/bootnode/bootstrap"
	"github.com/tikafog/of/dbc/bootnode/predicate"
)

// BootstrapUpdate is the builder for updating Bootstrap entities.
type BootstrapUpdate struct {
	config
	hooks    []Hook
	mutation *BootstrapMutation
}

// Where appends a list predicates to the BootstrapUpdate builder.
func (bu *BootstrapUpdate) Where(ps ...predicate.Bootstrap) *BootstrapUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetPid sets the "pid" field.
func (bu *BootstrapUpdate) SetPid(s string) *BootstrapUpdate {
	bu.mutation.SetPid(s)
	return bu
}

// SetAddrs sets the "addrs" field.
func (bu *BootstrapUpdate) SetAddrs(s []string) *BootstrapUpdate {
	bu.mutation.SetAddrs(s)
	return bu
}

// SetExpired sets the "expired" field.
func (bu *BootstrapUpdate) SetExpired(b bool) *BootstrapUpdate {
	bu.mutation.SetExpired(b)
	return bu
}

// SetNillableExpired sets the "expired" field if the given value is not nil.
func (bu *BootstrapUpdate) SetNillableExpired(b *bool) *BootstrapUpdate {
	if b != nil {
		bu.SetExpired(*b)
	}
	return bu
}

// SetLevel sets the "level" field.
func (bu *BootstrapUpdate) SetLevel(b bootstrap.Level) *BootstrapUpdate {
	bu.mutation.SetLevel(b)
	return bu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (bu *BootstrapUpdate) SetNillableLevel(b *bootstrap.Level) *BootstrapUpdate {
	if b != nil {
		bu.SetLevel(*b)
	}
	return bu
}

// SetServicePort sets the "service_port" field.
func (bu *BootstrapUpdate) SetServicePort(i int) *BootstrapUpdate {
	bu.mutation.ResetServicePort()
	bu.mutation.SetServicePort(i)
	return bu
}

// SetNillableServicePort sets the "service_port" field if the given value is not nil.
func (bu *BootstrapUpdate) SetNillableServicePort(i *int) *BootstrapUpdate {
	if i != nil {
		bu.SetServicePort(*i)
	}
	return bu
}

// AddServicePort adds i to the "service_port" field.
func (bu *BootstrapUpdate) AddServicePort(i int) *BootstrapUpdate {
	bu.mutation.AddServicePort(i)
	return bu
}

// SetFailCounts sets the "fail_counts" field.
func (bu *BootstrapUpdate) SetFailCounts(i int) *BootstrapUpdate {
	bu.mutation.ResetFailCounts()
	bu.mutation.SetFailCounts(i)
	return bu
}

// SetNillableFailCounts sets the "fail_counts" field if the given value is not nil.
func (bu *BootstrapUpdate) SetNillableFailCounts(i *int) *BootstrapUpdate {
	if i != nil {
		bu.SetFailCounts(*i)
	}
	return bu
}

// AddFailCounts adds i to the "fail_counts" field.
func (bu *BootstrapUpdate) AddFailCounts(i int) *BootstrapUpdate {
	bu.mutation.AddFailCounts(i)
	return bu
}

// SetSign sets the "sign" field.
func (bu *BootstrapUpdate) SetSign(s string) *BootstrapUpdate {
	bu.mutation.SetSign(s)
	return bu
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (bu *BootstrapUpdate) SetNillableSign(s *string) *BootstrapUpdate {
	if s != nil {
		bu.SetSign(*s)
	}
	return bu
}

// ClearSign clears the value of the "sign" field.
func (bu *BootstrapUpdate) ClearSign() *BootstrapUpdate {
	bu.mutation.ClearSign()
	return bu
}

// SetUpdatedUnix sets the "updated_unix" field.
func (bu *BootstrapUpdate) SetUpdatedUnix(i int64) *BootstrapUpdate {
	bu.mutation.ResetUpdatedUnix()
	bu.mutation.SetUpdatedUnix(i)
	return bu
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (bu *BootstrapUpdate) SetNillableUpdatedUnix(i *int64) *BootstrapUpdate {
	if i != nil {
		bu.SetUpdatedUnix(*i)
	}
	return bu
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (bu *BootstrapUpdate) AddUpdatedUnix(i int64) *BootstrapUpdate {
	bu.mutation.AddUpdatedUnix(i)
	return bu
}

// Mutation returns the BootstrapMutation object of the builder.
func (bu *BootstrapUpdate) Mutation() *BootstrapMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BootstrapUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BootstrapMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("bootnode: uninitialized hook (forgotten import bootnode/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BootstrapUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BootstrapUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BootstrapUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BootstrapUpdate) check() error {
	if v, ok := bu.mutation.Level(); ok {
		if err := bootstrap.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`bootnode: validator failed for field "Bootstrap.level": %w`, err)}
		}
	}
	return nil
}

func (bu *BootstrapUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bootstrap.Table,
			Columns: bootstrap.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bootstrap.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Pid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bootstrap.FieldPid,
		})
	}
	if value, ok := bu.mutation.Addrs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: bootstrap.FieldAddrs,
		})
	}
	if value, ok := bu.mutation.Expired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: bootstrap.FieldExpired,
		})
	}
	if value, ok := bu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: bootstrap.FieldLevel,
		})
	}
	if value, ok := bu.mutation.ServicePort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bootstrap.FieldServicePort,
		})
	}
	if value, ok := bu.mutation.AddedServicePort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bootstrap.FieldServicePort,
		})
	}
	if value, ok := bu.mutation.FailCounts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bootstrap.FieldFailCounts,
		})
	}
	if value, ok := bu.mutation.AddedFailCounts(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bootstrap.FieldFailCounts,
		})
	}
	if value, ok := bu.mutation.Sign(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bootstrap.FieldSign,
		})
	}
	if bu.mutation.SignCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: bootstrap.FieldSign,
		})
	}
	if value, ok := bu.mutation.UpdatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: bootstrap.FieldUpdatedUnix,
		})
	}
	if value, ok := bu.mutation.AddedUpdatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: bootstrap.FieldUpdatedUnix,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bootstrap.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BootstrapUpdateOne is the builder for updating a single Bootstrap entity.
type BootstrapUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BootstrapMutation
}

// SetPid sets the "pid" field.
func (buo *BootstrapUpdateOne) SetPid(s string) *BootstrapUpdateOne {
	buo.mutation.SetPid(s)
	return buo
}

// SetAddrs sets the "addrs" field.
func (buo *BootstrapUpdateOne) SetAddrs(s []string) *BootstrapUpdateOne {
	buo.mutation.SetAddrs(s)
	return buo
}

// SetExpired sets the "expired" field.
func (buo *BootstrapUpdateOne) SetExpired(b bool) *BootstrapUpdateOne {
	buo.mutation.SetExpired(b)
	return buo
}

// SetNillableExpired sets the "expired" field if the given value is not nil.
func (buo *BootstrapUpdateOne) SetNillableExpired(b *bool) *BootstrapUpdateOne {
	if b != nil {
		buo.SetExpired(*b)
	}
	return buo
}

// SetLevel sets the "level" field.
func (buo *BootstrapUpdateOne) SetLevel(b bootstrap.Level) *BootstrapUpdateOne {
	buo.mutation.SetLevel(b)
	return buo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (buo *BootstrapUpdateOne) SetNillableLevel(b *bootstrap.Level) *BootstrapUpdateOne {
	if b != nil {
		buo.SetLevel(*b)
	}
	return buo
}

// SetServicePort sets the "service_port" field.
func (buo *BootstrapUpdateOne) SetServicePort(i int) *BootstrapUpdateOne {
	buo.mutation.ResetServicePort()
	buo.mutation.SetServicePort(i)
	return buo
}

// SetNillableServicePort sets the "service_port" field if the given value is not nil.
func (buo *BootstrapUpdateOne) SetNillableServicePort(i *int) *BootstrapUpdateOne {
	if i != nil {
		buo.SetServicePort(*i)
	}
	return buo
}

// AddServicePort adds i to the "service_port" field.
func (buo *BootstrapUpdateOne) AddServicePort(i int) *BootstrapUpdateOne {
	buo.mutation.AddServicePort(i)
	return buo
}

// SetFailCounts sets the "fail_counts" field.
func (buo *BootstrapUpdateOne) SetFailCounts(i int) *BootstrapUpdateOne {
	buo.mutation.ResetFailCounts()
	buo.mutation.SetFailCounts(i)
	return buo
}

// SetNillableFailCounts sets the "fail_counts" field if the given value is not nil.
func (buo *BootstrapUpdateOne) SetNillableFailCounts(i *int) *BootstrapUpdateOne {
	if i != nil {
		buo.SetFailCounts(*i)
	}
	return buo
}

// AddFailCounts adds i to the "fail_counts" field.
func (buo *BootstrapUpdateOne) AddFailCounts(i int) *BootstrapUpdateOne {
	buo.mutation.AddFailCounts(i)
	return buo
}

// SetSign sets the "sign" field.
func (buo *BootstrapUpdateOne) SetSign(s string) *BootstrapUpdateOne {
	buo.mutation.SetSign(s)
	return buo
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (buo *BootstrapUpdateOne) SetNillableSign(s *string) *BootstrapUpdateOne {
	if s != nil {
		buo.SetSign(*s)
	}
	return buo
}

// ClearSign clears the value of the "sign" field.
func (buo *BootstrapUpdateOne) ClearSign() *BootstrapUpdateOne {
	buo.mutation.ClearSign()
	return buo
}

// SetUpdatedUnix sets the "updated_unix" field.
func (buo *BootstrapUpdateOne) SetUpdatedUnix(i int64) *BootstrapUpdateOne {
	buo.mutation.ResetUpdatedUnix()
	buo.mutation.SetUpdatedUnix(i)
	return buo
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (buo *BootstrapUpdateOne) SetNillableUpdatedUnix(i *int64) *BootstrapUpdateOne {
	if i != nil {
		buo.SetUpdatedUnix(*i)
	}
	return buo
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (buo *BootstrapUpdateOne) AddUpdatedUnix(i int64) *BootstrapUpdateOne {
	buo.mutation.AddUpdatedUnix(i)
	return buo
}

// Mutation returns the BootstrapMutation object of the builder.
func (buo *BootstrapUpdateOne) Mutation() *BootstrapMutation {
	return buo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BootstrapUpdateOne) Select(field string, fields ...string) *BootstrapUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bootstrap entity.
func (buo *BootstrapUpdateOne) Save(ctx context.Context) (*Bootstrap, error) {
	var (
		err  error
		node *Bootstrap
	)
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BootstrapMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("bootnode: uninitialized hook (forgotten import bootnode/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Bootstrap)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BootstrapMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BootstrapUpdateOne) SaveX(ctx context.Context) *Bootstrap {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BootstrapUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BootstrapUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BootstrapUpdateOne) check() error {
	if v, ok := buo.mutation.Level(); ok {
		if err := bootstrap.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`bootnode: validator failed for field "Bootstrap.level": %w`, err)}
		}
	}
	return nil
}

func (buo *BootstrapUpdateOne) sqlSave(ctx context.Context) (_node *Bootstrap, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bootstrap.Table,
			Columns: bootstrap.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bootstrap.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`bootnode: missing "Bootstrap.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bootstrap.FieldID)
		for _, f := range fields {
			if !bootstrap.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("bootnode: invalid field %q for query", f)}
			}
			if f != bootstrap.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Pid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bootstrap.FieldPid,
		})
	}
	if value, ok := buo.mutation.Addrs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: bootstrap.FieldAddrs,
		})
	}
	if value, ok := buo.mutation.Expired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: bootstrap.FieldExpired,
		})
	}
	if value, ok := buo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: bootstrap.FieldLevel,
		})
	}
	if value, ok := buo.mutation.ServicePort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bootstrap.FieldServicePort,
		})
	}
	if value, ok := buo.mutation.AddedServicePort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bootstrap.FieldServicePort,
		})
	}
	if value, ok := buo.mutation.FailCounts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bootstrap.FieldFailCounts,
		})
	}
	if value, ok := buo.mutation.AddedFailCounts(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bootstrap.FieldFailCounts,
		})
	}
	if value, ok := buo.mutation.Sign(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bootstrap.FieldSign,
		})
	}
	if buo.mutation.SignCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: bootstrap.FieldSign,
		})
	}
	if value, ok := buo.mutation.UpdatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: bootstrap.FieldUpdatedUnix,
		})
	}
	if value, ok := buo.mutation.AddedUpdatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: bootstrap.FieldUpdatedUnix,
		})
	}
	_node = &Bootstrap{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bootstrap.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
