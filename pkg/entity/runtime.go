// Code generated by ent, DO NOT EDIT.

package entity

import (
	"time"

	"github.com/lostsnow/keqing/db/schema"
	"github.com/lostsnow/keqing/pkg/entity/chat"
	"github.com/lostsnow/keqing/pkg/entity/chatoption"
	"github.com/lostsnow/keqing/pkg/entity/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatFields := schema.Chat{}.Fields()
	_ = chatFields
	// chatDescType is the schema descriptor for type field.
	chatDescType := chatFields[2].Descriptor()
	// chat.DefaultType holds the default value on creation for the type field.
	chat.DefaultType = chatDescType.Default.(string)
	// chatDescIsForum is the schema descriptor for is_forum field.
	chatDescIsForum := chatFields[3].Descriptor()
	// chat.DefaultIsForum holds the default value on creation for the is_forum field.
	chat.DefaultIsForum = chatDescIsForum.Default.(bool)
	// chatDescTitle is the schema descriptor for title field.
	chatDescTitle := chatFields[4].Descriptor()
	// chat.DefaultTitle holds the default value on creation for the title field.
	chat.DefaultTitle = chatDescTitle.Default.(string)
	// chatDescUserName is the schema descriptor for user_name field.
	chatDescUserName := chatFields[5].Descriptor()
	// chat.DefaultUserName holds the default value on creation for the user_name field.
	chat.DefaultUserName = chatDescUserName.Default.(string)
	// chatDescFirstName is the schema descriptor for first_name field.
	chatDescFirstName := chatFields[6].Descriptor()
	// chat.DefaultFirstName holds the default value on creation for the first_name field.
	chat.DefaultFirstName = chatDescFirstName.Default.(string)
	// chatDescLastName is the schema descriptor for last_name field.
	chatDescLastName := chatFields[7].Descriptor()
	// chat.DefaultLastName holds the default value on creation for the last_name field.
	chat.DefaultLastName = chatDescLastName.Default.(string)
	// chatDescDescription is the schema descriptor for description field.
	chatDescDescription := chatFields[8].Descriptor()
	// chat.DefaultDescription holds the default value on creation for the description field.
	chat.DefaultDescription = chatDescDescription.Default.(string)
	// chatDescCreateAt is the schema descriptor for create_at field.
	chatDescCreateAt := chatFields[9].Descriptor()
	// chat.DefaultCreateAt holds the default value on creation for the create_at field.
	chat.DefaultCreateAt = chatDescCreateAt.Default.(func() time.Time)
	// chatDescUpdateAt is the schema descriptor for update_at field.
	chatDescUpdateAt := chatFields[10].Descriptor()
	// chat.DefaultUpdateAt holds the default value on creation for the update_at field.
	chat.DefaultUpdateAt = chatDescUpdateAt.Default.(func() time.Time)
	// chat.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	chat.UpdateDefaultUpdateAt = chatDescUpdateAt.UpdateDefault.(func() time.Time)
	chatoptionFields := schema.ChatOption{}.Fields()
	_ = chatoptionFields
	// chatoptionDescKey is the schema descriptor for key field.
	chatoptionDescKey := chatoptionFields[2].Descriptor()
	// chatoption.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	chatoption.KeyValidator = chatoptionDescKey.Validators[0].(func(string) error)
	// chatoptionDescValue is the schema descriptor for value field.
	chatoptionDescValue := chatoptionFields[3].Descriptor()
	// chatoption.DefaultValue holds the default value on creation for the value field.
	chatoption.DefaultValue = chatoptionDescValue.Default.(string)
	// chatoptionDescCreateAt is the schema descriptor for create_at field.
	chatoptionDescCreateAt := chatoptionFields[4].Descriptor()
	// chatoption.DefaultCreateAt holds the default value on creation for the create_at field.
	chatoption.DefaultCreateAt = chatoptionDescCreateAt.Default.(func() time.Time)
	// chatoptionDescUpdateAt is the schema descriptor for update_at field.
	chatoptionDescUpdateAt := chatoptionFields[5].Descriptor()
	// chatoption.DefaultUpdateAt holds the default value on creation for the update_at field.
	chatoption.DefaultUpdateAt = chatoptionDescUpdateAt.Default.(func() time.Time)
	// chatoption.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	chatoption.UpdateDefaultUpdateAt = chatoptionDescUpdateAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescIsBot is the schema descriptor for is_bot field.
	userDescIsBot := userFields[2].Descriptor()
	// user.DefaultIsBot holds the default value on creation for the is_bot field.
	user.DefaultIsBot = userDescIsBot.Default.(bool)
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[3].Descriptor()
	// user.DefaultUserName holds the default value on creation for the user_name field.
	user.DefaultUserName = userDescUserName.Default.(string)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[4].Descriptor()
	// user.DefaultFirstName holds the default value on creation for the first_name field.
	user.DefaultFirstName = userDescFirstName.Default.(string)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[5].Descriptor()
	// user.DefaultLastName holds the default value on creation for the last_name field.
	user.DefaultLastName = userDescLastName.Default.(string)
	// userDescCreateAt is the schema descriptor for create_at field.
	userDescCreateAt := userFields[6].Descriptor()
	// user.DefaultCreateAt holds the default value on creation for the create_at field.
	user.DefaultCreateAt = userDescCreateAt.Default.(func() time.Time)
	// userDescUpdateAt is the schema descriptor for update_at field.
	userDescUpdateAt := userFields[7].Descriptor()
	// user.DefaultUpdateAt holds the default value on creation for the update_at field.
	user.DefaultUpdateAt = userDescUpdateAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	user.UpdateDefaultUpdateAt = userDescUpdateAt.UpdateDefault.(func() time.Time)
}
