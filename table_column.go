package gosql

import "strings"

// { column_name data_type [ COMPRESSION compression_method ] [ COLLATE collation ] [ column_constraint [ ... ] ] }
type column struct {
	// column name
	name string
	// data type
	dataType string
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

// SetType set column type
func (c *column) SetType(dataType string) *column {
	c.dataType = dataType
	return c
}

// GetType get column type
func (c *column) GetType() string {
	return c.dataType
}

// ResetType reset type
func (c *column) ResetType() *column {
	c.dataType = ""
	return c
}

// Constraint get constraint
func (c *column) Constraint() *constraintColumn {
	return &c.constraint
}

// IsEmpty check if empty
func (c *column) IsEmpty() bool {
	return c == nil || (c.name == "" && c.dataType == "" && c.collate == "" && c.compression == "" && c.constraint.IsEmpty())
}

// String render column
func (c *column) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if c.name != "" {
		b.WriteString(c.name)
	}
	if c.dataType != "" {
		b.WriteString(" " + c.dataType)
	}
	if c.compression != "" {
		b.WriteString(" COMPRESSION " + c.compression)
	}
	if c.collate != "" {
		b.WriteString(" COLLATE " + c.compression)
	}
	if !c.constraint.IsEmpty() {
		b.WriteString(c.constraint.String())
	}
	return b.String()
}

// NewColumn init column
func NewColumn() *column {
	return &column{}
}
