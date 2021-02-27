// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemgroup"
	"github.com/ogataka50/ent-test/ent/predicate"
)

// ItemGroupUpdate is the builder for updating ItemGroup entities.
type ItemGroupUpdate struct {
	config
	hooks    []Hook
	mutation *ItemGroupMutation
}

// Where adds a new predicate for the ItemGroupUpdate builder.
func (igu *ItemGroupUpdate) Where(ps ...predicate.ItemGroup) *ItemGroupUpdate {
	igu.mutation.predicates = append(igu.mutation.predicates, ps...)
	return igu
}

// SetName sets the "name" field.
func (igu *ItemGroupUpdate) SetName(s string) *ItemGroupUpdate {
	igu.mutation.SetName(s)
	return igu
}

// AddGroupItemIDs adds the "group_item" edge to the Item entity by IDs.
func (igu *ItemGroupUpdate) AddGroupItemIDs(ids ...int) *ItemGroupUpdate {
	igu.mutation.AddGroupItemIDs(ids...)
	return igu
}

// AddGroupItem adds the "group_item" edges to the Item entity.
func (igu *ItemGroupUpdate) AddGroupItem(i ...*Item) *ItemGroupUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return igu.AddGroupItemIDs(ids...)
}

// Mutation returns the ItemGroupMutation object of the builder.
func (igu *ItemGroupUpdate) Mutation() *ItemGroupMutation {
	return igu.mutation
}

// ClearGroupItem clears all "group_item" edges to the Item entity.
func (igu *ItemGroupUpdate) ClearGroupItem() *ItemGroupUpdate {
	igu.mutation.ClearGroupItem()
	return igu
}

// RemoveGroupItemIDs removes the "group_item" edge to Item entities by IDs.
func (igu *ItemGroupUpdate) RemoveGroupItemIDs(ids ...int) *ItemGroupUpdate {
	igu.mutation.RemoveGroupItemIDs(ids...)
	return igu
}

// RemoveGroupItem removes "group_item" edges to Item entities.
func (igu *ItemGroupUpdate) RemoveGroupItem(i ...*Item) *ItemGroupUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return igu.RemoveGroupItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (igu *ItemGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(igu.hooks) == 0 {
		affected, err = igu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			igu.mutation = mutation
			affected, err = igu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(igu.hooks) - 1; i >= 0; i-- {
			mut = igu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, igu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (igu *ItemGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := igu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (igu *ItemGroupUpdate) Exec(ctx context.Context) error {
	_, err := igu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (igu *ItemGroupUpdate) ExecX(ctx context.Context) {
	if err := igu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (igu *ItemGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemgroup.Table,
			Columns: itemgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemgroup.FieldID,
			},
		},
	}
	if ps := igu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := igu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: itemgroup.FieldName,
		})
	}
	if igu.mutation.GroupItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   itemgroup.GroupItemTable,
			Columns: itemgroup.GroupItemPrimaryKey,
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
	if nodes := igu.mutation.RemovedGroupItemIDs(); len(nodes) > 0 && !igu.mutation.GroupItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   itemgroup.GroupItemTable,
			Columns: itemgroup.GroupItemPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := igu.mutation.GroupItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   itemgroup.GroupItemTable,
			Columns: itemgroup.GroupItemPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, igu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemgroup.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ItemGroupUpdateOne is the builder for updating a single ItemGroup entity.
type ItemGroupUpdateOne struct {
	config
	hooks    []Hook
	mutation *ItemGroupMutation
}

// SetName sets the "name" field.
func (iguo *ItemGroupUpdateOne) SetName(s string) *ItemGroupUpdateOne {
	iguo.mutation.SetName(s)
	return iguo
}

// AddGroupItemIDs adds the "group_item" edge to the Item entity by IDs.
func (iguo *ItemGroupUpdateOne) AddGroupItemIDs(ids ...int) *ItemGroupUpdateOne {
	iguo.mutation.AddGroupItemIDs(ids...)
	return iguo
}

// AddGroupItem adds the "group_item" edges to the Item entity.
func (iguo *ItemGroupUpdateOne) AddGroupItem(i ...*Item) *ItemGroupUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iguo.AddGroupItemIDs(ids...)
}

// Mutation returns the ItemGroupMutation object of the builder.
func (iguo *ItemGroupUpdateOne) Mutation() *ItemGroupMutation {
	return iguo.mutation
}

// ClearGroupItem clears all "group_item" edges to the Item entity.
func (iguo *ItemGroupUpdateOne) ClearGroupItem() *ItemGroupUpdateOne {
	iguo.mutation.ClearGroupItem()
	return iguo
}

// RemoveGroupItemIDs removes the "group_item" edge to Item entities by IDs.
func (iguo *ItemGroupUpdateOne) RemoveGroupItemIDs(ids ...int) *ItemGroupUpdateOne {
	iguo.mutation.RemoveGroupItemIDs(ids...)
	return iguo
}

// RemoveGroupItem removes "group_item" edges to Item entities.
func (iguo *ItemGroupUpdateOne) RemoveGroupItem(i ...*Item) *ItemGroupUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iguo.RemoveGroupItemIDs(ids...)
}

// Save executes the query and returns the updated ItemGroup entity.
func (iguo *ItemGroupUpdateOne) Save(ctx context.Context) (*ItemGroup, error) {
	var (
		err  error
		node *ItemGroup
	)
	if len(iguo.hooks) == 0 {
		node, err = iguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iguo.mutation = mutation
			node, err = iguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iguo.hooks) - 1; i >= 0; i-- {
			mut = iguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iguo *ItemGroupUpdateOne) SaveX(ctx context.Context) *ItemGroup {
	node, err := iguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iguo *ItemGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := iguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iguo *ItemGroupUpdateOne) ExecX(ctx context.Context) {
	if err := iguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iguo *ItemGroupUpdateOne) sqlSave(ctx context.Context) (_node *ItemGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemgroup.Table,
			Columns: itemgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemgroup.FieldID,
			},
		},
	}
	id, ok := iguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ItemGroup.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := iguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iguo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: itemgroup.FieldName,
		})
	}
	if iguo.mutation.GroupItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   itemgroup.GroupItemTable,
			Columns: itemgroup.GroupItemPrimaryKey,
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
	if nodes := iguo.mutation.RemovedGroupItemIDs(); len(nodes) > 0 && !iguo.mutation.GroupItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   itemgroup.GroupItemTable,
			Columns: itemgroup.GroupItemPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iguo.mutation.GroupItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   itemgroup.GroupItemTable,
			Columns: itemgroup.GroupItemPrimaryKey,
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
	_node = &ItemGroup{config: iguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemgroup.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}