package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ItemVariation holds the schema definition for the ItemVariation entity.
type ItemVariation struct {
	ent.Schema
}

// Fields of the ItemVariation.
func (ItemVariation) Fields() []ent.Field {
	return []ent.Field{
		field.String("variant_name"),
	}
}

// Edges of the ItemVariation.
func (ItemVariation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("origin_item", Item.Type).
			Ref("item_variation").
			Required().Unique(),
	}
}
