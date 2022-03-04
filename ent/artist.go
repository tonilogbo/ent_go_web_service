// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entrds/ent/artist"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Artist is the model entity for the Artist schema.
type Artist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DateOfBirth holds the value of the "date_of_birth" field.
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtistQuery when eager-loading is set.
	Edges ArtistEdges `json:"edges"`
}

// ArtistEdges holds the relations/edges for other nodes in the graph.
type ArtistEdges struct {
	// Bands holds the value of the bands edge.
	Bands []*Band `json:"bands,omitempty"`
	// Albums holds the value of the albums edge.
	Albums []*Album `json:"albums,omitempty"`
	// AssociatedArtists holds the value of the associated_artists edge.
	AssociatedArtists []*Artist `json:"associated_artists,omitempty"`
	// Label holds the value of the label edge.
	Label []*Label `json:"label,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// BandsOrErr returns the Bands value or an error if the edge
// was not loaded in eager-loading.
func (e ArtistEdges) BandsOrErr() ([]*Band, error) {
	if e.loadedTypes[0] {
		return e.Bands, nil
	}
	return nil, &NotLoadedError{edge: "bands"}
}

// AlbumsOrErr returns the Albums value or an error if the edge
// was not loaded in eager-loading.
func (e ArtistEdges) AlbumsOrErr() ([]*Album, error) {
	if e.loadedTypes[1] {
		return e.Albums, nil
	}
	return nil, &NotLoadedError{edge: "albums"}
}

// AssociatedArtistsOrErr returns the AssociatedArtists value or an error if the edge
// was not loaded in eager-loading.
func (e ArtistEdges) AssociatedArtistsOrErr() ([]*Artist, error) {
	if e.loadedTypes[2] {
		return e.AssociatedArtists, nil
	}
	return nil, &NotLoadedError{edge: "associated_artists"}
}

// LabelOrErr returns the Label value or an error if the edge
// was not loaded in eager-loading.
func (e ArtistEdges) LabelOrErr() ([]*Label, error) {
	if e.loadedTypes[3] {
		return e.Label, nil
	}
	return nil, &NotLoadedError{edge: "label"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Artist) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case artist.FieldID:
			values[i] = new(sql.NullInt64)
		case artist.FieldName, artist.FieldGender:
			values[i] = new(sql.NullString)
		case artist.FieldDateOfBirth:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Artist", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Artist fields.
func (a *Artist) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case artist.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case artist.FieldDateOfBirth:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_of_birth", values[i])
			} else if value.Valid {
				a.DateOfBirth = value.Time
			}
		case artist.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				a.Gender = value.String
			}
		}
	}
	return nil
}

// QueryBands queries the "bands" edge of the Artist entity.
func (a *Artist) QueryBands() *BandQuery {
	return (&ArtistClient{config: a.config}).QueryBands(a)
}

// QueryAlbums queries the "albums" edge of the Artist entity.
func (a *Artist) QueryAlbums() *AlbumQuery {
	return (&ArtistClient{config: a.config}).QueryAlbums(a)
}

// QueryAssociatedArtists queries the "associated_artists" edge of the Artist entity.
func (a *Artist) QueryAssociatedArtists() *ArtistQuery {
	return (&ArtistClient{config: a.config}).QueryAssociatedArtists(a)
}

// QueryLabel queries the "label" edge of the Artist entity.
func (a *Artist) QueryLabel() *LabelQuery {
	return (&ArtistClient{config: a.config}).QueryLabel(a)
}

// Update returns a builder for updating this Artist.
// Note that you need to call Artist.Unwrap() before calling this method if this Artist
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Artist) Update() *ArtistUpdateOne {
	return (&ArtistClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Artist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Artist) Unwrap() *Artist {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Artist is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Artist) String() string {
	var builder strings.Builder
	builder.WriteString("Artist(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", date_of_birth=")
	builder.WriteString(a.DateOfBirth.Format(time.ANSIC))
	builder.WriteString(", gender=")
	builder.WriteString(a.Gender)
	builder.WriteByte(')')
	return builder.String()
}

// Artists is a parsable slice of Artist.
type Artists []*Artist

func (a Artists) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
