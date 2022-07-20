// Code generated by entc, DO NOT EDIT.

package media

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/discovery"
)

// DiscoveryCreate is the builder for creating a Discovery entity.
type DiscoveryCreate struct {
	config
	mutation *DiscoveryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedUnix sets the "created_unix" field.
func (dc *DiscoveryCreate) SetCreatedUnix(i int64) *DiscoveryCreate {
	dc.mutation.SetCreatedUnix(i)
	return dc
}

// SetNillableCreatedUnix sets the "created_unix" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableCreatedUnix(i *int64) *DiscoveryCreate {
	if i != nil {
		dc.SetCreatedUnix(*i)
	}
	return dc
}

// SetUpdatedUnix sets the "updated_unix" field.
func (dc *DiscoveryCreate) SetUpdatedUnix(i int64) *DiscoveryCreate {
	dc.mutation.SetUpdatedUnix(i)
	return dc
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableUpdatedUnix(i *int64) *DiscoveryCreate {
	if i != nil {
		dc.SetUpdatedUnix(*i)
	}
	return dc
}

// SetDeletedUnix sets the "deleted_unix" field.
func (dc *DiscoveryCreate) SetDeletedUnix(i int64) *DiscoveryCreate {
	dc.mutation.SetDeletedUnix(i)
	return dc
}

// SetNillableDeletedUnix sets the "deleted_unix" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableDeletedUnix(i *int64) *DiscoveryCreate {
	if i != nil {
		dc.SetDeletedUnix(*i)
	}
	return dc
}

// SetDate sets the "date" field.
func (dc *DiscoveryCreate) SetDate(s string) *DiscoveryCreate {
	dc.mutation.SetDate(s)
	return dc
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableDate(s *string) *DiscoveryCreate {
	if s != nil {
		dc.SetDate(*s)
	}
	return dc
}

// SetRid sets the "rid" field.
func (dc *DiscoveryCreate) SetRid(s string) *DiscoveryCreate {
	dc.mutation.SetRid(s)
	return dc
}

// SetNillableRid sets the "rid" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableRid(s *string) *DiscoveryCreate {
	if s != nil {
		dc.SetRid(*s)
	}
	return dc
}

// SetTitle sets the "title" field.
func (dc *DiscoveryCreate) SetTitle(s string) *DiscoveryCreate {
	dc.mutation.SetTitle(s)
	return dc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableTitle(s *string) *DiscoveryCreate {
	if s != nil {
		dc.SetTitle(*s)
	}
	return dc
}

// SetDetail sets the "detail" field.
func (dc *DiscoveryCreate) SetDetail(s string) *DiscoveryCreate {
	dc.mutation.SetDetail(s)
	return dc
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableDetail(s *string) *DiscoveryCreate {
	if s != nil {
		dc.SetDetail(*s)
	}
	return dc
}

// SetMtype sets the "mtype" field.
func (dc *DiscoveryCreate) SetMtype(d discovery.Mtype) *DiscoveryCreate {
	dc.mutation.SetMtype(d)
	return dc
}

// SetNillableMtype sets the "mtype" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableMtype(d *discovery.Mtype) *DiscoveryCreate {
	if d != nil {
		dc.SetMtype(*d)
	}
	return dc
}

// SetLinks sets the "links" field.
func (dc *DiscoveryCreate) SetLinks(s []string) *DiscoveryCreate {
	dc.mutation.SetLinks(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DiscoveryCreate) SetID(u uuid.UUID) *DiscoveryCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DiscoveryCreate) SetNillableID(u *uuid.UUID) *DiscoveryCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// Mutation returns the DiscoveryMutation object of the builder.
func (dc *DiscoveryCreate) Mutation() *DiscoveryMutation {
	return dc.mutation
}

// Save creates the Discovery in the database.
func (dc *DiscoveryCreate) Save(ctx context.Context) (*Discovery, error) {
	var (
		err  error
		node *Discovery
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscoveryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("media: uninitialized hook (forgotten import media/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DiscoveryCreate) SaveX(ctx context.Context) *Discovery {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DiscoveryCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DiscoveryCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DiscoveryCreate) defaults() {
	if _, ok := dc.mutation.CreatedUnix(); !ok {
		v := discovery.DefaultCreatedUnix
		dc.mutation.SetCreatedUnix(v)
	}
	if _, ok := dc.mutation.UpdatedUnix(); !ok {
		v := discovery.DefaultUpdatedUnix
		dc.mutation.SetUpdatedUnix(v)
	}
	if _, ok := dc.mutation.Date(); !ok {
		v := discovery.DefaultDate
		dc.mutation.SetDate(v)
	}
	if _, ok := dc.mutation.Rid(); !ok {
		v := discovery.DefaultRid
		dc.mutation.SetRid(v)
	}
	if _, ok := dc.mutation.Title(); !ok {
		v := discovery.DefaultTitle
		dc.mutation.SetTitle(v)
	}
	if _, ok := dc.mutation.Detail(); !ok {
		v := discovery.DefaultDetail
		dc.mutation.SetDetail(v)
	}
	if _, ok := dc.mutation.Mtype(); !ok {
		v := discovery.DefaultMtype
		dc.mutation.SetMtype(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := discovery.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DiscoveryCreate) check() error {
	if _, ok := dc.mutation.CreatedUnix(); !ok {
		return &ValidationError{Name: "created_unix", err: errors.New(`media: missing required field "Discovery.created_unix"`)}
	}
	if _, ok := dc.mutation.UpdatedUnix(); !ok {
		return &ValidationError{Name: "updated_unix", err: errors.New(`media: missing required field "Discovery.updated_unix"`)}
	}
	if _, ok := dc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`media: missing required field "Discovery.date"`)}
	}
	if _, ok := dc.mutation.Rid(); !ok {
		return &ValidationError{Name: "rid", err: errors.New(`media: missing required field "Discovery.rid"`)}
	}
	if _, ok := dc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`media: missing required field "Discovery.title"`)}
	}
	if _, ok := dc.mutation.Detail(); !ok {
		return &ValidationError{Name: "detail", err: errors.New(`media: missing required field "Discovery.detail"`)}
	}
	if _, ok := dc.mutation.Mtype(); !ok {
		return &ValidationError{Name: "mtype", err: errors.New(`media: missing required field "Discovery.mtype"`)}
	}
	if v, ok := dc.mutation.Mtype(); ok {
		if err := discovery.MtypeValidator(v); err != nil {
			return &ValidationError{Name: "mtype", err: fmt.Errorf(`media: validator failed for field "Discovery.mtype": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Links(); !ok {
		return &ValidationError{Name: "links", err: errors.New(`media: missing required field "Discovery.links"`)}
	}
	return nil
}

func (dc *DiscoveryCreate) sqlSave(ctx context.Context) (*Discovery, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (dc *DiscoveryCreate) createSpec() (*Discovery, *sqlgraph.CreateSpec) {
	var (
		_node = &Discovery{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: discovery.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: discovery.FieldID,
			},
		}
	)
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.CreatedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: discovery.FieldCreatedUnix,
		})
		_node.CreatedUnix = value
	}
	if value, ok := dc.mutation.UpdatedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: discovery.FieldUpdatedUnix,
		})
		_node.UpdatedUnix = value
	}
	if value, ok := dc.mutation.DeletedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: discovery.FieldDeletedUnix,
		})
		_node.DeletedUnix = value
	}
	if value, ok := dc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discovery.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := dc.mutation.Rid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discovery.FieldRid,
		})
		_node.Rid = value
	}
	if value, ok := dc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discovery.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := dc.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discovery.FieldDetail,
		})
		_node.Detail = value
	}
	if value, ok := dc.mutation.Mtype(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: discovery.FieldMtype,
		})
		_node.Mtype = value
	}
	if value, ok := dc.mutation.Links(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: discovery.FieldLinks,
		})
		_node.Links = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Discovery.Create().
//		SetCreatedUnix(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscoveryUpsert) {
//			SetCreatedUnix(v+v).
//		}).
//		Exec(ctx)
//
func (dc *DiscoveryCreate) OnConflict(opts ...sql.ConflictOption) *DiscoveryUpsertOne {
	dc.conflict = opts
	return &DiscoveryUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Discovery.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dc *DiscoveryCreate) OnConflictColumns(columns ...string) *DiscoveryUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DiscoveryUpsertOne{
		create: dc,
	}
}

type (
	// DiscoveryUpsertOne is the builder for "upsert"-ing
	//  one Discovery node.
	DiscoveryUpsertOne struct {
		create *DiscoveryCreate
	}

	// DiscoveryUpsert is the "OnConflict" setter.
	DiscoveryUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedUnix sets the "created_unix" field.
func (u *DiscoveryUpsert) SetCreatedUnix(v int64) *DiscoveryUpsert {
	u.Set(discovery.FieldCreatedUnix, v)
	return u
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateCreatedUnix() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldCreatedUnix)
	return u
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *DiscoveryUpsert) AddCreatedUnix(v int64) *DiscoveryUpsert {
	u.Add(discovery.FieldCreatedUnix, v)
	return u
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *DiscoveryUpsert) SetUpdatedUnix(v int64) *DiscoveryUpsert {
	u.Set(discovery.FieldUpdatedUnix, v)
	return u
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateUpdatedUnix() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldUpdatedUnix)
	return u
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *DiscoveryUpsert) AddUpdatedUnix(v int64) *DiscoveryUpsert {
	u.Add(discovery.FieldUpdatedUnix, v)
	return u
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *DiscoveryUpsert) SetDeletedUnix(v int64) *DiscoveryUpsert {
	u.Set(discovery.FieldDeletedUnix, v)
	return u
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateDeletedUnix() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldDeletedUnix)
	return u
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *DiscoveryUpsert) AddDeletedUnix(v int64) *DiscoveryUpsert {
	u.Add(discovery.FieldDeletedUnix, v)
	return u
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *DiscoveryUpsert) ClearDeletedUnix() *DiscoveryUpsert {
	u.SetNull(discovery.FieldDeletedUnix)
	return u
}

// SetDate sets the "date" field.
func (u *DiscoveryUpsert) SetDate(v string) *DiscoveryUpsert {
	u.Set(discovery.FieldDate, v)
	return u
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateDate() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldDate)
	return u
}

// SetRid sets the "rid" field.
func (u *DiscoveryUpsert) SetRid(v string) *DiscoveryUpsert {
	u.Set(discovery.FieldRid, v)
	return u
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateRid() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldRid)
	return u
}

// SetTitle sets the "title" field.
func (u *DiscoveryUpsert) SetTitle(v string) *DiscoveryUpsert {
	u.Set(discovery.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateTitle() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldTitle)
	return u
}

// SetDetail sets the "detail" field.
func (u *DiscoveryUpsert) SetDetail(v string) *DiscoveryUpsert {
	u.Set(discovery.FieldDetail, v)
	return u
}

// UpdateDetail sets the "detail" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateDetail() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldDetail)
	return u
}

// SetMtype sets the "mtype" field.
func (u *DiscoveryUpsert) SetMtype(v discovery.Mtype) *DiscoveryUpsert {
	u.Set(discovery.FieldMtype, v)
	return u
}

// UpdateMtype sets the "mtype" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateMtype() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldMtype)
	return u
}

// SetLinks sets the "links" field.
func (u *DiscoveryUpsert) SetLinks(v []string) *DiscoveryUpsert {
	u.Set(discovery.FieldLinks, v)
	return u
}

// UpdateLinks sets the "links" field to the value that was provided on create.
func (u *DiscoveryUpsert) UpdateLinks() *DiscoveryUpsert {
	u.SetExcluded(discovery.FieldLinks)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Discovery.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discovery.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DiscoveryUpsertOne) UpdateNewValues() *DiscoveryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(discovery.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Discovery.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DiscoveryUpsertOne) Ignore() *DiscoveryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscoveryUpsertOne) DoNothing() *DiscoveryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscoveryCreate.OnConflict
// documentation for more info.
func (u *DiscoveryUpsertOne) Update(set func(*DiscoveryUpsert)) *DiscoveryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscoveryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedUnix sets the "created_unix" field.
func (u *DiscoveryUpsertOne) SetCreatedUnix(v int64) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetCreatedUnix(v)
	})
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *DiscoveryUpsertOne) AddCreatedUnix(v int64) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.AddCreatedUnix(v)
	})
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateCreatedUnix() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateCreatedUnix()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *DiscoveryUpsertOne) SetUpdatedUnix(v int64) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *DiscoveryUpsertOne) AddUpdatedUnix(v int64) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateUpdatedUnix() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *DiscoveryUpsertOne) SetDeletedUnix(v int64) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetDeletedUnix(v)
	})
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *DiscoveryUpsertOne) AddDeletedUnix(v int64) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.AddDeletedUnix(v)
	})
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateDeletedUnix() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateDeletedUnix()
	})
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *DiscoveryUpsertOne) ClearDeletedUnix() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.ClearDeletedUnix()
	})
}

// SetDate sets the "date" field.
func (u *DiscoveryUpsertOne) SetDate(v string) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetDate(v)
	})
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateDate() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateDate()
	})
}

// SetRid sets the "rid" field.
func (u *DiscoveryUpsertOne) SetRid(v string) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetRid(v)
	})
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateRid() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateRid()
	})
}

// SetTitle sets the "title" field.
func (u *DiscoveryUpsertOne) SetTitle(v string) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateTitle() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateTitle()
	})
}

// SetDetail sets the "detail" field.
func (u *DiscoveryUpsertOne) SetDetail(v string) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetDetail(v)
	})
}

// UpdateDetail sets the "detail" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateDetail() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateDetail()
	})
}

// SetMtype sets the "mtype" field.
func (u *DiscoveryUpsertOne) SetMtype(v discovery.Mtype) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetMtype(v)
	})
}

// UpdateMtype sets the "mtype" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateMtype() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateMtype()
	})
}

// SetLinks sets the "links" field.
func (u *DiscoveryUpsertOne) SetLinks(v []string) *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetLinks(v)
	})
}

// UpdateLinks sets the "links" field to the value that was provided on create.
func (u *DiscoveryUpsertOne) UpdateLinks() *DiscoveryUpsertOne {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateLinks()
	})
}

// Exec executes the query.
func (u *DiscoveryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("media: missing options for DiscoveryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscoveryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DiscoveryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("media: DiscoveryUpsertOne.ID is not supported by MySQL driver. Use DiscoveryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DiscoveryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DiscoveryCreateBulk is the builder for creating many Discovery entities in bulk.
type DiscoveryCreateBulk struct {
	config
	builders []*DiscoveryCreate
	conflict []sql.ConflictOption
}

// Save creates the Discovery entities in the database.
func (dcb *DiscoveryCreateBulk) Save(ctx context.Context) ([]*Discovery, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Discovery, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscoveryMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DiscoveryCreateBulk) SaveX(ctx context.Context) []*Discovery {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DiscoveryCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DiscoveryCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Discovery.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscoveryUpsert) {
//			SetCreatedUnix(v+v).
//		}).
//		Exec(ctx)
//
func (dcb *DiscoveryCreateBulk) OnConflict(opts ...sql.ConflictOption) *DiscoveryUpsertBulk {
	dcb.conflict = opts
	return &DiscoveryUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Discovery.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dcb *DiscoveryCreateBulk) OnConflictColumns(columns ...string) *DiscoveryUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DiscoveryUpsertBulk{
		create: dcb,
	}
}

// DiscoveryUpsertBulk is the builder for "upsert"-ing
// a bulk of Discovery nodes.
type DiscoveryUpsertBulk struct {
	create *DiscoveryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Discovery.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discovery.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DiscoveryUpsertBulk) UpdateNewValues() *DiscoveryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(discovery.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Discovery.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DiscoveryUpsertBulk) Ignore() *DiscoveryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscoveryUpsertBulk) DoNothing() *DiscoveryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscoveryCreateBulk.OnConflict
// documentation for more info.
func (u *DiscoveryUpsertBulk) Update(set func(*DiscoveryUpsert)) *DiscoveryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscoveryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedUnix sets the "created_unix" field.
func (u *DiscoveryUpsertBulk) SetCreatedUnix(v int64) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetCreatedUnix(v)
	})
}

// AddCreatedUnix adds v to the "created_unix" field.
func (u *DiscoveryUpsertBulk) AddCreatedUnix(v int64) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.AddCreatedUnix(v)
	})
}

// UpdateCreatedUnix sets the "created_unix" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateCreatedUnix() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateCreatedUnix()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *DiscoveryUpsertBulk) SetUpdatedUnix(v int64) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *DiscoveryUpsertBulk) AddUpdatedUnix(v int64) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateUpdatedUnix() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// SetDeletedUnix sets the "deleted_unix" field.
func (u *DiscoveryUpsertBulk) SetDeletedUnix(v int64) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetDeletedUnix(v)
	})
}

// AddDeletedUnix adds v to the "deleted_unix" field.
func (u *DiscoveryUpsertBulk) AddDeletedUnix(v int64) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.AddDeletedUnix(v)
	})
}

// UpdateDeletedUnix sets the "deleted_unix" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateDeletedUnix() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateDeletedUnix()
	})
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (u *DiscoveryUpsertBulk) ClearDeletedUnix() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.ClearDeletedUnix()
	})
}

// SetDate sets the "date" field.
func (u *DiscoveryUpsertBulk) SetDate(v string) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetDate(v)
	})
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateDate() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateDate()
	})
}

// SetRid sets the "rid" field.
func (u *DiscoveryUpsertBulk) SetRid(v string) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetRid(v)
	})
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateRid() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateRid()
	})
}

// SetTitle sets the "title" field.
func (u *DiscoveryUpsertBulk) SetTitle(v string) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateTitle() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateTitle()
	})
}

// SetDetail sets the "detail" field.
func (u *DiscoveryUpsertBulk) SetDetail(v string) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetDetail(v)
	})
}

// UpdateDetail sets the "detail" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateDetail() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateDetail()
	})
}

// SetMtype sets the "mtype" field.
func (u *DiscoveryUpsertBulk) SetMtype(v discovery.Mtype) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetMtype(v)
	})
}

// UpdateMtype sets the "mtype" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateMtype() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateMtype()
	})
}

// SetLinks sets the "links" field.
func (u *DiscoveryUpsertBulk) SetLinks(v []string) *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.SetLinks(v)
	})
}

// UpdateLinks sets the "links" field to the value that was provided on create.
func (u *DiscoveryUpsertBulk) UpdateLinks() *DiscoveryUpsertBulk {
	return u.Update(func(s *DiscoveryUpsert) {
		s.UpdateLinks()
	})
}

// Exec executes the query.
func (u *DiscoveryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("media: OnConflict was set for builder %d. Set it on the DiscoveryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("media: missing options for DiscoveryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscoveryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}