// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemdescription"
	"github.com/ogataka50/ent-test/ent/predicate"
)

// ItemDescriptionQuery is the builder for querying ItemDescription entities.
type ItemDescriptionQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.ItemDescription
	// eager-loading edges.
	withOwner *ItemQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemDescriptionQuery builder.
func (idq *ItemDescriptionQuery) Where(ps ...predicate.ItemDescription) *ItemDescriptionQuery {
	idq.predicates = append(idq.predicates, ps...)
	return idq
}

// Limit adds a limit step to the query.
func (idq *ItemDescriptionQuery) Limit(limit int) *ItemDescriptionQuery {
	idq.limit = &limit
	return idq
}

// Offset adds an offset step to the query.
func (idq *ItemDescriptionQuery) Offset(offset int) *ItemDescriptionQuery {
	idq.offset = &offset
	return idq
}

// Order adds an order step to the query.
func (idq *ItemDescriptionQuery) Order(o ...OrderFunc) *ItemDescriptionQuery {
	idq.order = append(idq.order, o...)
	return idq
}

// QueryOwner chains the current query on the "owner" edge.
func (idq *ItemDescriptionQuery) QueryOwner() *ItemQuery {
	query := &ItemQuery{config: idq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := idq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := idq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemdescription.Table, itemdescription.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, itemdescription.OwnerTable, itemdescription.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(idq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ItemDescription entity from the query.
// Returns a *NotFoundError when no ItemDescription was found.
func (idq *ItemDescriptionQuery) First(ctx context.Context) (*ItemDescription, error) {
	nodes, err := idq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{itemdescription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (idq *ItemDescriptionQuery) FirstX(ctx context.Context) *ItemDescription {
	node, err := idq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ItemDescription ID from the query.
// Returns a *NotFoundError when no ItemDescription ID was found.
func (idq *ItemDescriptionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = idq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{itemdescription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (idq *ItemDescriptionQuery) FirstIDX(ctx context.Context) int {
	id, err := idq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ItemDescription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ItemDescription entity is not found.
// Returns a *NotFoundError when no ItemDescription entities are found.
func (idq *ItemDescriptionQuery) Only(ctx context.Context) (*ItemDescription, error) {
	nodes, err := idq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{itemdescription.Label}
	default:
		return nil, &NotSingularError{itemdescription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (idq *ItemDescriptionQuery) OnlyX(ctx context.Context) *ItemDescription {
	node, err := idq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ItemDescription ID in the query.
// Returns a *NotSingularError when exactly one ItemDescription ID is not found.
// Returns a *NotFoundError when no entities are found.
func (idq *ItemDescriptionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = idq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = &NotSingularError{itemdescription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (idq *ItemDescriptionQuery) OnlyIDX(ctx context.Context) int {
	id, err := idq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ItemDescriptions.
func (idq *ItemDescriptionQuery) All(ctx context.Context) ([]*ItemDescription, error) {
	if err := idq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return idq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (idq *ItemDescriptionQuery) AllX(ctx context.Context) []*ItemDescription {
	nodes, err := idq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ItemDescription IDs.
func (idq *ItemDescriptionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := idq.Select(itemdescription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (idq *ItemDescriptionQuery) IDsX(ctx context.Context) []int {
	ids, err := idq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (idq *ItemDescriptionQuery) Count(ctx context.Context) (int, error) {
	if err := idq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return idq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (idq *ItemDescriptionQuery) CountX(ctx context.Context) int {
	count, err := idq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (idq *ItemDescriptionQuery) Exist(ctx context.Context) (bool, error) {
	if err := idq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return idq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (idq *ItemDescriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := idq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemDescriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (idq *ItemDescriptionQuery) Clone() *ItemDescriptionQuery {
	if idq == nil {
		return nil
	}
	return &ItemDescriptionQuery{
		config:     idq.config,
		limit:      idq.limit,
		offset:     idq.offset,
		order:      append([]OrderFunc{}, idq.order...),
		predicates: append([]predicate.ItemDescription{}, idq.predicates...),
		withOwner:  idq.withOwner.Clone(),
		// clone intermediate query.
		sql:  idq.sql.Clone(),
		path: idq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (idq *ItemDescriptionQuery) WithOwner(opts ...func(*ItemQuery)) *ItemDescriptionQuery {
	query := &ItemQuery{config: idq.config}
	for _, opt := range opts {
		opt(query)
	}
	idq.withOwner = query
	return idq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ItemDescription.Query().
//		GroupBy(itemdescription.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (idq *ItemDescriptionQuery) GroupBy(field string, fields ...string) *ItemDescriptionGroupBy {
	group := &ItemDescriptionGroupBy{config: idq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := idq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return idq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.ItemDescription.Query().
//		Select(itemdescription.FieldDescription).
//		Scan(ctx, &v)
//
func (idq *ItemDescriptionQuery) Select(field string, fields ...string) *ItemDescriptionSelect {
	idq.fields = append([]string{field}, fields...)
	return &ItemDescriptionSelect{ItemDescriptionQuery: idq}
}

func (idq *ItemDescriptionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range idq.fields {
		if !itemdescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if idq.path != nil {
		prev, err := idq.path(ctx)
		if err != nil {
			return err
		}
		idq.sql = prev
	}
	return nil
}

func (idq *ItemDescriptionQuery) sqlAll(ctx context.Context) ([]*ItemDescription, error) {
	var (
		nodes       = []*ItemDescription{}
		withFKs     = idq.withFKs
		_spec       = idq.querySpec()
		loadedTypes = [1]bool{
			idq.withOwner != nil,
		}
	)
	if idq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, itemdescription.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ItemDescription{config: idq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, idq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := idq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ItemDescription)
		for i := range nodes {
			if fk := nodes[i].item_item_description; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(item.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "item_item_description" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	return nodes, nil
}

func (idq *ItemDescriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := idq.querySpec()
	return sqlgraph.CountNodes(ctx, idq.driver, _spec)
}

func (idq *ItemDescriptionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := idq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (idq *ItemDescriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemdescription.Table,
			Columns: itemdescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemdescription.FieldID,
			},
		},
		From:   idq.sql,
		Unique: true,
	}
	if fields := idq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itemdescription.FieldID)
		for i := range fields {
			if fields[i] != itemdescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := idq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := idq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := idq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := idq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, itemdescription.ValidColumn)
			}
		}
	}
	return _spec
}

func (idq *ItemDescriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(idq.driver.Dialect())
	t1 := builder.Table(itemdescription.Table)
	selector := builder.Select(t1.Columns(itemdescription.Columns...)...).From(t1)
	if idq.sql != nil {
		selector = idq.sql
		selector.Select(selector.Columns(itemdescription.Columns...)...)
	}
	for _, p := range idq.predicates {
		p(selector)
	}
	for _, p := range idq.order {
		p(selector, itemdescription.ValidColumn)
	}
	if offset := idq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := idq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ItemDescriptionGroupBy is the group-by builder for ItemDescription entities.
type ItemDescriptionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (idgb *ItemDescriptionGroupBy) Aggregate(fns ...AggregateFunc) *ItemDescriptionGroupBy {
	idgb.fns = append(idgb.fns, fns...)
	return idgb
}

// Scan applies the group-by query and scans the result into the given value.
func (idgb *ItemDescriptionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := idgb.path(ctx)
	if err != nil {
		return err
	}
	idgb.sql = query
	return idgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := idgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (idgb *ItemDescriptionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(idgb.fields) > 1 {
		return nil, errors.New("ent: ItemDescriptionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := idgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) StringsX(ctx context.Context) []string {
	v, err := idgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (idgb *ItemDescriptionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = idgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = fmt.Errorf("ent: ItemDescriptionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) StringX(ctx context.Context) string {
	v, err := idgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (idgb *ItemDescriptionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(idgb.fields) > 1 {
		return nil, errors.New("ent: ItemDescriptionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := idgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) IntsX(ctx context.Context) []int {
	v, err := idgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (idgb *ItemDescriptionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = idgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = fmt.Errorf("ent: ItemDescriptionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) IntX(ctx context.Context) int {
	v, err := idgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (idgb *ItemDescriptionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(idgb.fields) > 1 {
		return nil, errors.New("ent: ItemDescriptionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := idgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := idgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (idgb *ItemDescriptionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = idgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = fmt.Errorf("ent: ItemDescriptionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := idgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (idgb *ItemDescriptionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(idgb.fields) > 1 {
		return nil, errors.New("ent: ItemDescriptionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := idgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := idgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (idgb *ItemDescriptionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = idgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = fmt.Errorf("ent: ItemDescriptionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (idgb *ItemDescriptionGroupBy) BoolX(ctx context.Context) bool {
	v, err := idgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (idgb *ItemDescriptionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range idgb.fields {
		if !itemdescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := idgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := idgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (idgb *ItemDescriptionGroupBy) sqlQuery() *sql.Selector {
	selector := idgb.sql
	columns := make([]string, 0, len(idgb.fields)+len(idgb.fns))
	columns = append(columns, idgb.fields...)
	for _, fn := range idgb.fns {
		columns = append(columns, fn(selector, itemdescription.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(idgb.fields...)
}

// ItemDescriptionSelect is the builder for selecting fields of ItemDescription entities.
type ItemDescriptionSelect struct {
	*ItemDescriptionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ids *ItemDescriptionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ids.prepareQuery(ctx); err != nil {
		return err
	}
	ids.sql = ids.ItemDescriptionQuery.sqlQuery(ctx)
	return ids.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ids *ItemDescriptionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ids.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ids *ItemDescriptionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ids.fields) > 1 {
		return nil, errors.New("ent: ItemDescriptionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ids *ItemDescriptionSelect) StringsX(ctx context.Context) []string {
	v, err := ids.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ids *ItemDescriptionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ids.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = fmt.Errorf("ent: ItemDescriptionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ids *ItemDescriptionSelect) StringX(ctx context.Context) string {
	v, err := ids.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ids *ItemDescriptionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ids.fields) > 1 {
		return nil, errors.New("ent: ItemDescriptionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ids *ItemDescriptionSelect) IntsX(ctx context.Context) []int {
	v, err := ids.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ids *ItemDescriptionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ids.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = fmt.Errorf("ent: ItemDescriptionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ids *ItemDescriptionSelect) IntX(ctx context.Context) int {
	v, err := ids.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ids *ItemDescriptionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ids.fields) > 1 {
		return nil, errors.New("ent: ItemDescriptionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ids *ItemDescriptionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ids.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ids *ItemDescriptionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ids.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = fmt.Errorf("ent: ItemDescriptionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ids *ItemDescriptionSelect) Float64X(ctx context.Context) float64 {
	v, err := ids.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ids *ItemDescriptionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ids.fields) > 1 {
		return nil, errors.New("ent: ItemDescriptionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ids *ItemDescriptionSelect) BoolsX(ctx context.Context) []bool {
	v, err := ids.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ids *ItemDescriptionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ids.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{itemdescription.Label}
	default:
		err = fmt.Errorf("ent: ItemDescriptionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ids *ItemDescriptionSelect) BoolX(ctx context.Context) bool {
	v, err := ids.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ids *ItemDescriptionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ids.sqlQuery().Query()
	if err := ids.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ids *ItemDescriptionSelect) sqlQuery() sql.Querier {
	selector := ids.sql
	selector.Select(selector.Columns(ids.fields...)...)
	return selector
}
