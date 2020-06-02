// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"github.com/aca/go-restapi-boilerplate/ent/schema"
	"github.com/aca/go-restapi-boilerplate/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[1].Descriptor()
	// user.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	user.UserNameValidator = userDescUserName.Validators[0].(func(string) error)
}
