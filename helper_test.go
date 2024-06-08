package gosql

import (
	"testing"
	"time"
)

func TestSorting_Parse(t *testing.T) {
	t.Run("sorting_valid", func(t *testing.T) {
		s := Sorting{"   createdAt:desc", "name", "qty:ASC"}
		result := s.Allowed(map[string]string{"createdAt": "created_at {dir} NULLS LAST", "name": "internal_name", "qty": "quantity"})
		if len(result) != 3 {
			t.Fatal("wrong sorting parser")
		}
		if result[0] != "created_at DESC NULLS LAST" {
			t.Fatal("wrong created parser")
		}
		if result[1] != "internal_name" {
			t.Fatal("wrong internal_name parser")
		}
		if result[2] != "quantity" {
			t.Fatal("wrong quantity parser")
		}
	})
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dimonrus/gosql
// BenchmarkSorting_Allowed
// BenchmarkSorting_Allowed-12    	 5917818	       198.0 ns/op	     160 B/op	       5 allocs/op
func BenchmarkSorting_Allowed(b *testing.B) {
	s := Sorting{"   createdAt:desc", "name", "qty:ASC"}
	for i := 0; i < b.N; i++ {
		s.Allowed(map[string]string{"createdAt": "created_at NULLS LAST", "name": "internal_name", "qty": "quantity"})
	}
	b.ReportAllocs()
}

func TestPeriodFilter(t *testing.T) {
	t.Run("full_cond", func(t *testing.T) {
		now := time.Now()
		period := PeriodFilter{
			Start: &now,
			End:   &now,
		}
		if !period.IsEmpty() {
			cond := period.FieldCondition("created_at")
			if len(cond.GetArguments()) != 2 {
				t.Fatal("wrong arg count")
			}
			if cond.String() != "(created_at >= ? AND created_at <= ?)" {
				t.Fatal("wrong period condition")
			}
		}
	})
}

func TestSearchString_PrepareLikeValue(t *testing.T) {
	t.Run("double_column", func(t *testing.T) {
		s := SearchString("Foo Bar")
		cond := s.PrepareLikeValue("fullname")
		if len(cond.GetArguments()) != 2 {
			t.Fatal("wrong argument count")
		}
		if cond.String() != "(fullname like lower(?) AND fullname like lower(?))" {
			t.Fatal("wrong condition render")
		}
		if v, ok := cond.GetArguments()[0].(string); !ok {
			t.Fatal("must be a string")
		} else {
			if v != "%foo%" {
				t.Fatal("wrong 1st argument")
			}
		}
		if v, ok := cond.GetArguments()[1].(string); !ok {
			t.Fatal("must be a string")
		} else {
			if v != "%bar%" {
				t.Fatal("wrong 1st argument")
			}
		}
	})
}

func TestSorting_Contains(t *testing.T) {
	t.Run("asc direction 1", func(t *testing.T) {
		s := Sorting{"foo", "bar:asc"}
		contains, dir := s.Contains("foo")
		if !contains || dir != nil {
			t.Fatal("wrong asc implementation")
		}
	})
	t.Run("asc direction 2", func(t *testing.T) {
		s := Sorting{"foo:asc", "bar:asc"}
		contains, dir := s.Contains("foo")
		if !contains || dir == nil || *dir != true {
			t.Fatal("wrong asc implementation")
		}
	})
	t.Run("desc direction 1", func(t *testing.T) {
		s := Sorting{"foo", "bar:DESC"}
		contains, dir := s.Contains("bar")
		if !contains || dir == nil || *dir {
			t.Fatal("wrong desc implementation")
		}
	})
	t.Run("not contained", func(t *testing.T) {
		s := Sorting{"foo", "bar:DESC"}
		contains, dir := s.Contains("bar1")
		if contains || dir != nil {
			t.Fatal("wrong desc implementation")
		}
	})
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dimonrus/gosql
// BenchmarkSorting_Contains
// BenchmarkSorting_Contains-12    	49791888	        24.01 ns/op	      16 B/op	       1 allocs/op
func BenchmarkSorting_Contains(b *testing.B) {
	s := Sorting{"foo", "bar:asc"}
	for i := 0; i < b.N; i++ {
		_, _ = s.Contains("foo")
	}
	b.ReportAllocs()
}
