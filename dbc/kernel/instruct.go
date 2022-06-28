// Code generated by entc, DO NOT EDIT.

package kernel

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/tikafog/of/dbc/kernel/instruct"
)

// Instruct is the model entity for the Instruct schema.
type Instruct struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CurrentUnix holds the value of the "current_unix" field.
	CurrentUnix int64 `json:"current_unix,omitempty"`
	// UpdatedUnix holds the value of the "updated_unix" field.
	UpdatedUnix int64 `json:"updated_unix,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Instruct) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case instruct.FieldID, instruct.FieldCurrentUnix, instruct.FieldUpdatedUnix:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Instruct", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Instruct fields.
func (i *Instruct) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case instruct.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case instruct.FieldCurrentUnix:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field current_unix", values[j])
			} else if value.Valid {
				i.CurrentUnix = value.Int64
			}
		case instruct.FieldUpdatedUnix:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_unix", values[j])
			} else if value.Valid {
				i.UpdatedUnix = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Instruct.
// Note that you need to call Instruct.Unwrap() before calling this method if this Instruct
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Instruct) Update() *InstructUpdateOne {
	return (&InstructClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Instruct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Instruct) Unwrap() *Instruct {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("kernel: Instruct is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Instruct) String() string {
	var builder strings.Builder
	builder.WriteString("Instruct(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", current_unix=")
	builder.WriteString(fmt.Sprintf("%v", i.CurrentUnix))
	builder.WriteString(", updated_unix=")
	builder.WriteString(fmt.Sprintf("%v", i.UpdatedUnix))
	builder.WriteByte(')')
	return builder.String()
}

// Instructs is a parsable slice of Instruct.
type Instructs []*Instruct

func (i Instructs) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}