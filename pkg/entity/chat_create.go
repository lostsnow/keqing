// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lostsnow/keqing/pkg/entity/chat"
)

// ChatCreate is the builder for creating a Chat entity.
type ChatCreate struct {
	config
	mutation *ChatMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetChatID sets the "chat_id" field.
func (cc *ChatCreate) SetChatID(i int64) *ChatCreate {
	cc.mutation.SetChatID(i)
	return cc
}

// SetType sets the "type" field.
func (cc *ChatCreate) SetType(s string) *ChatCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cc *ChatCreate) SetNillableType(s *string) *ChatCreate {
	if s != nil {
		cc.SetType(*s)
	}

	return cc
}

// SetIsForum sets the "is_forum" field.
func (cc *ChatCreate) SetIsForum(b bool) *ChatCreate {
	cc.mutation.SetIsForum(b)
	return cc
}

// SetNillableIsForum sets the "is_forum" field if the given value is not nil.
func (cc *ChatCreate) SetNillableIsForum(b *bool) *ChatCreate {
	if b != nil {
		cc.SetIsForum(*b)
	}

	return cc
}

// SetTitle sets the "title" field.
func (cc *ChatCreate) SetTitle(s string) *ChatCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cc *ChatCreate) SetNillableTitle(s *string) *ChatCreate {
	if s != nil {
		cc.SetTitle(*s)
	}

	return cc
}

// SetUserName sets the "user_name" field.
func (cc *ChatCreate) SetUserName(s string) *ChatCreate {
	cc.mutation.SetUserName(s)
	return cc
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (cc *ChatCreate) SetNillableUserName(s *string) *ChatCreate {
	if s != nil {
		cc.SetUserName(*s)
	}

	return cc
}

// SetFirstName sets the "first_name" field.
func (cc *ChatCreate) SetFirstName(s string) *ChatCreate {
	cc.mutation.SetFirstName(s)
	return cc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (cc *ChatCreate) SetNillableFirstName(s *string) *ChatCreate {
	if s != nil {
		cc.SetFirstName(*s)
	}

	return cc
}

// SetLastName sets the "last_name" field.
func (cc *ChatCreate) SetLastName(s string) *ChatCreate {
	cc.mutation.SetLastName(s)
	return cc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (cc *ChatCreate) SetNillableLastName(s *string) *ChatCreate {
	if s != nil {
		cc.SetLastName(*s)
	}

	return cc
}

// SetDescription sets the "description" field.
func (cc *ChatCreate) SetDescription(s string) *ChatCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *ChatCreate) SetNillableDescription(s *string) *ChatCreate {
	if s != nil {
		cc.SetDescription(*s)
	}

	return cc
}

// SetCreateAt sets the "create_at" field.
func (cc *ChatCreate) SetCreateAt(t time.Time) *ChatCreate {
	cc.mutation.SetCreateAt(t)
	return cc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cc *ChatCreate) SetNillableCreateAt(t *time.Time) *ChatCreate {
	if t != nil {
		cc.SetCreateAt(*t)
	}

	return cc
}

// SetUpdateAt sets the "update_at" field.
func (cc *ChatCreate) SetUpdateAt(t time.Time) *ChatCreate {
	cc.mutation.SetUpdateAt(t)
	return cc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (cc *ChatCreate) SetNillableUpdateAt(t *time.Time) *ChatCreate {
	if t != nil {
		cc.SetUpdateAt(*t)
	}

	return cc
}

// SetID sets the "id" field.
func (cc *ChatCreate) SetID(i int64) *ChatCreate {
	cc.mutation.SetID(i)
	return cc
}

// Mutation returns the ChatMutation object of the builder.
func (cc *ChatCreate) Mutation() *ChatMutation {
	return cc.mutation
}

// Save creates the Chat in the database.
func (cc *ChatCreate) Save(ctx context.Context) (*Chat, error) {
	cc.defaults()
	return withHooks[*Chat, ChatMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChatCreate) SaveX(ctx context.Context) *Chat {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}

	return v
}

// Exec executes the query.
func (cc *ChatCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChatCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChatCreate) defaults() {
	if _, ok := cc.mutation.GetType(); !ok {
		v := chat.DefaultType
		cc.mutation.SetType(v)
	}

	if _, ok := cc.mutation.IsForum(); !ok {
		v := chat.DefaultIsForum
		cc.mutation.SetIsForum(v)
	}

	if _, ok := cc.mutation.Title(); !ok {
		v := chat.DefaultTitle
		cc.mutation.SetTitle(v)
	}

	if _, ok := cc.mutation.UserName(); !ok {
		v := chat.DefaultUserName
		cc.mutation.SetUserName(v)
	}

	if _, ok := cc.mutation.FirstName(); !ok {
		v := chat.DefaultFirstName
		cc.mutation.SetFirstName(v)
	}

	if _, ok := cc.mutation.LastName(); !ok {
		v := chat.DefaultLastName
		cc.mutation.SetLastName(v)
	}

	if _, ok := cc.mutation.Description(); !ok {
		v := chat.DefaultDescription
		cc.mutation.SetDescription(v)
	}

	if _, ok := cc.mutation.CreateAt(); !ok {
		v := chat.DefaultCreateAt()
		cc.mutation.SetCreateAt(v)
	}

	if _, ok := cc.mutation.UpdateAt(); !ok {
		v := chat.DefaultUpdateAt()
		cc.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChatCreate) check() error {
	if _, ok := cc.mutation.ChatID(); !ok {
		return &ValidationError{Name: "chat_id", err: errors.New(`entity: missing required field "Chat.chat_id"`)}
	}

	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`entity: missing required field "Chat.type"`)}
	}

	if _, ok := cc.mutation.IsForum(); !ok {
		return &ValidationError{Name: "is_forum", err: errors.New(`entity: missing required field "Chat.is_forum"`)}
	}

	if _, ok := cc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`entity: missing required field "Chat.title"`)}
	}

	if _, ok := cc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`entity: missing required field "Chat.user_name"`)}
	}

	if _, ok := cc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`entity: missing required field "Chat.first_name"`)}
	}

	if _, ok := cc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`entity: missing required field "Chat.last_name"`)}
	}

	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`entity: missing required field "Chat.description"`)}
	}

	if _, ok := cc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`entity: missing required field "Chat.create_at"`)}
	}

	if _, ok := cc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`entity: missing required field "Chat.update_at"`)}
	}

	return nil
}

func (cc *ChatCreate) sqlSave(ctx context.Context) (*Chat, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}

	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}

		return nil, err
	}

	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}

	cc.mutation.id = &_node.ID
	cc.mutation.done = true

	return _node, nil
}

func (cc *ChatCreate) createSpec() (*Chat, *sqlgraph.CreateSpec) {
	var (
		_node = &Chat{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(chat.Table, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt64))
	)

	_spec.OnConflict = cc.conflict

	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}

	if value, ok := cc.mutation.ChatID(); ok {
		_spec.SetField(chat.FieldChatID, field.TypeInt64, value)
		_node.ChatID = value
	}

	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(chat.FieldType, field.TypeString, value)
		_node.Type = value
	}

	if value, ok := cc.mutation.IsForum(); ok {
		_spec.SetField(chat.FieldIsForum, field.TypeBool, value)
		_node.IsForum = value
	}

	if value, ok := cc.mutation.Title(); ok {
		_spec.SetField(chat.FieldTitle, field.TypeString, value)
		_node.Title = value
	}

	if value, ok := cc.mutation.UserName(); ok {
		_spec.SetField(chat.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}

	if value, ok := cc.mutation.FirstName(); ok {
		_spec.SetField(chat.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}

	if value, ok := cc.mutation.LastName(); ok {
		_spec.SetField(chat.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}

	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(chat.FieldDescription, field.TypeString, value)
		_node.Description = value
	}

	if value, ok := cc.mutation.CreateAt(); ok {
		_spec.SetField(chat.FieldCreateAt, field.TypeTime, value)
		_node.CreateAt = value
	}

	if value, ok := cc.mutation.UpdateAt(); ok {
		_spec.SetField(chat.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}

	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Chat.Create().
//		SetChatID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChatUpsert) {
//			SetChatID(v+v).
//		}).
//		Exec(ctx)
func (cc *ChatCreate) OnConflict(opts ...sql.ConflictOption) *ChatUpsertOne {
	cc.conflict = opts

	return &ChatUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Chat.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *ChatCreate) OnConflictColumns(columns ...string) *ChatUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))

	return &ChatUpsertOne{
		create: cc,
	}
}

type (
	// ChatUpsertOne is the builder for "upsert"-ing
	//  one Chat node.
	ChatUpsertOne struct {
		create *ChatCreate
	}

	// ChatUpsert is the "OnConflict" setter.
	ChatUpsert struct {
		*sql.UpdateSet
	}
)

// SetChatID sets the "chat_id" field.
func (u *ChatUpsert) SetChatID(v int64) *ChatUpsert {
	u.Set(chat.FieldChatID, v)
	return u
}

// UpdateChatID sets the "chat_id" field to the value that was provided on create.
func (u *ChatUpsert) UpdateChatID() *ChatUpsert {
	u.SetExcluded(chat.FieldChatID)
	return u
}

// AddChatID adds v to the "chat_id" field.
func (u *ChatUpsert) AddChatID(v int64) *ChatUpsert {
	u.Add(chat.FieldChatID, v)
	return u
}

// SetType sets the "type" field.
func (u *ChatUpsert) SetType(v string) *ChatUpsert {
	u.Set(chat.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ChatUpsert) UpdateType() *ChatUpsert {
	u.SetExcluded(chat.FieldType)
	return u
}

// SetIsForum sets the "is_forum" field.
func (u *ChatUpsert) SetIsForum(v bool) *ChatUpsert {
	u.Set(chat.FieldIsForum, v)
	return u
}

// UpdateIsForum sets the "is_forum" field to the value that was provided on create.
func (u *ChatUpsert) UpdateIsForum() *ChatUpsert {
	u.SetExcluded(chat.FieldIsForum)
	return u
}

// SetTitle sets the "title" field.
func (u *ChatUpsert) SetTitle(v string) *ChatUpsert {
	u.Set(chat.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ChatUpsert) UpdateTitle() *ChatUpsert {
	u.SetExcluded(chat.FieldTitle)
	return u
}

// SetUserName sets the "user_name" field.
func (u *ChatUpsert) SetUserName(v string) *ChatUpsert {
	u.Set(chat.FieldUserName, v)
	return u
}

// UpdateUserName sets the "user_name" field to the value that was provided on create.
func (u *ChatUpsert) UpdateUserName() *ChatUpsert {
	u.SetExcluded(chat.FieldUserName)
	return u
}

// SetFirstName sets the "first_name" field.
func (u *ChatUpsert) SetFirstName(v string) *ChatUpsert {
	u.Set(chat.FieldFirstName, v)
	return u
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *ChatUpsert) UpdateFirstName() *ChatUpsert {
	u.SetExcluded(chat.FieldFirstName)
	return u
}

// SetLastName sets the "last_name" field.
func (u *ChatUpsert) SetLastName(v string) *ChatUpsert {
	u.Set(chat.FieldLastName, v)
	return u
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *ChatUpsert) UpdateLastName() *ChatUpsert {
	u.SetExcluded(chat.FieldLastName)
	return u
}

// SetDescription sets the "description" field.
func (u *ChatUpsert) SetDescription(v string) *ChatUpsert {
	u.Set(chat.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ChatUpsert) UpdateDescription() *ChatUpsert {
	u.SetExcluded(chat.FieldDescription)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *ChatUpsert) SetUpdateAt(v time.Time) *ChatUpsert {
	u.Set(chat.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ChatUpsert) UpdateUpdateAt() *ChatUpsert {
	u.SetExcluded(chat.FieldUpdateAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Chat.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(chat.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChatUpsertOne) UpdateNewValues() *ChatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(chat.FieldID)
		}

		if _, exists := u.create.mutation.CreateAt(); exists {
			s.SetIgnore(chat.FieldCreateAt)
		}
	}))

	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Chat.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ChatUpsertOne) Ignore() *ChatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChatUpsertOne) DoNothing() *ChatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChatCreate.OnConflict
// documentation for more info.
func (u *ChatUpsertOne) Update(set func(*ChatUpsert)) *ChatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChatUpsert{UpdateSet: update})
	}))

	return u
}

// SetChatID sets the "chat_id" field.
func (u *ChatUpsertOne) SetChatID(v int64) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetChatID(v)
	})
}

// AddChatID adds v to the "chat_id" field.
func (u *ChatUpsertOne) AddChatID(v int64) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.AddChatID(v)
	})
}

// UpdateChatID sets the "chat_id" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateChatID() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateChatID()
	})
}

// SetType sets the "type" field.
func (u *ChatUpsertOne) SetType(v string) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateType() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateType()
	})
}

// SetIsForum sets the "is_forum" field.
func (u *ChatUpsertOne) SetIsForum(v bool) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetIsForum(v)
	})
}

// UpdateIsForum sets the "is_forum" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateIsForum() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateIsForum()
	})
}

// SetTitle sets the "title" field.
func (u *ChatUpsertOne) SetTitle(v string) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateTitle() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateTitle()
	})
}

// SetUserName sets the "user_name" field.
func (u *ChatUpsertOne) SetUserName(v string) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetUserName(v)
	})
}

// UpdateUserName sets the "user_name" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateUserName() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateUserName()
	})
}

// SetFirstName sets the "first_name" field.
func (u *ChatUpsertOne) SetFirstName(v string) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateFirstName() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *ChatUpsertOne) SetLastName(v string) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateLastName() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateLastName()
	})
}

// SetDescription sets the "description" field.
func (u *ChatUpsertOne) SetDescription(v string) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateDescription() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateDescription()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *ChatUpsertOne) SetUpdateAt(v time.Time) *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ChatUpsertOne) UpdateUpdateAt() *ChatUpsertOne {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateUpdateAt()
	})
}

// Exec executes the query.
func (u *ChatUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entity: missing options for ChatCreate.OnConflict")
	}

	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChatUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ChatUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}

	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ChatUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}

	return id
}

// ChatCreateBulk is the builder for creating many Chat entities in bulk.
type ChatCreateBulk struct {
	config
	builders []*ChatCreate
	conflict []sql.ConflictOption
}

// Save creates the Chat entities in the database.
func (ccb *ChatCreateBulk) Save(ctx context.Context) ([]*Chat, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chat, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))

	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()

			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChatMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})

			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}

			mutators[i] = mut
		}(i, ctx)
	}

	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}

	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChatCreateBulk) SaveX(ctx context.Context) []*Chat {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}

	return v
}

// Exec executes the query.
func (ccb *ChatCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChatCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Chat.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChatUpsert) {
//			SetChatID(v+v).
//		}).
//		Exec(ctx)
func (ccb *ChatCreateBulk) OnConflict(opts ...sql.ConflictOption) *ChatUpsertBulk {
	ccb.conflict = opts

	return &ChatUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Chat.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *ChatCreateBulk) OnConflictColumns(columns ...string) *ChatUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))

	return &ChatUpsertBulk{
		create: ccb,
	}
}

// ChatUpsertBulk is the builder for "upsert"-ing
// a bulk of Chat nodes.
type ChatUpsertBulk struct {
	create *ChatCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Chat.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(chat.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChatUpsertBulk) UpdateNewValues() *ChatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(chat.FieldID)
			}

			if _, exists := b.mutation.CreateAt(); exists {
				s.SetIgnore(chat.FieldCreateAt)
			}
		}
	}))

	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Chat.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ChatUpsertBulk) Ignore() *ChatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChatUpsertBulk) DoNothing() *ChatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChatCreateBulk.OnConflict
// documentation for more info.
func (u *ChatUpsertBulk) Update(set func(*ChatUpsert)) *ChatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChatUpsert{UpdateSet: update})
	}))

	return u
}

// SetChatID sets the "chat_id" field.
func (u *ChatUpsertBulk) SetChatID(v int64) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetChatID(v)
	})
}

// AddChatID adds v to the "chat_id" field.
func (u *ChatUpsertBulk) AddChatID(v int64) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.AddChatID(v)
	})
}

// UpdateChatID sets the "chat_id" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateChatID() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateChatID()
	})
}

// SetType sets the "type" field.
func (u *ChatUpsertBulk) SetType(v string) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateType() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateType()
	})
}

// SetIsForum sets the "is_forum" field.
func (u *ChatUpsertBulk) SetIsForum(v bool) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetIsForum(v)
	})
}

// UpdateIsForum sets the "is_forum" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateIsForum() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateIsForum()
	})
}

// SetTitle sets the "title" field.
func (u *ChatUpsertBulk) SetTitle(v string) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateTitle() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateTitle()
	})
}

// SetUserName sets the "user_name" field.
func (u *ChatUpsertBulk) SetUserName(v string) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetUserName(v)
	})
}

// UpdateUserName sets the "user_name" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateUserName() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateUserName()
	})
}

// SetFirstName sets the "first_name" field.
func (u *ChatUpsertBulk) SetFirstName(v string) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateFirstName() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *ChatUpsertBulk) SetLastName(v string) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateLastName() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateLastName()
	})
}

// SetDescription sets the "description" field.
func (u *ChatUpsertBulk) SetDescription(v string) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateDescription() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateDescription()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *ChatUpsertBulk) SetUpdateAt(v time.Time) *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ChatUpsertBulk) UpdateUpdateAt() *ChatUpsertBulk {
	return u.Update(func(s *ChatUpsert) {
		s.UpdateUpdateAt()
	})
}

// Exec executes the query.
func (u *ChatUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entity: OnConflict was set for builder %d. Set it on the ChatCreateBulk instead", i)
		}
	}

	if len(u.create.conflict) == 0 {
		return errors.New("entity: missing options for ChatCreateBulk.OnConflict")
	}

	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChatUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
