package gosql

// ISQL each query should implement the interface
type ISQL interface {
	// SQL Get query as string with all params
	SQL() (query string, params []any, returning []any)
}
