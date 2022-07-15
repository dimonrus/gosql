package gosql

import "strings"

// [ CONSTRAINT constraint_name ]
// { CHECK ( expression ) [ NO INHERIT ] |
//  UNIQUE ( column_name [, ... ] ) index_parameters |
//  PRIMARY KEY ( column_name [, ... ] ) index_parameters |
//  EXCLUDE [ USING index_method ] ( exclude_element WITH operator [, ... ] ) index_parameters [ WHERE ( predicate ) ] |
//  FOREIGN KEY ( column_name [, ... ] ) REFERENCES reftable [ ( refcolumn [, ... ] ) ]
//    [ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ] [ ON DELETE referential_action ] [ ON UPDATE referential_action ] }
// [ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
type constraintTable struct {
	// name
	name string
	// check expression
	check detailedExpression
	// unique index
	unique columnIndexParameters
	// primary key
	primary columnIndexParameters
	// exclude
	exclude excludeTable
	// foreign key
	foreignKey foreignKey
	// deferrable
	deferrable *bool
	// initially
	initially string
}

// SetName set name
func (c *constraintTable) SetName(name string) *constraintTable {
	c.name = name
	return c
}

// GetName get name
func (c *constraintTable) GetName() string {
	return c.name
}

// ResetName reset name
func (c *constraintTable) ResetName() *constraintTable {
	c.name = ""
	return c
}

// Check get check expression
func (c *constraintTable) Check() *detailedExpression {
	return &c.check
}

// Unique get unique expression
func (c *constraintTable) Unique() *columnIndexParameters {
	return &c.unique
}

// PrimaryKey get primary key expression
func (c *constraintTable) PrimaryKey() *columnIndexParameters {
	return &c.primary
}

// Exclude get exclude expression
func (c *constraintTable) Exclude() *excludeTable {
	return &c.exclude
}

// ForeignKey get foreign key expression
func (c *constraintTable) ForeignKey() *foreignKey {
	return &c.foreignKey
}

// SetDeferrable set deferrable
func (c *constraintTable) SetDeferrable(deferrable *bool) *constraintTable {
	c.deferrable = deferrable
	return c
}

// SetInitially set initially
func (c *constraintTable) SetInitially(initially string) *constraintTable {
	c.initially = initially
	return c
}

// GetInitially get initially
func (c *constraintTable) GetInitially() string {
	return c.initially
}

// ResetInitially reset initially
func (c *constraintTable) ResetInitially() *constraintTable {
	c.initially = ""
	return c
}

// IsEmpty check if constraint is empty
func (c *constraintTable) IsEmpty() bool {
	return c == nil || (c.name == "" &&
		c.check.IsEmpty() &&
		c.unique.IsEmpty() &&
		c.primary.IsEmpty() &&
		c.exclude.IsEmpty() &&
		c.foreignKey.IsEmpty() &&
		c.deferrable == nil &&
		c.initially == "")
}

// String render table constraint
func (c *constraintTable) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if c.name != "" {
		b.WriteString(" CONSTRAINT " + c.name)
	}
	if !c.check.IsEmpty() {
		b.WriteString(" CHECK " + c.check.String())
	} else if !c.unique.IsEmpty() {
		b.WriteString(" UNIQUE " + c.unique.String())
	} else if !c.primary.IsEmpty() {
		b.WriteString(" PRIMARY KEY " + c.primary.String())
	} else if !c.exclude.IsEmpty() {
		b.WriteString(c.exclude.String())
	} else if !c.foreignKey.IsEmpty() {
		b.WriteString(c.foreignKey.String())
	}
	if c.deferrable != nil {
		if *c.deferrable {
			b.WriteString(" DEFERRABLE")
		} else {
			b.WriteString(" NOT DEFERRABLE")
		}
	}
	if c.initially != "" {
		b.WriteString(" INITIALLY " + c.initially)
	}
	return b.String()
}

// NewConstraintTable init table constraint
func NewConstraintTable() *constraintTable {
	return &constraintTable{}
}
