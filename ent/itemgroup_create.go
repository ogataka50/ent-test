// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemgroup"
)

// ItemGroupCreate is the builder for creating a ItemGroup entity.
type ItemGroupCreate struct {
	config
	mutation *ItemGroupMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (igc *ItemGroupCreate) SetName(s string) *ItemGroupCreate {
	igc.mutation.SetName(s)
	return igc
}

// AddGroupItemIDs adds the "group_item" edge to the Item entity by IDs.
func (igc *ItemGroupCreate) AddGroupItemIDs(ids ...int) *ItemGroupCreate {
	igc.mutation.AddGroupItemIDs(ids...)
	return igc
}

// AddGroupItem adds the "group_item" edges to the Item entity.
func (igc *ItemGroupCreate) AddGroupItem(i ...*Item) *ItemGroupCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return igc.AddGroupItemIDs(ids...)
}

// Mutation returns the ItemGroupMutation object of the builder.
func (igc *ItemGroupCreate) Mutation() *ItemGroupMutation {
	return igc.mutation
}

// Save creates the ItemGroup in the database.
func (igc *ItemGroupCreate) Save(ctx context.Context) (*ItemGroup, error) {
	var (
		err  error
		node *ItemGroup
	)
	if len(igc.hooks) == 0 {
		if err = igc.check(); err != nil {
			return nil, err
		}
		node, err = igc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = igc.check(); err != nil {
				return nil, err
			}
			igc.mutation = mutation
			node, err = igc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(igc.hooks) - 1; i >= 0; i-- {
			mut = igc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, igc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (igc *ItemGroupCreate) SaveX(ctx context.Context) *ItemGroup {
	v, err := igc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (igc *ItemGroupCreate) check() error {
	if _, ok := igc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (igc *ItemGroupCreate) sqlSave(ctx context.Context) (*ItemGroup, error) {
	_node, _spec := igc.createSpec()
	if err := sqlgraph.CreateNode(ctx, igc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (igc *ItemGroupCreate) createSpec() (*ItemGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &ItemGroup{config: igc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: itemgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemgroup.FieldID,
			},
		}
	)
	if value, ok := igc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: itemgroup.FieldName,
		})
		_node.Name = value
	}
	if nodes := igc.mutation.GroupItemIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ItemGroupCreateBulk is the builder for creating many ItemGroup entities in bulk.
type ItemGroupCreateBulk struct {
	config
	builders []*ItemGroupCreate
}

// Save creates the ItemGroup entities in the database.
func (igcb *ItemGroupCreateBulk) Save(ctx context.Context) ([]*ItemGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(igcb.builders))
	nodes := make([]*ItemGroup, len(igcb.builders))
	mutators := make([]Mutator, len(igcb.builders))
	for i := range igcb.builders {
		func(i int, root context.Context) {
			builder := igcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, igcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, igcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, igcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (igcb *ItemGroupCreateBulk) SaveX(ctx context.Context) []*ItemGroup {
	v, err := igcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}