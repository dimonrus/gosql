package gosql

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
