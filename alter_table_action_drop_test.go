package gosql

import "testing"

func Test_alterTableActionDrop(t *testing.T) {
	t.Run("drop_column_action", func(t *testing.T) {
		drop := alterTableActionDrop{}
		drop.Column("created_at").Restrict().IfExists()
		if drop.String() != "DROP COLUMN IF EXISTS created_at RESTRICT" {
			t.Fatal("wrong expression")
		}
	})
	t.Run("drop_constraint", func(t *testing.T) {
		drop := alterTableActionDrop{}
		drop.Constraint("ref_to_table").Cascade().IfExists()
		if drop.String() != "DROP CONSTRAINT IF EXISTS ref_to_table CASCADE" {
			t.Fatal("wrong expression")
		}
	})
}
