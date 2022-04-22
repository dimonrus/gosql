package gosql

import (
	"strconv"
	"strings"
)

const (
	// ConflictActionNothing On conflict action do nothing
	ConflictActionNothing = "NOTHING"
	// ConflictActionUpdate On conflict action do update nothing
	ConflictActionUpdate = "UPDATE"
)

// Insert query builder
type Insert struct {
	// with query
	with with
	// into table name
	into string
	// from insert
	from []string
	// list of columns
	columns []string
	// list of values
	values []any
	// conflict expression
	conflict conflict
	// returning
	returning expression
}

// Get sql insert query
func (i *Insert) String() string {
	b := strings.Builder{}
	if i.with.Len() > 0 {
		b.WriteString(i.with.String() + " ")
	}
	if i.into != "" {
		b.WriteString("INSERT INTO " + i.into)
	}
	if len(i.columns) > 0 {
		b.WriteString(" (" + strings.Join(i.columns, ", ") + ")")
	}
	if len(i.from) > 0 {
		b.WriteString(" " + strings.Join(i.from, ", "))
	} else if len(i.values) > 0 {
		b.WriteString(" VALUES ")
		for j := 1; j <= len(i.values); {
			b.WriteString("(")
			for u := 0; u < len(i.columns); u++ {
				b.WriteString("$" + strconv.Itoa(j))
				j++
				if u == len(i.columns)-1 {
					break
				}
				b.WriteString(", ")
			}
			b.WriteString(")")
			if j >= len(i.values) {
				break
			}
			b.WriteString(", ")
		}
	}
	if !i.conflict.IsEmpty() {
		b.WriteString(" " + i.Conflict().String())
	}
	if i.returning.Len() > 0 {
		b.WriteString(" RETURNING " + i.returning.String(", "))
	}
	b.WriteString(";")
	return b.String()
}

// IsEmpty return true if all parts not filed
func (i *Insert) IsEmpty() bool {
	return i == nil || (i.with.Len() == 0 &&
		i.into == "" &&
		len(i.from) == 0 &&
		len(i.columns) == 0 &&
		len(i.values) == 0 &&
		i.returning.Len() == 0 &&
		i.conflict.IsEmpty())
}

// SQL Get sql query
func (i *Insert) SQL() (query string, params []any, returning []any) {
	return i.String(), i.GetArguments(), i.returning.Params()
}

// GetArguments get all arguments
func (i *Insert) GetArguments() []any {
	return append(append(i.with.Values(), i.values...), i.conflict.GetArguments()...)
}

// SetConflict set conflict
func (i *Insert) SetConflict(conflict conflict) *Insert {
	i.conflict = conflict
	return i
}

// Conflict get conflict expression
func (i *Insert) Conflict() *conflict {
	return &i.conflict
}

// ResetConflict reset conflict expression
func (i *Insert) ResetConflict() *Insert {
	i.conflict = conflict{}
	return i
}

// AddReturning Add returning expression
func (i *Insert) AddReturning(returning string, args ...any) *Insert {
	i.returning.Add(returning, args...)
	return i
}

// ResetReturning Reset returning expressions
func (i *Insert) ResetReturning() *Insert {
	i.returning.Reset()
	return i
}

// GetReturningParams Get returning params
func (i *Insert) GetReturningParams() []any {
	return i.returning.Params()
}

// From insert from
func (i *Insert) From(from ...string) *Insert {
	i.from = append(i.from, from...)
	return i
}

// ResetFrom clear from
func (i *Insert) ResetFrom() *Insert {
	i.from = i.from[:0]
	return i
}

// AddValues Add row values
func (i *Insert) AddValues(values ...any) *Insert {
	i.values = append(i.values, values...)
	return i
}

// SetValues Set row values by index
func (i *Insert) SetValues(index int, value any) *Insert {
	i.values[index] = value
	return i
}

// ResetValues Reset all values
func (i *Insert) ResetValues() *Insert {
	i.values = make([]any, 0)
	return i
}

// GetValues get values by indexes
func (i *Insert) GetValues(start, end int) []any {
	return i.values[start:end]
}

// GetAllValues get all values
func (i *Insert) GetAllValues() []any {
	return i.values
}

// Into Set into value
func (i *Insert) Into(into string) *Insert {
	i.into = into
	return i
}

// ResetInto Set into empty string
func (i *Insert) ResetInto() *Insert {
	i.into = ""
	return i
}

// Columns Set columns
func (i *Insert) Columns(column ...string) *Insert {
	i.columns = append(i.columns, column...)
	return i
}

// ResetColumns reset columns
func (i *Insert) ResetColumns() *Insert {
	i.columns = make([]string, 0)
	return i
}

// GetWith Get with query
func (i *Insert) GetWith(name string) *Select {
	return i.with.Get(name)
}

// With add with query
func (i *Insert) With(name string, qb *Select) *Insert {
	i.with.Add(name, qb)
	return i
}

// ResetWith Reset With query
func (i *Insert) ResetWith() *Insert {
	i.with.Reset()
	return i
}

// WithValues get with values
func (i *Insert) WithValues() []any {
	return i.with.Values()
}

// NewInsert new insert query builder
func NewInsert() *Insert {
	return &Insert{
		with: with{
			keys: make(map[int]string),
		},
	}
}
