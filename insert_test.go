package gosql

import (
	"testing"
)

func TestInsert_String(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		i := NewInsert()
		i.Into("user")
		i.Columns().Add("name", "entity_id", "created_at")
		i.Returning().Add("id", "created_at")
		i.Columns().Arg("foo", 10, "2021-01-01T10:10:00Z")
		i.Columns().Arg("bar", 20, "2021-01-01T10:10:00Z")
		t.Log(i.String())
		if i.String() != "INSERT INTO user (name, entity_id, created_at) VALUES (?, ?, ?), (?, ?, ?) RETURNING id, created_at;" {
			t.Fatal("wrong build")
		}
	})
	t.Run("with", func(t *testing.T) {
		i := NewInsert()
		i.Into("user")
		i.Columns().Add("name", "entity_id", "created_at")
		i.Returning().Add("id", "created_at")
		q := NewSelect()
		q.From("dictionary d")
		q.Columns().Add("*")
		q.Where().AddExpression("some = ?", 1)
		q.Relate("JOIN relation r ON r.dictionary_id = d.id")
		i.With().Add("dict", q)
		t.Log(i.String())
		if i.String() != "WITH dict AS (SELECT * FROM dictionary d JOIN relation r ON r.dictionary_id = d.id WHERE (some = ?)) INSERT INTO user (name, entity_id, created_at) RETURNING id, created_at;" {
			t.Fatal("wrong with")
		}
	})
	t.Run("conflict", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns().Add("did", "dname")
		i.Columns().Arg(5, "Gizmo Transglobal")
		i.Columns().Arg(6, "Associated Computing, Inc")
		i.Conflict().Object("did").Action("UPDATE").Set().Add("dname = EXCLUDED.dname")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES (?, ?), (?, ?) ON CONFLICT (did) DO UPDATE SET dname = EXCLUDED.dname;" {
			t.Fatal("wrong conflict builder")
		}
		if len(i.GetArguments()) != 4 {
			t.Fatal("wrong arg count")
		}
	})
	t.Run("nothing_on_conflict", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns().Add("did", "dname")
		i.Columns().Arg(7, "Redline GmbH")
		i.Conflict().Object("did").Action("NOTHING")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES (?, ?) ON CONFLICT (did) DO NOTHING;" {
			t.Fatal("wrong nothing_on_conflict")
		}
	})

	t.Run("on_conflict_set_with_condition", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors AS d")
		i.Columns().Add("did", "dname")
		i.Columns().Arg(8, "Anvil Distribution")
		i.Conflict().Object("did").Action("UPDATE").Set().Add("dname = EXCLUDED.dname || ' (formerly ' || d.dname || ')'")
		i.Conflict().Where().AddExpression("d.zipcode <> '21201'")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors AS d (did, dname) VALUES (?, ?) ON CONFLICT (did) DO UPDATE SET dname = EXCLUDED.dname || ' (formerly ' || d.dname || ')' WHERE (d.zipcode <> '21201');" {
			t.Fatal("wrong on_conflict_set_with_condition")
		}
	})
	t.Run("on_conflict_constraint", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns().Add("did", "dname")
		i.Columns().Arg(9, "Antwerp Design")
		i.Conflict().Constraint("distributors_pkey").Action("NOTHING")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES (?, ?) ON CONFLICT ON CONSTRAINT distributors_pkey DO NOTHING;" {
			t.Fatal("wrong on_conflict_constraint")
		}
	})
	t.Run("on_conflict_constraint", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns().Add("did", "dname")
		i.Columns().Arg(10, "Conrad International")
		i.Conflict().Object("did").Action(ConflictActionNothing)
		i.Conflict().Where().AddExpression("is_active")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES (?, ?) ON CONFLICT (did) WHERE (is_active) DO NOTHING;" {
			t.Fatal("wrong on_conflict_constraint")
		}
	})
	t.Run("returning", func(t *testing.T) {
		i := NewInsert()
		i.Into("distributors")
		i.Columns().Add("did", "dname")
		i.Columns().Arg(1, "XYZ Widgets")
		i.Returning().Add("did")
		t.Log(i.String())
		if i.String() != "INSERT INTO distributors (did, dname) VALUES (?, ?) RETURNING did;" {
			t.Fatal("wrong on_conflict_constraint")
		}
	})
}
