package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ItemGroup holds the schema definition for the ItemGroup entity.
type ItemGroup struct {
	ent.Schema
}

// Fields of the ItemGroup.
func (ItemGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the ItemGroup.
func (ItemGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group_item", Item.Type).
			Ref("item_group"),
	}
}
