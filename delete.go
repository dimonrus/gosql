package gosql

import "strings"

// Delete query
type Delete struct {
	// with query
	with sqlWith
	// from
	from string
	// using
	using []string
	// where condition
	where Condition
	// returning
	returning expression
}

// IsEmpty check if query is empty
func (d *Delete) IsEmpty() bool {
	return d == nil || (d.with.Len() == 0 &&
		d.from == "" &&
		len(d.using) == 0 &&
		d.where.IsEmpty() &&
		d.returning.Len() == 0)
}

// From Set from value
func (d *Delete) From(from string) *Delete {
	d.from = from
	return d
}

// ResetFrom Set from empty string
func (d *Delete) ResetFrom() *Delete {
	d.from = ""
	return d
}

// Using add using
func (d *Delete) Using(using ...string) *Delete {
	d.using = append(d.using, using...)
	return d
}

// ResetUsing clear using
func (d *Delete) ResetUsing() *Delete {
	d.using = d.using[:0]
	return d
}

// Condition set condition
func (d *Delete) Condition(cond Condition) *Delete {
	d.where = cond
	return d
}

// ResetCondition reset condition
func (d *Delete) ResetCondition() *Delete {
	d.where = Condition{}
	return d
}

// AddReturning Add returning expression
func (d *Delete) AddReturning(returning string, args ...any) *Delete {
	d.returning.Add(returning, args...)
	return d
}

// ResetReturning Reset returning expressions
func (d *Delete) ResetReturning() *Delete {
	d.returning.Reset()
	return d
}

// GetReturningParams Get returning params
func (d *Delete) GetReturningParams() []any {
	return d.returning.Params()
}

// String return result query
func (d *Delete) String() string {
	if d.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if d.with.Len() > 0 {
		b.WriteString(d.with.String() + " ")
	}
	b.WriteString("DELETE")
	if d.from != "" {
		b.WriteString(" FROM " + d.from)
	}
	if len(d.using) > 0 {
		b.WriteString(" USING " + strings.Join(d.using, ", "))
	}
	if !d.where.IsEmpty() {
		b.WriteString(" WHERE " + d.where.String())
	}
	if d.returning.Len() > 0 {
		b.WriteString(" RETURNING " + d.returning.String(", "))
	}
	b.WriteString(";")
	return b.String()
}

func NewDelete() *Delete {
	return &Delete{
		with: sqlWith{
			keys: make(map[int]string),
		}}
}
