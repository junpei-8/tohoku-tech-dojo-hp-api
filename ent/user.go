// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// CreatedNewsList holds the value of the created_news_list edge.
	CreatedNewsList []*News `json:"created_news_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedCreatedNewsList map[string][]*News
}

// CreatedNewsListOrErr returns the CreatedNewsList value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedNewsListOrErr() ([]*News, error) {
	if e.loadedTypes[0] {
		return e.CreatedNewsList, nil
	}
	return nil, &NotLoadedError{edge: "created_news_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldEmail:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = new(time.Time)
				*u.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryCreatedNewsList queries the "created_news_list" edge of the User entity.
func (u *User) QueryCreatedNewsList() *NewsQuery {
	return (&UserClient{config: u.config}).QueryCreatedNewsList(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := u.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedCreatedNewsList returns the CreatedNewsList named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCreatedNewsList(name string) ([]*News, error) {
	if u.Edges.namedCreatedNewsList == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCreatedNewsList[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCreatedNewsList(name string, edges ...*News) {
	if u.Edges.namedCreatedNewsList == nil {
		u.Edges.namedCreatedNewsList = make(map[string][]*News)
	}
	if len(edges) == 0 {
		u.Edges.namedCreatedNewsList[name] = []*News{}
	} else {
		u.Edges.namedCreatedNewsList[name] = append(u.Edges.namedCreatedNewsList[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
