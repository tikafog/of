// Code generated by ent, DO NOT EDIT.

package upgrade

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/upgrade/version"
)

// VersionCreate is the builder for creating a Version entity.
type VersionCreate struct {
	config
	mutation *VersionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCurrent sets the "Current" field.
func (vc *VersionCreate) SetCurrent(i int) *VersionCreate {
	vc.mutation.SetCurrent(i)
	return vc
}

// SetNillableCurrent sets the "Current" field if the given value is not nil.
func (vc *VersionCreate) SetNillableCurrent(i *int) *VersionCreate {
	if i != nil {
		vc.SetCurrent(*i)
	}
	return vc
}

// SetLast sets the "Last" field.
func (vc *VersionCreate) SetLast(i int) *VersionCreate {
	vc.mutation.SetLast(i)
	return vc
}

// SetNillableLast sets the "Last" field if the given value is not nil.
func (vc *VersionCreate) SetNillableLast(i *int) *VersionCreate {
	if i != nil {
		vc.SetLast(*i)
	}
	return vc
}

// Mutation returns the VersionMutation object of the builder.
func (vc *VersionCreate) Mutation() *VersionMutation {
	return vc.mutation
}

// Save creates the Version in the database.
func (vc *VersionCreate) Save(ctx context.Context) (*Version, error) {
	var (
		err  error
		node *Version
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VersionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("upgrade: uninitialized hook (forgotten import upgrade/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, vc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (vc *VersionCreate) SaveX(ctx context.Context) *Version {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VersionCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VersionCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VersionCreate) defaults() {
	if _, ok := vc.mutation.Current(); !ok {
		v := version.DefaultCurrent
		vc.mutation.SetCurrent(v)
	}
	if _, ok := vc.mutation.Last(); !ok {
		v := version.DefaultLast
		vc.mutation.SetLast(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VersionCreate) check() error {
	if _, ok := vc.mutation.Current(); !ok {
		return &ValidationError{Name: "Current", err: errors.New(`upgrade: missing required field "Version.Current"`)}
	}
	if v, ok := vc.mutation.Current(); ok {
		if err := version.CurrentValidator(v); err != nil {
			return &ValidationError{Name: "Current", err: fmt.Errorf(`upgrade: validator failed for field "Version.Current": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Last(); !ok {
		return &ValidationError{Name: "Last", err: errors.New(`upgrade: missing required field "Version.Last"`)}
	}
	if v, ok := vc.mutation.Last(); ok {
		if err := version.LastValidator(v); err != nil {
			return &ValidationError{Name: "Last", err: fmt.Errorf(`upgrade: validator failed for field "Version.Last": %w`, err)}
		}
	}
	return nil
}

func (vc *VersionCreate) sqlSave(ctx context.Context) (*Version, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VersionCreate) createSpec() (*Version, *sqlgraph.CreateSpec) {
	var (
		_node = &Version{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: version.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: version.FieldID,
			},
		}
	)
	_spec.OnConflict = vc.conflict
	if value, ok := vc.mutation.Current(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldCurrent,
		})
		_node.Current = value
	}
	if value, ok := vc.mutation.Last(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: version.FieldLast,
		})
		_node.Last = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Version.Create().
//		SetCurrent(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VersionUpsert) {
//			SetCurrent(v+v).
//		}).
//		Exec(ctx)
func (vc *VersionCreate) OnConflict(opts ...sql.ConflictOption) *VersionUpsertOne {
	vc.conflict = opts
	return &VersionUpsertOne{
		create: vc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Version.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vc *VersionCreate) OnConflictColumns(columns ...string) *VersionUpsertOne {
	vc.conflict = append(vc.conflict, sql.ConflictColumns(columns...))
	return &VersionUpsertOne{
		create: vc,
	}
}

type (
	// VersionUpsertOne is the builder for "upsert"-ing
	//  one Version node.
	VersionUpsertOne struct {
		create *VersionCreate
	}

	// VersionUpsert is the "OnConflict" setter.
	VersionUpsert struct {
		*sql.UpdateSet
	}
)

// SetCurrent sets the "Current" field.
func (u *VersionUpsert) SetCurrent(v int) *VersionUpsert {
	u.Set(version.FieldCurrent, v)
	return u
}

// UpdateCurrent sets the "Current" field to the value that was provided on create.
func (u *VersionUpsert) UpdateCurrent() *VersionUpsert {
	u.SetExcluded(version.FieldCurrent)
	return u
}

// AddCurrent adds v to the "Current" field.
func (u *VersionUpsert) AddCurrent(v int) *VersionUpsert {
	u.Add(version.FieldCurrent, v)
	return u
}

// SetLast sets the "Last" field.
func (u *VersionUpsert) SetLast(v int) *VersionUpsert {
	u.Set(version.FieldLast, v)
	return u
}

// UpdateLast sets the "Last" field to the value that was provided on create.
func (u *VersionUpsert) UpdateLast() *VersionUpsert {
	u.SetExcluded(version.FieldLast)
	return u
}

// AddLast adds v to the "Last" field.
func (u *VersionUpsert) AddLast(v int) *VersionUpsert {
	u.Add(version.FieldLast, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Version.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *VersionUpsertOne) UpdateNewValues() *VersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Version.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *VersionUpsertOne) Ignore() *VersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VersionUpsertOne) DoNothing() *VersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VersionCreate.OnConflict
// documentation for more info.
func (u *VersionUpsertOne) Update(set func(*VersionUpsert)) *VersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VersionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCurrent sets the "Current" field.
func (u *VersionUpsertOne) SetCurrent(v int) *VersionUpsertOne {
	return u.Update(func(s *VersionUpsert) {
		s.SetCurrent(v)
	})
}

// AddCurrent adds v to the "Current" field.
func (u *VersionUpsertOne) AddCurrent(v int) *VersionUpsertOne {
	return u.Update(func(s *VersionUpsert) {
		s.AddCurrent(v)
	})
}

// UpdateCurrent sets the "Current" field to the value that was provided on create.
func (u *VersionUpsertOne) UpdateCurrent() *VersionUpsertOne {
	return u.Update(func(s *VersionUpsert) {
		s.UpdateCurrent()
	})
}

// SetLast sets the "Last" field.
func (u *VersionUpsertOne) SetLast(v int) *VersionUpsertOne {
	return u.Update(func(s *VersionUpsert) {
		s.SetLast(v)
	})
}

// AddLast adds v to the "Last" field.
func (u *VersionUpsertOne) AddLast(v int) *VersionUpsertOne {
	return u.Update(func(s *VersionUpsert) {
		s.AddLast(v)
	})
}

// UpdateLast sets the "Last" field to the value that was provided on create.
func (u *VersionUpsertOne) UpdateLast() *VersionUpsertOne {
	return u.Update(func(s *VersionUpsert) {
		s.UpdateLast()
	})
}

// Exec executes the query.
func (u *VersionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("upgrade: missing options for VersionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VersionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *VersionUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *VersionUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// VersionCreateBulk is the builder for creating many Version entities in bulk.
type VersionCreateBulk struct {
	config
	builders []*VersionCreate
	conflict []sql.ConflictOption
}

// Save creates the Version entities in the database.
func (vcb *VersionCreateBulk) Save(ctx context.Context) ([]*Version, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Version, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VersionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = vcb.conflict
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
func (vcb *VersionCreateBulk) SaveX(ctx context.Context) []*Version {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VersionCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VersionCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Version.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VersionUpsert) {
//			SetCurrent(v+v).
//		}).
//		Exec(ctx)
func (vcb *VersionCreateBulk) OnConflict(opts ...sql.ConflictOption) *VersionUpsertBulk {
	vcb.conflict = opts
	return &VersionUpsertBulk{
		create: vcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Version.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vcb *VersionCreateBulk) OnConflictColumns(columns ...string) *VersionUpsertBulk {
	vcb.conflict = append(vcb.conflict, sql.ConflictColumns(columns...))
	return &VersionUpsertBulk{
		create: vcb,
	}
}

// VersionUpsertBulk is the builder for "upsert"-ing
// a bulk of Version nodes.
type VersionUpsertBulk struct {
	create *VersionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Version.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *VersionUpsertBulk) UpdateNewValues() *VersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Version.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *VersionUpsertBulk) Ignore() *VersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VersionUpsertBulk) DoNothing() *VersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VersionCreateBulk.OnConflict
// documentation for more info.
func (u *VersionUpsertBulk) Update(set func(*VersionUpsert)) *VersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VersionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCurrent sets the "Current" field.
func (u *VersionUpsertBulk) SetCurrent(v int) *VersionUpsertBulk {
	return u.Update(func(s *VersionUpsert) {
		s.SetCurrent(v)
	})
}

// AddCurrent adds v to the "Current" field.
func (u *VersionUpsertBulk) AddCurrent(v int) *VersionUpsertBulk {
	return u.Update(func(s *VersionUpsert) {
		s.AddCurrent(v)
	})
}

// UpdateCurrent sets the "Current" field to the value that was provided on create.
func (u *VersionUpsertBulk) UpdateCurrent() *VersionUpsertBulk {
	return u.Update(func(s *VersionUpsert) {
		s.UpdateCurrent()
	})
}

// SetLast sets the "Last" field.
func (u *VersionUpsertBulk) SetLast(v int) *VersionUpsertBulk {
	return u.Update(func(s *VersionUpsert) {
		s.SetLast(v)
	})
}

// AddLast adds v to the "Last" field.
func (u *VersionUpsertBulk) AddLast(v int) *VersionUpsertBulk {
	return u.Update(func(s *VersionUpsert) {
		s.AddLast(v)
	})
}

// UpdateLast sets the "Last" field to the value that was provided on create.
func (u *VersionUpsertBulk) UpdateLast() *VersionUpsertBulk {
	return u.Update(func(s *VersionUpsert) {
		s.UpdateLast()
	})
}

// Exec executes the query.
func (u *VersionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("upgrade: OnConflict was set for builder %d. Set it on the VersionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("upgrade: missing options for VersionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VersionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
