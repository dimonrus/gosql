package gosql

import "testing"

func TestTable_AddDefinition(t *testing.T) {
	tt := NewTable("test_case")
	tt.AddColumn("id").SetType("serial").Constraint().SetPrimary()
	tt.AddColumn("name").SetType("text").Constraint().NotNull(true)
	tt.AddColumn("type_id").SetType("int").
		Constraint().NotNull(true).References().SetRefTable("dictionary").SetOnUpdate(ActionCascade).SetOnDelete(ActionRestrict).Columns().Add("id")
	tt.AddColumn("created_at").SetType("TIMESTAMP WITH TIME ZONE").Constraint().SetDefault("localtimestamp").NotNull(true)
	tt.AddColumn("updated_at").SetType("TIMESTAMP WITH TIME ZONE")
	t.Log(tt.String())
}
