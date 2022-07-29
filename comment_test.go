package gosql

import "testing"

func TestComment_String(t *testing.T) {
	t.Run("column", func(t *testing.T) {
		c := Comment().Column("table_name.column", "The column comment")
		t.Log(c.String())
		if c.String() != "COMMENT ON COLUMN table_name.column IS 'The column comment';" {
			t.Fatal("wrong comment column query")
		}
	})
	t.Run("table", func(t *testing.T) {
		c := Comment().Table("table_name", "The table comment")
		t.Log(c.String())
		if c.String() != "COMMENT ON TABLE table_name IS 'The table comment';" {
			t.Fatal("wrong comment table query")
		}
	})
}
