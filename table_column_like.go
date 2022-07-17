package gosql

import "strings"

// LIKE source_table [ like_option ... ] }
type likeTable struct {
	// source table
	source string
	// like options
	options expression
}

// SetSource set source
func (l *likeTable) SetSource(source string) *likeTable {
	l.source = source
	return l
}

// GetSource get source
func (l *likeTable) GetSource() string {
	return l.source
}

// ResetSource reset source
func (l *likeTable) ResetSource() *likeTable {
	l.source = ""
	return l
}

// IsEmpty check is like is empty
func (l *likeTable) IsEmpty() bool {
	return l == nil || (l.source == "" && l.options.Len() == 0)
}

// Options get options
func (l *likeTable) Options() *expression {
	return &l.options
}

// String render columnDefinition
func (l *likeTable) String() string {
	if l.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("LIKE")
	if l.source != "" {
		b.WriteString(" " + l.source)
	}
	if l.options.Len() > 0 {
		b.WriteString(" " + l.options.String(" "))
	}
	return b.String()
}
