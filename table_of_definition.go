package gosql

import "strings"

// list of column definition
type ofDefinitions []*ofDefinition

// String render all column definitions
func (c ofDefinitions) String() string {
	b := strings.Builder{}
	for i, definition := range c {
		if i == 0 {
			b.WriteString(definition.String())
		} else {
			b.WriteString(", " + definition.String())
		}
	}
	return b.String()
}

// Len count of definition
func (c ofDefinitions) Len() int {
	return len(c)
}

//  { column_name [ WITH OPTIONS ] [ column_constraint [ ... ] ]
//    | table_constraint }
//    [, ... ]
type ofDefinition struct {
	// column definition
	ofColumn ofColumn
	// constrain definition
	constraintTable constraintTable
}

// Column get of column
func (c *ofDefinition) Column() *ofColumn {
	return &c.ofColumn
}

// Constraint get constraint
func (c *ofDefinition) Constraint() *constraintTable {
	return &c.constraintTable
}

// IsEmpty check if empty
func (c *ofDefinition) IsEmpty() bool {
	return c == nil || (c.ofColumn.IsEmpty() && c.constraintTable.IsEmpty())
}

// String render column
func (c *ofDefinition) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if !c.ofColumn.IsEmpty() {
		b.WriteString(c.ofColumn.String())
	} else if !c.constraintTable.IsEmpty() {
		b.WriteString(" " + c.constraintTable.String())
	}
	return b.String()
}

// NewOfTypeDefinition init ofDefinition
func NewOfTypeDefinition() *ofDefinition {
	return &ofDefinition{}
}

// column_name [ WITH OPTIONS ] [ column_constraint [ ... ] ]
type ofColumn struct {
	// column name
	name string
	// WITH OPTIONS
	withOptions bool
	// constraint
	constraint constraintColumn
}

// SetName set name
func (c *ofColumn) SetName(name string) *ofColumn {
	c.name = name
	return c
}

// GetName get name
func (c *ofColumn) GetName() string {
	return c.name
}

// ResetName reset name
func (c *ofColumn) ResetName() *ofColumn {
	c.name = ""
	return c
}

// WithOptions reset name
func (c *ofColumn) WithOptions() *ofColumn {
	c.withOptions = true
	return c
}

// Constraint get constraint
func (c *ofColumn) Constraint() *constraintColumn {
	return &c.constraint
}

// IsEmpty check if empty
func (c *ofColumn) IsEmpty() bool {
	return c == nil || (c.name == "" && c.constraint.IsEmpty())
}

// String render column
func (c *ofColumn) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if c.name != "" {
		b.WriteString(c.name)
	}
	if c.withOptions {
		b.WriteString(" WITH OPTIONS")
	}
	if !c.constraint.IsEmpty() {
		b.WriteString(c.constraint.String())
	}
	return b.String()
}

// NewOfTypeColumn init of column
func NewOfTypeColumn() *ofColumn {
	return &ofColumn{}
}
