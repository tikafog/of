// Code generated by ent, DO NOT EDIT.

package bootnode

import (
	"context"
	"fmt"
	"log"

	"github.com/tikafog/of/dbc/bootnode/migrate"

	"github.com/tikafog/of/dbc/bootnode/bootstrap"
	"github.com/tikafog/of/dbc/bootnode/version"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Bootstrap is the client for interacting with the Bootstrap builders.
	Bootstrap *BootstrapClient
	// Version is the client for interacting with the Version builders.
	Version *VersionClient
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
	c.Bootstrap = NewBootstrapClient(c.config)
	c.Version = NewVersionClient(c.config)
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
		return nil, fmt.Errorf("bootnode: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("bootnode: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Bootstrap: NewBootstrapClient(cfg),
		Version:   NewVersionClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		Bootstrap: NewBootstrapClient(cfg),
		Version:   NewVersionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Bootstrap.
//		Query().
//		Count(ctx)
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
	c.Bootstrap.Use(hooks...)
	c.Version.Use(hooks...)
}

// BootstrapClient is a client for the Bootstrap schema.
type BootstrapClient struct {
	config
}

// NewBootstrapClient returns a client for the Bootstrap from the given config.
func NewBootstrapClient(c config) *BootstrapClient {
	return &BootstrapClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bootstrap.Hooks(f(g(h())))`.
func (c *BootstrapClient) Use(hooks ...Hook) {
	c.hooks.Bootstrap = append(c.hooks.Bootstrap, hooks...)
}

// Create returns a builder for creating a Bootstrap entity.
func (c *BootstrapClient) Create() *BootstrapCreate {
	mutation := newBootstrapMutation(c.config, OpCreate)
	return &BootstrapCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Bootstrap entities.
func (c *BootstrapClient) CreateBulk(builders ...*BootstrapCreate) *BootstrapCreateBulk {
	return &BootstrapCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Bootstrap.
func (c *BootstrapClient) Update() *BootstrapUpdate {
	mutation := newBootstrapMutation(c.config, OpUpdate)
	return &BootstrapUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BootstrapClient) UpdateOne(b *Bootstrap) *BootstrapUpdateOne {
	mutation := newBootstrapMutation(c.config, OpUpdateOne, withBootstrap(b))
	return &BootstrapUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BootstrapClient) UpdateOneID(id int) *BootstrapUpdateOne {
	mutation := newBootstrapMutation(c.config, OpUpdateOne, withBootstrapID(id))
	return &BootstrapUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bootstrap.
func (c *BootstrapClient) Delete() *BootstrapDelete {
	mutation := newBootstrapMutation(c.config, OpDelete)
	return &BootstrapDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BootstrapClient) DeleteOne(b *Bootstrap) *BootstrapDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BootstrapClient) DeleteOneID(id int) *BootstrapDeleteOne {
	builder := c.Delete().Where(bootstrap.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BootstrapDeleteOne{builder}
}

// Query returns a query builder for Bootstrap.
func (c *BootstrapClient) Query() *BootstrapQuery {
	return &BootstrapQuery{
		config: c.config,
	}
}

// Get returns a Bootstrap entity by its id.
func (c *BootstrapClient) Get(ctx context.Context, id int) (*Bootstrap, error) {
	return c.Query().Where(bootstrap.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BootstrapClient) GetX(ctx context.Context, id int) *Bootstrap {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BootstrapClient) Hooks() []Hook {
	return c.hooks.Bootstrap
}

// VersionClient is a client for the Version schema.
type VersionClient struct {
	config
}

// NewVersionClient returns a client for the Version from the given config.
func NewVersionClient(c config) *VersionClient {
	return &VersionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `version.Hooks(f(g(h())))`.
func (c *VersionClient) Use(hooks ...Hook) {
	c.hooks.Version = append(c.hooks.Version, hooks...)
}

// Create returns a builder for creating a Version entity.
func (c *VersionClient) Create() *VersionCreate {
	mutation := newVersionMutation(c.config, OpCreate)
	return &VersionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Version entities.
func (c *VersionClient) CreateBulk(builders ...*VersionCreate) *VersionCreateBulk {
	return &VersionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Version.
func (c *VersionClient) Update() *VersionUpdate {
	mutation := newVersionMutation(c.config, OpUpdate)
	return &VersionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VersionClient) UpdateOne(v *Version) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersion(v))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VersionClient) UpdateOneID(id int) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersionID(id))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Version.
func (c *VersionClient) Delete() *VersionDelete {
	mutation := newVersionMutation(c.config, OpDelete)
	return &VersionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VersionClient) DeleteOne(v *Version) *VersionDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *VersionClient) DeleteOneID(id int) *VersionDeleteOne {
	builder := c.Delete().Where(version.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VersionDeleteOne{builder}
}

// Query returns a query builder for Version.
func (c *VersionClient) Query() *VersionQuery {
	return &VersionQuery{
		config: c.config,
	}
}

// Get returns a Version entity by its id.
func (c *VersionClient) Get(ctx context.Context, id int) (*Version, error) {
	return c.Query().Where(version.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VersionClient) GetX(ctx context.Context, id int) *Version {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VersionClient) Hooks() []Hook {
	return c.hooks.Version
}
