// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemdescription"
)

// ItemDescriptionCreate is the builder for creating a ItemDescription entity.
type ItemDescriptionCreate struct {
	config
	mutation *ItemDescriptionMutation
	hooks    []Hook
}

// SetDescription sets the "description" field.
func (idc *ItemDescriptionCreate) SetDescription(s string) *ItemDescriptionCreate {
	idc.mutation.SetDescription(s)
	return idc
}

// SetOwnerID sets the "owner" edge to the Item entity by ID.
func (idc *ItemDescriptionCreate) SetOwnerID(id int) *ItemDescriptionCreate {
	idc.mutation.SetOwnerID(id)
	return idc
}

// SetOwner sets the "owner" edge to the Item entity.
func (idc *ItemDescriptionCreate) SetOwner(i *Item) *ItemDescriptionCreate {
	return idc.SetOwnerID(i.ID)
}

// Mutation returns the ItemDescriptionMutation object of the builder.
func (idc *ItemDescriptionCreate) Mutation() *ItemDescriptionMutation {
	return idc.mutation
}

// Save creates the ItemDescription in the database.
func (idc *ItemDescriptionCreate) Save(ctx context.Context) (*ItemDescription, error) {
	var (
		err  error
		node *ItemDescription
	)
	if len(idc.hooks) == 0 {
		if err = idc.check(); err != nil {
			return nil, err
		}
		node, err = idc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = idc.check(); err != nil {
				return nil, err
			}
			idc.mutation = mutation
			node, err = idc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(idc.hooks) - 1; i >= 0; i-- {
			mut = idc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, idc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (idc *ItemDescriptionCreate) SaveX(ctx context.Context) *ItemDescription {
	v, err := idc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (idc *ItemDescriptionCreate) check() error {
	if _, ok := idc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := idc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New("ent: missing required edge \"owner\"")}
	}
	return nil
}

func (idc *ItemDescriptionCreate) sqlSave(ctx context.Context) (*ItemDescription, error) {
	_node, _spec := idc.createSpec()
	if err := sqlgraph.CreateNode(ctx, idc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (idc *ItemDescriptionCreate) createSpec() (*ItemDescription, *sqlgraph.CreateSpec) {
	var (
		_node = &ItemDescription{config: idc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: itemdescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemdescription.FieldID,
			},
		}
	)
	if value, ok := idc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: itemdescription.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := idc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   itemdescription.OwnerTable,
			Columns: []string{itemdescription.OwnerColumn},
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

// ItemDescriptionCreateBulk is the builder for creating many ItemDescription entities in bulk.
type ItemDescriptionCreateBulk struct {
	config
	builders []*ItemDescriptionCreate
}

// Save creates the ItemDescription entities in the database.
func (idcb *ItemDescriptionCreateBulk) Save(ctx context.Context) ([]*ItemDescription, error) {
	specs := make([]*sqlgraph.CreateSpec, len(idcb.builders))
	nodes := make([]*ItemDescription, len(idcb.builders))
	mutators := make([]Mutator, len(idcb.builders))
	for i := range idcb.builders {
		func(i int, root context.Context) {
			builder := idcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemDescriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, idcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, idcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, idcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (idcb *ItemDescriptionCreateBulk) SaveX(ctx context.Context) []*ItemDescription {
	v, err := idcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}