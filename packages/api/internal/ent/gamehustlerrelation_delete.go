// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/gamehustlerrelation"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/predicate"
)

// GameHustlerRelationDelete is the builder for deleting a GameHustlerRelation entity.
type GameHustlerRelationDelete struct {
	config
	hooks    []Hook
	mutation *GameHustlerRelationMutation
}

// Where appends a list predicates to the GameHustlerRelationDelete builder.
func (ghrd *GameHustlerRelationDelete) Where(ps ...predicate.GameHustlerRelation) *GameHustlerRelationDelete {
	ghrd.mutation.Where(ps...)
	return ghrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ghrd *GameHustlerRelationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ghrd.hooks) == 0 {
		affected, err = ghrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameHustlerRelationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ghrd.mutation = mutation
			affected, err = ghrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ghrd.hooks) - 1; i >= 0; i-- {
			if ghrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ghrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ghrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghrd *GameHustlerRelationDelete) ExecX(ctx context.Context) int {
	n, err := ghrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ghrd *GameHustlerRelationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: gamehustlerrelation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: gamehustlerrelation.FieldID,
			},
		},
	}
	if ps := ghrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ghrd.driver, _spec)
}

// GameHustlerRelationDeleteOne is the builder for deleting a single GameHustlerRelation entity.
type GameHustlerRelationDeleteOne struct {
	ghrd *GameHustlerRelationDelete
}

// Exec executes the deletion query.
func (ghrdo *GameHustlerRelationDeleteOne) Exec(ctx context.Context) error {
	n, err := ghrdo.ghrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gamehustlerrelation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ghrdo *GameHustlerRelationDeleteOne) ExecX(ctx context.Context) {
	ghrdo.ghrd.ExecX(ctx)
}