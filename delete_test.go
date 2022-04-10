package gosql

import "testing"

func TestDelete_String(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		d := NewDelete()
		d.From("films")
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("kind <> ?", "Musical")
		d.Condition(*cond)
		t.Log(d.String())
		if d.String() != "DELETE FROM films WHERE (kind <> ?);" || len(cond.GetArguments()) != 1 {
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
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("status = ?", "DONE")
		d := NewDelete()
		d.From("tasks").AddReturning("*")
		d.Condition(*cond)
		t.Log(d.String())
		if d.String() != "DELETE FROM tasks WHERE (status = ?) RETURNING *;" || len(cond.GetArguments()) != 1 {
			t.Fatal("wrong simple_2")
		}
	})

	t.Run("current", func(t *testing.T) {
		sub := NewSelect()
		sub.Columns("id")
		sub.From("producers")
		sub.Where().AddExpression("name = ?", "foo")
		sub.SubQuery = true

		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("producer_id IN "+sub.String(), sub.GetArguments()...)
		d := NewDelete()
		d.From("tasks")
		d.Condition(*cond)
		t.Log(d.String())
		if d.String() != "DELETE FROM tasks WHERE (producer_id IN (SELECT id FROM producers WHERE (name = ?)));" || len(cond.GetArguments()) != 1 {
			t.Fatal("wrong simple_2")
		}
	})
}
