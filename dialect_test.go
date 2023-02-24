package gosql

import "testing"

// goos: darwin
// goarch: amd64
// pkg: github.com/dimonrus/gosql
// cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
// BenchmarkPostgresQueryParamHook
// BenchmarkPostgresQueryParamHook/normal
// BenchmarkPostgresQueryParamHook/normal-8         	 4647177	       252.3 ns/op	     304 B/op	       2 allocs/op
// BenchmarkPostgresQueryParamHook/no_need
// BenchmarkPostgresQueryParamHook/no_need-8        	 9483421	       134.5 ns/op	     256 B/op	       1 allocs/op
func BenchmarkPostgresQueryParamHook(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		u := NewUpdate().Table("apple_attribute")
		u.Set().Add("code = 'name_test_update'")
		u.Where().AddExpression("id = ?")
		u.Where().AddExpression("ab = ?")
		u.Where().AddExpression("ad = ?")
		u.Where().AddExpression("aa = ANY(?)")
		q, _, _ := u.SQL()
		for i := 0; i < b.N; i++ {
			PostgresQueryParamHook(q)
		}
		b.ReportAllocs()
	})
	b.Run("no_need", func(b *testing.B) {
		u := NewUpdate().Table("apple_attribute")
		u.Set().Add("code = 'name_test_update'")
		u.Where().AddExpression("id = 1")
		u.Where().AddExpression("ab = '2'")
		u.Where().AddExpression("ad = 'foo'")
		u.Where().AddExpression("aa = ANY(ARRAY[1,2,3])")
		q, _, _ := u.SQL()
		for i := 0; i < b.N; i++ {
			PostgresQueryParamHook(q)
		}
		b.ReportAllocs()
	})
}

func TestPreparePositionalArgsQuery(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		q := "update apple_attribute set code = 'name_test_update' where id = ? AND ab = ? OR ad = ? AND aa = ANY(ARRAY[1,2,3])"
		r := PostgresQueryParamHook(q)
		t.Log(r)
		if r != "update apple_attribute set code = 'name_test_update' where id = $1 AND ab = $2 OR ad = $3 AND aa = ANY(ARRAY[1,2,3])" {
			t.Fatal("wrong normal")
		}
	})
	t.Run("serial", func(t *testing.T) {
		q := "?????"
		r := PostgresQueryParamHook(q)
		if r != "$1$2$3$4$5" {
			t.Fatal("wrong serial")
		}
	})
	t.Run("no_need_transform", func(t *testing.T) {
		q := "update apple_attribute set code = 'name_test_update' where id = 1 AND ab = '2' OR ad = 'adad' AND aa = ANY(ARRAY[1,2,3])"
		r := PostgresQueryParamHook(q)
		if r != q {
			t.Fatal("wrong no_need_transform")
		}
	})
	t.Run("max_args", func(t *testing.T) {
		q := "INSERT INTO test_table (id, name, date, value, count, type_id, created_at, updated_at) VALUES "
		for i := 0; i < 255*255; {
			for j := 0; j < 8; j++ {
				if i == 0 {
					q += "(?, ?, ?, ?, ?, ?, ?, ?)"
				} else {
					q += ", (?, ?, ?, ?, ?, ?, ?, ?)"
				}
				i++
			}
		}
		r := PostgresQueryParamHook(q)
		if r[len(r)-1] != ')' {
			t.Fatal("wrong max_args")
		}
	})
	t.Run("pgsql", func(t *testing.T) {
		q := NewSelect().From("some_table")
		q.Columns().Add("foo", "bar")
		q.Where().AddExpression("field_1 = ?", 1)
		q.Where().AddExpression("field_2 = ?", true)
		q.Where().AddExpression("field_3 = ?", "some")
		query, params, returnings := PGSQL(q)
		t.Log(query)
		if query != "SELECT foo, bar FROM some_table WHERE (field_1 = $1 AND field_2 = $2 AND field_3 = $3)" {
			t.Fatal("wrong postgres query")
		}
		if len(params) != 3 {
			t.Fatal("wrong params count")
		}
		if len(returnings) != 0 {
			t.Fatal("wrong returnings count")
		}
	})
}
