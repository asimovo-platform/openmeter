// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationchannel"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationrule"
	"github.com/openmeterio/openmeter/internal/ent/db/predicate"
)

// NotificationRuleQuery is the builder for querying NotificationRule entities.
type NotificationRuleQuery struct {
	config
	ctx          *QueryContext
	order        []notificationrule.OrderOption
	inters       []Interceptor
	predicates   []predicate.NotificationRule
	withChannels *NotificationChannelQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotificationRuleQuery builder.
func (nrq *NotificationRuleQuery) Where(ps ...predicate.NotificationRule) *NotificationRuleQuery {
	nrq.predicates = append(nrq.predicates, ps...)
	return nrq
}

// Limit the number of records to be returned by this query.
func (nrq *NotificationRuleQuery) Limit(limit int) *NotificationRuleQuery {
	nrq.ctx.Limit = &limit
	return nrq
}

// Offset to start from.
func (nrq *NotificationRuleQuery) Offset(offset int) *NotificationRuleQuery {
	nrq.ctx.Offset = &offset
	return nrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nrq *NotificationRuleQuery) Unique(unique bool) *NotificationRuleQuery {
	nrq.ctx.Unique = &unique
	return nrq
}

// Order specifies how the records should be ordered.
func (nrq *NotificationRuleQuery) Order(o ...notificationrule.OrderOption) *NotificationRuleQuery {
	nrq.order = append(nrq.order, o...)
	return nrq
}

// QueryChannels chains the current query on the "channels" edge.
func (nrq *NotificationRuleQuery) QueryChannels() *NotificationChannelQuery {
	query := (&NotificationChannelClient{config: nrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notificationrule.Table, notificationrule.FieldID, selector),
			sqlgraph.To(notificationchannel.Table, notificationchannel.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, notificationrule.ChannelsTable, notificationrule.ChannelsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NotificationRule entity from the query.
// Returns a *NotFoundError when no NotificationRule was found.
func (nrq *NotificationRuleQuery) First(ctx context.Context) (*NotificationRule, error) {
	nodes, err := nrq.Limit(1).All(setContextOp(ctx, nrq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notificationrule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nrq *NotificationRuleQuery) FirstX(ctx context.Context) *NotificationRule {
	node, err := nrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NotificationRule ID from the query.
// Returns a *NotFoundError when no NotificationRule ID was found.
func (nrq *NotificationRuleQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = nrq.Limit(1).IDs(setContextOp(ctx, nrq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notificationrule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nrq *NotificationRuleQuery) FirstIDX(ctx context.Context) string {
	id, err := nrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NotificationRule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NotificationRule entity is found.
// Returns a *NotFoundError when no NotificationRule entities are found.
func (nrq *NotificationRuleQuery) Only(ctx context.Context) (*NotificationRule, error) {
	nodes, err := nrq.Limit(2).All(setContextOp(ctx, nrq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notificationrule.Label}
	default:
		return nil, &NotSingularError{notificationrule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nrq *NotificationRuleQuery) OnlyX(ctx context.Context) *NotificationRule {
	node, err := nrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NotificationRule ID in the query.
// Returns a *NotSingularError when more than one NotificationRule ID is found.
// Returns a *NotFoundError when no entities are found.
func (nrq *NotificationRuleQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = nrq.Limit(2).IDs(setContextOp(ctx, nrq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notificationrule.Label}
	default:
		err = &NotSingularError{notificationrule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nrq *NotificationRuleQuery) OnlyIDX(ctx context.Context) string {
	id, err := nrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NotificationRules.
func (nrq *NotificationRuleQuery) All(ctx context.Context) ([]*NotificationRule, error) {
	ctx = setContextOp(ctx, nrq.ctx, ent.OpQueryAll)
	if err := nrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NotificationRule, *NotificationRuleQuery]()
	return withInterceptors[[]*NotificationRule](ctx, nrq, qr, nrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nrq *NotificationRuleQuery) AllX(ctx context.Context) []*NotificationRule {
	nodes, err := nrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NotificationRule IDs.
func (nrq *NotificationRuleQuery) IDs(ctx context.Context) (ids []string, err error) {
	if nrq.ctx.Unique == nil && nrq.path != nil {
		nrq.Unique(true)
	}
	ctx = setContextOp(ctx, nrq.ctx, ent.OpQueryIDs)
	if err = nrq.Select(notificationrule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nrq *NotificationRuleQuery) IDsX(ctx context.Context) []string {
	ids, err := nrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nrq *NotificationRuleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nrq.ctx, ent.OpQueryCount)
	if err := nrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nrq, querierCount[*NotificationRuleQuery](), nrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nrq *NotificationRuleQuery) CountX(ctx context.Context) int {
	count, err := nrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nrq *NotificationRuleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nrq.ctx, ent.OpQueryExist)
	switch _, err := nrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nrq *NotificationRuleQuery) ExistX(ctx context.Context) bool {
	exist, err := nrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotificationRuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nrq *NotificationRuleQuery) Clone() *NotificationRuleQuery {
	if nrq == nil {
		return nil
	}
	return &NotificationRuleQuery{
		config:       nrq.config,
		ctx:          nrq.ctx.Clone(),
		order:        append([]notificationrule.OrderOption{}, nrq.order...),
		inters:       append([]Interceptor{}, nrq.inters...),
		predicates:   append([]predicate.NotificationRule{}, nrq.predicates...),
		withChannels: nrq.withChannels.Clone(),
		// clone intermediate query.
		sql:  nrq.sql.Clone(),
		path: nrq.path,
	}
}

// WithChannels tells the query-builder to eager-load the nodes that are connected to
// the "channels" edge. The optional arguments are used to configure the query builder of the edge.
func (nrq *NotificationRuleQuery) WithChannels(opts ...func(*NotificationChannelQuery)) *NotificationRuleQuery {
	query := (&NotificationChannelClient{config: nrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nrq.withChannels = query
	return nrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NotificationRule.Query().
//		GroupBy(notificationrule.FieldNamespace).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (nrq *NotificationRuleQuery) GroupBy(field string, fields ...string) *NotificationRuleGroupBy {
	nrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NotificationRuleGroupBy{build: nrq}
	grbuild.flds = &nrq.ctx.Fields
	grbuild.label = notificationrule.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//	}
//
//	client.NotificationRule.Query().
//		Select(notificationrule.FieldNamespace).
//		Scan(ctx, &v)
func (nrq *NotificationRuleQuery) Select(fields ...string) *NotificationRuleSelect {
	nrq.ctx.Fields = append(nrq.ctx.Fields, fields...)
	sbuild := &NotificationRuleSelect{NotificationRuleQuery: nrq}
	sbuild.label = notificationrule.Label
	sbuild.flds, sbuild.scan = &nrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NotificationRuleSelect configured with the given aggregations.
func (nrq *NotificationRuleQuery) Aggregate(fns ...AggregateFunc) *NotificationRuleSelect {
	return nrq.Select().Aggregate(fns...)
}

func (nrq *NotificationRuleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nrq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nrq); err != nil {
				return err
			}
		}
	}
	for _, f := range nrq.ctx.Fields {
		if !notificationrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if nrq.path != nil {
		prev, err := nrq.path(ctx)
		if err != nil {
			return err
		}
		nrq.sql = prev
	}
	return nil
}

func (nrq *NotificationRuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NotificationRule, error) {
	var (
		nodes       = []*NotificationRule{}
		_spec       = nrq.querySpec()
		loadedTypes = [1]bool{
			nrq.withChannels != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NotificationRule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NotificationRule{config: nrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(nrq.modifiers) > 0 {
		_spec.Modifiers = nrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nrq.withChannels; query != nil {
		if err := nrq.loadChannels(ctx, query, nodes,
			func(n *NotificationRule) { n.Edges.Channels = []*NotificationChannel{} },
			func(n *NotificationRule, e *NotificationChannel) { n.Edges.Channels = append(n.Edges.Channels, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nrq *NotificationRuleQuery) loadChannels(ctx context.Context, query *NotificationChannelQuery, nodes []*NotificationRule, init func(*NotificationRule), assign func(*NotificationRule, *NotificationChannel)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*NotificationRule)
	nids := make(map[string]map[*NotificationRule]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(notificationrule.ChannelsTable)
		s.Join(joinT).On(s.C(notificationchannel.FieldID), joinT.C(notificationrule.ChannelsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(notificationrule.ChannelsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(notificationrule.ChannelsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*NotificationRule]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*NotificationChannel](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "channels" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (nrq *NotificationRuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nrq.querySpec()
	if len(nrq.modifiers) > 0 {
		_spec.Modifiers = nrq.modifiers
	}
	_spec.Node.Columns = nrq.ctx.Fields
	if len(nrq.ctx.Fields) > 0 {
		_spec.Unique = nrq.ctx.Unique != nil && *nrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nrq.driver, _spec)
}

func (nrq *NotificationRuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(notificationrule.Table, notificationrule.Columns, sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString))
	_spec.From = nrq.sql
	if unique := nrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nrq.path != nil {
		_spec.Unique = true
	}
	if fields := nrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationrule.FieldID)
		for i := range fields {
			if fields[i] != notificationrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nrq *NotificationRuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nrq.driver.Dialect())
	t1 := builder.Table(notificationrule.Table)
	columns := nrq.ctx.Fields
	if len(columns) == 0 {
		columns = notificationrule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nrq.sql != nil {
		selector = nrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nrq.ctx.Unique != nil && *nrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range nrq.modifiers {
		m(selector)
	}
	for _, p := range nrq.predicates {
		p(selector)
	}
	for _, p := range nrq.order {
		p(selector)
	}
	if offset := nrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (nrq *NotificationRuleQuery) ForUpdate(opts ...sql.LockOption) *NotificationRuleQuery {
	if nrq.driver.Dialect() == dialect.Postgres {
		nrq.Unique(false)
	}
	nrq.modifiers = append(nrq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return nrq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (nrq *NotificationRuleQuery) ForShare(opts ...sql.LockOption) *NotificationRuleQuery {
	if nrq.driver.Dialect() == dialect.Postgres {
		nrq.Unique(false)
	}
	nrq.modifiers = append(nrq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return nrq
}

// NotificationRuleGroupBy is the group-by builder for NotificationRule entities.
type NotificationRuleGroupBy struct {
	selector
	build *NotificationRuleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nrgb *NotificationRuleGroupBy) Aggregate(fns ...AggregateFunc) *NotificationRuleGroupBy {
	nrgb.fns = append(nrgb.fns, fns...)
	return nrgb
}

// Scan applies the selector query and scans the result into the given value.
func (nrgb *NotificationRuleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nrgb.build.ctx, ent.OpQueryGroupBy)
	if err := nrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationRuleQuery, *NotificationRuleGroupBy](ctx, nrgb.build, nrgb, nrgb.build.inters, v)
}

func (nrgb *NotificationRuleGroupBy) sqlScan(ctx context.Context, root *NotificationRuleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nrgb.fns))
	for _, fn := range nrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nrgb.flds)+len(nrgb.fns))
		for _, f := range *nrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NotificationRuleSelect is the builder for selecting fields of NotificationRule entities.
type NotificationRuleSelect struct {
	*NotificationRuleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nrs *NotificationRuleSelect) Aggregate(fns ...AggregateFunc) *NotificationRuleSelect {
	nrs.fns = append(nrs.fns, fns...)
	return nrs
}

// Scan applies the selector query and scans the result into the given value.
func (nrs *NotificationRuleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nrs.ctx, ent.OpQuerySelect)
	if err := nrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationRuleQuery, *NotificationRuleSelect](ctx, nrs.NotificationRuleQuery, nrs, nrs.inters, v)
}

func (nrs *NotificationRuleSelect) sqlScan(ctx context.Context, root *NotificationRuleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nrs.fns))
	for _, fn := range nrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
