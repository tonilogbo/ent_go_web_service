// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AlbumsColumns holds the columns for the "albums" table.
	AlbumsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "release_year", Type: field.TypeInt},
	}
	// AlbumsTable holds the schema information for the "albums" table.
	AlbumsTable = &schema.Table{
		Name:       "albums",
		Columns:    AlbumsColumns,
		PrimaryKey: []*schema.Column{AlbumsColumns[0]},
	}
	// ArtistsColumns holds the columns for the "artists" table.
	ArtistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "date_of_birth", Type: field.TypeTime, Nullable: true},
		{Name: "gender", Type: field.TypeString, Nullable: true},
	}
	// ArtistsTable holds the schema information for the "artists" table.
	ArtistsTable = &schema.Table{
		Name:       "artists",
		Columns:    ArtistsColumns,
		PrimaryKey: []*schema.Column{ArtistsColumns[0]},
	}
	// BandsColumns holds the columns for the "bands" table.
	BandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "year_formed", Type: field.TypeInt, Nullable: true},
	}
	// BandsTable holds the schema information for the "bands" table.
	BandsTable = &schema.Table{
		Name:       "bands",
		Columns:    BandsColumns,
		PrimaryKey: []*schema.Column{BandsColumns[0]},
	}
	// LabelsColumns holds the columns for the "labels" table.
	LabelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "year_established", Type: field.TypeInt, Nullable: true},
	}
	// LabelsTable holds the schema information for the "labels" table.
	LabelsTable = &schema.Table{
		Name:       "labels",
		Columns:    LabelsColumns,
		PrimaryKey: []*schema.Column{LabelsColumns[0]},
	}
	// ArtistAlbumsColumns holds the columns for the "artist_albums" table.
	ArtistAlbumsColumns = []*schema.Column{
		{Name: "artist_id", Type: field.TypeInt},
		{Name: "album_id", Type: field.TypeInt},
	}
	// ArtistAlbumsTable holds the schema information for the "artist_albums" table.
	ArtistAlbumsTable = &schema.Table{
		Name:       "artist_albums",
		Columns:    ArtistAlbumsColumns,
		PrimaryKey: []*schema.Column{ArtistAlbumsColumns[0], ArtistAlbumsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "artist_albums_artist_id",
				Columns:    []*schema.Column{ArtistAlbumsColumns[0]},
				RefColumns: []*schema.Column{ArtistsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "artist_albums_album_id",
				Columns:    []*schema.Column{ArtistAlbumsColumns[1]},
				RefColumns: []*schema.Column{AlbumsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ArtistAssociatedArtistsColumns holds the columns for the "artist_associated_artists" table.
	ArtistAssociatedArtistsColumns = []*schema.Column{
		{Name: "artist_id", Type: field.TypeInt},
		{Name: "associated_artist_id", Type: field.TypeInt},
	}
	// ArtistAssociatedArtistsTable holds the schema information for the "artist_associated_artists" table.
	ArtistAssociatedArtistsTable = &schema.Table{
		Name:       "artist_associated_artists",
		Columns:    ArtistAssociatedArtistsColumns,
		PrimaryKey: []*schema.Column{ArtistAssociatedArtistsColumns[0], ArtistAssociatedArtistsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "artist_associated_artists_artist_id",
				Columns:    []*schema.Column{ArtistAssociatedArtistsColumns[0]},
				RefColumns: []*schema.Column{ArtistsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "artist_associated_artists_associated_artist_id",
				Columns:    []*schema.Column{ArtistAssociatedArtistsColumns[1]},
				RefColumns: []*schema.Column{ArtistsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// BandMembersColumns holds the columns for the "band_members" table.
	BandMembersColumns = []*schema.Column{
		{Name: "band_id", Type: field.TypeInt},
		{Name: "artist_id", Type: field.TypeInt},
	}
	// BandMembersTable holds the schema information for the "band_members" table.
	BandMembersTable = &schema.Table{
		Name:       "band_members",
		Columns:    BandMembersColumns,
		PrimaryKey: []*schema.Column{BandMembersColumns[0], BandMembersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "band_members_band_id",
				Columns:    []*schema.Column{BandMembersColumns[0]},
				RefColumns: []*schema.Column{BandsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "band_members_artist_id",
				Columns:    []*schema.Column{BandMembersColumns[1]},
				RefColumns: []*schema.Column{ArtistsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// BandAlbumsColumns holds the columns for the "band_albums" table.
	BandAlbumsColumns = []*schema.Column{
		{Name: "band_id", Type: field.TypeInt},
		{Name: "album_id", Type: field.TypeInt},
	}
	// BandAlbumsTable holds the schema information for the "band_albums" table.
	BandAlbumsTable = &schema.Table{
		Name:       "band_albums",
		Columns:    BandAlbumsColumns,
		PrimaryKey: []*schema.Column{BandAlbumsColumns[0], BandAlbumsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "band_albums_band_id",
				Columns:    []*schema.Column{BandAlbumsColumns[0]},
				RefColumns: []*schema.Column{BandsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "band_albums_album_id",
				Columns:    []*schema.Column{BandAlbumsColumns[1]},
				RefColumns: []*schema.Column{AlbumsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// BandAssociatedBandsColumns holds the columns for the "band_associated_bands" table.
	BandAssociatedBandsColumns = []*schema.Column{
		{Name: "band_id", Type: field.TypeInt},
		{Name: "associated_band_id", Type: field.TypeInt},
	}
	// BandAssociatedBandsTable holds the schema information for the "band_associated_bands" table.
	BandAssociatedBandsTable = &schema.Table{
		Name:       "band_associated_bands",
		Columns:    BandAssociatedBandsColumns,
		PrimaryKey: []*schema.Column{BandAssociatedBandsColumns[0], BandAssociatedBandsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "band_associated_bands_band_id",
				Columns:    []*schema.Column{BandAssociatedBandsColumns[0]},
				RefColumns: []*schema.Column{BandsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "band_associated_bands_associated_band_id",
				Columns:    []*schema.Column{BandAssociatedBandsColumns[1]},
				RefColumns: []*schema.Column{BandsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LabelIndividualArtistsColumns holds the columns for the "label_individual_artists" table.
	LabelIndividualArtistsColumns = []*schema.Column{
		{Name: "label_id", Type: field.TypeInt},
		{Name: "artist_id", Type: field.TypeInt},
	}
	// LabelIndividualArtistsTable holds the schema information for the "label_individual_artists" table.
	LabelIndividualArtistsTable = &schema.Table{
		Name:       "label_individual_artists",
		Columns:    LabelIndividualArtistsColumns,
		PrimaryKey: []*schema.Column{LabelIndividualArtistsColumns[0], LabelIndividualArtistsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "label_individual_artists_label_id",
				Columns:    []*schema.Column{LabelIndividualArtistsColumns[0]},
				RefColumns: []*schema.Column{LabelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "label_individual_artists_artist_id",
				Columns:    []*schema.Column{LabelIndividualArtistsColumns[1]},
				RefColumns: []*schema.Column{ArtistsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LabelBandsColumns holds the columns for the "label_bands" table.
	LabelBandsColumns = []*schema.Column{
		{Name: "label_id", Type: field.TypeInt},
		{Name: "band_id", Type: field.TypeInt},
	}
	// LabelBandsTable holds the schema information for the "label_bands" table.
	LabelBandsTable = &schema.Table{
		Name:       "label_bands",
		Columns:    LabelBandsColumns,
		PrimaryKey: []*schema.Column{LabelBandsColumns[0], LabelBandsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "label_bands_label_id",
				Columns:    []*schema.Column{LabelBandsColumns[0]},
				RefColumns: []*schema.Column{LabelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "label_bands_band_id",
				Columns:    []*schema.Column{LabelBandsColumns[1]},
				RefColumns: []*schema.Column{BandsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AlbumsTable,
		ArtistsTable,
		BandsTable,
		LabelsTable,
		ArtistAlbumsTable,
		ArtistAssociatedArtistsTable,
		BandMembersTable,
		BandAlbumsTable,
		BandAssociatedBandsTable,
		LabelIndividualArtistsTable,
		LabelBandsTable,
	}
)

func init() {
	ArtistAlbumsTable.ForeignKeys[0].RefTable = ArtistsTable
	ArtistAlbumsTable.ForeignKeys[1].RefTable = AlbumsTable
	ArtistAssociatedArtistsTable.ForeignKeys[0].RefTable = ArtistsTable
	ArtistAssociatedArtistsTable.ForeignKeys[1].RefTable = ArtistsTable
	BandMembersTable.ForeignKeys[0].RefTable = BandsTable
	BandMembersTable.ForeignKeys[1].RefTable = ArtistsTable
	BandAlbumsTable.ForeignKeys[0].RefTable = BandsTable
	BandAlbumsTable.ForeignKeys[1].RefTable = AlbumsTable
	BandAssociatedBandsTable.ForeignKeys[0].RefTable = BandsTable
	BandAssociatedBandsTable.ForeignKeys[1].RefTable = BandsTable
	LabelIndividualArtistsTable.ForeignKeys[0].RefTable = LabelsTable
	LabelIndividualArtistsTable.ForeignKeys[1].RefTable = ArtistsTable
	LabelBandsTable.ForeignKeys[0].RefTable = LabelsTable
	LabelBandsTable.ForeignKeys[1].RefTable = BandsTable
}