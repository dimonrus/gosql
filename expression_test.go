package gosql

import "testing"

func TestExpression(t *testing.T) {
	ex := expression{}
	ex.Append("foo = ?", 1)
	ex.Append("bar = ?", "simple")
	if s, p := ex.Get(", "); s != "foo = ?, bar = ?" || len(p) != 2 {
		t.Fatal("wrong add logic")
	}
	if ex.ArgLen() != 2 {
		t.Fatal("wrong arg len")
	}
	ex.Reset()
	if s, p := ex.Get(", "); s != "" || len(p) != 0 || ex.Len() != 0 {
		t.Fatal("wrong reset logic")
	}
	var exp *expression
	if exp.Len() != 0 {
		t.Fatal("wrong len")
	}
	if exp.String("|") != "" {
		t.Fatal("wrong string")
	}
	if exp.GetArguments() != nil {
		t.Fatal("wrong params")
	}
	exp.Reset()
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dimonrus/gosql
// BenchmarkExpression_Get
// BenchmarkExpression_Get-12    	36178676	        32.18 ns/op	      16 B/op	       1 allocs/op
func BenchmarkExpression_Get(b *testing.B) {
	ex := expression{}
	ex.Append("foo = ?", 1)
	ex.Append("foo = ?", "one")
	for i := 0; i < b.N; i++ {
		_, _ = ex.Get(", ")
	}
	b.ReportAllocs()
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dimonrus/gosql
// BenchmarkExpression_Add
// BenchmarkExpression_Add-12    	13390683	        86.47 ns/op	     320 B/op	       3 allocs/op
func BenchmarkExpression_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex := NewExpression()
		ex.Append("foo = ?", "one", 6)
	}
	b.ReportAllocs()
}
