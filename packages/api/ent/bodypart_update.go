// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dopedao/dope-monorepo/packages/api/ent/bodypart"
	"github.com/dopedao/dope-monorepo/packages/api/ent/hustler"
	"github.com/dopedao/dope-monorepo/packages/api/ent/predicate"
)

// BodyPartUpdate is the builder for updating BodyPart entities.
type BodyPartUpdate struct {
	config
	hooks    []Hook
	mutation *BodyPartMutation
}

// Where appends a list predicates to the BodyPartUpdate builder.
func (bpu *BodyPartUpdate) Where(ps ...predicate.BodyPart) *BodyPartUpdate {
	bpu.mutation.Where(ps...)
	return bpu
}

// SetHustlerID sets the "hustler" edge to the Hustler entity by ID.
func (bpu *BodyPartUpdate) SetHustlerID(id string) *BodyPartUpdate {
	bpu.mutation.SetHustlerID(id)
	return bpu
}

// SetNillableHustlerID sets the "hustler" edge to the Hustler entity by ID if the given value is not nil.
func (bpu *BodyPartUpdate) SetNillableHustlerID(id *string) *BodyPartUpdate {
	if id != nil {
		bpu = bpu.SetHustlerID(*id)
	}
	return bpu
}

// SetHustler sets the "hustler" edge to the Hustler entity.
func (bpu *BodyPartUpdate) SetHustler(h *Hustler) *BodyPartUpdate {
	return bpu.SetHustlerID(h.ID)
}

// Mutation returns the BodyPartMutation object of the builder.
func (bpu *BodyPartUpdate) Mutation() *BodyPartMutation {
	return bpu.mutation
}

// ClearHustler clears the "hustler" edge to the Hustler entity.
func (bpu *BodyPartUpdate) ClearHustler() *BodyPartUpdate {
	bpu.mutation.ClearHustler()
	return bpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bpu *BodyPartUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bpu.hooks) == 0 {
		affected, err = bpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BodyPartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bpu.mutation = mutation
			affected, err = bpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bpu.hooks) - 1; i >= 0; i-- {
			if bpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bpu *BodyPartUpdate) SaveX(ctx context.Context) int {
	affected, err := bpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bpu *BodyPartUpdate) Exec(ctx context.Context) error {
	_, err := bpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpu *BodyPartUpdate) ExecX(ctx context.Context) {
	if err := bpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bpu *BodyPartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bodypart.Table,
			Columns: bodypart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: bodypart.FieldID,
			},
		},
	}
	if ps := bpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if bpu.mutation.HustlerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bodypart.HustlerTable,
			Columns: []string{bodypart.HustlerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: hustler.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bpu.mutation.HustlerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bodypart.HustlerTable,
			Columns: []string{bodypart.HustlerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: hustler.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bodypart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BodyPartUpdateOne is the builder for updating a single BodyPart entity.
type BodyPartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BodyPartMutation
}

// SetHustlerID sets the "hustler" edge to the Hustler entity by ID.
func (bpuo *BodyPartUpdateOne) SetHustlerID(id string) *BodyPartUpdateOne {
	bpuo.mutation.SetHustlerID(id)
	return bpuo
}

// SetNillableHustlerID sets the "hustler" edge to the Hustler entity by ID if the given value is not nil.
func (bpuo *BodyPartUpdateOne) SetNillableHustlerID(id *string) *BodyPartUpdateOne {
	if id != nil {
		bpuo = bpuo.SetHustlerID(*id)
	}
	return bpuo
}

// SetHustler sets the "hustler" edge to the Hustler entity.
func (bpuo *BodyPartUpdateOne) SetHustler(h *Hustler) *BodyPartUpdateOne {
	return bpuo.SetHustlerID(h.ID)
}

// Mutation returns the BodyPartMutation object of the builder.
func (bpuo *BodyPartUpdateOne) Mutation() *BodyPartMutation {
	return bpuo.mutation
}

// ClearHustler clears the "hustler" edge to the Hustler entity.
func (bpuo *BodyPartUpdateOne) ClearHustler() *BodyPartUpdateOne {
	bpuo.mutation.ClearHustler()
	return bpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bpuo *BodyPartUpdateOne) Select(field string, fields ...string) *BodyPartUpdateOne {
	bpuo.fields = append([]string{field}, fields...)
	return bpuo
}

// Save executes the query and returns the updated BodyPart entity.
func (bpuo *BodyPartUpdateOne) Save(ctx context.Context) (*BodyPart, error) {
	var (
		err  error
		node *BodyPart
	)
	if len(bpuo.hooks) == 0 {
		node, err = bpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BodyPartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bpuo.mutation = mutation
			node, err = bpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bpuo.hooks) - 1; i >= 0; i-- {
			if bpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bpuo *BodyPartUpdateOne) SaveX(ctx context.Context) *BodyPart {
	node, err := bpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bpuo *BodyPartUpdateOne) Exec(ctx context.Context) error {
	_, err := bpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpuo *BodyPartUpdateOne) ExecX(ctx context.Context) {
	if err := bpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bpuo *BodyPartUpdateOne) sqlSave(ctx context.Context) (_node *BodyPart, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bodypart.Table,
			Columns: bodypart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: bodypart.FieldID,
			},
		},
	}
	id, ok := bpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BodyPart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bodypart.FieldID)
		for _, f := range fields {
			if !bodypart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bodypart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if bpuo.mutation.HustlerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bodypart.HustlerTable,
			Columns: []string{bodypart.HustlerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: hustler.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bpuo.mutation.HustlerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bodypart.HustlerTable,
			Columns: []string{bodypart.HustlerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: hustler.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BodyPart{config: bpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bodypart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
