package gosql

import "strings"

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
