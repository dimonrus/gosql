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

	// LikeIncluding INCLUDING
	LikeIncluding = "INCLUDING"
	// LikeExcluding EXCLUDING
	LikeExcluding = "EXCLUDING"
	// LikeComments COMMENTS
	LikeComments = "COMMENTS"
	// LikeCompression COMPRESSION
	LikeCompression = "COMPRESSION"
	// LikeConstraints CONSTRAINTS
	LikeConstraints = "CONSTRAINTS"
	// LikeDefaults DEFAULTS
	LikeDefaults = "DEFAULTS"
	// LikeGenerated GENERATED
	LikeGenerated = "GENERATED"
	// LikeIdentity IDENTITY
	LikeIdentity = "IDENTITY"
	// LikeIndexes INDEXES
	LikeIndexes = "INDEXES"
	// LikeStatistics STATISTICS
	LikeStatistics = "STATISTICS"
	// LikeStorage STORAGE
	LikeStorage = "STORAGE"
	// LikeAll ALL
	LikeAll = "ALL"

	// WithWithoutOIDS WITHOUT OIDS
	WithWithoutOIDS = "WITHOUT OIDS"

	// PartitionByRange RANGE
	PartitionByRange = "RANGE"
	// PartitionByList LIST
	PartitionByList = "LIST"
	// PartitionByHash HASH
	PartitionByHash = "HASH"

	// PartitionBoundFromMin MINVALUE
	PartitionBoundFromMin = "MINVALUE"
	// PartitionBoundFromMax MAXVALUE
	PartitionBoundFromMax = "MAXVALUE"
	// PartitionBoundWithModulus MODULUS
	PartitionBoundWithModulus = "MODULUS"
	// PartitionBoundWithRemainder REMAINDER
	PartitionBoundWithRemainder = "REMAINDER"

	// PartitionOfWithOptions WITH OPTIONS
	PartitionOfWithOptions = "WITH OPTIONS"
)

// Table create table query builder
// CREATE [ [ GLOBAL | LOCAL ] { TEMPORARY | TEMP } | UNLOGGED ] TABLE [ IF NOT EXISTS ] table_name ( [
//  { column_name data_type [ COMPRESSION compression_method ] [ COLLATE collation ] [ column_constraint [ ... ] ]
//    | table_constraint
//    | LIKE source_table [ like_option ... ] }
//    [, ... ]
// ] )
// [ INHERITS ( parent_table [, ... ] ) ]
// [ PARTITION BY { RANGE | LIST | HASH } ( { column_name | ( expression ) } [ COLLATE collation ] [ opclass ] [, ... ] ) ]
// [ USING method ]
// [ WITH ( storage_parameter [= value] [, ... ] ) | WITHOUT OIDS ]
// [ ON COMMIT { PRESERVE ROWS | DELETE ROWS | DROP } ]
// [ TABLESPACE tablespace_name ]
type Table struct {
	// scope GLOBAL | LOCAL
	scope string
	// table name
	name string
	// definitions
	definitions columnDefinitions
	// of type
	ofType string
	// of definitions
	ofDefinitions ofDefinitions
	// inherits
	inherits expression
	// using method
	using string
	// Partition
	partition partitionTable
	// with
	with detailedExpression
	// tablespace
	tablespace string
	// on commit params
	onCommit string
	// is table temporary
	temp bool
	// unLogged table
	unLogged bool
	// If not exists
	ifNotExists bool
}

// String render table
func (t *Table) String() string {
	if t.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("CREATE")
	if t.scope != "" {
		b.WriteString(" " + t.scope)
	}
	if t.temp {
		b.WriteString(" TEMPORARY")
	} else if t.unLogged {
		b.WriteString(" UNLOGGED")
	}
	if t.name != "" {
		b.WriteString(" TABLE " + t.name)
	}
	if t.definitions.Len() > 0 {
		b.WriteString(" (" + t.definitions.String() + ")")
	} else if t.ofType != "" {
		b.WriteString(" OF " + t.ofType)
		if t.ofDefinitions.Len() > 0 {
			b.WriteString(" (" + t.ofDefinitions.String() + ")")
		}
	}
	if t.inherits.Len() > 0 {
		b.WriteString(" INHERITS " + t.inherits.String(", "))
	}
	if !t.partition.IsEmpty() {
		b.WriteString(" PARTITION BY " + t.partition.String())
	}
	if t.using != "" {
		b.WriteString(" USING " + t.using)
	}
	if !t.with.IsEmpty() {
		if t.with.GetDetail() != "" {
			b.WriteString(" " + t.with.GetDetail())
		} else {
			b.WriteString(" WITH " + t.with.String())
		}
	}
	if t.onCommit != "" {
		b.WriteString(" ON COMMIT " + t.onCommit)
	}
	if t.tablespace != "" {
		b.WriteString(" TABLESPACE " + t.tablespace)
	}
	return b.String() + ";"
}

// SetUsing set using
func (t *Table) SetUsing(using string) *Table {
	t.using = using
	return t
}

// GetUsing get using
func (t *Table) GetUsing() string {
	return t.using
}

// ResetUsing reset using
func (t *Table) ResetUsing() *Table {
	t.using = ""
	return t
}

// AddColumn add column
func (t *Table) AddColumn(name string) *column {
	def, _ := t.NewDefinition()
	return def.Column().SetName(name)
}

// AddForeignKey add foreign key
func (t *Table) AddForeignKey(target string, columns ...string) *foreignKey {
	def, _ := t.NewDefinition()
	def.Constraint().ForeignKey().Columns().Add(columns...)
	def.Constraint().ForeignKey().References().SetRefTable(target)
	def.Constraint().ForeignKey().References().Columns().Add(columns...)
	return def.Constraint().ForeignKey()
}

// AddConstraint add constraint
func (t *Table) AddConstraint() *constraintTable {
	def, _ := t.NewDefinition()
	return def.Constraint()
}

// With expression
func (t *Table) With(expr ...string) *detailedExpression {
	t.with.Expression().Add(expr...)
	return &t.with
}

// WithOutOIDS expression
func (t *Table) WithOutOIDS() *detailedExpression {
	t.with.SetDetail(WithWithoutOIDS)
	return &t.with
}

// Partition expression
func (t *Table) Partition() *partitionTable {
	return &t.partition
}

// Definitions implement definitions
func (t *Table) Definitions(definition ...*columnDefinition) *Table {
	t.definitions = definition
	return t
}

// NewDefinition add column definition
func (t *Table) NewDefinition() (def *columnDefinition, n int) {
	def = NewColumnDefinition()
	t.definitions = append(t.definitions, def)
	return def, len(t.definitions) - 1
}

// RemoveDefinition remove definition by n
func (t *Table) RemoveDefinition(n int) *Table {
	t.definitions = append(t.definitions[:n], t.definitions[n+1:]...)
	return t
}

// ClearDefinition remove all definitions
func (t *Table) ClearDefinition() *Table {
	t.definitions = t.definitions[:0]
	return t
}

// IsEmpty check if table is empty
func (t *Table) IsEmpty() bool {
	return t == nil || (t.scope == "" &&
		t.name == "" &&
		len(t.definitions) == 0 &&
		t.ofType == "" &&
		t.ofDefinitions.Len() == 0 &&
		t.inherits.Len() == 0 &&
		t.using == "" &&
		t.with.IsEmpty() &&
		t.tablespace == "" &&
		t.onCommit == "")
}

// SetOnCommit set onCommit
func (t *Table) SetOnCommit(onCommit string) *Table {
	t.onCommit = onCommit
	return t
}

// GetOnCommit get onCommit
func (t *Table) GetOnCommit() string {
	return t.onCommit
}

// ResetOnCommit reset onCommit
func (t *Table) ResetOnCommit() *Table {
	t.onCommit = ""
	return t
}

// Flags Set create flags
func (t *Table) Flags(ifNotExists, temp, unLogged bool) *Table {
	t.ifNotExists = ifNotExists
	t.unLogged = unLogged
	t.temp = temp
	return t
}

// SetScope set scope
func (t *Table) SetScope(scope string) *Table {
	t.scope = scope
	return t
}

// GetScope get scope
func (t *Table) GetScope() string {
	return t.scope
}

// ResetScope reset scope
func (t *Table) ResetScope() *Table {
	t.scope = ""
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

// NewOfTypeDefinition add of type definition
func (t *Table) NewOfTypeDefinition() (def *ofDefinition, n int) {
	def = NewOfTypeDefinition()
	t.ofDefinitions = append(t.ofDefinitions, def)
	return def, len(t.definitions) - 1
}

// RemoveOfTypeDefinition remove of type definition by n
func (t *Table) RemoveOfTypeDefinition(n int) *Table {
	t.ofDefinitions = append(t.ofDefinitions[:n], t.ofDefinitions[n+1:]...)
	return t
}

// ClearOfTypeDefinition remove all of type definitions
func (t *Table) ClearOfTypeDefinition() *Table {
	t.ofDefinitions = t.ofDefinitions[:0]
	return t
}

// AddOfTypeConstraint add of type constraint
func (t *Table) AddOfTypeConstraint() *constraintTable {
	def, _ := t.NewOfTypeDefinition()
	return def.Constraint()
}

// AddOfTypeColumn add of type column
func (t *Table) AddOfTypeColumn(name string) *ofColumn {
	def, _ := t.NewOfTypeDefinition()
	return def.Column().SetName(name)
}

// SetOfType set ofType
func (t *Table) SetOfType(ofType string) *Table {
	t.ofType = ofType
	return t
}

// GetOfType get ofType
func (t *Table) GetOfType() string {
	return t.ofType
}

// ResetOfType reset ofType
func (t *Table) ResetOfType() *Table {
	t.ofType = ""
	return t
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

// NewTable init table
func NewTable(name string) *Table {
	return &Table{name: name}
}
