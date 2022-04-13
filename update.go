package gosql

import "strings"

// Update query builder
type Update struct {
	// with query
	with sqlWith
	// table name
	table string
	// set of changes
	set []string
	// values
	values []any
	// from source
	from []string
	// condition
	condition Condition
	// returning
	returning expression
}

// IsEmpty check if query is empty
func (u *Update) IsEmpty() bool {
	return u == nil || (u.with.Len() == 0 &&
		u.table == "" &&
		len(u.from) == 0 &&
		u.condition.IsEmpty() &&
		len(u.set) == 0 &&
		len(u.values) == 0 &&
		u.returning.Len() == 0)
}

// From update
func (u *Update) From(from ...string) *Update {
	u.from = append(u.from, from...)
	return u
}

// ResetFrom clear from
func (u *Update) ResetFrom() *Update {
	u.from = u.from[:0]
	return u
}

// Condition set conflict condition
func (u *Update) Condition(cond Condition) *Update {
	u.condition = cond
	return u
}

// ResetCondition reset condition
func (u *Update) ResetCondition() *Update {
	u.condition = Condition{}
	return u
}

// String return result query
func (u *Update) String() string {
	if u.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if u.with.Len() > 0 {
		b.WriteString(u.with.String() + " ")
	}
	if u.table != "" {
		b.WriteString("UPDATE " + u.table)
	}
	if len(u.set) > 0 {
		b.WriteString(" SET " + strings.Join(u.set, ", "))
	}
	if len(u.from) != 0 {
		b.WriteString(" FROM " + strings.Join(u.from, ", "))
	}
	if !u.condition.IsEmpty() {
		b.WriteString(" WHERE " + u.condition.String())
	}
	if u.returning.Len() > 0 {
		b.WriteString(" RETURNING " + u.returning.String(", "))
	}
	b.WriteString(";")
	return b.String()
}

// GetValues get values
func (u *Update) GetValues() []any {
	return u.values
}

// ResetValues reset values
func (u *Update) ResetValues() *Update {
	u.values = u.values[:0]
	return u
}

// AddValues add values
func (u *Update) AddValues(values ...any) *Update {
	u.values = append(u.values, values...)
	return u
}

// Table Set table
func (u *Update) Table(table string) *Update {
	u.table = table
	return u
}

// ResetTable reset table
func (u *Update) ResetTable() *Update {
	u.table = ""
	return u
}

// AddReturning Add returning expression
func (u *Update) AddReturning(returning string, args ...any) *Update {
	u.returning.Add(returning, args...)
	return u
}

// ResetReturning Reset returning expressions
func (u *Update) ResetReturning() *Update {
	u.returning.Reset()
	return u
}

// GetReturningParams Get returning params
func (u *Update) GetReturningParams() []any {
	return u.returning.Params()
}

// Set expression
func (u *Update) Set(expression string, args ...any) *Update {
	u.set = append(u.set, expression)
	u.values = append(u.values, args...)
	return u
}

// UnSetAll from expressions
func (u *Update) UnSetAll() *Update {
	u.set = make([]string, 0)
	u.values = make([]any, 0)
	return u
}

// GetWith Get with query
func (u *Update) GetWith(name string) *Select {
	return u.with.Get(name)
}

// With Add with query
func (u *Update) With(name string, s *Select) *Update {
	u.with.Add(name, s)
	return u
}

// WithValues get with values
func (u *Update) WithValues() []any {
	return u.with.Values()
}

// ResetWith Reset With query
func (u *Update) ResetWith() *Update {
	u.with.Reset()
	return u
}

// NewUpdate Update Query Builder
func NewUpdate() *Update {
	return &Update{
		with: sqlWith{
			keys: make(map[int]string),
		},
	}
}
