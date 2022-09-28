// Code generated by ent, DO NOT EDIT.

package kernel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/kernel/resource"
)

// ResourceCreate is the builder for creating a Resource entity.
type ResourceCreate struct {
	config
	mutation *ResourceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRid sets the "rid" field.
func (rc *ResourceCreate) SetRid(s string) *ResourceCreate {
	rc.mutation.SetRid(s)
	return rc
}

// SetStatus sets the "status" field.
func (rc *ResourceCreate) SetStatus(s string) *ResourceCreate {
	rc.mutation.SetStatus(s)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableStatus(s *string) *ResourceCreate {
	if s != nil {
		rc.SetStatus(*s)
	}
	return rc
}

// SetStep sets the "step" field.
func (rc *ResourceCreate) SetStep(s string) *ResourceCreate {
	rc.mutation.SetStep(s)
	return rc
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableStep(s *string) *ResourceCreate {
	if s != nil {
		rc.SetStep(*s)
	}
	return rc
}

// SetPriority sets the "priority" field.
func (rc *ResourceCreate) SetPriority(i int) *ResourceCreate {
	rc.mutation.SetPriority(i)
	return rc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (rc *ResourceCreate) SetNillablePriority(i *int) *ResourceCreate {
	if i != nil {
		rc.SetPriority(*i)
	}
	return rc
}

// SetRelate sets the "relate" field.
func (rc *ResourceCreate) SetRelate(s string) *ResourceCreate {
	rc.mutation.SetRelate(s)
	return rc
}

// SetNillableRelate sets the "relate" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableRelate(s *string) *ResourceCreate {
	if s != nil {
		rc.SetRelate(*s)
	}
	return rc
}

// SetUpdatedUnix sets the "updated_unix" field.
func (rc *ResourceCreate) SetUpdatedUnix(i int64) *ResourceCreate {
	rc.mutation.SetUpdatedUnix(i)
	return rc
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableUpdatedUnix(i *int64) *ResourceCreate {
	if i != nil {
		rc.SetUpdatedUnix(*i)
	}
	return rc
}

// SetComment sets the "comment" field.
func (rc *ResourceCreate) SetComment(s string) *ResourceCreate {
	rc.mutation.SetComment(s)
	return rc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableComment(s *string) *ResourceCreate {
	if s != nil {
		rc.SetComment(*s)
	}
	return rc
}

// Mutation returns the ResourceMutation object of the builder.
func (rc *ResourceCreate) Mutation() *ResourceMutation {
	return rc.mutation
}

// Save creates the Resource in the database.
func (rc *ResourceCreate) Save(ctx context.Context) (*Resource, error) {
	var (
		err  error
		node *Resource
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("kernel: uninitialized hook (forgotten import kernel/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Resource)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ResourceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResourceCreate) SaveX(ctx context.Context) *Resource {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ResourceCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ResourceCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ResourceCreate) defaults() {
	if _, ok := rc.mutation.Status(); !ok {
		v := resource.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.Step(); !ok {
		v := resource.DefaultStep
		rc.mutation.SetStep(v)
	}
	if _, ok := rc.mutation.Priority(); !ok {
		v := resource.DefaultPriority
		rc.mutation.SetPriority(v)
	}
	if _, ok := rc.mutation.Relate(); !ok {
		v := resource.DefaultRelate
		rc.mutation.SetRelate(v)
	}
	if _, ok := rc.mutation.UpdatedUnix(); !ok {
		v := resource.DefaultUpdatedUnix
		rc.mutation.SetUpdatedUnix(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ResourceCreate) check() error {
	if _, ok := rc.mutation.Rid(); !ok {
		return &ValidationError{Name: "rid", err: errors.New(`kernel: missing required field "Resource.rid"`)}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`kernel: missing required field "Resource.status"`)}
	}
	if _, ok := rc.mutation.Step(); !ok {
		return &ValidationError{Name: "step", err: errors.New(`kernel: missing required field "Resource.step"`)}
	}
	if _, ok := rc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`kernel: missing required field "Resource.priority"`)}
	}
	if _, ok := rc.mutation.Relate(); !ok {
		return &ValidationError{Name: "relate", err: errors.New(`kernel: missing required field "Resource.relate"`)}
	}
	if _, ok := rc.mutation.UpdatedUnix(); !ok {
		return &ValidationError{Name: "updated_unix", err: errors.New(`kernel: missing required field "Resource.updated_unix"`)}
	}
	return nil
}

func (rc *ResourceCreate) sqlSave(ctx context.Context) (*Resource, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *ResourceCreate) createSpec() (*Resource, *sqlgraph.CreateSpec) {
	var (
		_node = &Resource{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: resource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resource.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.Rid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldRid,
		})
		_node.Rid = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := rc.mutation.Step(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldStep,
		})
		_node.Step = value
	}
	if value, ok := rc.mutation.Priority(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldPriority,
		})
		_node.Priority = value
	}
	if value, ok := rc.mutation.Relate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldRelate,
		})
		_node.Relate = value
	}
	if value, ok := rc.mutation.UpdatedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: resource.FieldUpdatedUnix,
		})
		_node.UpdatedUnix = value
	}
	if value, ok := rc.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldComment,
		})
		_node.Comment = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Resource.Create().
//		SetRid(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ResourceUpsert) {
//			SetRid(v+v).
//		}).
//		Exec(ctx)
func (rc *ResourceCreate) OnConflict(opts ...sql.ConflictOption) *ResourceUpsertOne {
	rc.conflict = opts
	return &ResourceUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Resource.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *ResourceCreate) OnConflictColumns(columns ...string) *ResourceUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &ResourceUpsertOne{
		create: rc,
	}
}

type (
	// ResourceUpsertOne is the builder for "upsert"-ing
	//  one Resource node.
	ResourceUpsertOne struct {
		create *ResourceCreate
	}

	// ResourceUpsert is the "OnConflict" setter.
	ResourceUpsert struct {
		*sql.UpdateSet
	}
)

// SetRid sets the "rid" field.
func (u *ResourceUpsert) SetRid(v string) *ResourceUpsert {
	u.Set(resource.FieldRid, v)
	return u
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *ResourceUpsert) UpdateRid() *ResourceUpsert {
	u.SetExcluded(resource.FieldRid)
	return u
}

// SetStatus sets the "status" field.
func (u *ResourceUpsert) SetStatus(v string) *ResourceUpsert {
	u.Set(resource.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ResourceUpsert) UpdateStatus() *ResourceUpsert {
	u.SetExcluded(resource.FieldStatus)
	return u
}

// SetStep sets the "step" field.
func (u *ResourceUpsert) SetStep(v string) *ResourceUpsert {
	u.Set(resource.FieldStep, v)
	return u
}

// UpdateStep sets the "step" field to the value that was provided on create.
func (u *ResourceUpsert) UpdateStep() *ResourceUpsert {
	u.SetExcluded(resource.FieldStep)
	return u
}

// SetPriority sets the "priority" field.
func (u *ResourceUpsert) SetPriority(v int) *ResourceUpsert {
	u.Set(resource.FieldPriority, v)
	return u
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *ResourceUpsert) UpdatePriority() *ResourceUpsert {
	u.SetExcluded(resource.FieldPriority)
	return u
}

// AddPriority adds v to the "priority" field.
func (u *ResourceUpsert) AddPriority(v int) *ResourceUpsert {
	u.Add(resource.FieldPriority, v)
	return u
}

// SetRelate sets the "relate" field.
func (u *ResourceUpsert) SetRelate(v string) *ResourceUpsert {
	u.Set(resource.FieldRelate, v)
	return u
}

// UpdateRelate sets the "relate" field to the value that was provided on create.
func (u *ResourceUpsert) UpdateRelate() *ResourceUpsert {
	u.SetExcluded(resource.FieldRelate)
	return u
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *ResourceUpsert) SetUpdatedUnix(v int64) *ResourceUpsert {
	u.Set(resource.FieldUpdatedUnix, v)
	return u
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *ResourceUpsert) UpdateUpdatedUnix() *ResourceUpsert {
	u.SetExcluded(resource.FieldUpdatedUnix)
	return u
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *ResourceUpsert) AddUpdatedUnix(v int64) *ResourceUpsert {
	u.Add(resource.FieldUpdatedUnix, v)
	return u
}

// SetComment sets the "comment" field.
func (u *ResourceUpsert) SetComment(v string) *ResourceUpsert {
	u.Set(resource.FieldComment, v)
	return u
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *ResourceUpsert) UpdateComment() *ResourceUpsert {
	u.SetExcluded(resource.FieldComment)
	return u
}

// ClearComment clears the value of the "comment" field.
func (u *ResourceUpsert) ClearComment() *ResourceUpsert {
	u.SetNull(resource.FieldComment)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Resource.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ResourceUpsertOne) UpdateNewValues() *ResourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Resource.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ResourceUpsertOne) Ignore() *ResourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ResourceUpsertOne) DoNothing() *ResourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ResourceCreate.OnConflict
// documentation for more info.
func (u *ResourceUpsertOne) Update(set func(*ResourceUpsert)) *ResourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ResourceUpsert{UpdateSet: update})
	}))
	return u
}

// SetRid sets the "rid" field.
func (u *ResourceUpsertOne) SetRid(v string) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.SetRid(v)
	})
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *ResourceUpsertOne) UpdateRid() *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateRid()
	})
}

// SetStatus sets the "status" field.
func (u *ResourceUpsertOne) SetStatus(v string) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ResourceUpsertOne) UpdateStatus() *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateStatus()
	})
}

// SetStep sets the "step" field.
func (u *ResourceUpsertOne) SetStep(v string) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.SetStep(v)
	})
}

// UpdateStep sets the "step" field to the value that was provided on create.
func (u *ResourceUpsertOne) UpdateStep() *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateStep()
	})
}

// SetPriority sets the "priority" field.
func (u *ResourceUpsertOne) SetPriority(v int) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.SetPriority(v)
	})
}

// AddPriority adds v to the "priority" field.
func (u *ResourceUpsertOne) AddPriority(v int) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.AddPriority(v)
	})
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *ResourceUpsertOne) UpdatePriority() *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdatePriority()
	})
}

// SetRelate sets the "relate" field.
func (u *ResourceUpsertOne) SetRelate(v string) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.SetRelate(v)
	})
}

// UpdateRelate sets the "relate" field to the value that was provided on create.
func (u *ResourceUpsertOne) UpdateRelate() *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateRelate()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *ResourceUpsertOne) SetUpdatedUnix(v int64) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *ResourceUpsertOne) AddUpdatedUnix(v int64) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *ResourceUpsertOne) UpdateUpdatedUnix() *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// SetComment sets the "comment" field.
func (u *ResourceUpsertOne) SetComment(v string) *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *ResourceUpsertOne) UpdateComment() *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateComment()
	})
}

// ClearComment clears the value of the "comment" field.
func (u *ResourceUpsertOne) ClearComment() *ResourceUpsertOne {
	return u.Update(func(s *ResourceUpsert) {
		s.ClearComment()
	})
}

// Exec executes the query.
func (u *ResourceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("kernel: missing options for ResourceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ResourceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ResourceUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ResourceUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ResourceCreateBulk is the builder for creating many Resource entities in bulk.
type ResourceCreateBulk struct {
	config
	builders []*ResourceCreate
	conflict []sql.ConflictOption
}

// Save creates the Resource entities in the database.
func (rcb *ResourceCreateBulk) Save(ctx context.Context) ([]*Resource, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Resource, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourceMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ResourceCreateBulk) SaveX(ctx context.Context) []*Resource {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ResourceCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ResourceCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Resource.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ResourceUpsert) {
//			SetRid(v+v).
//		}).
//		Exec(ctx)
func (rcb *ResourceCreateBulk) OnConflict(opts ...sql.ConflictOption) *ResourceUpsertBulk {
	rcb.conflict = opts
	return &ResourceUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Resource.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *ResourceCreateBulk) OnConflictColumns(columns ...string) *ResourceUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &ResourceUpsertBulk{
		create: rcb,
	}
}

// ResourceUpsertBulk is the builder for "upsert"-ing
// a bulk of Resource nodes.
type ResourceUpsertBulk struct {
	create *ResourceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Resource.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ResourceUpsertBulk) UpdateNewValues() *ResourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Resource.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ResourceUpsertBulk) Ignore() *ResourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ResourceUpsertBulk) DoNothing() *ResourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ResourceCreateBulk.OnConflict
// documentation for more info.
func (u *ResourceUpsertBulk) Update(set func(*ResourceUpsert)) *ResourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ResourceUpsert{UpdateSet: update})
	}))
	return u
}

// SetRid sets the "rid" field.
func (u *ResourceUpsertBulk) SetRid(v string) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.SetRid(v)
	})
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *ResourceUpsertBulk) UpdateRid() *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateRid()
	})
}

// SetStatus sets the "status" field.
func (u *ResourceUpsertBulk) SetStatus(v string) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ResourceUpsertBulk) UpdateStatus() *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateStatus()
	})
}

// SetStep sets the "step" field.
func (u *ResourceUpsertBulk) SetStep(v string) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.SetStep(v)
	})
}

// UpdateStep sets the "step" field to the value that was provided on create.
func (u *ResourceUpsertBulk) UpdateStep() *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateStep()
	})
}

// SetPriority sets the "priority" field.
func (u *ResourceUpsertBulk) SetPriority(v int) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.SetPriority(v)
	})
}

// AddPriority adds v to the "priority" field.
func (u *ResourceUpsertBulk) AddPriority(v int) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.AddPriority(v)
	})
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *ResourceUpsertBulk) UpdatePriority() *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdatePriority()
	})
}

// SetRelate sets the "relate" field.
func (u *ResourceUpsertBulk) SetRelate(v string) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.SetRelate(v)
	})
}

// UpdateRelate sets the "relate" field to the value that was provided on create.
func (u *ResourceUpsertBulk) UpdateRelate() *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateRelate()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *ResourceUpsertBulk) SetUpdatedUnix(v int64) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *ResourceUpsertBulk) AddUpdatedUnix(v int64) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *ResourceUpsertBulk) UpdateUpdatedUnix() *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// SetComment sets the "comment" field.
func (u *ResourceUpsertBulk) SetComment(v string) *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *ResourceUpsertBulk) UpdateComment() *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.UpdateComment()
	})
}

// ClearComment clears the value of the "comment" field.
func (u *ResourceUpsertBulk) ClearComment() *ResourceUpsertBulk {
	return u.Update(func(s *ResourceUpsert) {
		s.ClearComment()
	})
}

// Exec executes the query.
func (u *ResourceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("kernel: OnConflict was set for builder %d. Set it on the ResourceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("kernel: missing options for ResourceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ResourceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}