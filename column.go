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
	unique string
	// primary key
	primary string
	// references
	references referencesColumn
	// deferrable
	deferrable *bool
	// initially
	initially string
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
