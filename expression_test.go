package gosql

import "testing"

func TestExpression(t *testing.T) {
	ex := expression{}
	ex.Append("foo = ?", 1)
	ex.Append("bar = ?", "simple")
	if s, p := ex.Get(", "); s != "foo = ?, bar = ?" || len(p) != 2 {
		t.Fatal("wrong add logic")
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

func BenchmarkExpression_Get(b *testing.B) {
	ex := expression{}
	ex.Append("foo = ?", 1)
	ex.Append("foo = ?", "one")
	for i := 0; i < b.N; i++ {
		_, _ = ex.Get(", ")
	}
	b.ReportAllocs()
}

func BenchmarkExpression_Add(b *testing.B) {
	ex := expression{}
	for i := 0; i < b.N; i++ {
		ex.Append("foo = ?", "one")
	}
	b.ReportAllocs()
}
