// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/news"
	"app/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// News is the model entity for the News schema.
type News struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID int `json:"uid,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// HTML holds the value of the "html" field.
	HTML string `json:"html,omitempty"`
	// Markdown holds the value of the "markdown" field.
	Markdown string `json:"markdown,omitempty"`
	// CreatorID holds the value of the "creator_id" field.
	CreatorID int `json:"creator_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NewsQuery when eager-loading is set.
	Edges NewsEdges `json:"edges"`
}

// NewsEdges holds the relations/edges for other nodes in the graph.
type NewsEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// NewsTagging holds the value of the news_tagging edge.
	NewsTagging []*NewsTagging `json:"news_tagging,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedTags        map[string][]*Tag
	namedNewsTagging map[string][]*NewsTagging
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NewsEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e NewsEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[1] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// NewsTaggingOrErr returns the NewsTagging value or an error if the edge
// was not loaded in eager-loading.
func (e NewsEdges) NewsTaggingOrErr() ([]*NewsTagging, error) {
	if e.loadedTypes[2] {
		return e.NewsTagging, nil
	}
	return nil, &NotLoadedError{edge: "news_tagging"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*News) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case news.FieldID, news.FieldUID, news.FieldCreatorID:
			values[i] = new(sql.NullInt64)
		case news.FieldTitle, news.FieldHTML, news.FieldMarkdown:
			values[i] = new(sql.NullString)
		case news.FieldCreatedAt, news.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type News", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the News fields.
func (n *News) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case news.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case news.FieldUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				n.UID = int(value.Int64)
			}
		case news.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				n.Title = value.String
			}
		case news.FieldHTML:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html", values[i])
			} else if value.Valid {
				n.HTML = value.String
			}
		case news.FieldMarkdown:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field markdown", values[i])
			} else if value.Valid {
				n.Markdown = value.String
			}
		case news.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				n.CreatorID = int(value.Int64)
			}
		case news.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case news.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = new(time.Time)
				*n.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the News entity.
func (n *News) QueryUser() *UserQuery {
	return (&NewsClient{config: n.config}).QueryUser(n)
}

// QueryTags queries the "tags" edge of the News entity.
func (n *News) QueryTags() *TagQuery {
	return (&NewsClient{config: n.config}).QueryTags(n)
}

// QueryNewsTagging queries the "news_tagging" edge of the News entity.
func (n *News) QueryNewsTagging() *NewsTaggingQuery {
	return (&NewsClient{config: n.config}).QueryNewsTagging(n)
}

// Update returns a builder for updating this News.
// Note that you need to call News.Unwrap() before calling this method if this News
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *News) Update() *NewsUpdateOne {
	return (&NewsClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the News entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *News) Unwrap() *News {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: News is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *News) String() string {
	var builder strings.Builder
	builder.WriteString("News(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("uid=")
	builder.WriteString(fmt.Sprintf("%v", n.UID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(n.Title)
	builder.WriteString(", ")
	builder.WriteString("html=")
	builder.WriteString(n.HTML)
	builder.WriteString(", ")
	builder.WriteString("markdown=")
	builder.WriteString(n.Markdown)
	builder.WriteString(", ")
	builder.WriteString("creator_id=")
	builder.WriteString(fmt.Sprintf("%v", n.CreatorID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := n.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedTags returns the Tags named value or an error if the edge was not
// loaded in eager-loading with this name.
func (n *News) NamedTags(name string) ([]*Tag, error) {
	if n.Edges.namedTags == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := n.Edges.namedTags[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (n *News) appendNamedTags(name string, edges ...*Tag) {
	if n.Edges.namedTags == nil {
		n.Edges.namedTags = make(map[string][]*Tag)
	}
	if len(edges) == 0 {
		n.Edges.namedTags[name] = []*Tag{}
	} else {
		n.Edges.namedTags[name] = append(n.Edges.namedTags[name], edges...)
	}
}

// NamedNewsTagging returns the NewsTagging named value or an error if the edge was not
// loaded in eager-loading with this name.
func (n *News) NamedNewsTagging(name string) ([]*NewsTagging, error) {
	if n.Edges.namedNewsTagging == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := n.Edges.namedNewsTagging[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (n *News) appendNamedNewsTagging(name string, edges ...*NewsTagging) {
	if n.Edges.namedNewsTagging == nil {
		n.Edges.namedNewsTagging = make(map[string][]*NewsTagging)
	}
	if len(edges) == 0 {
		n.Edges.namedNewsTagging[name] = []*NewsTagging{}
	} else {
		n.Edges.namedNewsTagging[name] = append(n.Edges.namedNewsTagging[name], edges...)
	}
}

// NewsSlice is a parsable slice of News.
type NewsSlice []*News

func (n NewsSlice) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
