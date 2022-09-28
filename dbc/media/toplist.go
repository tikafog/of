// Code generated by ent, DO NOT EDIT.

package media

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/informationv1"
	"github.com/tikafog/of/dbc/media/page"
	"github.com/tikafog/of/dbc/media/toplist"
)

// TopList is the model entity for the TopList schema.
type TopList struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedUnix holds the value of the "created_unix" field.
	CreatedUnix int64 `json:"created_unix,omitempty"`
	// UpdatedUnix holds the value of the "updated_unix" field.
	UpdatedUnix int64 `json:"updated_unix,omitempty"`
	// DeletedUnix holds the value of the "deleted_unix" field.
	DeletedUnix int64 `json:"deleted_unix,omitempty"`
	// InformationID holds the value of the "information_id" field.
	InformationID uuid.UUID `json:"information_id,omitempty"`
	// PageID holds the value of the "page_id" field.
	PageID uuid.UUID `json:"page_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Intro holds the value of the "intro" field.
	Intro string `json:"intro,omitempty"`
	// Category holds the value of the "category" field.
	Category toplist.Category `json:"category,omitempty"`
	// LowerBanner holds the value of the "lower_banner" field.
	LowerBanner toplist.LowerBanner `json:"lower_banner,omitempty"`
	// TopRight holds the value of the "top_right" field.
	TopRight toplist.TopRight `json:"top_right,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TopListQuery when eager-loading is set.
	Edges TopListEdges `json:"edges"`
}

// TopListEdges holds the relations/edges for other nodes in the graph.
type TopListEdges struct {
	// Information holds the value of the information edge.
	Information *InformationV1 `json:"information,omitempty"`
	// Page holds the value of the page edge.
	Page *Page `json:"page,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// InformationOrErr returns the Information value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TopListEdges) InformationOrErr() (*InformationV1, error) {
	if e.loadedTypes[0] {
		if e.Information == nil {
			// The edge information was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: informationv1.Label}
		}
		return e.Information, nil
	}
	return nil, &NotLoadedError{edge: "information"}
}

// PageOrErr returns the Page value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TopListEdges) PageOrErr() (*Page, error) {
	if e.loadedTypes[1] {
		if e.Page == nil {
			// The edge page was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: page.Label}
		}
		return e.Page, nil
	}
	return nil, &NotLoadedError{edge: "page"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TopList) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case toplist.FieldCreatedUnix, toplist.FieldUpdatedUnix, toplist.FieldDeletedUnix:
			values[i] = new(sql.NullInt64)
		case toplist.FieldTitle, toplist.FieldIntro, toplist.FieldCategory, toplist.FieldLowerBanner, toplist.FieldTopRight:
			values[i] = new(sql.NullString)
		case toplist.FieldID, toplist.FieldInformationID, toplist.FieldPageID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TopList", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TopList fields.
func (tl *TopList) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case toplist.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tl.ID = *value
			}
		case toplist.FieldCreatedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_unix", values[i])
			} else if value.Valid {
				tl.CreatedUnix = value.Int64
			}
		case toplist.FieldUpdatedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_unix", values[i])
			} else if value.Valid {
				tl.UpdatedUnix = value.Int64
			}
		case toplist.FieldDeletedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_unix", values[i])
			} else if value.Valid {
				tl.DeletedUnix = value.Int64
			}
		case toplist.FieldInformationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field information_id", values[i])
			} else if value != nil {
				tl.InformationID = *value
			}
		case toplist.FieldPageID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field page_id", values[i])
			} else if value != nil {
				tl.PageID = *value
			}
		case toplist.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				tl.Title = value.String
			}
		case toplist.FieldIntro:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field intro", values[i])
			} else if value.Valid {
				tl.Intro = value.String
			}
		case toplist.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				tl.Category = toplist.Category(value.String)
			}
		case toplist.FieldLowerBanner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lower_banner", values[i])
			} else if value.Valid {
				tl.LowerBanner = toplist.LowerBanner(value.String)
			}
		case toplist.FieldTopRight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field top_right", values[i])
			} else if value.Valid {
				tl.TopRight = toplist.TopRight(value.String)
			}
		}
	}
	return nil
}

// QueryInformation queries the "information" edge of the TopList entity.
func (tl *TopList) QueryInformation() *InformationV1Query {
	return (&TopListClient{config: tl.config}).QueryInformation(tl)
}

// QueryPage queries the "page" edge of the TopList entity.
func (tl *TopList) QueryPage() *PageQuery {
	return (&TopListClient{config: tl.config}).QueryPage(tl)
}

// Update returns a builder for updating this TopList.
// Note that you need to call TopList.Unwrap() before calling this method if this TopList
// was returned from a transaction, and the transaction was committed or rolled back.
func (tl *TopList) Update() *TopListUpdateOne {
	return (&TopListClient{config: tl.config}).UpdateOne(tl)
}

// Unwrap unwraps the TopList entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tl *TopList) Unwrap() *TopList {
	_tx, ok := tl.config.driver.(*txDriver)
	if !ok {
		panic("media: TopList is not a transactional entity")
	}
	tl.config.driver = _tx.drv
	return tl
}

// String implements the fmt.Stringer.
func (tl *TopList) String() string {
	var builder strings.Builder
	builder.WriteString("TopList(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tl.ID))
	builder.WriteString("created_unix=")
	builder.WriteString(fmt.Sprintf("%v", tl.CreatedUnix))
	builder.WriteString(", ")
	builder.WriteString("updated_unix=")
	builder.WriteString(fmt.Sprintf("%v", tl.UpdatedUnix))
	builder.WriteString(", ")
	builder.WriteString("deleted_unix=")
	builder.WriteString(fmt.Sprintf("%v", tl.DeletedUnix))
	builder.WriteString(", ")
	builder.WriteString("information_id=")
	builder.WriteString(fmt.Sprintf("%v", tl.InformationID))
	builder.WriteString(", ")
	builder.WriteString("page_id=")
	builder.WriteString(fmt.Sprintf("%v", tl.PageID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(tl.Title)
	builder.WriteString(", ")
	builder.WriteString("intro=")
	builder.WriteString(tl.Intro)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", tl.Category))
	builder.WriteString(", ")
	builder.WriteString("lower_banner=")
	builder.WriteString(fmt.Sprintf("%v", tl.LowerBanner))
	builder.WriteString(", ")
	builder.WriteString("top_right=")
	builder.WriteString(fmt.Sprintf("%v", tl.TopRight))
	builder.WriteByte(')')
	return builder.String()
}

// TopLists is a parsable slice of TopList.
type TopLists []*TopList

func (tl TopLists) config(cfg config) {
	for _i := range tl {
		tl[_i].config = cfg
	}
}
