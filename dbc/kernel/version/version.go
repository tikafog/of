// Code generated by ent, DO NOT EDIT.

package version

const (
	// Label holds the string label denoting the version type in the database.
	Label = "version"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCurrent holds the string denoting the current field in the database.
	FieldCurrent = "current"
	// FieldLast holds the string denoting the last field in the database.
	FieldLast = "last"
	// Table holds the table name of the version in the database.
	Table = "versions"
)

// Columns holds all SQL columns for version fields.
var Columns = []string{
	FieldID,
	FieldCurrent,
	FieldLast,
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
	// DefaultCurrent holds the default value on creation for the "Current" field.
	DefaultCurrent int
	// CurrentValidator is a validator for the "Current" field. It is called by the builders before save.
	CurrentValidator func(int) error
	// DefaultLast holds the default value on creation for the "Last" field.
	DefaultLast int
	// LastValidator is a validator for the "Last" field. It is called by the builders before save.
	LastValidator func(int) error
)
