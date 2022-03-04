// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entrds/ent/band"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Band is the model entity for the Band schema.
type Band struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// YearFormed holds the value of the "year_formed" field.
	YearFormed int `json:"year_formed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BandQuery when eager-loading is set.
	Edges BandEdges `json:"edges"`
}

// BandEdges holds the relations/edges for other nodes in the graph.
type BandEdges struct {
	// Members holds the value of the members edge.
	Members []*Artist `json:"members,omitempty"`
	// Albums holds the value of the albums edge.
	Albums []*Album `json:"albums,omitempty"`
	// AssociatedBands holds the value of the associated_bands edge.
	AssociatedBands []*Band `json:"associated_bands,omitempty"`
	// Label holds the value of the label edge.
	Label []*Label `json:"label,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e BandEdges) MembersOrErr() ([]*Artist, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// AlbumsOrErr returns the Albums value or an error if the edge
// was not loaded in eager-loading.
func (e BandEdges) AlbumsOrErr() ([]*Album, error) {
	if e.loadedTypes[1] {
		return e.Albums, nil
	}
	return nil, &NotLoadedError{edge: "albums"}
}

// AssociatedBandsOrErr returns the AssociatedBands value or an error if the edge
// was not loaded in eager-loading.
func (e BandEdges) AssociatedBandsOrErr() ([]*Band, error) {
	if e.loadedTypes[2] {
		return e.AssociatedBands, nil
	}
	return nil, &NotLoadedError{edge: "associated_bands"}
}

// LabelOrErr returns the Label value or an error if the edge
// was not loaded in eager-loading.
func (e BandEdges) LabelOrErr() ([]*Label, error) {
	if e.loadedTypes[3] {
		return e.Label, nil
	}
	return nil, &NotLoadedError{edge: "label"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Band) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case band.FieldID, band.FieldYearFormed:
			values[i] = new(sql.NullInt64)
		case band.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Band", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Band fields.
func (b *Band) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case band.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case band.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case band.FieldYearFormed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field year_formed", values[i])
			} else if value.Valid {
				b.YearFormed = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMembers queries the "members" edge of the Band entity.
func (b *Band) QueryMembers() *ArtistQuery {
	return (&BandClient{config: b.config}).QueryMembers(b)
}

// QueryAlbums queries the "albums" edge of the Band entity.
func (b *Band) QueryAlbums() *AlbumQuery {
	return (&BandClient{config: b.config}).QueryAlbums(b)
}

// QueryAssociatedBands queries the "associated_bands" edge of the Band entity.
func (b *Band) QueryAssociatedBands() *BandQuery {
	return (&BandClient{config: b.config}).QueryAssociatedBands(b)
}

// QueryLabel queries the "label" edge of the Band entity.
func (b *Band) QueryLabel() *LabelQuery {
	return (&BandClient{config: b.config}).QueryLabel(b)
}

// Update returns a builder for updating this Band.
// Note that you need to call Band.Unwrap() before calling this method if this Band
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Band) Update() *BandUpdateOne {
	return (&BandClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Band entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Band) Unwrap() *Band {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Band is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Band) String() string {
	var builder strings.Builder
	builder.WriteString("Band(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", name=")
	builder.WriteString(b.Name)
	builder.WriteString(", year_formed=")
	builder.WriteString(fmt.Sprintf("%v", b.YearFormed))
	builder.WriteByte(')')
	return builder.String()
}

// Bands is a parsable slice of Band.
type Bands []*Band

func (b Bands) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
