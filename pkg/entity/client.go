// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/lostsnow/keqing/pkg/entity/migrate"

	"github.com/lostsnow/keqing/pkg/entity/chat"
	"github.com/lostsnow/keqing/pkg/entity/chatoption"
	"github.com/lostsnow/keqing/pkg/entity/gameaccount"
	"github.com/lostsnow/keqing/pkg/entity/gamerole"
	"github.com/lostsnow/keqing/pkg/entity/gameroleattribute"
	"github.com/lostsnow/keqing/pkg/entity/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Chat is the client for interacting with the Chat builders.
	Chat *ChatClient
	// ChatOption is the client for interacting with the ChatOption builders.
	ChatOption *ChatOptionClient
	// GameAccount is the client for interacting with the GameAccount builders.
	GameAccount *GameAccountClient
	// GameRole is the client for interacting with the GameRole builders.
	GameRole *GameRoleClient
	// GameRoleAttribute is the client for interacting with the GameRoleAttribute builders.
	GameRoleAttribute *GameRoleAttributeClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()

	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Chat = NewChatClient(c.config)
	c.ChatOption = NewChatOptionClient(c.config)
	c.GameAccount = NewGameAccountClient(c.config)
	c.GameRole = NewGameRoleClient(c.config)
	c.GameRoleAttribute = NewGameRoleAttributeClient(c.config)
	c.User = NewUserClient(c.config)
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
		return nil, errors.New("entity: cannot start a transaction within a transaction")
	}

	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("entity: starting a transaction: %w", err)
	}

	cfg := c.config
	cfg.driver = tx

	return &Tx{
		ctx:               ctx,
		config:            cfg,
		Chat:              NewChatClient(cfg),
		ChatOption:        NewChatOptionClient(cfg),
		GameAccount:       NewGameAccountClient(cfg),
		GameRole:          NewGameRoleClient(cfg),
		GameRoleAttribute: NewGameRoleAttributeClient(cfg),
		User:              NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:               ctx,
		config:            cfg,
		Chat:              NewChatClient(cfg),
		ChatOption:        NewChatOptionClient(cfg),
		GameAccount:       NewGameAccountClient(cfg),
		GameRole:          NewGameRoleClient(cfg),
		GameRoleAttribute: NewGameRoleAttributeClient(cfg),
		User:              NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Chat.
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
	c.Chat.Use(hooks...)
	c.ChatOption.Use(hooks...)
	c.GameAccount.Use(hooks...)
	c.GameRole.Use(hooks...)
	c.GameRoleAttribute.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Chat.Intercept(interceptors...)
	c.ChatOption.Intercept(interceptors...)
	c.GameAccount.Intercept(interceptors...)
	c.GameRole.Intercept(interceptors...)
	c.GameRoleAttribute.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ChatMutation:
		return c.Chat.mutate(ctx, m)
	case *ChatOptionMutation:
		return c.ChatOption.mutate(ctx, m)
	case *GameAccountMutation:
		return c.GameAccount.mutate(ctx, m)
	case *GameRoleMutation:
		return c.GameRole.mutate(ctx, m)
	case *GameRoleAttributeMutation:
		return c.GameRoleAttribute.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("entity: unknown mutation type %T", m)
	}
}

// ChatClient is a client for the Chat schema.
type ChatClient struct {
	config
}

// NewChatClient returns a client for the Chat from the given config.
func NewChatClient(c config) *ChatClient {
	return &ChatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chat.Hooks(f(g(h())))`.
func (c *ChatClient) Use(hooks ...Hook) {
	c.hooks.Chat = append(c.hooks.Chat, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chat.Intercept(f(g(h())))`.
func (c *ChatClient) Intercept(interceptors ...Interceptor) {
	c.inters.Chat = append(c.inters.Chat, interceptors...)
}

// Create returns a builder for creating a Chat entity.
func (c *ChatClient) Create() *ChatCreate {
	mutation := newChatMutation(c.config, OpCreate)
	return &ChatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Chat entities.
func (c *ChatClient) CreateBulk(builders ...*ChatCreate) *ChatCreateBulk {
	return &ChatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Chat.
func (c *ChatClient) Update() *ChatUpdate {
	mutation := newChatMutation(c.config, OpUpdate)
	return &ChatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatClient) UpdateOne(ch *Chat) *ChatUpdateOne {
	mutation := newChatMutation(c.config, OpUpdateOne, withChat(ch))
	return &ChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatClient) UpdateOneID(id int64) *ChatUpdateOne {
	mutation := newChatMutation(c.config, OpUpdateOne, withChatID(id))
	return &ChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Chat.
func (c *ChatClient) Delete() *ChatDelete {
	mutation := newChatMutation(c.config, OpDelete)
	return &ChatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChatClient) DeleteOne(ch *Chat) *ChatDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChatClient) DeleteOneID(id int64) *ChatDeleteOne {
	builder := c.Delete().Where(chat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne

	return &ChatDeleteOne{builder}
}

// Query returns a query builder for Chat.
func (c *ChatClient) Query() *ChatQuery {
	return &ChatQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChat},
		inters: c.Interceptors(),
	}
}

// Get returns a Chat entity by its id.
func (c *ChatClient) Get(ctx context.Context, id int64) (*Chat, error) {
	return c.Query().Where(chat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatClient) GetX(ctx context.Context, id int64) *Chat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	return obj
}

// Hooks returns the client hooks.
func (c *ChatClient) Hooks() []Hook {
	return c.hooks.Chat
}

// Interceptors returns the client interceptors.
func (c *ChatClient) Interceptors() []Interceptor {
	return c.inters.Chat
}

func (c *ChatClient) mutate(ctx context.Context, m *ChatMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChatCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChatUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChatDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown Chat mutation op: %q", m.Op())
	}
}

// ChatOptionClient is a client for the ChatOption schema.
type ChatOptionClient struct {
	config
}

// NewChatOptionClient returns a client for the ChatOption from the given config.
func NewChatOptionClient(c config) *ChatOptionClient {
	return &ChatOptionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chatoption.Hooks(f(g(h())))`.
func (c *ChatOptionClient) Use(hooks ...Hook) {
	c.hooks.ChatOption = append(c.hooks.ChatOption, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chatoption.Intercept(f(g(h())))`.
func (c *ChatOptionClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChatOption = append(c.inters.ChatOption, interceptors...)
}

// Create returns a builder for creating a ChatOption entity.
func (c *ChatOptionClient) Create() *ChatOptionCreate {
	mutation := newChatOptionMutation(c.config, OpCreate)
	return &ChatOptionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChatOption entities.
func (c *ChatOptionClient) CreateBulk(builders ...*ChatOptionCreate) *ChatOptionCreateBulk {
	return &ChatOptionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChatOption.
func (c *ChatOptionClient) Update() *ChatOptionUpdate {
	mutation := newChatOptionMutation(c.config, OpUpdate)
	return &ChatOptionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatOptionClient) UpdateOne(co *ChatOption) *ChatOptionUpdateOne {
	mutation := newChatOptionMutation(c.config, OpUpdateOne, withChatOption(co))
	return &ChatOptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatOptionClient) UpdateOneID(id int64) *ChatOptionUpdateOne {
	mutation := newChatOptionMutation(c.config, OpUpdateOne, withChatOptionID(id))
	return &ChatOptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChatOption.
func (c *ChatOptionClient) Delete() *ChatOptionDelete {
	mutation := newChatOptionMutation(c.config, OpDelete)
	return &ChatOptionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChatOptionClient) DeleteOne(co *ChatOption) *ChatOptionDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChatOptionClient) DeleteOneID(id int64) *ChatOptionDeleteOne {
	builder := c.Delete().Where(chatoption.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne

	return &ChatOptionDeleteOne{builder}
}

// Query returns a query builder for ChatOption.
func (c *ChatOptionClient) Query() *ChatOptionQuery {
	return &ChatOptionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChatOption},
		inters: c.Interceptors(),
	}
}

// Get returns a ChatOption entity by its id.
func (c *ChatOptionClient) Get(ctx context.Context, id int64) (*ChatOption, error) {
	return c.Query().Where(chatoption.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatOptionClient) GetX(ctx context.Context, id int64) *ChatOption {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	return obj
}

// Hooks returns the client hooks.
func (c *ChatOptionClient) Hooks() []Hook {
	return c.hooks.ChatOption
}

// Interceptors returns the client interceptors.
func (c *ChatOptionClient) Interceptors() []Interceptor {
	return c.inters.ChatOption
}

func (c *ChatOptionClient) mutate(ctx context.Context, m *ChatOptionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChatOptionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChatOptionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChatOptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChatOptionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown ChatOption mutation op: %q", m.Op())
	}
}

// GameAccountClient is a client for the GameAccount schema.
type GameAccountClient struct {
	config
}

// NewGameAccountClient returns a client for the GameAccount from the given config.
func NewGameAccountClient(c config) *GameAccountClient {
	return &GameAccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `gameaccount.Hooks(f(g(h())))`.
func (c *GameAccountClient) Use(hooks ...Hook) {
	c.hooks.GameAccount = append(c.hooks.GameAccount, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `gameaccount.Intercept(f(g(h())))`.
func (c *GameAccountClient) Intercept(interceptors ...Interceptor) {
	c.inters.GameAccount = append(c.inters.GameAccount, interceptors...)
}

// Create returns a builder for creating a GameAccount entity.
func (c *GameAccountClient) Create() *GameAccountCreate {
	mutation := newGameAccountMutation(c.config, OpCreate)
	return &GameAccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GameAccount entities.
func (c *GameAccountClient) CreateBulk(builders ...*GameAccountCreate) *GameAccountCreateBulk {
	return &GameAccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GameAccount.
func (c *GameAccountClient) Update() *GameAccountUpdate {
	mutation := newGameAccountMutation(c.config, OpUpdate)
	return &GameAccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GameAccountClient) UpdateOne(ga *GameAccount) *GameAccountUpdateOne {
	mutation := newGameAccountMutation(c.config, OpUpdateOne, withGameAccount(ga))
	return &GameAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GameAccountClient) UpdateOneID(id int64) *GameAccountUpdateOne {
	mutation := newGameAccountMutation(c.config, OpUpdateOne, withGameAccountID(id))
	return &GameAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GameAccount.
func (c *GameAccountClient) Delete() *GameAccountDelete {
	mutation := newGameAccountMutation(c.config, OpDelete)
	return &GameAccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GameAccountClient) DeleteOne(ga *GameAccount) *GameAccountDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GameAccountClient) DeleteOneID(id int64) *GameAccountDeleteOne {
	builder := c.Delete().Where(gameaccount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne

	return &GameAccountDeleteOne{builder}
}

// Query returns a query builder for GameAccount.
func (c *GameAccountClient) Query() *GameAccountQuery {
	return &GameAccountQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGameAccount},
		inters: c.Interceptors(),
	}
}

// Get returns a GameAccount entity by its id.
func (c *GameAccountClient) Get(ctx context.Context, id int64) (*GameAccount, error) {
	return c.Query().Where(gameaccount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GameAccountClient) GetX(ctx context.Context, id int64) *GameAccount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	return obj
}

// Hooks returns the client hooks.
func (c *GameAccountClient) Hooks() []Hook {
	return c.hooks.GameAccount
}

// Interceptors returns the client interceptors.
func (c *GameAccountClient) Interceptors() []Interceptor {
	return c.inters.GameAccount
}

func (c *GameAccountClient) mutate(ctx context.Context, m *GameAccountMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GameAccountCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GameAccountUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GameAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GameAccountDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown GameAccount mutation op: %q", m.Op())
	}
}

// GameRoleClient is a client for the GameRole schema.
type GameRoleClient struct {
	config
}

// NewGameRoleClient returns a client for the GameRole from the given config.
func NewGameRoleClient(c config) *GameRoleClient {
	return &GameRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `gamerole.Hooks(f(g(h())))`.
func (c *GameRoleClient) Use(hooks ...Hook) {
	c.hooks.GameRole = append(c.hooks.GameRole, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `gamerole.Intercept(f(g(h())))`.
func (c *GameRoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.GameRole = append(c.inters.GameRole, interceptors...)
}

// Create returns a builder for creating a GameRole entity.
func (c *GameRoleClient) Create() *GameRoleCreate {
	mutation := newGameRoleMutation(c.config, OpCreate)
	return &GameRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GameRole entities.
func (c *GameRoleClient) CreateBulk(builders ...*GameRoleCreate) *GameRoleCreateBulk {
	return &GameRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GameRole.
func (c *GameRoleClient) Update() *GameRoleUpdate {
	mutation := newGameRoleMutation(c.config, OpUpdate)
	return &GameRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GameRoleClient) UpdateOne(gr *GameRole) *GameRoleUpdateOne {
	mutation := newGameRoleMutation(c.config, OpUpdateOne, withGameRole(gr))
	return &GameRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GameRoleClient) UpdateOneID(id int64) *GameRoleUpdateOne {
	mutation := newGameRoleMutation(c.config, OpUpdateOne, withGameRoleID(id))
	return &GameRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GameRole.
func (c *GameRoleClient) Delete() *GameRoleDelete {
	mutation := newGameRoleMutation(c.config, OpDelete)
	return &GameRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GameRoleClient) DeleteOne(gr *GameRole) *GameRoleDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GameRoleClient) DeleteOneID(id int64) *GameRoleDeleteOne {
	builder := c.Delete().Where(gamerole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne

	return &GameRoleDeleteOne{builder}
}

// Query returns a query builder for GameRole.
func (c *GameRoleClient) Query() *GameRoleQuery {
	return &GameRoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGameRole},
		inters: c.Interceptors(),
	}
}

// Get returns a GameRole entity by its id.
func (c *GameRoleClient) Get(ctx context.Context, id int64) (*GameRole, error) {
	return c.Query().Where(gamerole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GameRoleClient) GetX(ctx context.Context, id int64) *GameRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	return obj
}

// Hooks returns the client hooks.
func (c *GameRoleClient) Hooks() []Hook {
	return c.hooks.GameRole
}

// Interceptors returns the client interceptors.
func (c *GameRoleClient) Interceptors() []Interceptor {
	return c.inters.GameRole
}

func (c *GameRoleClient) mutate(ctx context.Context, m *GameRoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GameRoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GameRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GameRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GameRoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown GameRole mutation op: %q", m.Op())
	}
}

// GameRoleAttributeClient is a client for the GameRoleAttribute schema.
type GameRoleAttributeClient struct {
	config
}

// NewGameRoleAttributeClient returns a client for the GameRoleAttribute from the given config.
func NewGameRoleAttributeClient(c config) *GameRoleAttributeClient {
	return &GameRoleAttributeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `gameroleattribute.Hooks(f(g(h())))`.
func (c *GameRoleAttributeClient) Use(hooks ...Hook) {
	c.hooks.GameRoleAttribute = append(c.hooks.GameRoleAttribute, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `gameroleattribute.Intercept(f(g(h())))`.
func (c *GameRoleAttributeClient) Intercept(interceptors ...Interceptor) {
	c.inters.GameRoleAttribute = append(c.inters.GameRoleAttribute, interceptors...)
}

// Create returns a builder for creating a GameRoleAttribute entity.
func (c *GameRoleAttributeClient) Create() *GameRoleAttributeCreate {
	mutation := newGameRoleAttributeMutation(c.config, OpCreate)
	return &GameRoleAttributeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GameRoleAttribute entities.
func (c *GameRoleAttributeClient) CreateBulk(builders ...*GameRoleAttributeCreate) *GameRoleAttributeCreateBulk {
	return &GameRoleAttributeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GameRoleAttribute.
func (c *GameRoleAttributeClient) Update() *GameRoleAttributeUpdate {
	mutation := newGameRoleAttributeMutation(c.config, OpUpdate)
	return &GameRoleAttributeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GameRoleAttributeClient) UpdateOne(gra *GameRoleAttribute) *GameRoleAttributeUpdateOne {
	mutation := newGameRoleAttributeMutation(c.config, OpUpdateOne, withGameRoleAttribute(gra))
	return &GameRoleAttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GameRoleAttributeClient) UpdateOneID(id int64) *GameRoleAttributeUpdateOne {
	mutation := newGameRoleAttributeMutation(c.config, OpUpdateOne, withGameRoleAttributeID(id))
	return &GameRoleAttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GameRoleAttribute.
func (c *GameRoleAttributeClient) Delete() *GameRoleAttributeDelete {
	mutation := newGameRoleAttributeMutation(c.config, OpDelete)
	return &GameRoleAttributeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GameRoleAttributeClient) DeleteOne(gra *GameRoleAttribute) *GameRoleAttributeDeleteOne {
	return c.DeleteOneID(gra.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GameRoleAttributeClient) DeleteOneID(id int64) *GameRoleAttributeDeleteOne {
	builder := c.Delete().Where(gameroleattribute.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne

	return &GameRoleAttributeDeleteOne{builder}
}

// Query returns a query builder for GameRoleAttribute.
func (c *GameRoleAttributeClient) Query() *GameRoleAttributeQuery {
	return &GameRoleAttributeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGameRoleAttribute},
		inters: c.Interceptors(),
	}
}

// Get returns a GameRoleAttribute entity by its id.
func (c *GameRoleAttributeClient) Get(ctx context.Context, id int64) (*GameRoleAttribute, error) {
	return c.Query().Where(gameroleattribute.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GameRoleAttributeClient) GetX(ctx context.Context, id int64) *GameRoleAttribute {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	return obj
}

// Hooks returns the client hooks.
func (c *GameRoleAttributeClient) Hooks() []Hook {
	return c.hooks.GameRoleAttribute
}

// Interceptors returns the client interceptors.
func (c *GameRoleAttributeClient) Interceptors() []Interceptor {
	return c.inters.GameRoleAttribute
}

func (c *GameRoleAttributeClient) mutate(ctx context.Context, m *GameRoleAttributeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GameRoleAttributeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GameRoleAttributeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GameRoleAttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GameRoleAttributeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown GameRoleAttribute mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int64) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int64) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne

	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int64) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int64) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown User mutation op: %q", m.Op())
	}
}
