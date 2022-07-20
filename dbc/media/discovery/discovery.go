// Code generated by entc, DO NOT EDIT.

package discovery

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the discovery type in the database.
	Label = "discovery"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedUnix holds the string denoting the created_unix field in the database.
	FieldCreatedUnix = "created_unix"
	// FieldUpdatedUnix holds the string denoting the updated_unix field in the database.
	FieldUpdatedUnix = "updated_unix"
	// FieldDeletedUnix holds the string denoting the deleted_unix field in the database.
	FieldDeletedUnix = "deleted_unix"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldRid holds the string denoting the rid field in the database.
	FieldRid = "rid"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldMtype holds the string denoting the mtype field in the database.
	FieldMtype = "mtype"
	// FieldLinks holds the string denoting the links field in the database.
	FieldLinks = "links"
	// Table holds the table name of the discovery in the database.
	Table = "discoveries"
)

// Columns holds all SQL columns for discovery fields.
var Columns = []string{
	FieldID,
	FieldCreatedUnix,
	FieldUpdatedUnix,
	FieldDeletedUnix,
	FieldDate,
	FieldRid,
	FieldTitle,
	FieldDetail,
	FieldMtype,
	FieldLinks,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedUnix holds the default value on creation for the "created_unix" field.
	DefaultCreatedUnix int64
	// DefaultUpdatedUnix holds the default value on creation for the "updated_unix" field.
	DefaultUpdatedUnix int64
	// DefaultDate holds the default value on creation for the "date" field.
	DefaultDate string
	// DefaultRid holds the default value on creation for the "rid" field.
	DefaultRid string
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultDetail holds the default value on creation for the "detail" field.
	DefaultDetail string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Mtype defines the type for the "mtype" enum field.
type Mtype string

// MtypeBoth is the default value of the Mtype enum.
const DefaultMtype = MtypeBoth

// Mtype values.
const (
	MtypeNone  Mtype = "none"
	MtypeVideo Mtype = "video"
	MtypePhoto Mtype = "photo"
	MtypeBoth  Mtype = "both"
)

func (m Mtype) String() string {
	return string(m)
}

// MtypeValidator is a validator for the "mtype" field enum values. It is called by the builders before save.
func MtypeValidator(m Mtype) error {
	switch m {
	case MtypeNone, MtypeVideo, MtypePhoto, MtypeBoth:
		return nil
	default:
		return fmt.Errorf("discovery: invalid enum value for mtype field: %q", m)
	}
}