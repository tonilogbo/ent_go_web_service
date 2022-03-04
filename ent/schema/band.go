package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Band holds the schema definition for the Band entity.
type Band struct {
	ent.Schema
}

// Fields of the Band.
func (Band) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("year_formed").
			Optional(),
	}
}

// Edges of the Band.
func (Band) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", Artist.Type).Required(),
		edge.To("albums", Album.Type),
		edge.To("associated_bands", Band.Type),
		edge.From("label", Label.Type).Ref("bands"),
	}
}
