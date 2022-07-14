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

// Table query builder
type Table struct {
	// table name
	name string
	// If not exists
	ifNotExists bool
	// is table temporary
	temp bool
	// unLogged table
	unLogged bool
	// inherits
	inherits expression
	// tablespace
	tablespace string
}

// Flags Set create flags
func (t *Table) Flags(ifNotExists, temp, unLogged bool) *Table {
	t.ifNotExists = ifNotExists
	t.unLogged = unLogged
	t.temp = temp
	return t
}

// SetName Set name
func (t *Table) SetName(name string) *Table {
	t.name = name
	return t
}

// GetName get name
func (t *Table) GetName() string {
	return t.name
}

// Inherits inherit form tables
func (t *Table) Inherits() *expression {
	return &t.inherits
}

// SetTableSpace set table space
func (t *Table) SetTableSpace(space string) *Table {
	t.tablespace = space
	return t
}

// GetTableSpace get table space
func (t *Table) GetTableSpace() string {
	return t.tablespace
}

// ResetTableSpace reset table space
func (t *Table) ResetTableSpace() *Table {
	t.tablespace = ""
	return t
}

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
	foreignKey string // TODO
	// deferrable
	deferrable *bool
	// initially
	initially string
}

// table constraint index parameters
type columnIndexParameters struct {
	// columns
	columns expression
	// index parameter
	indexParameters indexParameters
}

// Columns get columns
func (i *columnIndexParameters) Columns() *expression {
	return &i.columns
}

// IndexParameters get index parameters
func (i *columnIndexParameters) IndexParameters() *indexParameters {
	return &i.indexParameters
}

// String render index parameters
func (i *columnIndexParameters) String() string {
	if i.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if i.columns.Len() > 0 {
		b.WriteString("(" + i.columns.String(", ") + ")")
	}
	if !i.indexParameters.IsEmpty() {
		b.WriteString(i.indexParameters.String())
	}
	return b.String()
}

// IsEmpty is index parameter empty
func (i *columnIndexParameters) IsEmpty() bool {
	return i == nil || (i.columns.Len() == 0 && i.indexParameters.IsEmpty())
}

// NewColumnIndexParameters new column index parameters
func NewColumnIndexParameters() *columnIndexParameters {
	return &columnIndexParameters{}
}

// EXCLUDE [ USING index_method ] ( exclude_element WITH operator [, ... ] ) index_parameters [ WHERE ( predicate ) ] |
type excludeTable struct {
	// using index method
	using string
	// exclude element
	excludeElement excludeElement
	// with expression
	with expression
	// index parameters
	indexParameters indexParameters
	// where condition
	where Condition
}

// SetUsing set using
func (e *excludeTable) SetUsing(using string) *excludeTable {
	e.using = using
	return e
}

// GetUsing get using
func (e *excludeTable) GetUsing() string {
	return e.using
}

// ResetUsing reset using
func (e *excludeTable) ResetUsing() *excludeTable {
	e.using = ""
	return e
}

// ExcludeElement return exclude element
func (e *excludeTable) ExcludeElement() *excludeElement {
	return &e.excludeElement
}

// With expression
func (e *excludeTable) With() *expression {
	return &e.with
}

// IndexParameters return index params
func (e *excludeTable) IndexParameters() *indexParameters {
	return &e.indexParameters
}

// Where condition
func (e *excludeTable) Where() *Condition {
	return &e.where
}

// IsEmpty check if empty return true
func (e *excludeTable) IsEmpty() bool {
	return e == nil || (e.using == "" && e.excludeElement.IsEmpty() && e.with.Len() == 0 && e.indexParameters.IsEmpty() && e.where.IsEmpty())
}

// String render to string
func (e *excludeTable) String() string {
	if e.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("EXCLUDE")
	if e.using != "" {
		b.WriteString(" USING " + e.using)
	}
	if !e.excludeElement.IsEmpty() {
		if e.with.Len() > 0 {
			b.WriteString(" (" + e.excludeElement.String() + " WITH " + e.with.String(", ") + ")")
		} else {
			b.WriteString(" (" + e.excludeElement.String() + ")")
		}
	} else if e.with.Len() > 0 {
		b.WriteString(" WITH (" + e.with.String(", ") + ")")
	}
	if !e.indexParameters.IsEmpty() {
		b.WriteString(" " + e.indexParameters.String())
	}
	if !e.where.IsEmpty() {
		b.WriteString(" WHERE " + e.where.String())
	}
	return b.String()
}

// exclude element
// { column_name | ( expression ) } [ opclass ] [ ASC | DESC ] [ NULLS { FIRST | LAST } ]
type excludeElement struct {
	// column name
	column string
	// expression
	expression expression
	// operator class
	opclass string
	// direction
	direction string
	// nulls expression
	nulls *string
}

// SetColumn set column name
func (e *excludeElement) SetColumn(column string) *excludeElement {
	e.column = column
	return e
}

// GetColumn get column name
func (e *excludeElement) GetColumn() string {
	return e.column
}

// ResetColumn set column to empty
func (e *excludeElement) ResetColumn() *excludeElement {
	e.column = ""
	return e
}

// SetOpClass set opclass
func (e *excludeElement) SetOpClass(opclass string) *excludeElement {
	e.opclass = opclass
	return e
}

// GetOpClass get opclass
func (e *excludeElement) GetOpClass() string {
	return e.opclass
}

// ResetOpClass set opclass to empty
func (e *excludeElement) ResetOpClass() *excludeElement {
	e.opclass = ""
	return e
}

// SetDirection set direction
func (e *excludeElement) SetDirection(direction string) *excludeElement {
	e.direction = direction
	return e
}

// GetDirection get direction
func (e *excludeElement) GetDirection() string {
	return e.direction
}

// ResetDirection set direction to empty
func (e *excludeElement) ResetDirection() *excludeElement {
	e.direction = ""
	return e
}

// SetNulls set nulls
func (e *excludeElement) SetNulls(nulls *string) *excludeElement {
	e.nulls = nulls
	return e
}

// GetNulls get nulls
func (e *excludeElement) GetNulls() *string {
	return e.nulls
}

// ResetNulls reset null
func (e *excludeElement) ResetNulls() *excludeElement {
	e.nulls = nil
	return e
}

// Expression get expression
func (e *excludeElement) Expression() *expression {
	return &e.expression
}

// String render exclude element
func (e *excludeElement) String() string {
	if e.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if e.column != "" {
		b.WriteString(e.column)
	} else if e.expression.Len() > 0 {
		b.WriteString("(" + e.expression.String(", ") + ")")
	}
	if e.opclass != "" {
		b.WriteString(" " + e.opclass)
	}
	if e.direction != "" {
		b.WriteString(" " + e.direction)
	}
	if e.nulls != nil {
		b.WriteString(" NULLS " + *e.nulls)
	}
	return b.String()
}

// IsEmpty is exclude element is empty
func (e *excludeElement) IsEmpty() bool {
	return e == nil || (e.column == "" && e.expression.Len() == 0 && e.opclass == "" && e.direction == " " && e.nulls == nil)
}

// NewExcludeElement init exclude element
func NewExcludeElement() *excludeElement {
	return &excludeElement{}
}
