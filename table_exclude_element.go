package gosql

import "strings"

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
	return e == nil || (e.column == "" && e.expression.Len() == 0 && e.opclass == "" && e.direction == "" && e.nulls == nil)
}

// NewExcludeElement init exclude element
func NewExcludeElement() *excludeElement {
	return &excludeElement{}
}
