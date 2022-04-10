package gosql

import "strings"

// With SQL
type sqlWith struct {
	// map of with names
	keys map[int]string
	// slice of queries
	queries []*Select
}

// Len of with queries
func (w *sqlWith) Len() int {
	if w == nil {
		return 0
	}
	return len(w.keys)
}

// Get With
func (w *sqlWith) Get(name string) *Select {
	for j, key := range w.keys {
		if key == name {
			return w.queries[j]
		}
	}
	return nil
}

// Add With
func (w *sqlWith) Add(name string, qb *Select) *sqlWith {
	if name != "" {
		w.queries = append(w.queries, qb)
		w.keys[len(w.queries)-1] = name
	}
	return w
}

// Reset With query
func (w *sqlWith) Reset() *sqlWith {
	w.queries = make([]*Select, 0)
	w.keys = make(map[int]string, 0)
	return w
}

// String for with
func (w *sqlWith) String() string {
	var b strings.Builder
	if w.Len() > 0 {
		b.WriteString("WITH ")
		for index, q := range w.queries {
			if index < len(w.queries)-1 {
				b.WriteString(w.keys[index] + " AS (" + q.String() + "),")
			} else {
				b.WriteString(w.keys[index] + " AS (" + q.String() + ")")
			}
		}
	}
	return b.String()
}

// Values get values from all queries
func (w *sqlWith) Values() []any {
	var params []any
	if w.Len() > 0 {
		for index := range w.queries {
			params = append(params, w.queries[index].GetArguments()...)
		}
	}
	return params
}
