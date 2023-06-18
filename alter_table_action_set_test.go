package gosql

import "testing"

func Test_alterTableActionSet_String(t *testing.T) {
	t.Run("check_order_set_action", func(t *testing.T) {
		set := alterTableActionSet{}
		set.Using("third")
		set.DataType("integer")
		set.Collate("collation")
		t.Log(set.String())
		if set.String() != "SET DATA TYPE integer COLLATE collation USING third" {
			t.Fatal("wrong render check_order_set_action")
		}
	})
	t.Run("check_order_set_action_storage_params", func(t *testing.T) {
		set := alterTableActionSet{}
		set.StorageParameters("foo = ?")
		set.StorageParameters("bar = ?")
		set.StorageParameters("baz1 = ?")
		set.StorageParameters("baz2 = ?")
		set.StorageParameters("baz3 = ?")
		t.Log(set.String())
		if set.String() != "SET (foo = ?, bar = ?, baz1 = ?, baz2 = ?, baz3 = ?)" {
			t.Fatal("wrong render check_order_set_action_storage_params")
		}
	})
	t.Run("check_attribute_options", func(t *testing.T) {
		set := alterTableActionSet{}
		set.AttributeOptions("foo = 1", "bar = 2", "baz = 3")
		t.Log(set.String())
		if set.String() != "SET (foo = 1, bar = 2, baz = 3)" {
			t.Fatal("wrong render check_attribute_options")
		}
	})
	t.Run("check_order_set_action_storage_params", func(t *testing.T) {
		set := alterTableActionSet{}
		set.Using("third")
		t.Log(set.String())
		if set.String() != "USING third" {
			t.Fatal("wrong render check_order_set_action_storage_params")
		}
	})
}

func BenchmarkAlterTableActionSet_String(b *testing.B) {
	// goos: darwin
	// goarch: arm64
	// pkg: github.com/dimonrus/gosql
	// BenchmarkAlterTableActionSet_String
	// BenchmarkAlterTableActionSet_String/simple
	// BenchmarkAlterTableActionSet_String/simple-12         	 8427409	       148.0 ns/op	     447 B/op	       3 allocs/op
	b.Run("simple", func(b *testing.B) {
		set := alterTableActionSet{}
		set.Grow(4)
		for i := 0; i < b.N; i++ {
			set.Using("third")
			set.DataType("integer")
			//set.String()
			set.Reset()
		}
		b.ReportAllocs()
	})
	// goos: darwin
	// goarch: arm64
	// pkg: github.com/dimonrus/gosql
	// BenchmarkAlterTableActionSet_String
	// BenchmarkAlterTableActionSet_String/attribute_options
	// BenchmarkAlterTableActionSet_String/attribute_options-12         	 6982264	       154.6 ns/op	     396 B/op	       3 allocs/op
	b.Run("attribute_options", func(b *testing.B) {
		set := alterTableActionSet{}
		set.Grow(4)
		for i := 0; i < b.N; i++ {
			set.AttributeOptions("foo = 1", "bar = 2", "baz = 3", "meow = ?")
			set.Reset()
		}
		b.ReportAllocs()
	})

}
