package gosql

import "testing"

func TestReferencesColumn_String(t *testing.T) {
	t.Run("classic", func(t *testing.T) {
		ref := &referencesColumn{}
		ref.RefTable("user")
		ref.OnDelete("CASCADE")
		ref.OnUpdate("CASCADE")
		ref.Columns().Add("dict_id", "other_id")
		ref.Match(ReferencesMatchFull)
		t.Log(ref.String())
		if "REFERENCES user (dict_id, other_id) MATCH FULL ON DELETE CASCADE ON UPDATE CASCADE" != ref.String() {
			t.Fatal("wrong classic")
		}
	})
	t.Run("restrict_delete", func(t *testing.T) {
		ref := &referencesColumn{}
		ref.RefTable("user")
		ref.OnDelete("RESTRICT")
		ref.Columns().Add("other_id")
		ref.Match(ReferencesMatchPartial)
		t.Log(ref.String())
		if "REFERENCES user (other_id) MATCH PARTIAL ON DELETE RESTRICT" != ref.String() {
			t.Fatal("wrong restrict_delete")
		}
	})
	t.Run("simple", func(t *testing.T) {
		ref := &referencesColumn{}
		ref.RefTable("user")
		ref.Columns().Add("other_id")
		t.Log(ref.String())
		if "REFERENCES user (other_id)" != ref.String() {
			t.Fatal("wrong simple")
		}
	})
}
