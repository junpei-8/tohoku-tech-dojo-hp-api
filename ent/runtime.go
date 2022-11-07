// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/news"
	"app/ent/schema"
	"app/ent/tag"
	"app/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	newsFields := schema.News{}.Fields()
	_ = newsFields
	// newsDescUID is the schema descriptor for uid field.
	newsDescUID := newsFields[0].Descriptor()
	// news.DefaultUID holds the default value on creation for the uid field.
	news.DefaultUID = newsDescUID.Default.(func() int)
	// newsDescTitle is the schema descriptor for title field.
	newsDescTitle := newsFields[1].Descriptor()
	// news.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	news.TitleValidator = newsDescTitle.Validators[0].(func(string) error)
	// newsDescHTML is the schema descriptor for html field.
	newsDescHTML := newsFields[2].Descriptor()
	// news.HTMLValidator is a validator for the "html" field. It is called by the builders before save.
	news.HTMLValidator = newsDescHTML.Validators[0].(func(string) error)
	// newsDescMarkdown is the schema descriptor for markdown field.
	newsDescMarkdown := newsFields[3].Descriptor()
	// news.MarkdownValidator is a validator for the "markdown" field. It is called by the builders before save.
	news.MarkdownValidator = newsDescMarkdown.Validators[0].(func(string) error)
	// newsDescCreatedAt is the schema descriptor for created_at field.
	newsDescCreatedAt := newsFields[5].Descriptor()
	// news.DefaultCreatedAt holds the default value on creation for the created_at field.
	news.DefaultCreatedAt = newsDescCreatedAt.Default.(func() time.Time)
	// newsDescUpdatedAt is the schema descriptor for updated_at field.
	newsDescUpdatedAt := newsFields[6].Descriptor()
	// news.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	news.UpdateDefaultUpdatedAt = newsDescUpdatedAt.UpdateDefault.(func() time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescTitle is the schema descriptor for title field.
	tagDescTitle := tagFields[0].Descriptor()
	// tag.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	tag.TitleValidator = tagDescTitle.Validators[0].(func(string) error)
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagFields[1].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}