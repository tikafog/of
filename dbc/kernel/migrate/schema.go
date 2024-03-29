// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// InstructsColumns holds the columns for the "instructs" table.
	InstructsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "current_unix", Type: field.TypeInt64, Default: 0},
		{Name: "updated_unix", Type: field.TypeInt64, Default: 0},
	}
	// InstructsTable holds the schema information for the "instructs" table.
	InstructsTable = &schema.Table{
		Name:       "instructs",
		Columns:    InstructsColumns,
		PrimaryKey: []*schema.Column{InstructsColumns[0]},
	}
	// ResourcesColumns holds the columns for the "resources" table.
	ResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rid", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeUint8, Default: 0},
		{Name: "step", Type: field.TypeUint8, Default: 0},
		{Name: "retried", Type: field.TypeInt, Default: 0},
		{Name: "priority", Type: field.TypeInt, Default: -1},
		{Name: "relate", Type: field.TypeString, Default: "none"},
		{Name: "updated_unix", Type: field.TypeInt64, Default: 0},
		{Name: "comment", Type: field.TypeString, Nullable: true},
	}
	// ResourcesTable holds the schema information for the "resources" table.
	ResourcesTable = &schema.Table{
		Name:       "resources",
		Columns:    ResourcesColumns,
		PrimaryKey: []*schema.Column{ResourcesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "resource_rid",
				Unique:  false,
				Columns: []*schema.Column{ResourcesColumns[1]},
			},
		},
	}
	// VersionsColumns holds the columns for the "versions" table.
	VersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "current", Type: field.TypeInt, Default: 2},
		{Name: "last", Type: field.TypeInt, Default: 2},
	}
	// VersionsTable holds the schema information for the "versions" table.
	VersionsTable = &schema.Table{
		Name:       "versions",
		Columns:    VersionsColumns,
		PrimaryKey: []*schema.Column{VersionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		InstructsTable,
		ResourcesTable,
		VersionsTable,
	}
)

func init() {
}
