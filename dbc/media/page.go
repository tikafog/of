// Code generated by ent, DO NOT EDIT.

package media

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/page"
)

// Page is the model entity for the Page schema.
type Page struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedUnix holds the value of the "created_unix" field.
	CreatedUnix int64 `json:"created_unix,omitempty"`
	// UpdatedUnix holds the value of the "updated_unix" field.
	UpdatedUnix int64 `json:"updated_unix,omitempty"`
	// DeletedUnix holds the value of the "deleted_unix" field.
	DeletedUnix int64 `json:"deleted_unix,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID uuid.UUID `json:"parent_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// FeaturedIndex holds the value of the "featured_index" field.
	FeaturedIndex int8 `json:"featured_index,omitempty"`
	// FeaturedContent holds the value of the "featured_content" field.
	FeaturedContent string `json:"featured_content,omitempty"`
	// Recommend holds the value of the "recommend" field.
	Recommend int64 `json:"recommend,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PageQuery when eager-loading is set.
	Edges PageEdges `json:"edges"`
}

// PageEdges holds the relations/edges for other nodes in the graph.
type PageEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Page `json:"parent,omitempty"`
	// Subpages holds the value of the subpages edge.
	Subpages []*Page `json:"subpages,omitempty"`
	// TopLists holds the value of the top_lists edge.
	TopLists []*TopList `json:"top_lists,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PageEdges) ParentOrErr() (*Page, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: page.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// SubpagesOrErr returns the Subpages value or an error if the edge
// was not loaded in eager-loading.
func (e PageEdges) SubpagesOrErr() ([]*Page, error) {
	if e.loadedTypes[1] {
		return e.Subpages, nil
	}
	return nil, &NotLoadedError{edge: "subpages"}
}

// TopListsOrErr returns the TopLists value or an error if the edge
// was not loaded in eager-loading.
func (e PageEdges) TopListsOrErr() ([]*TopList, error) {
	if e.loadedTypes[2] {
		return e.TopLists, nil
	}
	return nil, &NotLoadedError{edge: "top_lists"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Page) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case page.FieldCreatedUnix, page.FieldUpdatedUnix, page.FieldDeletedUnix, page.FieldFeaturedIndex, page.FieldRecommend:
			values[i] = new(sql.NullInt64)
		case page.FieldTitle, page.FieldFeaturedContent:
			values[i] = new(sql.NullString)
		case page.FieldID, page.FieldParentID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Page", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Page fields.
func (pa *Page) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pa.ID = *value
			}
		case page.FieldCreatedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_unix", values[i])
			} else if value.Valid {
				pa.CreatedUnix = value.Int64
			}
		case page.FieldUpdatedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_unix", values[i])
			} else if value.Valid {
				pa.UpdatedUnix = value.Int64
			}
		case page.FieldDeletedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_unix", values[i])
			} else if value.Valid {
				pa.DeletedUnix = value.Int64
			}
		case page.FieldParentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value != nil {
				pa.ParentID = *value
			}
		case page.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pa.Title = value.String
			}
		case page.FieldFeaturedIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field featured_index", values[i])
			} else if value.Valid {
				pa.FeaturedIndex = int8(value.Int64)
			}
		case page.FieldFeaturedContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field featured_content", values[i])
			} else if value.Valid {
				pa.FeaturedContent = value.String
			}
		case page.FieldRecommend:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field recommend", values[i])
			} else if value.Valid {
				pa.Recommend = value.Int64
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the Page entity.
func (pa *Page) QueryParent() *PageQuery {
	return (&PageClient{config: pa.config}).QueryParent(pa)
}

// QuerySubpages queries the "subpages" edge of the Page entity.
func (pa *Page) QuerySubpages() *PageQuery {
	return (&PageClient{config: pa.config}).QuerySubpages(pa)
}

// QueryTopLists queries the "top_lists" edge of the Page entity.
func (pa *Page) QueryTopLists() *TopListQuery {
	return (&PageClient{config: pa.config}).QueryTopLists(pa)
}

// Update returns a builder for updating this Page.
// Note that you need to call Page.Unwrap() before calling this method if this Page
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Page) Update() *PageUpdateOne {
	return (&PageClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Page entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Page) Unwrap() *Page {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("media: Page is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Page) String() string {
	var builder strings.Builder
	builder.WriteString("Page(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("created_unix=")
	builder.WriteString(fmt.Sprintf("%v", pa.CreatedUnix))
	builder.WriteString(", ")
	builder.WriteString("updated_unix=")
	builder.WriteString(fmt.Sprintf("%v", pa.UpdatedUnix))
	builder.WriteString(", ")
	builder.WriteString("deleted_unix=")
	builder.WriteString(fmt.Sprintf("%v", pa.DeletedUnix))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.ParentID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pa.Title)
	builder.WriteString(", ")
	builder.WriteString("featured_index=")
	builder.WriteString(fmt.Sprintf("%v", pa.FeaturedIndex))
	builder.WriteString(", ")
	builder.WriteString("featured_content=")
	builder.WriteString(pa.FeaturedContent)
	builder.WriteString(", ")
	builder.WriteString("recommend=")
	builder.WriteString(fmt.Sprintf("%v", pa.Recommend))
	builder.WriteByte(')')
	return builder.String()
}

// Pages is a parsable slice of Page.
type Pages []*Page

func (pa Pages) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
