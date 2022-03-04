package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

// Fields of the Artist.
// Artists have a name, date of birth, gender
func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("date_of_birth").Optional(),
		field.String("gender").Optional(),
	}
}

// Edges of the Artist.
// Artists can be a member of bands, can have songs, can have albums, can be associated with other artists or bands
func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bands",Band.Type).Ref("members"),
		edge.To("albums",Album.Type),
		edge.To("associated_artists",Artist.Type),
		edge.From("label",Label.Type).Ref("individual_artists"),
	}
}
