// Code generated by ent, DO NOT EDIT.

package media

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/channel"
)

// Channel is the model entity for the Channel schema.
type Channel struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedUnix holds the value of the "created_unix" field.
	CreatedUnix int64 `json:"created_unix,omitempty"`
	// UpdatedUnix holds the value of the "updated_unix" field.
	UpdatedUnix int64 `json:"updated_unix,omitempty"`
	// DeletedUnix holds the value of the "deleted_unix" field.
	DeletedUnix int64 `json:"deleted_unix,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChannelQuery when eager-loading is set.
	Edges ChannelEdges `json:"edges"`
}

// ChannelEdges holds the relations/edges for other nodes in the graph.
type ChannelEdges struct {
	// Informations holds the value of the informations edge.
	Informations []*InformationV1 `json:"informations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InformationsOrErr returns the Informations value or an error if the edge
// was not loaded in eager-loading.
func (e ChannelEdges) InformationsOrErr() ([]*InformationV1, error) {
	if e.loadedTypes[0] {
		return e.Informations, nil
	}
	return nil, &NotLoadedError{edge: "informations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Channel) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case channel.FieldCreatedUnix, channel.FieldUpdatedUnix, channel.FieldDeletedUnix:
			values[i] = new(sql.NullInt64)
		case channel.FieldComment:
			values[i] = new(sql.NullString)
		case channel.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Channel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Channel fields.
func (c *Channel) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case channel.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case channel.FieldCreatedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_unix", values[i])
			} else if value.Valid {
				c.CreatedUnix = value.Int64
			}
		case channel.FieldUpdatedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_unix", values[i])
			} else if value.Valid {
				c.UpdatedUnix = value.Int64
			}
		case channel.FieldDeletedUnix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_unix", values[i])
			} else if value.Valid {
				c.DeletedUnix = value.Int64
			}
		case channel.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				c.Comment = value.String
			}
		}
	}
	return nil
}

// QueryInformations queries the "informations" edge of the Channel entity.
func (c *Channel) QueryInformations() *InformationV1Query {
	return (&ChannelClient{config: c.config}).QueryInformations(c)
}

// Update returns a builder for updating this Channel.
// Note that you need to call Channel.Unwrap() before calling this method if this Channel
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Channel) Update() *ChannelUpdateOne {
	return (&ChannelClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Channel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Channel) Unwrap() *Channel {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("media: Channel is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Channel) String() string {
	var builder strings.Builder
	builder.WriteString("Channel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_unix=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedUnix))
	builder.WriteString(", ")
	builder.WriteString("updated_unix=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedUnix))
	builder.WriteString(", ")
	builder.WriteString("deleted_unix=")
	builder.WriteString(fmt.Sprintf("%v", c.DeletedUnix))
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(c.Comment)
	builder.WriteByte(')')
	return builder.String()
}

// Channels is a parsable slice of Channel.
type Channels []*Channel

func (c Channels) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
