package gosql

import "strings"

// ( column_name [, ... ] ) index_parameters |
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
