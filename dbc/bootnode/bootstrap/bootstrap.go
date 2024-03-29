// Code generated by ent, DO NOT EDIT.

package bootstrap

import (
	"fmt"
)

const (
	// Label holds the string label denoting the bootstrap type in the database.
	Label = "bootstrap"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldAddrs holds the string denoting the addrs field in the database.
	FieldAddrs = "addrs"
	// FieldExpired holds the string denoting the expired field in the database.
	FieldExpired = "expired"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldServicePort holds the string denoting the service_port field in the database.
	FieldServicePort = "service_port"
	// FieldFailCounts holds the string denoting the fail_counts field in the database.
	FieldFailCounts = "fail_counts"
	// FieldSign holds the string denoting the sign field in the database.
	FieldSign = "sign"
	// FieldUpdatedUnix holds the string denoting the updated_unix field in the database.
	FieldUpdatedUnix = "updated_unix"
	// Table holds the table name of the bootstrap in the database.
	Table = "bootstraps"
)

// Columns holds all SQL columns for bootstrap fields.
var Columns = []string{
	FieldID,
	FieldPid,
	FieldAddrs,
	FieldExpired,
	FieldLevel,
	FieldServicePort,
	FieldFailCounts,
	FieldSign,
	FieldUpdatedUnix,
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
	// DefaultExpired holds the default value on creation for the "expired" field.
	DefaultExpired bool
	// DefaultServicePort holds the default value on creation for the "service_port" field.
	DefaultServicePort int
	// DefaultFailCounts holds the default value on creation for the "fail_counts" field.
	DefaultFailCounts int
	// DefaultUpdatedUnix holds the default value on creation for the "updated_unix" field.
	DefaultUpdatedUnix int64
)

// Level defines the type for the "level" enum field.
type Level string

// LevelCore is the default value of the Level enum.
const DefaultLevel = LevelCore

// Level values.
const (
	LevelCore   Level = "core"
	LevelSpeed  Level = "speed"
	LevelNormal Level = "normal"
)

func (l Level) String() string {
	return string(l)
}

// LevelValidator is a validator for the "level" field enum values. It is called by the builders before save.
func LevelValidator(l Level) error {
	switch l {
	case LevelCore, LevelSpeed, LevelNormal:
		return nil
	default:
		return fmt.Errorf("bootstrap: invalid enum value for level field: %q", l)
	}
}
