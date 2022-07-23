package gosql

import "strings"

// PARTITION OF parent_table [ (
//  { column_name [ WITH OPTIONS ] [ column_constraint [ ... ] ]
//    | table_constraint }
//    [, ... ]
//) ] { FOR VALUES partition_bound_spec | DEFAULT }
type ofPartition struct {
	// name of parent table
	parent string
	// column definitions
	definitions ofDefinitions
	// bounds
	values partitionBound
}

// SetParent set parent
func (p *ofPartition) SetParent(parent string) *ofPartition {
	p.parent = parent
	return p
}

// GetParent get parent
func (p *ofPartition) GetParent() string {
	return p.parent
}

// ResetParent reset parent
func (p *ofPartition) ResetParent() *ofPartition {
	p.parent = ""
	return p
}

// Columns get columns definitions
func (p *ofPartition) Columns() *ofDefinitions {
	return &p.definitions
}

// Values get values bounds
func (p *ofPartition) Values() *partitionBound {
	return &p.values
}

// String render ofPartition sub query
func (p *ofPartition) String() string {
	if p.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if p.parent != "" {
		b.WriteString(" PARTITION OF " + p.parent)
	}
	if p.definitions.Len() > 0 {
		b.WriteString(" (" + p.definitions.String() + ")")
	}
	if !p.values.IsEmpty() {
		b.WriteString(" FOR VALUES " + p.values.String())
	} else {
		b.WriteString(" FOR VALUES DEFAULT")
	}
	return b.String()
}

// IsEmpty is ofPartition empty
func (p *ofPartition) IsEmpty() bool {
	return p == nil || (p.parent == "" && p.definitions.Len() == 0 && p.values.IsEmpty())
}
