// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"entrds/ent/artist"
	"entrds/ent/band"
	"entrds/ent/label"
	"entrds/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LabelUpdate is the builder for updating Label entities.
type LabelUpdate struct {
	config
	hooks    []Hook
	mutation *LabelMutation
}

// Where appends a list predicates to the LabelUpdate builder.
func (lu *LabelUpdate) Where(ps ...predicate.Label) *LabelUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "name" field.
func (lu *LabelUpdate) SetName(s string) *LabelUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetYearEstablished sets the "year_established" field.
func (lu *LabelUpdate) SetYearEstablished(i int) *LabelUpdate {
	lu.mutation.ResetYearEstablished()
	lu.mutation.SetYearEstablished(i)
	return lu
}

// SetNillableYearEstablished sets the "year_established" field if the given value is not nil.
func (lu *LabelUpdate) SetNillableYearEstablished(i *int) *LabelUpdate {
	if i != nil {
		lu.SetYearEstablished(*i)
	}
	return lu
}

// AddYearEstablished adds i to the "year_established" field.
func (lu *LabelUpdate) AddYearEstablished(i int) *LabelUpdate {
	lu.mutation.AddYearEstablished(i)
	return lu
}

// ClearYearEstablished clears the value of the "year_established" field.
func (lu *LabelUpdate) ClearYearEstablished() *LabelUpdate {
	lu.mutation.ClearYearEstablished()
	return lu
}

// AddIndividualArtistIDs adds the "individual_artists" edge to the Artist entity by IDs.
func (lu *LabelUpdate) AddIndividualArtistIDs(ids ...int) *LabelUpdate {
	lu.mutation.AddIndividualArtistIDs(ids...)
	return lu
}

// AddIndividualArtists adds the "individual_artists" edges to the Artist entity.
func (lu *LabelUpdate) AddIndividualArtists(a ...*Artist) *LabelUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.AddIndividualArtistIDs(ids...)
}

// AddBandIDs adds the "bands" edge to the Band entity by IDs.
func (lu *LabelUpdate) AddBandIDs(ids ...int) *LabelUpdate {
	lu.mutation.AddBandIDs(ids...)
	return lu
}

// AddBands adds the "bands" edges to the Band entity.
func (lu *LabelUpdate) AddBands(b ...*Band) *LabelUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return lu.AddBandIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (lu *LabelUpdate) Mutation() *LabelMutation {
	return lu.mutation
}

// ClearIndividualArtists clears all "individual_artists" edges to the Artist entity.
func (lu *LabelUpdate) ClearIndividualArtists() *LabelUpdate {
	lu.mutation.ClearIndividualArtists()
	return lu
}

// RemoveIndividualArtistIDs removes the "individual_artists" edge to Artist entities by IDs.
func (lu *LabelUpdate) RemoveIndividualArtistIDs(ids ...int) *LabelUpdate {
	lu.mutation.RemoveIndividualArtistIDs(ids...)
	return lu
}

// RemoveIndividualArtists removes "individual_artists" edges to Artist entities.
func (lu *LabelUpdate) RemoveIndividualArtists(a ...*Artist) *LabelUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.RemoveIndividualArtistIDs(ids...)
}

// ClearBands clears all "bands" edges to the Band entity.
func (lu *LabelUpdate) ClearBands() *LabelUpdate {
	lu.mutation.ClearBands()
	return lu
}

// RemoveBandIDs removes the "bands" edge to Band entities by IDs.
func (lu *LabelUpdate) RemoveBandIDs(ids ...int) *LabelUpdate {
	lu.mutation.RemoveBandIDs(ids...)
	return lu
}

// RemoveBands removes "bands" edges to Band entities.
func (lu *LabelUpdate) RemoveBands(b ...*Band) *LabelUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return lu.RemoveBandIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LabelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LabelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LabelUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LabelUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LabelUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LabelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   label.Table,
			Columns: label.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: label.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: label.FieldName,
		})
	}
	if value, ok := lu.mutation.YearEstablished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: label.FieldYearEstablished,
		})
	}
	if value, ok := lu.mutation.AddedYearEstablished(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: label.FieldYearEstablished,
		})
	}
	if lu.mutation.YearEstablishedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: label.FieldYearEstablished,
		})
	}
	if lu.mutation.IndividualArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.IndividualArtistsTable,
			Columns: label.IndividualArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedIndividualArtistsIDs(); len(nodes) > 0 && !lu.mutation.IndividualArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.IndividualArtistsTable,
			Columns: label.IndividualArtistsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.IndividualArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.IndividualArtistsTable,
			Columns: label.IndividualArtistsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.BandsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.BandsTable,
			Columns: label.BandsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: band.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedBandsIDs(); len(nodes) > 0 && !lu.mutation.BandsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.BandsTable,
			Columns: label.BandsPrimaryKey,
			Bidi:    false,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.BandsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.BandsTable,
			Columns: label.BandsPrimaryKey,
			Bidi:    false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{label.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LabelUpdateOne is the builder for updating a single Label entity.
type LabelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LabelMutation
}

// SetName sets the "name" field.
func (luo *LabelUpdateOne) SetName(s string) *LabelUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetYearEstablished sets the "year_established" field.
func (luo *LabelUpdateOne) SetYearEstablished(i int) *LabelUpdateOne {
	luo.mutation.ResetYearEstablished()
	luo.mutation.SetYearEstablished(i)
	return luo
}

// SetNillableYearEstablished sets the "year_established" field if the given value is not nil.
func (luo *LabelUpdateOne) SetNillableYearEstablished(i *int) *LabelUpdateOne {
	if i != nil {
		luo.SetYearEstablished(*i)
	}
	return luo
}

// AddYearEstablished adds i to the "year_established" field.
func (luo *LabelUpdateOne) AddYearEstablished(i int) *LabelUpdateOne {
	luo.mutation.AddYearEstablished(i)
	return luo
}

// ClearYearEstablished clears the value of the "year_established" field.
func (luo *LabelUpdateOne) ClearYearEstablished() *LabelUpdateOne {
	luo.mutation.ClearYearEstablished()
	return luo
}

// AddIndividualArtistIDs adds the "individual_artists" edge to the Artist entity by IDs.
func (luo *LabelUpdateOne) AddIndividualArtistIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.AddIndividualArtistIDs(ids...)
	return luo
}

// AddIndividualArtists adds the "individual_artists" edges to the Artist entity.
func (luo *LabelUpdateOne) AddIndividualArtists(a ...*Artist) *LabelUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.AddIndividualArtistIDs(ids...)
}

// AddBandIDs adds the "bands" edge to the Band entity by IDs.
func (luo *LabelUpdateOne) AddBandIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.AddBandIDs(ids...)
	return luo
}

// AddBands adds the "bands" edges to the Band entity.
func (luo *LabelUpdateOne) AddBands(b ...*Band) *LabelUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return luo.AddBandIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (luo *LabelUpdateOne) Mutation() *LabelMutation {
	return luo.mutation
}

// ClearIndividualArtists clears all "individual_artists" edges to the Artist entity.
func (luo *LabelUpdateOne) ClearIndividualArtists() *LabelUpdateOne {
	luo.mutation.ClearIndividualArtists()
	return luo
}

// RemoveIndividualArtistIDs removes the "individual_artists" edge to Artist entities by IDs.
func (luo *LabelUpdateOne) RemoveIndividualArtistIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.RemoveIndividualArtistIDs(ids...)
	return luo
}

// RemoveIndividualArtists removes "individual_artists" edges to Artist entities.
func (luo *LabelUpdateOne) RemoveIndividualArtists(a ...*Artist) *LabelUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.RemoveIndividualArtistIDs(ids...)
}

// ClearBands clears all "bands" edges to the Band entity.
func (luo *LabelUpdateOne) ClearBands() *LabelUpdateOne {
	luo.mutation.ClearBands()
	return luo
}

// RemoveBandIDs removes the "bands" edge to Band entities by IDs.
func (luo *LabelUpdateOne) RemoveBandIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.RemoveBandIDs(ids...)
	return luo
}

// RemoveBands removes "bands" edges to Band entities.
func (luo *LabelUpdateOne) RemoveBands(b ...*Band) *LabelUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return luo.RemoveBandIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LabelUpdateOne) Select(field string, fields ...string) *LabelUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Label entity.
func (luo *LabelUpdateOne) Save(ctx context.Context) (*Label, error) {
	var (
		err  error
		node *Label
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LabelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LabelUpdateOne) SaveX(ctx context.Context) *Label {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LabelUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LabelUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LabelUpdateOne) sqlSave(ctx context.Context) (_node *Label, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   label.Table,
			Columns: label.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: label.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Label.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, label.FieldID)
		for _, f := range fields {
			if !label.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != label.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: label.FieldName,
		})
	}
	if value, ok := luo.mutation.YearEstablished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: label.FieldYearEstablished,
		})
	}
	if value, ok := luo.mutation.AddedYearEstablished(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: label.FieldYearEstablished,
		})
	}
	if luo.mutation.YearEstablishedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: label.FieldYearEstablished,
		})
	}
	if luo.mutation.IndividualArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.IndividualArtistsTable,
			Columns: label.IndividualArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedIndividualArtistsIDs(); len(nodes) > 0 && !luo.mutation.IndividualArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.IndividualArtistsTable,
			Columns: label.IndividualArtistsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.IndividualArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.IndividualArtistsTable,
			Columns: label.IndividualArtistsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.BandsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.BandsTable,
			Columns: label.BandsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: band.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedBandsIDs(); len(nodes) > 0 && !luo.mutation.BandsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.BandsTable,
			Columns: label.BandsPrimaryKey,
			Bidi:    false,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.BandsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.BandsTable,
			Columns: label.BandsPrimaryKey,
			Bidi:    false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Label{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{label.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
