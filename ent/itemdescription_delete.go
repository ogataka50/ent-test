// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogataka50/ent-test/ent/itemdescription"
	"github.com/ogataka50/ent-test/ent/predicate"
)

// ItemDescriptionDelete is the builder for deleting a ItemDescription entity.
type ItemDescriptionDelete struct {
	config
	hooks    []Hook
	mutation *ItemDescriptionMutation
}

// Where adds a new predicate to the ItemDescriptionDelete builder.
func (idd *ItemDescriptionDelete) Where(ps ...predicate.ItemDescription) *ItemDescriptionDelete {
	idd.mutation.predicates = append(idd.mutation.predicates, ps...)
	return idd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (idd *ItemDescriptionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(idd.hooks) == 0 {
		affected, err = idd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			idd.mutation = mutation
			affected, err = idd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(idd.hooks) - 1; i >= 0; i-- {
			mut = idd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, idd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (idd *ItemDescriptionDelete) ExecX(ctx context.Context) int {
	n, err := idd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (idd *ItemDescriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: itemdescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemdescription.FieldID,
			},
		},
	}
	if ps := idd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, idd.driver, _spec)
}

// ItemDescriptionDeleteOne is the builder for deleting a single ItemDescription entity.
type ItemDescriptionDeleteOne struct {
	idd *ItemDescriptionDelete
}

// Exec executes the deletion query.
func (iddo *ItemDescriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := iddo.idd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{itemdescription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iddo *ItemDescriptionDeleteOne) ExecX(ctx context.Context) {
	iddo.idd.ExecX(ctx)
}
