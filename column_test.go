package gosql

import "testing"

func TestReferencesColumn_String(t *testing.T) {
	t.Run("classic", func(t *testing.T) {
		ref := NewReferenceColumn()
		ref.SetRefTable("user")
		ref.SetOnDelete("CASCADE")
		ref.SetOnUpdate("CASCADE")
		ref.Columns().Add("dict_id", "other_id")
		ref.SetMatch(ReferencesMatchFull)
		t.Log(ref.String())
		if "REFERENCES user (dict_id, other_id) MATCH FULL ON DELETE CASCADE ON UPDATE CASCADE" != ref.String() {
			t.Fatal("wrong classic")
		}
	})
	t.Run("restrict_delete", func(t *testing.T) {
		ref := NewReferenceColumn()
		ref.SetRefTable("user")
		ref.SetOnDelete("RESTRICT")
		ref.Columns().Add("other_id")
		ref.SetMatch(ReferencesMatchPartial)
		t.Log(ref.String())
		if "REFERENCES user (other_id) MATCH PARTIAL ON DELETE RESTRICT" != ref.String() {
			t.Fatal("wrong restrict_delete")
		}
	})
	t.Run("simple", func(t *testing.T) {
		ref := NewReferenceColumn()
		ref.SetRefTable("user")
		ref.Columns().Add("other_id")
		t.Log(ref.String())
		if "REFERENCES user (other_id)" != ref.String() {
			t.Fatal("wrong simple")
		}
	})
	t.Run("check_empty", func(t *testing.T) {
		ref := NewReferenceColumn()
		ref.SetRefTable("user")
		ref.SetOnDelete("CASCADE")
		ref.SetOnUpdate("CASCADE")
		ref.Columns().Add("dict_id", "other_id")
		ref.SetMatch(ReferencesMatchFull)
		t.Log(ref.String())
		ref.ResetRefTable()
		ref.ResetMatch()
		ref.Columns().Reset()
		ref.ResetOnDelete()
		ref.ResetOnUpdate()
		if !ref.IsEmpty() {
			t.Fatal("must be empty")
		}
		if ref.GetOnDelete() != "" || ref.GetMatch() != "" || ref.GetOnUpdate() != "" || ref.GetRefTable() != "" {
			t.Fatal("must be empty string")
		}
		ref = nil
		if ref.String() != "" {
			t.Fatal("must be empty")
		}
	})
}
