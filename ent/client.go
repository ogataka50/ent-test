// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/ogataka50/ent-test/ent/migrate"

	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemdescription"
	"github.com/ogataka50/ent-test/ent/itemgroup"
	"github.com/ogataka50/ent-test/ent/itemvariation"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// ItemDescription is the client for interacting with the ItemDescription builders.
	ItemDescription *ItemDescriptionClient
	// ItemGroup is the client for interacting with the ItemGroup builders.
	ItemGroup *ItemGroupClient
	// ItemVariation is the client for interacting with the ItemVariation builders.
	ItemVariation *ItemVariationClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Item = NewItemClient(c.config)
	c.ItemDescription = NewItemDescriptionClient(c.config)
	c.ItemGroup = NewItemGroupClient(c.config)
	c.ItemVariation = NewItemVariationClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Item:            NewItemClient(cfg),
		ItemDescription: NewItemDescriptionClient(cfg),
		ItemGroup:       NewItemGroupClient(cfg),
		ItemVariation:   NewItemVariationClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:          cfg,
		Item:            NewItemClient(cfg),
		ItemDescription: NewItemDescriptionClient(cfg),
		ItemGroup:       NewItemGroupClient(cfg),
		ItemVariation:   NewItemVariationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Item.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Item.Use(hooks...)
	c.ItemDescription.Use(hooks...)
	c.ItemGroup.Use(hooks...)
	c.ItemVariation.Use(hooks...)
}

// ItemClient is a client for the Item schema.
type ItemClient struct {
	config
}

// NewItemClient returns a client for the Item from the given config.
func NewItemClient(c config) *ItemClient {
	return &ItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `item.Hooks(f(g(h())))`.
func (c *ItemClient) Use(hooks ...Hook) {
	c.hooks.Item = append(c.hooks.Item, hooks...)
}

// Create returns a create builder for Item.
func (c *ItemClient) Create() *ItemCreate {
	mutation := newItemMutation(c.config, OpCreate)
	return &ItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Item entities.
func (c *ItemClient) CreateBulk(builders ...*ItemCreate) *ItemCreateBulk {
	return &ItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Item.
func (c *ItemClient) Update() *ItemUpdate {
	mutation := newItemMutation(c.config, OpUpdate)
	return &ItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemClient) UpdateOne(i *Item) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItem(i))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemClient) UpdateOneID(id int) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItemID(id))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Item.
func (c *ItemClient) Delete() *ItemDelete {
	mutation := newItemMutation(c.config, OpDelete)
	return &ItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemClient) DeleteOne(i *Item) *ItemDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemClient) DeleteOneID(id int) *ItemDeleteOne {
	builder := c.Delete().Where(item.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemDeleteOne{builder}
}

// Query returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{config: c.config}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id int) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id int) *Item {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryItemDescription queries the item_description edge of a Item.
func (c *ItemClient) QueryItemDescription(i *Item) *ItemDescriptionQuery {
	query := &ItemDescriptionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(itemdescription.Table, itemdescription.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, item.ItemDescriptionTable, item.ItemDescriptionColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryItemVariation queries the item_variation edge of a Item.
func (c *ItemClient) QueryItemVariation(i *Item) *ItemVariationQuery {
	query := &ItemVariationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(itemvariation.Table, itemvariation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, item.ItemVariationTable, item.ItemVariationColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryItemGroup queries the item_group edge of a Item.
func (c *ItemClient) QueryItemGroup(i *Item) *ItemGroupQuery {
	query := &ItemGroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(itemgroup.Table, itemgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, item.ItemGroupTable, item.ItemGroupPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemClient) Hooks() []Hook {
	return c.hooks.Item
}

// ItemDescriptionClient is a client for the ItemDescription schema.
type ItemDescriptionClient struct {
	config
}

// NewItemDescriptionClient returns a client for the ItemDescription from the given config.
func NewItemDescriptionClient(c config) *ItemDescriptionClient {
	return &ItemDescriptionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `itemdescription.Hooks(f(g(h())))`.
func (c *ItemDescriptionClient) Use(hooks ...Hook) {
	c.hooks.ItemDescription = append(c.hooks.ItemDescription, hooks...)
}

// Create returns a create builder for ItemDescription.
func (c *ItemDescriptionClient) Create() *ItemDescriptionCreate {
	mutation := newItemDescriptionMutation(c.config, OpCreate)
	return &ItemDescriptionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ItemDescription entities.
func (c *ItemDescriptionClient) CreateBulk(builders ...*ItemDescriptionCreate) *ItemDescriptionCreateBulk {
	return &ItemDescriptionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ItemDescription.
func (c *ItemDescriptionClient) Update() *ItemDescriptionUpdate {
	mutation := newItemDescriptionMutation(c.config, OpUpdate)
	return &ItemDescriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemDescriptionClient) UpdateOne(id *ItemDescription) *ItemDescriptionUpdateOne {
	mutation := newItemDescriptionMutation(c.config, OpUpdateOne, withItemDescription(id))
	return &ItemDescriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemDescriptionClient) UpdateOneID(id int) *ItemDescriptionUpdateOne {
	mutation := newItemDescriptionMutation(c.config, OpUpdateOne, withItemDescriptionID(id))
	return &ItemDescriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ItemDescription.
func (c *ItemDescriptionClient) Delete() *ItemDescriptionDelete {
	mutation := newItemDescriptionMutation(c.config, OpDelete)
	return &ItemDescriptionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemDescriptionClient) DeleteOne(id *ItemDescription) *ItemDescriptionDeleteOne {
	return c.DeleteOneID(id.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemDescriptionClient) DeleteOneID(id int) *ItemDescriptionDeleteOne {
	builder := c.Delete().Where(itemdescription.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemDescriptionDeleteOne{builder}
}

// Query returns a query builder for ItemDescription.
func (c *ItemDescriptionClient) Query() *ItemDescriptionQuery {
	return &ItemDescriptionQuery{config: c.config}
}

// Get returns a ItemDescription entity by its id.
func (c *ItemDescriptionClient) Get(ctx context.Context, id int) (*ItemDescription, error) {
	return c.Query().Where(itemdescription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemDescriptionClient) GetX(ctx context.Context, id int) *ItemDescription {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a ItemDescription.
func (c *ItemDescriptionClient) QueryOwner(node *ItemDescription) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := node.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(itemdescription.Table, itemdescription.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, itemdescription.OwnerTable, itemdescription.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(node.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemDescriptionClient) Hooks() []Hook {
	return c.hooks.ItemDescription
}

// ItemGroupClient is a client for the ItemGroup schema.
type ItemGroupClient struct {
	config
}

// NewItemGroupClient returns a client for the ItemGroup from the given config.
func NewItemGroupClient(c config) *ItemGroupClient {
	return &ItemGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `itemgroup.Hooks(f(g(h())))`.
func (c *ItemGroupClient) Use(hooks ...Hook) {
	c.hooks.ItemGroup = append(c.hooks.ItemGroup, hooks...)
}

// Create returns a create builder for ItemGroup.
func (c *ItemGroupClient) Create() *ItemGroupCreate {
	mutation := newItemGroupMutation(c.config, OpCreate)
	return &ItemGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ItemGroup entities.
func (c *ItemGroupClient) CreateBulk(builders ...*ItemGroupCreate) *ItemGroupCreateBulk {
	return &ItemGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ItemGroup.
func (c *ItemGroupClient) Update() *ItemGroupUpdate {
	mutation := newItemGroupMutation(c.config, OpUpdate)
	return &ItemGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemGroupClient) UpdateOne(ig *ItemGroup) *ItemGroupUpdateOne {
	mutation := newItemGroupMutation(c.config, OpUpdateOne, withItemGroup(ig))
	return &ItemGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemGroupClient) UpdateOneID(id int) *ItemGroupUpdateOne {
	mutation := newItemGroupMutation(c.config, OpUpdateOne, withItemGroupID(id))
	return &ItemGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ItemGroup.
func (c *ItemGroupClient) Delete() *ItemGroupDelete {
	mutation := newItemGroupMutation(c.config, OpDelete)
	return &ItemGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemGroupClient) DeleteOne(ig *ItemGroup) *ItemGroupDeleteOne {
	return c.DeleteOneID(ig.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemGroupClient) DeleteOneID(id int) *ItemGroupDeleteOne {
	builder := c.Delete().Where(itemgroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemGroupDeleteOne{builder}
}

// Query returns a query builder for ItemGroup.
func (c *ItemGroupClient) Query() *ItemGroupQuery {
	return &ItemGroupQuery{config: c.config}
}

// Get returns a ItemGroup entity by its id.
func (c *ItemGroupClient) Get(ctx context.Context, id int) (*ItemGroup, error) {
	return c.Query().Where(itemgroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemGroupClient) GetX(ctx context.Context, id int) *ItemGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroupItem queries the group_item edge of a ItemGroup.
func (c *ItemGroupClient) QueryGroupItem(ig *ItemGroup) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ig.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(itemgroup.Table, itemgroup.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, itemgroup.GroupItemTable, itemgroup.GroupItemPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ig.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemGroupClient) Hooks() []Hook {
	return c.hooks.ItemGroup
}

// ItemVariationClient is a client for the ItemVariation schema.
type ItemVariationClient struct {
	config
}

// NewItemVariationClient returns a client for the ItemVariation from the given config.
func NewItemVariationClient(c config) *ItemVariationClient {
	return &ItemVariationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `itemvariation.Hooks(f(g(h())))`.
func (c *ItemVariationClient) Use(hooks ...Hook) {
	c.hooks.ItemVariation = append(c.hooks.ItemVariation, hooks...)
}

// Create returns a create builder for ItemVariation.
func (c *ItemVariationClient) Create() *ItemVariationCreate {
	mutation := newItemVariationMutation(c.config, OpCreate)
	return &ItemVariationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ItemVariation entities.
func (c *ItemVariationClient) CreateBulk(builders ...*ItemVariationCreate) *ItemVariationCreateBulk {
	return &ItemVariationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ItemVariation.
func (c *ItemVariationClient) Update() *ItemVariationUpdate {
	mutation := newItemVariationMutation(c.config, OpUpdate)
	return &ItemVariationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemVariationClient) UpdateOne(iv *ItemVariation) *ItemVariationUpdateOne {
	mutation := newItemVariationMutation(c.config, OpUpdateOne, withItemVariation(iv))
	return &ItemVariationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemVariationClient) UpdateOneID(id int) *ItemVariationUpdateOne {
	mutation := newItemVariationMutation(c.config, OpUpdateOne, withItemVariationID(id))
	return &ItemVariationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ItemVariation.
func (c *ItemVariationClient) Delete() *ItemVariationDelete {
	mutation := newItemVariationMutation(c.config, OpDelete)
	return &ItemVariationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemVariationClient) DeleteOne(iv *ItemVariation) *ItemVariationDeleteOne {
	return c.DeleteOneID(iv.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemVariationClient) DeleteOneID(id int) *ItemVariationDeleteOne {
	builder := c.Delete().Where(itemvariation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemVariationDeleteOne{builder}
}

// Query returns a query builder for ItemVariation.
func (c *ItemVariationClient) Query() *ItemVariationQuery {
	return &ItemVariationQuery{config: c.config}
}

// Get returns a ItemVariation entity by its id.
func (c *ItemVariationClient) Get(ctx context.Context, id int) (*ItemVariation, error) {
	return c.Query().Where(itemvariation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemVariationClient) GetX(ctx context.Context, id int) *ItemVariation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOriginItem queries the origin_item edge of a ItemVariation.
func (c *ItemVariationClient) QueryOriginItem(iv *ItemVariation) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := iv.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(itemvariation.Table, itemvariation.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, itemvariation.OriginItemTable, itemvariation.OriginItemColumn),
		)
		fromV = sqlgraph.Neighbors(iv.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemVariationClient) Hooks() []Hook {
	return c.hooks.ItemVariation
}
