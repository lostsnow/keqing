// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lostsnow/keqing/pkg/entity/gameaccount"
	"github.com/lostsnow/keqing/pkg/entity/predicate"
)

// GameAccountQuery is the builder for querying GameAccount entities.
type GameAccountQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.GameAccount
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GameAccountQuery builder.
func (gaq *GameAccountQuery) Where(ps ...predicate.GameAccount) *GameAccountQuery {
	gaq.predicates = append(gaq.predicates, ps...)
	return gaq
}

// Limit the number of records to be returned by this query.
func (gaq *GameAccountQuery) Limit(limit int) *GameAccountQuery {
	gaq.ctx.Limit = &limit
	return gaq
}

// Offset to start from.
func (gaq *GameAccountQuery) Offset(offset int) *GameAccountQuery {
	gaq.ctx.Offset = &offset
	return gaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gaq *GameAccountQuery) Unique(unique bool) *GameAccountQuery {
	gaq.ctx.Unique = &unique
	return gaq
}

// Order specifies how the records should be ordered.
func (gaq *GameAccountQuery) Order(o ...OrderFunc) *GameAccountQuery {
	gaq.order = append(gaq.order, o...)
	return gaq
}

// First returns the first GameAccount entity from the query.
// Returns a *NotFoundError when no GameAccount was found.
func (gaq *GameAccountQuery) First(ctx context.Context) (*GameAccount, error) {
	nodes, err := gaq.Limit(1).All(setContextOp(ctx, gaq.ctx, "First"))
	if err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nil, &NotFoundError{gameaccount.Label}
	}

	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gaq *GameAccountQuery) FirstX(ctx context.Context) *GameAccount {
	node, err := gaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}

	return node
}

// FirstID returns the first GameAccount ID from the query.
// Returns a *NotFoundError when no GameAccount ID was found.
func (gaq *GameAccountQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64

	if ids, err = gaq.Limit(1).IDs(setContextOp(ctx, gaq.ctx, "FirstID")); err != nil {
		return
	}

	if len(ids) == 0 {
		err = &NotFoundError{gameaccount.Label}
		return
	}

	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gaq *GameAccountQuery) FirstIDX(ctx context.Context) int64 {
	id, err := gaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}

	return id
}

// Only returns a single GameAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GameAccount entity is found.
// Returns a *NotFoundError when no GameAccount entities are found.
func (gaq *GameAccountQuery) Only(ctx context.Context) (*GameAccount, error) {
	nodes, err := gaq.Limit(2).All(setContextOp(ctx, gaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}

	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gameaccount.Label}
	default:
		return nil, &NotSingularError{gameaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gaq *GameAccountQuery) OnlyX(ctx context.Context) *GameAccount {
	node, err := gaq.Only(ctx)
	if err != nil {
		panic(err)
	}

	return node
}

// OnlyID is like Only, but returns the only GameAccount ID in the query.
// Returns a *NotSingularError when more than one GameAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (gaq *GameAccountQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64

	if ids, err = gaq.Limit(2).IDs(setContextOp(ctx, gaq.ctx, "OnlyID")); err != nil {
		return
	}

	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = &NotSingularError{gameaccount.Label}
	}

	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gaq *GameAccountQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := gaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}

	return id
}

// All executes the query and returns a list of GameAccounts.
func (gaq *GameAccountQuery) All(ctx context.Context) ([]*GameAccount, error) {
	ctx = setContextOp(ctx, gaq.ctx, "All")
	if err := gaq.prepareQuery(ctx); err != nil {
		return nil, err
	}

	qr := querierAll[[]*GameAccount, *GameAccountQuery]()

	return withInterceptors[[]*GameAccount](ctx, gaq, qr, gaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gaq *GameAccountQuery) AllX(ctx context.Context) []*GameAccount {
	nodes, err := gaq.All(ctx)
	if err != nil {
		panic(err)
	}

	return nodes
}

// IDs executes the query and returns a list of GameAccount IDs.
func (gaq *GameAccountQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if gaq.ctx.Unique == nil && gaq.path != nil {
		gaq.Unique(true)
	}

	ctx = setContextOp(ctx, gaq.ctx, "IDs")
	if err = gaq.Select(gameaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gaq *GameAccountQuery) IDsX(ctx context.Context) []int64 {
	ids, err := gaq.IDs(ctx)
	if err != nil {
		panic(err)
	}

	return ids
}

// Count returns the count of the given query.
func (gaq *GameAccountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gaq.ctx, "Count")
	if err := gaq.prepareQuery(ctx); err != nil {
		return 0, err
	}

	return withInterceptors[int](ctx, gaq, querierCount[*GameAccountQuery](), gaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gaq *GameAccountQuery) CountX(ctx context.Context) int {
	count, err := gaq.Count(ctx)
	if err != nil {
		panic(err)
	}

	return count
}

// Exist returns true if the query has elements in the graph.
func (gaq *GameAccountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gaq.ctx, "Exist")

	switch _, err := gaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entity: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gaq *GameAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := gaq.Exist(ctx)
	if err != nil {
		panic(err)
	}

	return exist
}

// Clone returns a duplicate of the GameAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gaq *GameAccountQuery) Clone() *GameAccountQuery {
	if gaq == nil {
		return nil
	}

	return &GameAccountQuery{
		config:     gaq.config,
		ctx:        gaq.ctx.Clone(),
		order:      append([]OrderFunc{}, gaq.order...),
		inters:     append([]Interceptor{}, gaq.inters...),
		predicates: append([]predicate.GameAccount{}, gaq.predicates...),
		// clone intermediate query.
		sql:  gaq.sql.Clone(),
		path: gaq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GameAccount.Query().
//		GroupBy(gameaccount.FieldUserID).
//		Aggregate(entity.Count()).
//		Scan(ctx, &v)
func (gaq *GameAccountQuery) GroupBy(field string, fields ...string) *GameAccountGroupBy {
	gaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GameAccountGroupBy{build: gaq}
	grbuild.flds = &gaq.ctx.Fields
	grbuild.label = gameaccount.Label
	grbuild.scan = grbuild.Scan

	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//	}
//
//	client.GameAccount.Query().
//		Select(gameaccount.FieldUserID).
//		Scan(ctx, &v)
func (gaq *GameAccountQuery) Select(fields ...string) *GameAccountSelect {
	gaq.ctx.Fields = append(gaq.ctx.Fields, fields...)
	sbuild := &GameAccountSelect{GameAccountQuery: gaq}
	sbuild.label = gameaccount.Label
	sbuild.flds, sbuild.scan = &gaq.ctx.Fields, sbuild.Scan

	return sbuild
}

// Aggregate returns a GameAccountSelect configured with the given aggregations.
func (gaq *GameAccountQuery) Aggregate(fns ...AggregateFunc) *GameAccountSelect {
	return gaq.Select().Aggregate(fns...)
}

func (gaq *GameAccountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gaq.inters {
		if inter == nil {
			return fmt.Errorf("entity: uninitialized interceptor (forgotten import entity/runtime?)")
		}

		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gaq); err != nil {
				return err
			}
		}
	}

	for _, f := range gaq.ctx.Fields {
		if !gameaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entity: invalid field %q for query", f)}
		}
	}

	if gaq.path != nil {
		prev, err := gaq.path(ctx)
		if err != nil {
			return err
		}

		gaq.sql = prev
	}

	return nil
}

func (gaq *GameAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GameAccount, error) {
	var (
		nodes = []*GameAccount{}
		_spec = gaq.querySpec()
	)

	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GameAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GameAccount{config: gaq.config}
		nodes = append(nodes, node)

		return node.assignValues(columns, values)
	}

	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}

	for i := range hooks {
		hooks[i](ctx, _spec)
	}

	if err := sqlgraph.QueryNodes(ctx, gaq.driver, _spec); err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nodes, nil
	}

	return nodes, nil
}

func (gaq *GameAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gaq.querySpec()
	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}

	_spec.Node.Columns = gaq.ctx.Fields
	if len(gaq.ctx.Fields) > 0 {
		_spec.Unique = gaq.ctx.Unique != nil && *gaq.ctx.Unique
	}

	return sqlgraph.CountNodes(ctx, gaq.driver, _spec)
}

func (gaq *GameAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(gameaccount.Table, gameaccount.Columns, sqlgraph.NewFieldSpec(gameaccount.FieldID, field.TypeInt64))
	_spec.From = gaq.sql

	if unique := gaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gaq.path != nil {
		_spec.Unique = true
	}

	if fields := gaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gameaccount.FieldID)

		for i := range fields {
			if fields[i] != gameaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}

	if ps := gaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}

	if limit := gaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}

	if offset := gaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}

	if ps := gaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}

	return _spec
}

func (gaq *GameAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gaq.driver.Dialect())
	t1 := builder.Table(gameaccount.Table)

	columns := gaq.ctx.Fields
	if len(columns) == 0 {
		columns = gameaccount.Columns
	}

	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gaq.sql != nil {
		selector = gaq.sql
		selector.Select(selector.Columns(columns...)...)
	}

	if gaq.ctx.Unique != nil && *gaq.ctx.Unique {
		selector.Distinct()
	}

	for _, m := range gaq.modifiers {
		m(selector)
	}

	for _, p := range gaq.predicates {
		p(selector)
	}

	for _, p := range gaq.order {
		p(selector)
	}

	if offset := gaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}

	if limit := gaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}

	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gaq *GameAccountQuery) ForUpdate(opts ...sql.LockOption) *GameAccountQuery {
	if gaq.driver.Dialect() == dialect.Postgres {
		gaq.Unique(false)
	}

	gaq.modifiers = append(gaq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})

	return gaq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gaq *GameAccountQuery) ForShare(opts ...sql.LockOption) *GameAccountQuery {
	if gaq.driver.Dialect() == dialect.Postgres {
		gaq.Unique(false)
	}

	gaq.modifiers = append(gaq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})

	return gaq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gaq *GameAccountQuery) Modify(modifiers ...func(s *sql.Selector)) *GameAccountSelect {
	gaq.modifiers = append(gaq.modifiers, modifiers...)
	return gaq.Select()
}

// GameAccountGroupBy is the group-by builder for GameAccount entities.
type GameAccountGroupBy struct {
	selector
	build *GameAccountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gagb *GameAccountGroupBy) Aggregate(fns ...AggregateFunc) *GameAccountGroupBy {
	gagb.fns = append(gagb.fns, fns...)
	return gagb
}

// Scan applies the selector query and scans the result into the given value.
func (gagb *GameAccountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gagb.build.ctx, "GroupBy")
	if err := gagb.build.prepareQuery(ctx); err != nil {
		return err
	}

	return scanWithInterceptors[*GameAccountQuery, *GameAccountGroupBy](ctx, gagb.build, gagb, gagb.build.inters, v)
}

func (gagb *GameAccountGroupBy) sqlScan(ctx context.Context, root *GameAccountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gagb.fns))

	for _, fn := range gagb.fns {
		aggregation = append(aggregation, fn(selector))
	}

	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gagb.flds)+len(gagb.fns))
		for _, f := range *gagb.flds {
			columns = append(columns, selector.C(f))
		}

		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}

	selector.GroupBy(selector.Columns(*gagb.flds...)...)

	if err := selector.Err(); err != nil {
		return err
	}

	rows := &sql.Rows{}
	query, args := selector.Query()

	if err := gagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}

	defer rows.Close()

	return sql.ScanSlice(rows, v)
}

// GameAccountSelect is the builder for selecting fields of GameAccount entities.
type GameAccountSelect struct {
	*GameAccountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gas *GameAccountSelect) Aggregate(fns ...AggregateFunc) *GameAccountSelect {
	gas.fns = append(gas.fns, fns...)
	return gas
}

// Scan applies the selector query and scans the result into the given value.
func (gas *GameAccountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gas.ctx, "Select")
	if err := gas.prepareQuery(ctx); err != nil {
		return err
	}

	return scanWithInterceptors[*GameAccountQuery, *GameAccountSelect](ctx, gas.GameAccountQuery, gas, gas.inters, v)
}

func (gas *GameAccountSelect) sqlScan(ctx context.Context, root *GameAccountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gas.fns))

	for _, fn := range gas.fns {
		aggregation = append(aggregation, fn(selector))
	}

	switch n := len(*gas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}

	rows := &sql.Rows{}
	query, args := selector.Query()

	if err := gas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}

	defer rows.Close()

	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gas *GameAccountSelect) Modify(modifiers ...func(s *sql.Selector)) *GameAccountSelect {
	gas.modifiers = append(gas.modifiers, modifiers...)
	return gas
}
