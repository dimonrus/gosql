package gosql

import "strings"

// list of partition columns
type partitionColumns []*partitionColumn

// String render all column definitions
func (c partitionColumns) String() string {
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

// Len count of partition columns
func (c partitionColumns) Len() int {
	return len(c)
}

//{ column_name | ( expression ) } [ COLLATE collation ] [ opclass ]
type partitionColumn struct {
	expression expression
	// collation
	collate string
	// opclass
	opclass string
}

// Expression get expression
func (p *partitionColumn) Expression() *expression {
	return &p.expression
}

// SetCollate set sort rule
func (p *partitionColumn) SetCollate(collation string) *partitionColumn {
	p.collate = collation
	return p
}

// GetCollate get sort rule
func (p *partitionColumn) GetCollate() string {
	return p.collate
}

// ResetCollate set sort rule to empty
func (p *partitionColumn) ResetCollate() *partitionColumn {
	p.collate = ""
	return p
}

// SetOpClass set opclass
func (p *partitionColumn) SetOpClass(opclass string) *partitionColumn {
	p.opclass = opclass
	return p
}

// GetOpClass get opclass
func (p *partitionColumn) GetOpClass() string {
	return p.opclass
}

// ResetOpClass set opclass to empty
func (p *partitionColumn) ResetOpClass() *partitionColumn {
	p.opclass = ""
	return p
}

// IsEmpty check is partition is empty
func (p *partitionColumn) IsEmpty() bool {
	return p == nil || (p.expression.Len() == 0 && p.collate == "" && p.opclass == "")
}

// String Render partition into string
func (p *partitionColumn) String() string {
	if p.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if p.expression.Len() > 0 {
		b.WriteString(" (" + p.expression.String(", ") + ")")
	}
	if p.collate != "" {
		b.WriteString(" COLLATE " + p.collate)
	}
	if p.opclass != "" {
		b.WriteString(" " + p.opclass)
	}
	return b.String()
}

// NewPartitionColumn init partition column
func NewPartitionColumn() *partitionColumn {
	return &partitionColumn{}
}

// [ PARTITION BY { RANGE | LIST | HASH } ( { column_name | ( expression ) } [ COLLATE collation ] [ opclass ] [, ... ] ) ]
type partitionTable struct {
	// type of partition
	partitionBy string
	// partition columns list
	partitionColumns partitionColumns
}

// By type of partition RANGE | LIST | HASH
func (p *partitionTable) By(class string) *partitionTable {
	p.partitionBy = class
	return p
}

// Add partition
//  { column_name | ( expression ) } [ COLLATE collation ] [ opclass ]
func (p *partitionTable) Add(expression, collation, opclass string) *partitionColumn {
	col := NewPartitionColumn()
	col.Expression().Add(expression)
	col.SetCollate(collation)
	col.SetOpClass(opclass)
	p.partitionColumns = append(p.partitionColumns, col)
	return col
}

// IsEmpty check is partition is empty
func (p *partitionTable) IsEmpty() bool {
	return p == nil || (p.partitionBy == "" && p.partitionColumns.Len() == 0)
}

// String Render partition into string
func (p *partitionTable) String() string {
	if p.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if p.partitionBy != "" {
		b.WriteString(" " + p.partitionBy)
	}
	if p.partitionColumns.Len() > 0 {
		b.WriteString(" " + p.partitionColumns.String())
	}
	return b.String()[1:]
}
