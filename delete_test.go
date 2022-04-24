package gosql

import "testing"

func TestDelete_String(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		d := NewDelete()
		d.From("films")
		d.Where().AddExpression("kind <> ?", "Musical")
		t.Log(d.String())
		if d.String() != "DELETE FROM films WHERE (kind <> ?);" || len(d.Where().GetArguments()) != 1 {
			t.Fatal("wrong simple")
		}
	})

	t.Run("simple_1", func(t *testing.T) {
		d := NewDelete()
		d.From("films")
		t.Log(d.String())
		if d.String() != "DELETE FROM films;" {
			t.Fatal("wrong simple_1")
		}
	})

	t.Run("simple_2", func(t *testing.T) {
		d := NewDelete()
		d.From("tasks").Returning().AddExpressions("*")
		d.Where().AddExpression("status = ?", "DONE")
		t.Log(d.String())
		if d.String() != "DELETE FROM tasks WHERE (status = ?) RETURNING *;" || len(d.Where().GetArguments()) != 1 {
			t.Fatal("wrong simple_2")
		}
	})

	t.Run("current", func(t *testing.T) {
		sub := NewSelect()
		sub.Columns("id")
		sub.From("producers")
		sub.Where().AddExpression("name = ?", "foo")
		sub.SubQuery = true

		d := NewDelete()
		d.From("tasks")
		d.Where().AddExpression("producer_id IN "+sub.String(), sub.GetArguments()...)
		t.Log(d.String())
		if d.String() != "DELETE FROM tasks WHERE (producer_id IN (SELECT id FROM producers WHERE (name = ?)));" || len(d.Where().GetArguments()) != 1 {
			t.Fatal("wrong simple_2")
		}
	})
}
