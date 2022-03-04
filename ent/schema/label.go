package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Label holds the schema definition for the Label entity.
type Label struct {
	ent.Schema
}

// Fields of the Label.
func (Label) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("year_established").
			Optional(),
	}
}

// Edges of the Label.
func (Label) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("individual_artists", Artist.Type),
		edge.To("bands", Band.Type),
	}
}
