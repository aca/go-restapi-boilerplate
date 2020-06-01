// Code generated (@generated) by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID       = "id"      // FieldUserID holds the string denoting the user_id vertex property in the database.
	FieldUserID   = "user_id" // FieldUserName holds the string denoting the user_name vertex property in the database.
	FieldUserName = "user_name"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldUserName,
}

var (
	// UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	UserNameValidator func(string) error
)
