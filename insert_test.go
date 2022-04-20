package gosql

import (
	"testing"
)

func TestInsert_String(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		i := NewInsert()
		i.Into("user")
		i.Columns("name", "entity_id", "created_at")
		i.AddReturning("id")
		i.AddReturning("created_at")
		i.AddValues("foo", 10, "2021-01-01T10:10:00Z")
		i.AddValues("bar", 20, "2021-01-01T10:10:00Z")
		t.Log(i.String())
		if i.String() != "INSERT INTO user (name, entity_id, created_at) VALUES ($1, $2, $3), ($4, $5, $6) RETURNING id, created_at;" {
			t.Fatal("wrong build")
		}
	})
	t.Run("with", func(t *testing.T) {
		i := NewInsert()
		i.Into("user")
		i.Columns("name", "entity_id", "created_at")
		i.AddReturning("id")
		i.AddReturning("created_at")
		q := NewSelect()
		q.From("dictionary d")
		q.Columns("*")
		q.Where().AddExpression("some = ?", 1)
		q.Relate("JOIN relation r ON r.dictionary_id = d.id")
		i.With("dict", q)
		t.Log(i.String())
		if i.String() != "WITH dict AS (SELECT * FROM dictionary d JOIN relation r ON r.dictionary_id = d.id WHERE (some = ?)) INSERT INTO user (name, entity_id, created_at) RETURNING id, created_at;" {
			t.Fatal("wrong with")
		}
	})
	t.Run("conflict", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns("did", "dname")
		i.AddValues(5, "Gizmo Transglobal")
		i.AddValues(6, "Associated Computing, Inc")
		i.Conflict().Object("did").Action("UPDATE").Set("dname = EXCLUDED.dname")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES ($1, $2), ($3, $4) ON CONFLICT (did) DO UPDATE SET dname = EXCLUDED.dname;" {
			t.Fatal("wrong conflict builder")
		}
	})
	t.Run("nothing_on_conflict", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns("did", "dname")
		i.AddValues(7, "Redline GmbH")
		i.Conflict().Object("did").Action("NOTHING")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES ($1, $2) ON CONFLICT (did) DO NOTHING;" {
			t.Fatal("wrong nothing_on_conflict")
		}
	})

	t.Run("on_conflict_set_with_condition", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors AS d")
		i.Columns("did", "dname")
		i.AddValues(8, "Anvil Distribution")
		i.Conflict().Object("did").Action("UPDATE").Set("dname = EXCLUDED.dname || ' (formerly ' || d.dname || ')'")
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("d.zipcode <> '21201'")
		i.Conflict().Condition(*cond)
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors AS d (did, dname) VALUES ($1, $2) ON CONFLICT (did) DO UPDATE SET dname = EXCLUDED.dname || ' (formerly ' || d.dname || ')' WHERE (d.zipcode <> '21201');" {
			t.Fatal("wrong on_conflict_set_with_condition")
		}
	})
	t.Run("on_conflict_constraint", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns("did", "dname")
		i.AddValues(9, "Antwerp Design")
		i.Conflict().Constraint("distributors_pkey").Action("NOTHING")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES ($1, $2) ON CONFLICT ON CONSTRAINT distributors_pkey DO NOTHING;" {
			t.Fatal("wrong on_conflict_constraint")
		}
	})
	t.Run("on_conflict_constraint", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns("did", "dname")
		i.AddValues(10, "Conrad International")
		i.Conflict().Object("did").Action(ConflictActionNothing)
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("is_active")
		i.Conflict().Condition(*cond)
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES ($1, $2) ON CONFLICT (did) WHERE (is_active) DO NOTHING;" {
			t.Fatal("wrong on_conflict_constraint")
		}
	})
	t.Run("returning", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns("did", "dname")
		i.AddValues(1, "XYZ Widgets")
		i.AddReturning("did")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES ($1, $2) RETURNING did;" {
			t.Fatal("wrong on_conflict_constraint")
		}
	})
}
