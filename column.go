package gosql

import "strings"

const (
	// ReferencesMatchFull MATCH FULL
	ReferencesMatchFull = "FULL"
	// ReferencesMatchPartial MATCH PARTIAL
	ReferencesMatchPartial = "PARTIAL"
	// ReferencesMatchSimple MATCH SIMPLE
	ReferencesMatchSimple = "SIMPLE"

	// CheckNoInherit NO INHERIT
	CheckNoInherit = "NO INHERIT"

	// GeneratedStored STORED
	GeneratedStored = "STORED"
	// GeneratedAlways ALWAYS
	GeneratedAlways = "ALWAYS"
	// GeneratedByDefault BY DEFAULT
	GeneratedByDefault = "BY DEFAULT"

	// Deferrable DEFERRABLE
	Deferrable = "DEFERRABLE"
	// NotDeferrable NOT DEFERRABLE
	NotDeferrable = "NOT DEFERRABLE"

	// InitiallyDeferred DEFERRED
	InitiallyDeferred = "DEFERRED"
	// InitiallyImmediate IMMEDIATE
	InitiallyImmediate = "IMMEDIATE"

	// ActionNoAction NO ACTION
	ActionNoAction = "NO ACTION"
	// ActionCascade CASCADE
	ActionCascade = "CASCADE"
	// ActionRestrict RESTRICT
	ActionRestrict = "RESTRICT"
	// ActionSetNull SET NULL
	ActionSetNull = "SET NULL"
	// ActionSetDefault SET DEFAULT
	ActionSetDefault = "SET DEFAULT"
)

// column query builder
type column struct {
	// table name
	name string
	// sort rule
	collate string
	// compression method
	compression string
	// column constraint
	constraint constraintColumn
}

// SetCollate set sort rule
func (c *column) SetCollate(collation string) *column {
	c.collate = collation
	return c
}

// GetCollate get sort rule
func (c *column) GetCollate() string {
	return c.collate
}

// ResetCollate set sort rule to empty
func (c *column) ResetCollate() *column {
	c.collate = ""
	return c
}

// SetCompression set compression method
func (c *column) SetCompression(method string) *column {
	c.compression = method
	return c
}

// GetCompression get compression method
func (c *column) GetCompression() string {
	return c.compression
}

// ResetCompression set compression method to empty
func (c *column) ResetCompression() *column {
	c.compression = ""
	return c
}

// SetName set column name
func (c *column) SetName(name string) *column {
	c.name = name
	return c
}

// GetName get column name
func (c *column) GetName() string {
	return c.name
}

// ResetName reset name
func (c *column) ResetName() *column {
	c.name = ""
	return c
}

// String render column
func (c *column) String() string {
	// TODO render
	return ""
}

// constraint column
type constraintColumn struct {
	// name
	name string
	// nullable
	nullable *bool
	// check expression
	check detailedExpression
	// default
	def string
	// generated always as expression
	generatedAlwaysAs expression
	// generated
	generated detailedExpression
	// unique index
	unique indexParameters
	// primary key
	primary indexParameters
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

// Nullable is constraint nullable
func (c *constraintColumn) Nullable(isNullable *bool) *constraintColumn {
	c.nullable = isNullable
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

// Unique get unique
func (c *constraintColumn) Unique() *indexParameters {
	return &c.unique
}

// Primary get primary
func (c *constraintColumn) Primary() *indexParameters {
	return &c.primary
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
	// TODO render
	return ""
}

// IsEmpty check if empty
func (c *constraintColumn) IsEmpty() bool {
	return c == nil || (c.name == "" &&
		c.nullable == nil &&
		c.check.IsEmpty() &&
		c.def == "" &&
		c.generatedAlwaysAs.Len() == 0 &&
		c.generated.IsEmpty() &&
		c.unique.IsEmpty() &&
		c.primary.IsEmpty() &&
		c.references.IsEmpty() &&
		c.deferrable == nil &&
		c.initially == "")
}

//[ CONSTRAINT constraint_name ]
//{ NOT NULL |
//  NULL |
//  CHECK ( expression ) [ NO INHERIT ] |
//  DEFAULT default_expr |
//  GENERATED ALWAYS AS ( generation_expr ) STORED |
//  GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY [ ( sequence_options ) ] |
//  UNIQUE index_parameters |
//  PRIMARY KEY index_parameters |
//  REFERENCES reftable [ ( refcolumn ) ] [ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ]
//    [ ON DELETE referential_action ] [ ON UPDATE referential_action ] }
//[ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]

//
//[ INCLUDE ( column_name [, ... ] ) ]
//[ WITH ( storage_parameter [ = value] [, ... ] ) ]
//[ USING INDEX TABLESPACE tablespace_name ]

// indexParameters parameters of index
type indexParameters struct {
	// include
	include expression
	// with
	with expression
	// using index tablespace
	tableSpace string
}

// String render index parameters
func (i *indexParameters) String() string {
	if i.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if i.include.Len() > 0 {
		b.WriteString(" INCLUDE (" + i.include.String(", ") + ")")
	}
	if i.with.Len() > 0 {
		b.WriteString(" WITH (" + i.with.String(", ") + ")")
	}
	if i.tableSpace != "" {
		b.WriteString(" USING INDEX TABLESPACE " + i.tableSpace)
	}
	return b.String()
}

// IsEmpty is index parameter empty
func (i *indexParameters) IsEmpty() bool {
	return i == nil || (i.include.Len() == 0 && i.with.Len() > 0 && i.tableSpace == "")
}

// references column
type referencesColumn struct {
	// target table
	target string
	// target columns
	column expression
	// match full partial simple
	match string
	// on update
	update string
	// on delete
	delete string
}

// SetRefTable set reference table
func (r *referencesColumn) SetRefTable(table string) *referencesColumn {
	r.target = table
	return r
}

// GetRefTable get reference table
func (r *referencesColumn) GetRefTable() string {
	return r.target
}

// ResetRefTable reset reference table
func (r *referencesColumn) ResetRefTable() *referencesColumn {
	r.target = ""
	return r
}

// Columns reference columns
func (r *referencesColumn) Columns() *expression {
	return &r.column
}

// SetMatch set match
func (r *referencesColumn) SetMatch(match string) *referencesColumn {
	r.match = match
	return r
}

// GetMatch get match
func (r *referencesColumn) GetMatch() string {
	return r.match
}

// ResetMatch reset match
func (r *referencesColumn) ResetMatch() *referencesColumn {
	r.match = ""
	return r
}

// SetOnUpdate set on update
func (r *referencesColumn) SetOnUpdate(update string) *referencesColumn {
	r.update = update
	return r
}

// GetOnUpdate get on update
func (r *referencesColumn) GetOnUpdate() string {
	return r.update
}

// ResetOnUpdate reset on update
func (r *referencesColumn) ResetOnUpdate() *referencesColumn {
	r.update = ""
	return r
}

// SetOnDelete set on delete
func (r *referencesColumn) SetOnDelete(delete string) *referencesColumn {
	r.delete = delete
	return r
}

// GetOnDelete get on delete
func (r *referencesColumn) GetOnDelete() string {
	return r.delete
}

// ResetOnDelete reset on delete
func (r *referencesColumn) ResetOnDelete() *referencesColumn {
	r.delete = ""
	return r
}

// IsEmpty check if empty
func (r *referencesColumn) IsEmpty() bool {
	return r == nil || (r.column.Len() == 0 &&
		r.target == "" &&
		r.update == "" &&
		r.delete == "" &&
		r.match == "")
}

// String create reference for column
func (r *referencesColumn) String() string {
	if r.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("REFERENCES " + r.target)
	if r.column.Len() > 0 {
		b.WriteString(" (" + r.column.String(", ") + ")")
	}
	if r.match != "" {
		b.WriteString(" MATCH " + r.match)
	}
	if r.delete != "" {
		b.WriteString(" ON DELETE " + r.delete)
	}
	if r.update != "" {
		b.WriteString(" ON UPDATE " + r.update)
	}
	return b.String()
}

// NewReferenceColumn init ref column
func NewReferenceColumn() *referencesColumn {
	return &referencesColumn{}
}
