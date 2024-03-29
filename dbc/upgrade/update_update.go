// Code generated by ent, DO NOT EDIT.

package upgrade

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/upgrade/predicate"
	"github.com/tikafog/of/dbc/upgrade/update"
)

// UpdateUpdate is the builder for updating Update entities.
type UpdateUpdate struct {
	config
	hooks    []Hook
	mutation *UpdateMutation
}

// Where appends a list predicates to the UpdateUpdate builder.
func (uu *UpdateUpdate) Where(ps ...predicate.Update) *UpdateUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetCreatedUnix sets the "created_unix" field.
func (uu *UpdateUpdate) SetCreatedUnix(i int64) *UpdateUpdate {
	uu.mutation.ResetCreatedUnix()
	uu.mutation.SetCreatedUnix(i)
	return uu
}

// SetNillableCreatedUnix sets the "created_unix" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableCreatedUnix(i *int64) *UpdateUpdate {
	if i != nil {
		uu.SetCreatedUnix(*i)
	}
	return uu
}

// AddCreatedUnix adds i to the "created_unix" field.
func (uu *UpdateUpdate) AddCreatedUnix(i int64) *UpdateUpdate {
	uu.mutation.AddCreatedUnix(i)
	return uu
}

// SetUpdatedUnix sets the "updated_unix" field.
func (uu *UpdateUpdate) SetUpdatedUnix(i int64) *UpdateUpdate {
	uu.mutation.ResetUpdatedUnix()
	uu.mutation.SetUpdatedUnix(i)
	return uu
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableUpdatedUnix(i *int64) *UpdateUpdate {
	if i != nil {
		uu.SetUpdatedUnix(*i)
	}
	return uu
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (uu *UpdateUpdate) AddUpdatedUnix(i int64) *UpdateUpdate {
	uu.mutation.AddUpdatedUnix(i)
	return uu
}

// SetDeletedUnix sets the "deleted_unix" field.
func (uu *UpdateUpdate) SetDeletedUnix(i int64) *UpdateUpdate {
	uu.mutation.ResetDeletedUnix()
	uu.mutation.SetDeletedUnix(i)
	return uu
}

// SetNillableDeletedUnix sets the "deleted_unix" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableDeletedUnix(i *int64) *UpdateUpdate {
	if i != nil {
		uu.SetDeletedUnix(*i)
	}
	return uu
}

// AddDeletedUnix adds i to the "deleted_unix" field.
func (uu *UpdateUpdate) AddDeletedUnix(i int64) *UpdateUpdate {
	uu.mutation.AddDeletedUnix(i)
	return uu
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (uu *UpdateUpdate) ClearDeletedUnix() *UpdateUpdate {
	uu.mutation.ClearDeletedUnix()
	return uu
}

// SetOs sets the "os" field.
func (uu *UpdateUpdate) SetOs(s string) *UpdateUpdate {
	uu.mutation.SetOs(s)
	return uu
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableOs(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetOs(*s)
	}
	return uu
}

// SetArch sets the "arch" field.
func (uu *UpdateUpdate) SetArch(s string) *UpdateUpdate {
	uu.mutation.SetArch(s)
	return uu
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableArch(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetArch(*s)
	}
	return uu
}

// SetVersion sets the "version" field.
func (uu *UpdateUpdate) SetVersion(s string) *UpdateUpdate {
	uu.mutation.SetVersion(s)
	return uu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableVersion(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetVersion(*s)
	}
	return uu
}

// SetRid sets the "rid" field.
func (uu *UpdateUpdate) SetRid(s string) *UpdateUpdate {
	uu.mutation.SetRid(s)
	return uu
}

// SetNillableRid sets the "rid" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableRid(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetRid(*s)
	}
	return uu
}

// SetCrc32 sets the "crc32" field.
func (uu *UpdateUpdate) SetCrc32(s string) *UpdateUpdate {
	uu.mutation.SetCrc32(s)
	return uu
}

// SetNillableCrc32 sets the "crc32" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableCrc32(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetCrc32(*s)
	}
	return uu
}

// SetAttr sets the "attr" field.
func (uu *UpdateUpdate) SetAttr(s string) *UpdateUpdate {
	uu.mutation.SetAttr(s)
	return uu
}

// SetNillableAttr sets the "attr" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableAttr(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetAttr(*s)
	}
	return uu
}

// SetForcibly sets the "forcibly" field.
func (uu *UpdateUpdate) SetForcibly(b bool) *UpdateUpdate {
	uu.mutation.SetForcibly(b)
	return uu
}

// SetNillableForcibly sets the "forcibly" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableForcibly(b *bool) *UpdateUpdate {
	if b != nil {
		uu.SetForcibly(*b)
	}
	return uu
}

// SetTruncate sets the "truncate" field.
func (uu *UpdateUpdate) SetTruncate(b bool) *UpdateUpdate {
	uu.mutation.SetTruncate(b)
	return uu
}

// SetNillableTruncate sets the "truncate" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableTruncate(b *bool) *UpdateUpdate {
	if b != nil {
		uu.SetTruncate(*b)
	}
	return uu
}

// SetTitle sets the "title" field.
func (uu *UpdateUpdate) SetTitle(s string) *UpdateUpdate {
	uu.mutation.SetTitle(s)
	return uu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableTitle(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetTitle(*s)
	}
	return uu
}

// SetDetail sets the "detail" field.
func (uu *UpdateUpdate) SetDetail(s string) *UpdateUpdate {
	uu.mutation.SetDetail(s)
	return uu
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableDetail(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetDetail(*s)
	}
	return uu
}

// SetSign sets the "sign" field.
func (uu *UpdateUpdate) SetSign(s string) *UpdateUpdate {
	uu.mutation.SetSign(s)
	return uu
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (uu *UpdateUpdate) SetNillableSign(s *string) *UpdateUpdate {
	if s != nil {
		uu.SetSign(*s)
	}
	return uu
}

// ClearSign clears the value of the "sign" field.
func (uu *UpdateUpdate) ClearSign() *UpdateUpdate {
	uu.mutation.ClearSign()
	return uu
}

// Mutation returns the UpdateMutation object of the builder.
func (uu *UpdateUpdate) Mutation() *UpdateMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UpdateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UpdateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("upgrade: uninitialized hook (forgotten import upgrade/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UpdateUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UpdateUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UpdateUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UpdateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   update.Table,
			Columns: update.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: update.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.CreatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldCreatedUnix,
		})
	}
	if value, ok := uu.mutation.AddedCreatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldCreatedUnix,
		})
	}
	if value, ok := uu.mutation.UpdatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldUpdatedUnix,
		})
	}
	if value, ok := uu.mutation.AddedUpdatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldUpdatedUnix,
		})
	}
	if value, ok := uu.mutation.DeletedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldDeletedUnix,
		})
	}
	if value, ok := uu.mutation.AddedDeletedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldDeletedUnix,
		})
	}
	if uu.mutation.DeletedUnixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: update.FieldDeletedUnix,
		})
	}
	if value, ok := uu.mutation.Os(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldOs,
		})
	}
	if value, ok := uu.mutation.Arch(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldArch,
		})
	}
	if value, ok := uu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldVersion,
		})
	}
	if value, ok := uu.mutation.Rid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldRid,
		})
	}
	if value, ok := uu.mutation.Crc32(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldCrc32,
		})
	}
	if value, ok := uu.mutation.Attr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldAttr,
		})
	}
	if value, ok := uu.mutation.Forcibly(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: update.FieldForcibly,
		})
	}
	if value, ok := uu.mutation.Truncate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: update.FieldTruncate,
		})
	}
	if value, ok := uu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldTitle,
		})
	}
	if value, ok := uu.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldDetail,
		})
	}
	if value, ok := uu.mutation.Sign(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldSign,
		})
	}
	if uu.mutation.SignCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: update.FieldSign,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{update.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UpdateUpdateOne is the builder for updating a single Update entity.
type UpdateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UpdateMutation
}

// SetCreatedUnix sets the "created_unix" field.
func (uuo *UpdateUpdateOne) SetCreatedUnix(i int64) *UpdateUpdateOne {
	uuo.mutation.ResetCreatedUnix()
	uuo.mutation.SetCreatedUnix(i)
	return uuo
}

// SetNillableCreatedUnix sets the "created_unix" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableCreatedUnix(i *int64) *UpdateUpdateOne {
	if i != nil {
		uuo.SetCreatedUnix(*i)
	}
	return uuo
}

// AddCreatedUnix adds i to the "created_unix" field.
func (uuo *UpdateUpdateOne) AddCreatedUnix(i int64) *UpdateUpdateOne {
	uuo.mutation.AddCreatedUnix(i)
	return uuo
}

// SetUpdatedUnix sets the "updated_unix" field.
func (uuo *UpdateUpdateOne) SetUpdatedUnix(i int64) *UpdateUpdateOne {
	uuo.mutation.ResetUpdatedUnix()
	uuo.mutation.SetUpdatedUnix(i)
	return uuo
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableUpdatedUnix(i *int64) *UpdateUpdateOne {
	if i != nil {
		uuo.SetUpdatedUnix(*i)
	}
	return uuo
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (uuo *UpdateUpdateOne) AddUpdatedUnix(i int64) *UpdateUpdateOne {
	uuo.mutation.AddUpdatedUnix(i)
	return uuo
}

// SetDeletedUnix sets the "deleted_unix" field.
func (uuo *UpdateUpdateOne) SetDeletedUnix(i int64) *UpdateUpdateOne {
	uuo.mutation.ResetDeletedUnix()
	uuo.mutation.SetDeletedUnix(i)
	return uuo
}

// SetNillableDeletedUnix sets the "deleted_unix" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableDeletedUnix(i *int64) *UpdateUpdateOne {
	if i != nil {
		uuo.SetDeletedUnix(*i)
	}
	return uuo
}

// AddDeletedUnix adds i to the "deleted_unix" field.
func (uuo *UpdateUpdateOne) AddDeletedUnix(i int64) *UpdateUpdateOne {
	uuo.mutation.AddDeletedUnix(i)
	return uuo
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (uuo *UpdateUpdateOne) ClearDeletedUnix() *UpdateUpdateOne {
	uuo.mutation.ClearDeletedUnix()
	return uuo
}

// SetOs sets the "os" field.
func (uuo *UpdateUpdateOne) SetOs(s string) *UpdateUpdateOne {
	uuo.mutation.SetOs(s)
	return uuo
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableOs(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetOs(*s)
	}
	return uuo
}

// SetArch sets the "arch" field.
func (uuo *UpdateUpdateOne) SetArch(s string) *UpdateUpdateOne {
	uuo.mutation.SetArch(s)
	return uuo
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableArch(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetArch(*s)
	}
	return uuo
}

// SetVersion sets the "version" field.
func (uuo *UpdateUpdateOne) SetVersion(s string) *UpdateUpdateOne {
	uuo.mutation.SetVersion(s)
	return uuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableVersion(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetVersion(*s)
	}
	return uuo
}

// SetRid sets the "rid" field.
func (uuo *UpdateUpdateOne) SetRid(s string) *UpdateUpdateOne {
	uuo.mutation.SetRid(s)
	return uuo
}

// SetNillableRid sets the "rid" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableRid(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetRid(*s)
	}
	return uuo
}

// SetCrc32 sets the "crc32" field.
func (uuo *UpdateUpdateOne) SetCrc32(s string) *UpdateUpdateOne {
	uuo.mutation.SetCrc32(s)
	return uuo
}

// SetNillableCrc32 sets the "crc32" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableCrc32(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetCrc32(*s)
	}
	return uuo
}

// SetAttr sets the "attr" field.
func (uuo *UpdateUpdateOne) SetAttr(s string) *UpdateUpdateOne {
	uuo.mutation.SetAttr(s)
	return uuo
}

// SetNillableAttr sets the "attr" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableAttr(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetAttr(*s)
	}
	return uuo
}

// SetForcibly sets the "forcibly" field.
func (uuo *UpdateUpdateOne) SetForcibly(b bool) *UpdateUpdateOne {
	uuo.mutation.SetForcibly(b)
	return uuo
}

// SetNillableForcibly sets the "forcibly" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableForcibly(b *bool) *UpdateUpdateOne {
	if b != nil {
		uuo.SetForcibly(*b)
	}
	return uuo
}

// SetTruncate sets the "truncate" field.
func (uuo *UpdateUpdateOne) SetTruncate(b bool) *UpdateUpdateOne {
	uuo.mutation.SetTruncate(b)
	return uuo
}

// SetNillableTruncate sets the "truncate" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableTruncate(b *bool) *UpdateUpdateOne {
	if b != nil {
		uuo.SetTruncate(*b)
	}
	return uuo
}

// SetTitle sets the "title" field.
func (uuo *UpdateUpdateOne) SetTitle(s string) *UpdateUpdateOne {
	uuo.mutation.SetTitle(s)
	return uuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableTitle(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetTitle(*s)
	}
	return uuo
}

// SetDetail sets the "detail" field.
func (uuo *UpdateUpdateOne) SetDetail(s string) *UpdateUpdateOne {
	uuo.mutation.SetDetail(s)
	return uuo
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableDetail(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetDetail(*s)
	}
	return uuo
}

// SetSign sets the "sign" field.
func (uuo *UpdateUpdateOne) SetSign(s string) *UpdateUpdateOne {
	uuo.mutation.SetSign(s)
	return uuo
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (uuo *UpdateUpdateOne) SetNillableSign(s *string) *UpdateUpdateOne {
	if s != nil {
		uuo.SetSign(*s)
	}
	return uuo
}

// ClearSign clears the value of the "sign" field.
func (uuo *UpdateUpdateOne) ClearSign() *UpdateUpdateOne {
	uuo.mutation.ClearSign()
	return uuo
}

// Mutation returns the UpdateMutation object of the builder.
func (uuo *UpdateUpdateOne) Mutation() *UpdateMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UpdateUpdateOne) Select(field string, fields ...string) *UpdateUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Update entity.
func (uuo *UpdateUpdateOne) Save(ctx context.Context) (*Update, error) {
	var (
		err  error
		node *Update
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UpdateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("upgrade: uninitialized hook (forgotten import upgrade/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Update)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UpdateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UpdateUpdateOne) SaveX(ctx context.Context) *Update {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UpdateUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UpdateUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UpdateUpdateOne) sqlSave(ctx context.Context) (_node *Update, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   update.Table,
			Columns: update.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: update.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`upgrade: missing "Update.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, update.FieldID)
		for _, f := range fields {
			if !update.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("upgrade: invalid field %q for query", f)}
			}
			if f != update.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.CreatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldCreatedUnix,
		})
	}
	if value, ok := uuo.mutation.AddedCreatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldCreatedUnix,
		})
	}
	if value, ok := uuo.mutation.UpdatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldUpdatedUnix,
		})
	}
	if value, ok := uuo.mutation.AddedUpdatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldUpdatedUnix,
		})
	}
	if value, ok := uuo.mutation.DeletedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldDeletedUnix,
		})
	}
	if value, ok := uuo.mutation.AddedDeletedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: update.FieldDeletedUnix,
		})
	}
	if uuo.mutation.DeletedUnixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: update.FieldDeletedUnix,
		})
	}
	if value, ok := uuo.mutation.Os(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldOs,
		})
	}
	if value, ok := uuo.mutation.Arch(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldArch,
		})
	}
	if value, ok := uuo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldVersion,
		})
	}
	if value, ok := uuo.mutation.Rid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldRid,
		})
	}
	if value, ok := uuo.mutation.Crc32(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldCrc32,
		})
	}
	if value, ok := uuo.mutation.Attr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldAttr,
		})
	}
	if value, ok := uuo.mutation.Forcibly(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: update.FieldForcibly,
		})
	}
	if value, ok := uuo.mutation.Truncate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: update.FieldTruncate,
		})
	}
	if value, ok := uuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldTitle,
		})
	}
	if value, ok := uuo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldDetail,
		})
	}
	if value, ok := uuo.mutation.Sign(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: update.FieldSign,
		})
	}
	if uuo.mutation.SignCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: update.FieldSign,
		})
	}
	_node = &Update{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{update.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
