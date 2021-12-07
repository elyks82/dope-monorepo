// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/dopedao/dope-monorepo/packages/api/ent/migrate"

	"github.com/dopedao/dope-monorepo/packages/api/ent/dope"
	"github.com/dopedao/dope-monorepo/packages/api/ent/hustler"
	"github.com/dopedao/dope-monorepo/packages/api/ent/item"
	"github.com/dopedao/dope-monorepo/packages/api/ent/syncstate"
	"github.com/dopedao/dope-monorepo/packages/api/ent/wallet"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Dope is the client for interacting with the Dope builders.
	Dope *DopeClient
	// Hustler is the client for interacting with the Hustler builders.
	Hustler *HustlerClient
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// SyncState is the client for interacting with the SyncState builders.
	SyncState *SyncStateClient
	// Wallet is the client for interacting with the Wallet builders.
	Wallet *WalletClient
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
	c.Dope = NewDopeClient(c.config)
	c.Hustler = NewHustlerClient(c.config)
	c.Item = NewItemClient(c.config)
	c.SyncState = NewSyncStateClient(c.config)
	c.Wallet = NewWalletClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Dope:      NewDopeClient(cfg),
		Hustler:   NewHustlerClient(cfg),
		Item:      NewItemClient(cfg),
		SyncState: NewSyncStateClient(cfg),
		Wallet:    NewWalletClient(cfg),
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
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:    cfg,
		Dope:      NewDopeClient(cfg),
		Hustler:   NewHustlerClient(cfg),
		Item:      NewItemClient(cfg),
		SyncState: NewSyncStateClient(cfg),
		Wallet:    NewWalletClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Dope.
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
	c.Dope.Use(hooks...)
	c.Hustler.Use(hooks...)
	c.Item.Use(hooks...)
	c.SyncState.Use(hooks...)
	c.Wallet.Use(hooks...)
}

// DopeClient is a client for the Dope schema.
type DopeClient struct {
	config
}

// NewDopeClient returns a client for the Dope from the given config.
func NewDopeClient(c config) *DopeClient {
	return &DopeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dope.Hooks(f(g(h())))`.
func (c *DopeClient) Use(hooks ...Hook) {
	c.hooks.Dope = append(c.hooks.Dope, hooks...)
}

// Create returns a create builder for Dope.
func (c *DopeClient) Create() *DopeCreate {
	mutation := newDopeMutation(c.config, OpCreate)
	return &DopeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dope entities.
func (c *DopeClient) CreateBulk(builders ...*DopeCreate) *DopeCreateBulk {
	return &DopeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dope.
func (c *DopeClient) Update() *DopeUpdate {
	mutation := newDopeMutation(c.config, OpUpdate)
	return &DopeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DopeClient) UpdateOne(d *Dope) *DopeUpdateOne {
	mutation := newDopeMutation(c.config, OpUpdateOne, withDope(d))
	return &DopeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DopeClient) UpdateOneID(id string) *DopeUpdateOne {
	mutation := newDopeMutation(c.config, OpUpdateOne, withDopeID(id))
	return &DopeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dope.
func (c *DopeClient) Delete() *DopeDelete {
	mutation := newDopeMutation(c.config, OpDelete)
	return &DopeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DopeClient) DeleteOne(d *Dope) *DopeDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DopeClient) DeleteOneID(id string) *DopeDeleteOne {
	builder := c.Delete().Where(dope.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DopeDeleteOne{builder}
}

// Query returns a query builder for Dope.
func (c *DopeClient) Query() *DopeQuery {
	return &DopeQuery{
		config: c.config,
	}
}

// Get returns a Dope entity by its id.
func (c *DopeClient) Get(ctx context.Context, id string) (*Dope, error) {
	return c.Query().Where(dope.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DopeClient) GetX(ctx context.Context, id string) *Dope {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWallet queries the wallet edge of a Dope.
func (c *DopeClient) QueryWallet(d *Dope) *WalletQuery {
	query := &WalletQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dope.Table, dope.FieldID, id),
			sqlgraph.To(wallet.Table, wallet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dope.WalletTable, dope.WalletColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryItems queries the items edge of a Dope.
func (c *DopeClient) QueryItems(d *Dope) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dope.Table, dope.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, dope.ItemsTable, dope.ItemsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DopeClient) Hooks() []Hook {
	return c.hooks.Dope
}

// HustlerClient is a client for the Hustler schema.
type HustlerClient struct {
	config
}

// NewHustlerClient returns a client for the Hustler from the given config.
func NewHustlerClient(c config) *HustlerClient {
	return &HustlerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `hustler.Hooks(f(g(h())))`.
func (c *HustlerClient) Use(hooks ...Hook) {
	c.hooks.Hustler = append(c.hooks.Hustler, hooks...)
}

// Create returns a create builder for Hustler.
func (c *HustlerClient) Create() *HustlerCreate {
	mutation := newHustlerMutation(c.config, OpCreate)
	return &HustlerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Hustler entities.
func (c *HustlerClient) CreateBulk(builders ...*HustlerCreate) *HustlerCreateBulk {
	return &HustlerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Hustler.
func (c *HustlerClient) Update() *HustlerUpdate {
	mutation := newHustlerMutation(c.config, OpUpdate)
	return &HustlerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HustlerClient) UpdateOne(h *Hustler) *HustlerUpdateOne {
	mutation := newHustlerMutation(c.config, OpUpdateOne, withHustler(h))
	return &HustlerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HustlerClient) UpdateOneID(id string) *HustlerUpdateOne {
	mutation := newHustlerMutation(c.config, OpUpdateOne, withHustlerID(id))
	return &HustlerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Hustler.
func (c *HustlerClient) Delete() *HustlerDelete {
	mutation := newHustlerMutation(c.config, OpDelete)
	return &HustlerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *HustlerClient) DeleteOne(h *Hustler) *HustlerDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *HustlerClient) DeleteOneID(id string) *HustlerDeleteOne {
	builder := c.Delete().Where(hustler.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HustlerDeleteOne{builder}
}

// Query returns a query builder for Hustler.
func (c *HustlerClient) Query() *HustlerQuery {
	return &HustlerQuery{
		config: c.config,
	}
}

// Get returns a Hustler entity by its id.
func (c *HustlerClient) Get(ctx context.Context, id string) (*Hustler, error) {
	return c.Query().Where(hustler.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HustlerClient) GetX(ctx context.Context, id string) *Hustler {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWallet queries the wallet edge of a Hustler.
func (c *HustlerClient) QueryWallet(h *Hustler) *WalletQuery {
	query := &WalletQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(hustler.Table, hustler.FieldID, id),
			sqlgraph.To(wallet.Table, wallet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hustler.WalletTable, hustler.WalletColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HustlerClient) Hooks() []Hook {
	return c.hooks.Hustler
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
func (c *ItemClient) UpdateOneID(id string) *ItemUpdateOne {
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
func (c *ItemClient) DeleteOneID(id string) *ItemDeleteOne {
	builder := c.Delete().Where(item.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemDeleteOne{builder}
}

// Query returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{
		config: c.config,
	}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id string) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id string) *Item {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWallet queries the wallet edge of a Item.
func (c *ItemClient) QueryWallet(i *Item) *WalletQuery {
	query := &WalletQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(wallet.Table, wallet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, item.WalletTable, item.WalletColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDopes queries the dopes edge of a Item.
func (c *ItemClient) QueryDopes(i *Item) *DopeQuery {
	query := &DopeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(dope.Table, dope.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, item.DopesTable, item.DopesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBase queries the base edge of a Item.
func (c *ItemClient) QueryBase(i *Item) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, item.BaseTable, item.BaseColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDerivative queries the derivative edge of a Item.
func (c *ItemClient) QueryDerivative(i *Item) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, item.DerivativeTable, item.DerivativeColumn),
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

// SyncStateClient is a client for the SyncState schema.
type SyncStateClient struct {
	config
}

// NewSyncStateClient returns a client for the SyncState from the given config.
func NewSyncStateClient(c config) *SyncStateClient {
	return &SyncStateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `syncstate.Hooks(f(g(h())))`.
func (c *SyncStateClient) Use(hooks ...Hook) {
	c.hooks.SyncState = append(c.hooks.SyncState, hooks...)
}

// Create returns a create builder for SyncState.
func (c *SyncStateClient) Create() *SyncStateCreate {
	mutation := newSyncStateMutation(c.config, OpCreate)
	return &SyncStateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SyncState entities.
func (c *SyncStateClient) CreateBulk(builders ...*SyncStateCreate) *SyncStateCreateBulk {
	return &SyncStateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SyncState.
func (c *SyncStateClient) Update() *SyncStateUpdate {
	mutation := newSyncStateMutation(c.config, OpUpdate)
	return &SyncStateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SyncStateClient) UpdateOne(ss *SyncState) *SyncStateUpdateOne {
	mutation := newSyncStateMutation(c.config, OpUpdateOne, withSyncState(ss))
	return &SyncStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SyncStateClient) UpdateOneID(id string) *SyncStateUpdateOne {
	mutation := newSyncStateMutation(c.config, OpUpdateOne, withSyncStateID(id))
	return &SyncStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SyncState.
func (c *SyncStateClient) Delete() *SyncStateDelete {
	mutation := newSyncStateMutation(c.config, OpDelete)
	return &SyncStateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SyncStateClient) DeleteOne(ss *SyncState) *SyncStateDeleteOne {
	return c.DeleteOneID(ss.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SyncStateClient) DeleteOneID(id string) *SyncStateDeleteOne {
	builder := c.Delete().Where(syncstate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SyncStateDeleteOne{builder}
}

// Query returns a query builder for SyncState.
func (c *SyncStateClient) Query() *SyncStateQuery {
	return &SyncStateQuery{
		config: c.config,
	}
}

// Get returns a SyncState entity by its id.
func (c *SyncStateClient) Get(ctx context.Context, id string) (*SyncState, error) {
	return c.Query().Where(syncstate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SyncStateClient) GetX(ctx context.Context, id string) *SyncState {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SyncStateClient) Hooks() []Hook {
	return c.hooks.SyncState
}

// WalletClient is a client for the Wallet schema.
type WalletClient struct {
	config
}

// NewWalletClient returns a client for the Wallet from the given config.
func NewWalletClient(c config) *WalletClient {
	return &WalletClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `wallet.Hooks(f(g(h())))`.
func (c *WalletClient) Use(hooks ...Hook) {
	c.hooks.Wallet = append(c.hooks.Wallet, hooks...)
}

// Create returns a create builder for Wallet.
func (c *WalletClient) Create() *WalletCreate {
	mutation := newWalletMutation(c.config, OpCreate)
	return &WalletCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Wallet entities.
func (c *WalletClient) CreateBulk(builders ...*WalletCreate) *WalletCreateBulk {
	return &WalletCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Wallet.
func (c *WalletClient) Update() *WalletUpdate {
	mutation := newWalletMutation(c.config, OpUpdate)
	return &WalletUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WalletClient) UpdateOne(w *Wallet) *WalletUpdateOne {
	mutation := newWalletMutation(c.config, OpUpdateOne, withWallet(w))
	return &WalletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WalletClient) UpdateOneID(id string) *WalletUpdateOne {
	mutation := newWalletMutation(c.config, OpUpdateOne, withWalletID(id))
	return &WalletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Wallet.
func (c *WalletClient) Delete() *WalletDelete {
	mutation := newWalletMutation(c.config, OpDelete)
	return &WalletDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WalletClient) DeleteOne(w *Wallet) *WalletDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WalletClient) DeleteOneID(id string) *WalletDeleteOne {
	builder := c.Delete().Where(wallet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WalletDeleteOne{builder}
}

// Query returns a query builder for Wallet.
func (c *WalletClient) Query() *WalletQuery {
	return &WalletQuery{
		config: c.config,
	}
}

// Get returns a Wallet entity by its id.
func (c *WalletClient) Get(ctx context.Context, id string) (*Wallet, error) {
	return c.Query().Where(wallet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WalletClient) GetX(ctx context.Context, id string) *Wallet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDopes queries the dopes edge of a Wallet.
func (c *WalletClient) QueryDopes(w *Wallet) *DopeQuery {
	query := &DopeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(wallet.Table, wallet.FieldID, id),
			sqlgraph.To(dope.Table, dope.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wallet.DopesTable, wallet.DopesColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryItems queries the items edge of a Wallet.
func (c *WalletClient) QueryItems(w *Wallet) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(wallet.Table, wallet.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wallet.ItemsTable, wallet.ItemsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHustlers queries the hustlers edge of a Wallet.
func (c *WalletClient) QueryHustlers(w *Wallet) *HustlerQuery {
	query := &HustlerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(wallet.Table, wallet.FieldID, id),
			sqlgraph.To(hustler.Table, hustler.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wallet.HustlersTable, wallet.HustlersColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WalletClient) Hooks() []Hook {
	return c.hooks.Wallet
}