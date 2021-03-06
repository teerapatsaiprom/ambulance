// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/ambu/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserEmail holds the value of the "user_email" field.
	UserEmail string `json:"user_email,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// UserPassword holds the value of the "user_password" field.
	UserPassword string `json:"user_password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Predicament holds the value of the predicament edge.
	Predicament []*Predicament
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PredicamentOrErr returns the Predicament value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PredicamentOrErr() ([]*Predicament, error) {
	if e.loadedTypes[0] {
		return e.Predicament, nil
	}
	return nil, &NotLoadedError{edge: "predicament"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // user_email
		&sql.NullString{}, // user_name
		&sql.NullString{}, // user_password
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user_email", values[0])
	} else if value.Valid {
		u.UserEmail = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user_name", values[1])
	} else if value.Valid {
		u.UserName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user_password", values[2])
	} else if value.Valid {
		u.UserPassword = value.String
	}
	return nil
}

// QueryPredicament queries the predicament edge of the User.
func (u *User) QueryPredicament() *PredicamentQuery {
	return (&UserClient{config: u.config}).QueryPredicament(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", user_email=")
	builder.WriteString(u.UserEmail)
	builder.WriteString(", user_name=")
	builder.WriteString(u.UserName)
	builder.WriteString(", user_password=")
	builder.WriteString(u.UserPassword)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
