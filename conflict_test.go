package gosql

import "testing"

func TestConflict_String(t *testing.T) {
	t.Run("conflict_update", func(t *testing.T) {
		con := conflict{}
		con.Object("id,name")
		con.Constraint("distributors_pkey")
		con.Action("UPDATE")
		con.Set("created_at = now()", "id = 1")
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("id = ?", 10)
		con.Condition(*cond)
		if con.String() != "ON CONFLICT (id,name) ON CONSTRAINT distributors_pkey DO UPDATE SET created_at = now(), id = 1 WHERE (id = ?)" {
			t.Fatal("wrong conflict expression")
		}
	})
	t.Run("conflict_do_nothing", func(t *testing.T) {
		con := conflict{}
		con.Object("id,name")
		con.Constraint("distributors_pkey")
		con.Action("NOTHING")
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("id = ?", 10)
		con.Condition(*cond)
		if con.String() != "ON CONFLICT (id,name) ON CONSTRAINT distributors_pkey WHERE (id = ?) DO NOTHING" {
			t.Fatal("wrong conflict expression")
		}
	})
}
