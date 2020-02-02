// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/aca/go-restapi-boilerplate/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	user_id   *string
	user_name *string
}

// SetUserID sets the user_id field.
func (uc *UserCreate) SetUserID(s string) *UserCreate {
	uc.user_id = &s
	return uc
}

// SetUserName sets the user_name field.
func (uc *UserCreate) SetUserName(s string) *UserCreate {
	uc.user_name = &s
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.user_id == nil {
		return nil, errors.New("ent: missing required field \"user_id\"")
	}
	if uc.user_name == nil {
		return nil, errors.New("ent: missing required field \"user_name\"")
	}
	if err := user.UserNameValidator(*uc.user_name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"user_name\": %v", err)
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value := uc.user_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldUserID,
		})
		u.UserID = *value
	}
	if value := uc.user_name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldUserName,
		})
		u.UserName = *value
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
