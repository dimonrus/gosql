package gosql

import "strings"

// EnumDelimiter for join strings
const EnumDelimiter = "#"

// Expression slice
type expression struct {
	// list of expressions
	list strings.Builder
	// params
	params []any
}

// Add expressions
func (e *expression) Add(expr string, args ...any) {
	e.AddExpressions(expr)
	e.AddParams(args...)
}

// Len len of expressions
func (e *expression) Len() int {
	if e == nil {
		return 0
	}
	return e.list.Len()
}

// Reset expressions
func (e *expression) Reset() {
	if e.Len() == 0 {
		return
	}
	e.list.Reset()
	e.params = e.params[:0]
}

// String all expressions
func (e *expression) String(delimiter string) string {
	if e.Len() == 0 {
		return ""
	}
	return strings.ReplaceAll(e.list.String(), EnumDelimiter, delimiter)
}

// AddParams add params
func (e *expression) AddParams(args ...any) *expression {
	e.params = append(e.params, args...)
	return e
}

// AddExpressions add expressions
func (e *expression) AddExpressions(args ...string) *expression {
	for _, arg := range args {
		if e.Len() == 0 {
			e.list.WriteString(arg)
		} else {
			e.list.WriteString(EnumDelimiter + arg)
		}
	}
	return e
}

// Params return params
func (e *expression) Params() []any {
	if e == nil {
		return nil
	}
	return e.params
}

// Get expressions
func (e *expression) Get(delimiter string) (string, []any) {
	return e.String(delimiter), e.Params()
}
