// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"entrds/ent/album"
	"entrds/ent/artist"
	"entrds/ent/band"
	"entrds/ent/label"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BandCreate is the builder for creating a Band entity.
type BandCreate struct {
	config
	mutation *BandMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bc *BandCreate) SetName(s string) *BandCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetYearFormed sets the "year_formed" field.
func (bc *BandCreate) SetYearFormed(i int) *BandCreate {
	bc.mutation.SetYearFormed(i)
	return bc
}

// SetNillableYearFormed sets the "year_formed" field if the given value is not nil.
func (bc *BandCreate) SetNillableYearFormed(i *int) *BandCreate {
	if i != nil {
		bc.SetYearFormed(*i)
	}
	return bc
}

// AddMemberIDs adds the "members" edge to the Artist entity by IDs.
func (bc *BandCreate) AddMemberIDs(ids ...int) *BandCreate {
	bc.mutation.AddMemberIDs(ids...)
	return bc
}

// AddMembers adds the "members" edges to the Artist entity.
func (bc *BandCreate) AddMembers(a ...*Artist) *BandCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bc.AddMemberIDs(ids...)
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (bc *BandCreate) AddAlbumIDs(ids ...int) *BandCreate {
	bc.mutation.AddAlbumIDs(ids...)
	return bc
}

// AddAlbums adds the "albums" edges to the Album entity.
func (bc *BandCreate) AddAlbums(a ...*Album) *BandCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bc.AddAlbumIDs(ids...)
}

// AddAssociatedBandIDs adds the "associated_bands" edge to the Band entity by IDs.
func (bc *BandCreate) AddAssociatedBandIDs(ids ...int) *BandCreate {
	bc.mutation.AddAssociatedBandIDs(ids...)
	return bc
}

// AddAssociatedBands adds the "associated_bands" edges to the Band entity.
func (bc *BandCreate) AddAssociatedBands(b ...*Band) *BandCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddAssociatedBandIDs(ids...)
}

// AddLabelIDs adds the "label" edge to the Label entity by IDs.
func (bc *BandCreate) AddLabelIDs(ids ...int) *BandCreate {
	bc.mutation.AddLabelIDs(ids...)
	return bc
}

// AddLabel adds the "label" edges to the Label entity.
func (bc *BandCreate) AddLabel(l ...*Label) *BandCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return bc.AddLabelIDs(ids...)
}

// Mutation returns the BandMutation object of the builder.
func (bc *BandCreate) Mutation() *BandMutation {
	return bc.mutation
}

// Save creates the Band in the database.
func (bc *BandCreate) Save(ctx context.Context) (*Band, error) {
	var (
		err  error
		node *Band
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BandCreate) SaveX(ctx context.Context) *Band {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BandCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BandCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BandCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Band.name"`)}
	}
	if len(bc.mutation.MembersIDs()) == 0 {
		return &ValidationError{Name: "members", err: errors.New(`ent: missing required edge "Band.members"`)}
	}
	return nil
}

func (bc *BandCreate) sqlSave(ctx context.Context) (*Band, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BandCreate) createSpec() (*Band, *sqlgraph.CreateSpec) {
	var (
		_node = &Band{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: band.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: band.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: band.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bc.mutation.YearFormed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: band.FieldYearFormed,
		})
		_node.YearFormed = value
	}
	if nodes := bc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   band.MembersTable,
			Columns: band.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   band.AlbumsTable,
			Columns: band.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.AssociatedBandsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   band.AssociatedBandsTable,
			Columns: band.AssociatedBandsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: band.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.LabelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   band.LabelTable,
			Columns: band.LabelPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: label.FieldID,
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

// BandCreateBulk is the builder for creating many Band entities in bulk.
type BandCreateBulk struct {
	config
	builders []*BandCreate
}

// Save creates the Band entities in the database.
func (bcb *BandCreateBulk) Save(ctx context.Context) ([]*Band, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Band, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BandMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BandCreateBulk) SaveX(ctx context.Context) []*Band {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BandCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BandCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
