// Code generated by entc, DO NOT EDIT.

package media

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/channel"
	"github.com/tikafog/of/dbc/media/informationv1"
	"github.com/tikafog/of/dbc/media/predicate"
)

// ChannelUpdate is the builder for updating Channel entities.
type ChannelUpdate struct {
	config
	hooks    []Hook
	mutation *ChannelMutation
}

// Where appends a list predicates to the ChannelUpdate builder.
func (cu *ChannelUpdate) Where(ps ...predicate.Channel) *ChannelUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedUnix sets the "created_unix" field.
func (cu *ChannelUpdate) SetCreatedUnix(i int64) *ChannelUpdate {
	cu.mutation.ResetCreatedUnix()
	cu.mutation.SetCreatedUnix(i)
	return cu
}

// SetNillableCreatedUnix sets the "created_unix" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableCreatedUnix(i *int64) *ChannelUpdate {
	if i != nil {
		cu.SetCreatedUnix(*i)
	}
	return cu
}

// AddCreatedUnix adds i to the "created_unix" field.
func (cu *ChannelUpdate) AddCreatedUnix(i int64) *ChannelUpdate {
	cu.mutation.AddCreatedUnix(i)
	return cu
}

// SetUpdatedUnix sets the "updated_unix" field.
func (cu *ChannelUpdate) SetUpdatedUnix(i int64) *ChannelUpdate {
	cu.mutation.ResetUpdatedUnix()
	cu.mutation.SetUpdatedUnix(i)
	return cu
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableUpdatedUnix(i *int64) *ChannelUpdate {
	if i != nil {
		cu.SetUpdatedUnix(*i)
	}
	return cu
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (cu *ChannelUpdate) AddUpdatedUnix(i int64) *ChannelUpdate {
	cu.mutation.AddUpdatedUnix(i)
	return cu
}

// SetDeletedUnix sets the "deleted_unix" field.
func (cu *ChannelUpdate) SetDeletedUnix(i int64) *ChannelUpdate {
	cu.mutation.ResetDeletedUnix()
	cu.mutation.SetDeletedUnix(i)
	return cu
}

// SetNillableDeletedUnix sets the "deleted_unix" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableDeletedUnix(i *int64) *ChannelUpdate {
	if i != nil {
		cu.SetDeletedUnix(*i)
	}
	return cu
}

// AddDeletedUnix adds i to the "deleted_unix" field.
func (cu *ChannelUpdate) AddDeletedUnix(i int64) *ChannelUpdate {
	cu.mutation.AddDeletedUnix(i)
	return cu
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (cu *ChannelUpdate) ClearDeletedUnix() *ChannelUpdate {
	cu.mutation.ClearDeletedUnix()
	return cu
}

// SetComment sets the "comment" field.
func (cu *ChannelUpdate) SetComment(s string) *ChannelUpdate {
	cu.mutation.SetComment(s)
	return cu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableComment(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetComment(*s)
	}
	return cu
}

// ClearComment clears the value of the "comment" field.
func (cu *ChannelUpdate) ClearComment() *ChannelUpdate {
	cu.mutation.ClearComment()
	return cu
}

// AddInformationIDs adds the "informations" edge to the InformationV1 entity by IDs.
func (cu *ChannelUpdate) AddInformationIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.AddInformationIDs(ids...)
	return cu
}

// AddInformations adds the "informations" edges to the InformationV1 entity.
func (cu *ChannelUpdate) AddInformations(i ...*InformationV1) *ChannelUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.AddInformationIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cu *ChannelUpdate) Mutation() *ChannelMutation {
	return cu.mutation
}

// ClearInformations clears all "informations" edges to the InformationV1 entity.
func (cu *ChannelUpdate) ClearInformations() *ChannelUpdate {
	cu.mutation.ClearInformations()
	return cu
}

// RemoveInformationIDs removes the "informations" edge to InformationV1 entities by IDs.
func (cu *ChannelUpdate) RemoveInformationIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.RemoveInformationIDs(ids...)
	return cu
}

// RemoveInformations removes "informations" edges to InformationV1 entities.
func (cu *ChannelUpdate) RemoveInformations(i ...*InformationV1) *ChannelUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.RemoveInformationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChannelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("media: uninitialized hook (forgotten import media/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChannelUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChannelUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: channel.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldCreatedUnix,
		})
	}
	if value, ok := cu.mutation.AddedCreatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldCreatedUnix,
		})
	}
	if value, ok := cu.mutation.UpdatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldUpdatedUnix,
		})
	}
	if value, ok := cu.mutation.AddedUpdatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldUpdatedUnix,
		})
	}
	if value, ok := cu.mutation.DeletedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldDeletedUnix,
		})
	}
	if value, ok := cu.mutation.AddedDeletedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldDeletedUnix,
		})
	}
	if cu.mutation.DeletedUnixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: channel.FieldDeletedUnix,
		})
	}
	if value, ok := cu.mutation.Comment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldComment,
		})
	}
	if cu.mutation.CommentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: channel.FieldComment,
		})
	}
	if cu.mutation.InformationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.InformationsTable,
			Columns: []string{channel.InformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: informationv1.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedInformationsIDs(); len(nodes) > 0 && !cu.mutation.InformationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.InformationsTable,
			Columns: []string{channel.InformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: informationv1.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.InformationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.InformationsTable,
			Columns: []string{channel.InformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: informationv1.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChannelUpdateOne is the builder for updating a single Channel entity.
type ChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChannelMutation
}

// SetCreatedUnix sets the "created_unix" field.
func (cuo *ChannelUpdateOne) SetCreatedUnix(i int64) *ChannelUpdateOne {
	cuo.mutation.ResetCreatedUnix()
	cuo.mutation.SetCreatedUnix(i)
	return cuo
}

// SetNillableCreatedUnix sets the "created_unix" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableCreatedUnix(i *int64) *ChannelUpdateOne {
	if i != nil {
		cuo.SetCreatedUnix(*i)
	}
	return cuo
}

// AddCreatedUnix adds i to the "created_unix" field.
func (cuo *ChannelUpdateOne) AddCreatedUnix(i int64) *ChannelUpdateOne {
	cuo.mutation.AddCreatedUnix(i)
	return cuo
}

// SetUpdatedUnix sets the "updated_unix" field.
func (cuo *ChannelUpdateOne) SetUpdatedUnix(i int64) *ChannelUpdateOne {
	cuo.mutation.ResetUpdatedUnix()
	cuo.mutation.SetUpdatedUnix(i)
	return cuo
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableUpdatedUnix(i *int64) *ChannelUpdateOne {
	if i != nil {
		cuo.SetUpdatedUnix(*i)
	}
	return cuo
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (cuo *ChannelUpdateOne) AddUpdatedUnix(i int64) *ChannelUpdateOne {
	cuo.mutation.AddUpdatedUnix(i)
	return cuo
}

// SetDeletedUnix sets the "deleted_unix" field.
func (cuo *ChannelUpdateOne) SetDeletedUnix(i int64) *ChannelUpdateOne {
	cuo.mutation.ResetDeletedUnix()
	cuo.mutation.SetDeletedUnix(i)
	return cuo
}

// SetNillableDeletedUnix sets the "deleted_unix" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableDeletedUnix(i *int64) *ChannelUpdateOne {
	if i != nil {
		cuo.SetDeletedUnix(*i)
	}
	return cuo
}

// AddDeletedUnix adds i to the "deleted_unix" field.
func (cuo *ChannelUpdateOne) AddDeletedUnix(i int64) *ChannelUpdateOne {
	cuo.mutation.AddDeletedUnix(i)
	return cuo
}

// ClearDeletedUnix clears the value of the "deleted_unix" field.
func (cuo *ChannelUpdateOne) ClearDeletedUnix() *ChannelUpdateOne {
	cuo.mutation.ClearDeletedUnix()
	return cuo
}

// SetComment sets the "comment" field.
func (cuo *ChannelUpdateOne) SetComment(s string) *ChannelUpdateOne {
	cuo.mutation.SetComment(s)
	return cuo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableComment(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetComment(*s)
	}
	return cuo
}

// ClearComment clears the value of the "comment" field.
func (cuo *ChannelUpdateOne) ClearComment() *ChannelUpdateOne {
	cuo.mutation.ClearComment()
	return cuo
}

// AddInformationIDs adds the "informations" edge to the InformationV1 entity by IDs.
func (cuo *ChannelUpdateOne) AddInformationIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.AddInformationIDs(ids...)
	return cuo
}

// AddInformations adds the "informations" edges to the InformationV1 entity.
func (cuo *ChannelUpdateOne) AddInformations(i ...*InformationV1) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.AddInformationIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cuo *ChannelUpdateOne) Mutation() *ChannelMutation {
	return cuo.mutation
}

// ClearInformations clears all "informations" edges to the InformationV1 entity.
func (cuo *ChannelUpdateOne) ClearInformations() *ChannelUpdateOne {
	cuo.mutation.ClearInformations()
	return cuo
}

// RemoveInformationIDs removes the "informations" edge to InformationV1 entities by IDs.
func (cuo *ChannelUpdateOne) RemoveInformationIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.RemoveInformationIDs(ids...)
	return cuo
}

// RemoveInformations removes "informations" edges to InformationV1 entities.
func (cuo *ChannelUpdateOne) RemoveInformations(i ...*InformationV1) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.RemoveInformationIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChannelUpdateOne) Select(field string, fields ...string) *ChannelUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Channel entity.
func (cuo *ChannelUpdateOne) Save(ctx context.Context) (*Channel, error) {
	var (
		err  error
		node *Channel
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("media: uninitialized hook (forgotten import media/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChannelUpdateOne) SaveX(ctx context.Context) *Channel {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChannelUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChannelUpdateOne) sqlSave(ctx context.Context) (_node *Channel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: channel.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`media: missing "Channel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channel.FieldID)
		for _, f := range fields {
			if !channel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("media: invalid field %q for query", f)}
			}
			if f != channel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldCreatedUnix,
		})
	}
	if value, ok := cuo.mutation.AddedCreatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldCreatedUnix,
		})
	}
	if value, ok := cuo.mutation.UpdatedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldUpdatedUnix,
		})
	}
	if value, ok := cuo.mutation.AddedUpdatedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldUpdatedUnix,
		})
	}
	if value, ok := cuo.mutation.DeletedUnix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldDeletedUnix,
		})
	}
	if value, ok := cuo.mutation.AddedDeletedUnix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: channel.FieldDeletedUnix,
		})
	}
	if cuo.mutation.DeletedUnixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: channel.FieldDeletedUnix,
		})
	}
	if value, ok := cuo.mutation.Comment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldComment,
		})
	}
	if cuo.mutation.CommentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: channel.FieldComment,
		})
	}
	if cuo.mutation.InformationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.InformationsTable,
			Columns: []string{channel.InformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: informationv1.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedInformationsIDs(); len(nodes) > 0 && !cuo.mutation.InformationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.InformationsTable,
			Columns: []string{channel.InformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: informationv1.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.InformationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.InformationsTable,
			Columns: []string{channel.InformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: informationv1.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Channel{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}