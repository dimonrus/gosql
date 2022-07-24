package gosql

import "strings"

// [ CONSTRAINT constraint_name ]
// { NOT NULL |
//  NULL |
//  CHECK ( expression ) [ NO INHERIT ] |
//  DEFAULT default_expr |
//  GENERATED ALWAYS AS ( generation_expr ) STORED |
//  GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY [ ( sequence_options ) ] |
//  UNIQUE index_parameters |
//  PRIMARY KEY index_parameters |
//  REFERENCES reftable [ ( refcolumn ) ] [ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ]
//    [ ON DELETE referential_action ] [ ON UPDATE referential_action ] }
// [ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
type constraintColumn struct {
	// name
	name string
	// nullable
	notNull bool
	// check expression
	check detailedExpression
	// default
	def string
	// generated always as expression
	generatedAlwaysAs expression
	// generated
	generated detailedExpression
	// unique index
	unique *indexParameters
	// primary key
	primary *indexParameters
	// references
	references referencesColumn
	// deferrable
	deferrable *bool
	// initially
	initially string
}

// SetName set name
func (c *constraintColumn) SetName(name string) *constraintColumn {
	c.name = name
	return c
}

// GetName get name
func (c *constraintColumn) GetName() string {
	return c.name
}

// ResetName reset name
func (c *constraintColumn) ResetName() *constraintColumn {
	c.name = ""
	return c
}

// NotNull is constraint nullable
func (c *constraintColumn) NotNull() *constraintColumn {
	c.notNull = true
	return c
}

// Check detailed expression
func (c *constraintColumn) Check() *detailedExpression {
	return &c.check
}

// SetDefault set default
func (c *constraintColumn) SetDefault(def string) *constraintColumn {
	c.def = def
	return c
}

// GetDefault get default
func (c *constraintColumn) GetDefault() string {
	return c.def
}

// ResetDefault reset default
func (c *constraintColumn) ResetDefault() *constraintColumn {
	c.def = ""
	return c
}

// GeneratedAlways expression
func (c *constraintColumn) GeneratedAlways() *expression {
	return &c.generatedAlwaysAs
}

// Generated expression
func (c *constraintColumn) Generated() *detailedExpression {
	return &c.generated
}

// SetUnique set unique
func (c *constraintColumn) SetUnique() *indexParameters {
	c.unique = NewIndexParameters()
	return c.unique
}

// SetPrimary set primary
func (c *constraintColumn) SetPrimary() *indexParameters {
	c.primary = NewIndexParameters()
	return c.primary
}

// References get references
func (c *constraintColumn) References() *referencesColumn {
	return &c.references
}

// SetDeferrable set deferrable
func (c *constraintColumn) SetDeferrable(deferrable *bool) *constraintColumn {
	c.deferrable = deferrable
	return c
}

// SetInitially set initially
func (c *constraintColumn) SetInitially(initially string) *constraintColumn {
	c.initially = initially
	return c
}

// GetInitially get initially
func (c *constraintColumn) GetInitially() string {
	return c.initially
}

// ResetInitially reset initially
func (c *constraintColumn) ResetInitially() *constraintColumn {
	c.initially = ""
	return c
}

// String render column constraint
func (c *constraintColumn) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if c.name != "" {
		b.WriteString(" CONSTRAINT " + c.name)
	}
	if c.notNull {
		b.WriteString(" NOT NULL")
	}
	if !c.check.IsEmpty() {
		b.WriteString(" CHECK " + c.check.String())
	}
	if c.def != "" {
		b.WriteString(" DEFAULT " + c.def)
	}
	if c.generatedAlwaysAs.Len() > 0 {
		b.WriteString(" GENERATED ALWAYS AS (" + c.generatedAlwaysAs.String(", ") + ") STORED")
	}
	if !c.generated.IsEmpty() {
		b.WriteString(" GENERATED " + c.generated.GetDetail() + " AS IDENTITY")
		if c.generated.Expression().Len() > 0 {
			b.WriteString(" " + c.generated.Expression().String(", "))
		}
	}
	if c.unique != nil {
		b.WriteString(" UNIQUE" + c.unique.String())
	}
	if c.primary != nil {
		b.WriteString(" PRIMARY KEY" + c.primary.String())
	}
	if !c.references.IsEmpty() {
		b.WriteString(" " + c.references.String())
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

// IsEmpty check if empty
func (c *constraintColumn) IsEmpty() bool {
	return c == nil || (c.name == "" &&
		!c.notNull &&
		c.check.IsEmpty() &&
		c.def == "" &&
		c.generatedAlwaysAs.Len() == 0 &&
		c.generated.IsEmpty() &&
		c.unique == nil &&
		c.primary == nil &&
		c.references.IsEmpty() &&
		c.deferrable == nil &&
		c.initially == "")
}

// NewConstraintColumn init column constraint
func NewConstraintColumn() *constraintColumn {
	return &constraintColumn{}
}
