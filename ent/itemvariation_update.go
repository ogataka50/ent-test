// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemvariation"
	"github.com/ogataka50/ent-test/ent/predicate"
)

// ItemVariationUpdate is the builder for updating ItemVariation entities.
type ItemVariationUpdate struct {
	config
	hooks    []Hook
	mutation *ItemVariationMutation
}

// Where adds a new predicate for the ItemVariationUpdate builder.
func (ivu *ItemVariationUpdate) Where(ps ...predicate.ItemVariation) *ItemVariationUpdate {
	ivu.mutation.predicates = append(ivu.mutation.predicates, ps...)
	return ivu
}

// SetVariantName sets the "variant_name" field.
func (ivu *ItemVariationUpdate) SetVariantName(s string) *ItemVariationUpdate {
	ivu.mutation.SetVariantName(s)
	return ivu
}

// SetOriginItemID sets the "origin_item" edge to the Item entity by ID.
func (ivu *ItemVariationUpdate) SetOriginItemID(id int) *ItemVariationUpdate {
	ivu.mutation.SetOriginItemID(id)
	return ivu
}

// SetOriginItem sets the "origin_item" edge to the Item entity.
func (ivu *ItemVariationUpdate) SetOriginItem(i *Item) *ItemVariationUpdate {
	return ivu.SetOriginItemID(i.ID)
}

// Mutation returns the ItemVariationMutation object of the builder.
func (ivu *ItemVariationUpdate) Mutation() *ItemVariationMutation {
	return ivu.mutation
}

// ClearOriginItem clears the "origin_item" edge to the Item entity.
func (ivu *ItemVariationUpdate) ClearOriginItem() *ItemVariationUpdate {
	ivu.mutation.ClearOriginItem()
	return ivu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ivu *ItemVariationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ivu.hooks) == 0 {
		if err = ivu.check(); err != nil {
			return 0, err
		}
		affected, err = ivu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemVariationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ivu.check(); err != nil {
				return 0, err
			}
			ivu.mutation = mutation
			affected, err = ivu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ivu.hooks) - 1; i >= 0; i-- {
			mut = ivu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ivu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ivu *ItemVariationUpdate) SaveX(ctx context.Context) int {
	affected, err := ivu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ivu *ItemVariationUpdate) Exec(ctx context.Context) error {
	_, err := ivu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ivu *ItemVariationUpdate) ExecX(ctx context.Context) {
	if err := ivu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ivu *ItemVariationUpdate) check() error {
	if _, ok := ivu.mutation.OriginItemID(); ivu.mutation.OriginItemCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"origin_item\"")
	}
	return nil
}

func (ivu *ItemVariationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemvariation.Table,
			Columns: itemvariation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemvariation.FieldID,
			},
		},
	}
	if ps := ivu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ivu.mutation.VariantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: itemvariation.FieldVariantName,
		})
	}
	if ivu.mutation.OriginItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemvariation.OriginItemTable,
			Columns: []string{itemvariation.OriginItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ivu.mutation.OriginItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemvariation.OriginItemTable,
			Columns: []string{itemvariation.OriginItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ivu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemvariation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ItemVariationUpdateOne is the builder for updating a single ItemVariation entity.
type ItemVariationUpdateOne struct {
	config
	hooks    []Hook
	mutation *ItemVariationMutation
}

// SetVariantName sets the "variant_name" field.
func (ivuo *ItemVariationUpdateOne) SetVariantName(s string) *ItemVariationUpdateOne {
	ivuo.mutation.SetVariantName(s)
	return ivuo
}

// SetOriginItemID sets the "origin_item" edge to the Item entity by ID.
func (ivuo *ItemVariationUpdateOne) SetOriginItemID(id int) *ItemVariationUpdateOne {
	ivuo.mutation.SetOriginItemID(id)
	return ivuo
}

// SetOriginItem sets the "origin_item" edge to the Item entity.
func (ivuo *ItemVariationUpdateOne) SetOriginItem(i *Item) *ItemVariationUpdateOne {
	return ivuo.SetOriginItemID(i.ID)
}

// Mutation returns the ItemVariationMutation object of the builder.
func (ivuo *ItemVariationUpdateOne) Mutation() *ItemVariationMutation {
	return ivuo.mutation
}

// ClearOriginItem clears the "origin_item" edge to the Item entity.
func (ivuo *ItemVariationUpdateOne) ClearOriginItem() *ItemVariationUpdateOne {
	ivuo.mutation.ClearOriginItem()
	return ivuo
}

// Save executes the query and returns the updated ItemVariation entity.
func (ivuo *ItemVariationUpdateOne) Save(ctx context.Context) (*ItemVariation, error) {
	var (
		err  error
		node *ItemVariation
	)
	if len(ivuo.hooks) == 0 {
		if err = ivuo.check(); err != nil {
			return nil, err
		}
		node, err = ivuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemVariationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ivuo.check(); err != nil {
				return nil, err
			}
			ivuo.mutation = mutation
			node, err = ivuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ivuo.hooks) - 1; i >= 0; i-- {
			mut = ivuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ivuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ivuo *ItemVariationUpdateOne) SaveX(ctx context.Context) *ItemVariation {
	node, err := ivuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ivuo *ItemVariationUpdateOne) Exec(ctx context.Context) error {
	_, err := ivuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ivuo *ItemVariationUpdateOne) ExecX(ctx context.Context) {
	if err := ivuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ivuo *ItemVariationUpdateOne) check() error {
	if _, ok := ivuo.mutation.OriginItemID(); ivuo.mutation.OriginItemCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"origin_item\"")
	}
	return nil
}

func (ivuo *ItemVariationUpdateOne) sqlSave(ctx context.Context) (_node *ItemVariation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemvariation.Table,
			Columns: itemvariation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemvariation.FieldID,
			},
		},
	}
	id, ok := ivuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ItemVariation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ivuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ivuo.mutation.VariantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: itemvariation.FieldVariantName,
		})
	}
	if ivuo.mutation.OriginItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemvariation.OriginItemTable,
			Columns: []string{itemvariation.OriginItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ivuo.mutation.OriginItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemvariation.OriginItemTable,
			Columns: []string{itemvariation.OriginItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ItemVariation{config: ivuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ivuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemvariation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
