// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dopedao/dope-monorepo/packages/api/ent/bodypart"
	"github.com/dopedao/dope-monorepo/packages/api/ent/hustler"
	"github.com/dopedao/dope-monorepo/packages/api/ent/item"
	"github.com/dopedao/dope-monorepo/packages/api/ent/wallet"
)

// HustlerCreate is the builder for creating a Hustler entity.
type HustlerCreate struct {
	config
	mutation *HustlerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetType sets the "type" field.
func (hc *HustlerCreate) SetType(h hustler.Type) *HustlerCreate {
	hc.mutation.SetType(h)
	return hc
}

// SetName sets the "name" field.
func (hc *HustlerCreate) SetName(s string) *HustlerCreate {
	hc.mutation.SetName(s)
	return hc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (hc *HustlerCreate) SetNillableName(s *string) *HustlerCreate {
	if s != nil {
		hc.SetName(*s)
	}
	return hc
}

// SetTitle sets the "title" field.
func (hc *HustlerCreate) SetTitle(s string) *HustlerCreate {
	hc.mutation.SetTitle(s)
	return hc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (hc *HustlerCreate) SetNillableTitle(s *string) *HustlerCreate {
	if s != nil {
		hc.SetTitle(*s)
	}
	return hc
}

// SetColor sets the "color" field.
func (hc *HustlerCreate) SetColor(s string) *HustlerCreate {
	hc.mutation.SetColor(s)
	return hc
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (hc *HustlerCreate) SetNillableColor(s *string) *HustlerCreate {
	if s != nil {
		hc.SetColor(*s)
	}
	return hc
}

// SetBackground sets the "background" field.
func (hc *HustlerCreate) SetBackground(s string) *HustlerCreate {
	hc.mutation.SetBackground(s)
	return hc
}

// SetNillableBackground sets the "background" field if the given value is not nil.
func (hc *HustlerCreate) SetNillableBackground(s *string) *HustlerCreate {
	if s != nil {
		hc.SetBackground(*s)
	}
	return hc
}

// SetAge sets the "age" field.
func (hc *HustlerCreate) SetAge(u uint64) *HustlerCreate {
	hc.mutation.SetAge(u)
	return hc
}

// SetSex sets the "sex" field.
func (hc *HustlerCreate) SetSex(h hustler.Sex) *HustlerCreate {
	hc.mutation.SetSex(h)
	return hc
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (hc *HustlerCreate) SetNillableSex(h *hustler.Sex) *HustlerCreate {
	if h != nil {
		hc.SetSex(*h)
	}
	return hc
}

// SetViewbox sets the "viewbox" field.
func (hc *HustlerCreate) SetViewbox(i []int) *HustlerCreate {
	hc.mutation.SetViewbox(i)
	return hc
}

// SetOrder sets the "order" field.
func (hc *HustlerCreate) SetOrder(i []int) *HustlerCreate {
	hc.mutation.SetOrder(i)
	return hc
}

// SetSvg sets the "svg" field.
func (hc *HustlerCreate) SetSvg(s string) *HustlerCreate {
	hc.mutation.SetSvg(s)
	return hc
}

// SetNillableSvg sets the "svg" field if the given value is not nil.
func (hc *HustlerCreate) SetNillableSvg(s *string) *HustlerCreate {
	if s != nil {
		hc.SetSvg(*s)
	}
	return hc
}

// SetID sets the "id" field.
func (hc *HustlerCreate) SetID(s string) *HustlerCreate {
	hc.mutation.SetID(s)
	return hc
}

// SetWalletID sets the "wallet" edge to the Wallet entity by ID.
func (hc *HustlerCreate) SetWalletID(id string) *HustlerCreate {
	hc.mutation.SetWalletID(id)
	return hc
}

// SetNillableWalletID sets the "wallet" edge to the Wallet entity by ID if the given value is not nil.
func (hc *HustlerCreate) SetNillableWalletID(id *string) *HustlerCreate {
	if id != nil {
		hc = hc.SetWalletID(*id)
	}
	return hc
}

// SetWallet sets the "wallet" edge to the Wallet entity.
func (hc *HustlerCreate) SetWallet(w *Wallet) *HustlerCreate {
	return hc.SetWalletID(w.ID)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (hc *HustlerCreate) AddItemIDs(ids ...string) *HustlerCreate {
	hc.mutation.AddItemIDs(ids...)
	return hc
}

// AddItems adds the "items" edges to the Item entity.
func (hc *HustlerCreate) AddItems(i ...*Item) *HustlerCreate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return hc.AddItemIDs(ids...)
}

// SetBodyID sets the "body" edge to the BodyPart entity by ID.
func (hc *HustlerCreate) SetBodyID(id string) *HustlerCreate {
	hc.mutation.SetBodyID(id)
	return hc
}

// SetNillableBodyID sets the "body" edge to the BodyPart entity by ID if the given value is not nil.
func (hc *HustlerCreate) SetNillableBodyID(id *string) *HustlerCreate {
	if id != nil {
		hc = hc.SetBodyID(*id)
	}
	return hc
}

// SetBody sets the "body" edge to the BodyPart entity.
func (hc *HustlerCreate) SetBody(b *BodyPart) *HustlerCreate {
	return hc.SetBodyID(b.ID)
}

// SetHairID sets the "hair" edge to the BodyPart entity by ID.
func (hc *HustlerCreate) SetHairID(id string) *HustlerCreate {
	hc.mutation.SetHairID(id)
	return hc
}

// SetNillableHairID sets the "hair" edge to the BodyPart entity by ID if the given value is not nil.
func (hc *HustlerCreate) SetNillableHairID(id *string) *HustlerCreate {
	if id != nil {
		hc = hc.SetHairID(*id)
	}
	return hc
}

// SetHair sets the "hair" edge to the BodyPart entity.
func (hc *HustlerCreate) SetHair(b *BodyPart) *HustlerCreate {
	return hc.SetHairID(b.ID)
}

// SetBeardID sets the "beard" edge to the BodyPart entity by ID.
func (hc *HustlerCreate) SetBeardID(id string) *HustlerCreate {
	hc.mutation.SetBeardID(id)
	return hc
}

// SetNillableBeardID sets the "beard" edge to the BodyPart entity by ID if the given value is not nil.
func (hc *HustlerCreate) SetNillableBeardID(id *string) *HustlerCreate {
	if id != nil {
		hc = hc.SetBeardID(*id)
	}
	return hc
}

// SetBeard sets the "beard" edge to the BodyPart entity.
func (hc *HustlerCreate) SetBeard(b *BodyPart) *HustlerCreate {
	return hc.SetBeardID(b.ID)
}

// Mutation returns the HustlerMutation object of the builder.
func (hc *HustlerCreate) Mutation() *HustlerMutation {
	return hc.mutation
}

// Save creates the Hustler in the database.
func (hc *HustlerCreate) Save(ctx context.Context) (*Hustler, error) {
	var (
		err  error
		node *Hustler
	)
	hc.defaults()
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HustlerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			if node, err = hc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			if hc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HustlerCreate) SaveX(ctx context.Context) *Hustler {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HustlerCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HustlerCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HustlerCreate) defaults() {
	if _, ok := hc.mutation.Sex(); !ok {
		v := hustler.DefaultSex
		hc.mutation.SetSex(v)
	}
	if _, ok := hc.mutation.Viewbox(); !ok {
		v := hustler.DefaultViewbox
		hc.mutation.SetViewbox(v)
	}
	if _, ok := hc.mutation.Order(); !ok {
		v := hustler.DefaultOrder
		hc.mutation.SetOrder(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HustlerCreate) check() error {
	if _, ok := hc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Hustler.type"`)}
	}
	if v, ok := hc.mutation.GetType(); ok {
		if err := hustler.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Hustler.type": %w`, err)}
		}
	}
	if _, ok := hc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Hustler.age"`)}
	}
	if _, ok := hc.mutation.Sex(); !ok {
		return &ValidationError{Name: "sex", err: errors.New(`ent: missing required field "Hustler.sex"`)}
	}
	if v, ok := hc.mutation.Sex(); ok {
		if err := hustler.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf(`ent: validator failed for field "Hustler.sex": %w`, err)}
		}
	}
	if _, ok := hc.mutation.Viewbox(); !ok {
		return &ValidationError{Name: "viewbox", err: errors.New(`ent: missing required field "Hustler.viewbox"`)}
	}
	if _, ok := hc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Hustler.order"`)}
	}
	return nil
}

func (hc *HustlerCreate) sqlSave(ctx context.Context) (*Hustler, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Hustler.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (hc *HustlerCreate) createSpec() (*Hustler, *sqlgraph.CreateSpec) {
	var (
		_node = &Hustler{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hustler.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: hustler.FieldID,
			},
		}
	)
	_spec.OnConflict = hc.conflict
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: hustler.FieldType,
		})
		_node.Type = value
	}
	if value, ok := hc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hustler.FieldName,
		})
		_node.Name = value
	}
	if value, ok := hc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hustler.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := hc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hustler.FieldColor,
		})
		_node.Color = value
	}
	if value, ok := hc.mutation.Background(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hustler.FieldBackground,
		})
		_node.Background = value
	}
	if value, ok := hc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: hustler.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := hc.mutation.Sex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: hustler.FieldSex,
		})
		_node.Sex = value
	}
	if value, ok := hc.mutation.Viewbox(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: hustler.FieldViewbox,
		})
		_node.Viewbox = value
	}
	if value, ok := hc.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: hustler.FieldOrder,
		})
		_node.Order = value
	}
	if value, ok := hc.mutation.Svg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hustler.FieldSvg,
		})
		_node.Svg = value
	}
	if nodes := hc.mutation.WalletIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hustler.WalletTable,
			Columns: []string{hustler.WalletColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: wallet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.wallet_hustlers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hustler.ItemsTable,
			Columns: []string{hustler.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.BodyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hustler.BodyTable,
			Columns: []string{hustler.BodyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: bodypart.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.body_part_hustler_bodies = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.HairIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hustler.HairTable,
			Columns: []string{hustler.HairColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: bodypart.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.body_part_hustler_hairs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.BeardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hustler.BeardTable,
			Columns: []string{hustler.BeardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: bodypart.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.body_part_hustler_beards = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Hustler.Create().
//		SetType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HustlerUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
//
func (hc *HustlerCreate) OnConflict(opts ...sql.ConflictOption) *HustlerUpsertOne {
	hc.conflict = opts
	return &HustlerUpsertOne{
		create: hc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Hustler.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (hc *HustlerCreate) OnConflictColumns(columns ...string) *HustlerUpsertOne {
	hc.conflict = append(hc.conflict, sql.ConflictColumns(columns...))
	return &HustlerUpsertOne{
		create: hc,
	}
}

type (
	// HustlerUpsertOne is the builder for "upsert"-ing
	//  one Hustler node.
	HustlerUpsertOne struct {
		create *HustlerCreate
	}

	// HustlerUpsert is the "OnConflict" setter.
	HustlerUpsert struct {
		*sql.UpdateSet
	}
)

// SetType sets the "type" field.
func (u *HustlerUpsert) SetType(v hustler.Type) *HustlerUpsert {
	u.Set(hustler.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateType() *HustlerUpsert {
	u.SetExcluded(hustler.FieldType)
	return u
}

// SetName sets the "name" field.
func (u *HustlerUpsert) SetName(v string) *HustlerUpsert {
	u.Set(hustler.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateName() *HustlerUpsert {
	u.SetExcluded(hustler.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *HustlerUpsert) ClearName() *HustlerUpsert {
	u.SetNull(hustler.FieldName)
	return u
}

// SetTitle sets the "title" field.
func (u *HustlerUpsert) SetTitle(v string) *HustlerUpsert {
	u.Set(hustler.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateTitle() *HustlerUpsert {
	u.SetExcluded(hustler.FieldTitle)
	return u
}

// ClearTitle clears the value of the "title" field.
func (u *HustlerUpsert) ClearTitle() *HustlerUpsert {
	u.SetNull(hustler.FieldTitle)
	return u
}

// SetColor sets the "color" field.
func (u *HustlerUpsert) SetColor(v string) *HustlerUpsert {
	u.Set(hustler.FieldColor, v)
	return u
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateColor() *HustlerUpsert {
	u.SetExcluded(hustler.FieldColor)
	return u
}

// ClearColor clears the value of the "color" field.
func (u *HustlerUpsert) ClearColor() *HustlerUpsert {
	u.SetNull(hustler.FieldColor)
	return u
}

// SetBackground sets the "background" field.
func (u *HustlerUpsert) SetBackground(v string) *HustlerUpsert {
	u.Set(hustler.FieldBackground, v)
	return u
}

// UpdateBackground sets the "background" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateBackground() *HustlerUpsert {
	u.SetExcluded(hustler.FieldBackground)
	return u
}

// ClearBackground clears the value of the "background" field.
func (u *HustlerUpsert) ClearBackground() *HustlerUpsert {
	u.SetNull(hustler.FieldBackground)
	return u
}

// SetAge sets the "age" field.
func (u *HustlerUpsert) SetAge(v uint64) *HustlerUpsert {
	u.Set(hustler.FieldAge, v)
	return u
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateAge() *HustlerUpsert {
	u.SetExcluded(hustler.FieldAge)
	return u
}

// AddAge adds v to the "age" field.
func (u *HustlerUpsert) AddAge(v uint64) *HustlerUpsert {
	u.Add(hustler.FieldAge, v)
	return u
}

// SetSex sets the "sex" field.
func (u *HustlerUpsert) SetSex(v hustler.Sex) *HustlerUpsert {
	u.Set(hustler.FieldSex, v)
	return u
}

// UpdateSex sets the "sex" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateSex() *HustlerUpsert {
	u.SetExcluded(hustler.FieldSex)
	return u
}

// SetViewbox sets the "viewbox" field.
func (u *HustlerUpsert) SetViewbox(v []int) *HustlerUpsert {
	u.Set(hustler.FieldViewbox, v)
	return u
}

// UpdateViewbox sets the "viewbox" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateViewbox() *HustlerUpsert {
	u.SetExcluded(hustler.FieldViewbox)
	return u
}

// SetOrder sets the "order" field.
func (u *HustlerUpsert) SetOrder(v []int) *HustlerUpsert {
	u.Set(hustler.FieldOrder, v)
	return u
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateOrder() *HustlerUpsert {
	u.SetExcluded(hustler.FieldOrder)
	return u
}

// SetSvg sets the "svg" field.
func (u *HustlerUpsert) SetSvg(v string) *HustlerUpsert {
	u.Set(hustler.FieldSvg, v)
	return u
}

// UpdateSvg sets the "svg" field to the value that was provided on create.
func (u *HustlerUpsert) UpdateSvg() *HustlerUpsert {
	u.SetExcluded(hustler.FieldSvg)
	return u
}

// ClearSvg clears the value of the "svg" field.
func (u *HustlerUpsert) ClearSvg() *HustlerUpsert {
	u.SetNull(hustler.FieldSvg)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Hustler.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(hustler.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *HustlerUpsertOne) UpdateNewValues() *HustlerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(hustler.FieldID)
		}
		if _, exists := u.create.mutation.GetType(); exists {
			s.SetIgnore(hustler.FieldType)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Hustler.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *HustlerUpsertOne) Ignore() *HustlerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HustlerUpsertOne) DoNothing() *HustlerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HustlerCreate.OnConflict
// documentation for more info.
func (u *HustlerUpsertOne) Update(set func(*HustlerUpsert)) *HustlerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HustlerUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *HustlerUpsertOne) SetType(v hustler.Type) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateType() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateType()
	})
}

// SetName sets the "name" field.
func (u *HustlerUpsertOne) SetName(v string) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateName() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *HustlerUpsertOne) ClearName() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearName()
	})
}

// SetTitle sets the "title" field.
func (u *HustlerUpsertOne) SetTitle(v string) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateTitle() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *HustlerUpsertOne) ClearTitle() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearTitle()
	})
}

// SetColor sets the "color" field.
func (u *HustlerUpsertOne) SetColor(v string) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetColor(v)
	})
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateColor() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateColor()
	})
}

// ClearColor clears the value of the "color" field.
func (u *HustlerUpsertOne) ClearColor() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearColor()
	})
}

// SetBackground sets the "background" field.
func (u *HustlerUpsertOne) SetBackground(v string) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetBackground(v)
	})
}

// UpdateBackground sets the "background" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateBackground() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateBackground()
	})
}

// ClearBackground clears the value of the "background" field.
func (u *HustlerUpsertOne) ClearBackground() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearBackground()
	})
}

// SetAge sets the "age" field.
func (u *HustlerUpsertOne) SetAge(v uint64) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *HustlerUpsertOne) AddAge(v uint64) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateAge() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateAge()
	})
}

// SetSex sets the "sex" field.
func (u *HustlerUpsertOne) SetSex(v hustler.Sex) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetSex(v)
	})
}

// UpdateSex sets the "sex" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateSex() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateSex()
	})
}

// SetViewbox sets the "viewbox" field.
func (u *HustlerUpsertOne) SetViewbox(v []int) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetViewbox(v)
	})
}

// UpdateViewbox sets the "viewbox" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateViewbox() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateViewbox()
	})
}

// SetOrder sets the "order" field.
func (u *HustlerUpsertOne) SetOrder(v []int) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateOrder() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateOrder()
	})
}

// SetSvg sets the "svg" field.
func (u *HustlerUpsertOne) SetSvg(v string) *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.SetSvg(v)
	})
}

// UpdateSvg sets the "svg" field to the value that was provided on create.
func (u *HustlerUpsertOne) UpdateSvg() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateSvg()
	})
}

// ClearSvg clears the value of the "svg" field.
func (u *HustlerUpsertOne) ClearSvg() *HustlerUpsertOne {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearSvg()
	})
}

// Exec executes the query.
func (u *HustlerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HustlerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HustlerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *HustlerUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: HustlerUpsertOne.ID is not supported by MySQL driver. Use HustlerUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *HustlerUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// HustlerCreateBulk is the builder for creating many Hustler entities in bulk.
type HustlerCreateBulk struct {
	config
	builders []*HustlerCreate
	conflict []sql.ConflictOption
}

// Save creates the Hustler entities in the database.
func (hcb *HustlerCreateBulk) Save(ctx context.Context) ([]*Hustler, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Hustler, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HustlerMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = hcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HustlerCreateBulk) SaveX(ctx context.Context) []*Hustler {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HustlerCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HustlerCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Hustler.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HustlerUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
//
func (hcb *HustlerCreateBulk) OnConflict(opts ...sql.ConflictOption) *HustlerUpsertBulk {
	hcb.conflict = opts
	return &HustlerUpsertBulk{
		create: hcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Hustler.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (hcb *HustlerCreateBulk) OnConflictColumns(columns ...string) *HustlerUpsertBulk {
	hcb.conflict = append(hcb.conflict, sql.ConflictColumns(columns...))
	return &HustlerUpsertBulk{
		create: hcb,
	}
}

// HustlerUpsertBulk is the builder for "upsert"-ing
// a bulk of Hustler nodes.
type HustlerUpsertBulk struct {
	create *HustlerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Hustler.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(hustler.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *HustlerUpsertBulk) UpdateNewValues() *HustlerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(hustler.FieldID)
				return
			}
			if _, exists := b.mutation.GetType(); exists {
				s.SetIgnore(hustler.FieldType)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Hustler.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *HustlerUpsertBulk) Ignore() *HustlerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HustlerUpsertBulk) DoNothing() *HustlerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HustlerCreateBulk.OnConflict
// documentation for more info.
func (u *HustlerUpsertBulk) Update(set func(*HustlerUpsert)) *HustlerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HustlerUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *HustlerUpsertBulk) SetType(v hustler.Type) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateType() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateType()
	})
}

// SetName sets the "name" field.
func (u *HustlerUpsertBulk) SetName(v string) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateName() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *HustlerUpsertBulk) ClearName() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearName()
	})
}

// SetTitle sets the "title" field.
func (u *HustlerUpsertBulk) SetTitle(v string) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateTitle() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *HustlerUpsertBulk) ClearTitle() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearTitle()
	})
}

// SetColor sets the "color" field.
func (u *HustlerUpsertBulk) SetColor(v string) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetColor(v)
	})
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateColor() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateColor()
	})
}

// ClearColor clears the value of the "color" field.
func (u *HustlerUpsertBulk) ClearColor() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearColor()
	})
}

// SetBackground sets the "background" field.
func (u *HustlerUpsertBulk) SetBackground(v string) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetBackground(v)
	})
}

// UpdateBackground sets the "background" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateBackground() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateBackground()
	})
}

// ClearBackground clears the value of the "background" field.
func (u *HustlerUpsertBulk) ClearBackground() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearBackground()
	})
}

// SetAge sets the "age" field.
func (u *HustlerUpsertBulk) SetAge(v uint64) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *HustlerUpsertBulk) AddAge(v uint64) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateAge() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateAge()
	})
}

// SetSex sets the "sex" field.
func (u *HustlerUpsertBulk) SetSex(v hustler.Sex) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetSex(v)
	})
}

// UpdateSex sets the "sex" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateSex() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateSex()
	})
}

// SetViewbox sets the "viewbox" field.
func (u *HustlerUpsertBulk) SetViewbox(v []int) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetViewbox(v)
	})
}

// UpdateViewbox sets the "viewbox" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateViewbox() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateViewbox()
	})
}

// SetOrder sets the "order" field.
func (u *HustlerUpsertBulk) SetOrder(v []int) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateOrder() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateOrder()
	})
}

// SetSvg sets the "svg" field.
func (u *HustlerUpsertBulk) SetSvg(v string) *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.SetSvg(v)
	})
}

// UpdateSvg sets the "svg" field to the value that was provided on create.
func (u *HustlerUpsertBulk) UpdateSvg() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.UpdateSvg()
	})
}

// ClearSvg clears the value of the "svg" field.
func (u *HustlerUpsertBulk) ClearSvg() *HustlerUpsertBulk {
	return u.Update(func(s *HustlerUpsert) {
		s.ClearSvg()
	})
}

// Exec executes the query.
func (u *HustlerUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the HustlerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HustlerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HustlerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
