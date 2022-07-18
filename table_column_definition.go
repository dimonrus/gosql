package gosql

import "strings"

// list of column definition
type columnDefinitions []*columnDefinition

// String render all column definitions
func (c columnDefinitions) String() string {
	b := strings.Builder{}
	for i, definition := range c {
		if i == 0 {
			b.WriteString(definition.String())
		} else {
			b.WriteString("," + definition.String())
		}
	}
	return b.String()
}

// table column definition
// { column_name data_type [ COMPRESSION compression_method ] [ COLLATE collation ] [ column_constraint [ ... ] ]
//    | table_constraint
//    | LIKE source_table [ like_option ... ] }
type columnDefinition struct {
	// column
	column column
	// or constraint
	constraintTable constraintTable
	// or like expression
	like likeTable
}

// Column return column
func (d *columnDefinition) Column() *column {
	return &d.column
}

// Constraint return constraint
func (d *columnDefinition) Constraint() *constraintTable {
	return &d.constraintTable
}

// Like return like table
func (d *columnDefinition) Like() *likeTable {
	return &d.like
}

// IsEmpty check is columnDefinition is empty
func (d *columnDefinition) IsEmpty() bool {
	return d == nil || (d.column.IsEmpty() && d.constraintTable.IsEmpty() && d.like.IsEmpty())
}

// String render columnDefinition
func (d *columnDefinition) String() string {
	if d.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if !d.column.IsEmpty() {
		b.WriteString(d.column.String())
	} else if !d.constraintTable.IsEmpty() {
		b.WriteString(d.constraintTable.String())
	} else if !d.like.IsEmpty() {
		b.WriteString(d.like.String())
	}
	return b.String()
}

// NewColumnDefinition init definition
func NewColumnDefinition() *columnDefinition {
	return &columnDefinition{}
}
