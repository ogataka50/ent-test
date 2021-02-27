package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ItemDescription holds the schema definition for the ItemDescription entity.
type ItemDescription struct {
	ent.Schema
}

// Fields of the ItemDescription.
func (ItemDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("description"),
	}
}

// Edges of the ItemDescription.
func (ItemDescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Item.Type).
			Ref("item_description").
			Unique().
			Required(),
	}
}
