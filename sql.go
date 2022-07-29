package gosql

import "strings"

// ISQL each query should implement the interface
type ISQL interface {
	// SQL Get query as string with all params
	SQL() (query string, params []any, returning []any)
}

// SQList Collection of SQL element
type SQList []ISQL

// Join Return query string joined in one string
func (s SQList) Join() (query string, params []any, returning []any) {
	b := strings.Builder{}
	for _, isql := range s {
		q, p, r := isql.SQL()
		if q[len(q)-1] != ';' {
			b.WriteString(q + ";")
		} else {
			b.WriteString(q)
		}
		params = append(params, p...)
		returning = append(returning, r...)
	}
	query = b.String()
	return
}
