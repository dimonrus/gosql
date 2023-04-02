package gosql

import (
	"testing"
	"time"
)

func TestSorting_Parse(t *testing.T) {
	t.Run("sorting_valid", func(t *testing.T) {
		s := Sorting{"   createdAt:desc", "name", "qty:ASC"}
		result := s.Allowed(map[string]string{"createdAt": "created_at", "name": "internal_name", "qty": "quantity"})
		if len(result) != 3 {
			t.Fatal("wrong sorting parser")
		}
		if result[0] != "created_at DESC" {
			t.Fatal("wrong created parser")
		}
		if result[1] != "internal_name" {
			t.Fatal("wrong internal_name parser")
		}
		if result[2] != "quantity" {
			t.Fatal("wrong quantity parser")
		}
	})
}

func TestPeriodFilter(t *testing.T) {
	t.Run("full_cond", func(t *testing.T) {
		now := time.Now()
		period := PeriodFilter{
			Start: &now,
			End:   &now,
		}
		if !period.IsEmpty() {
			cond := period.FieldCondition("created_at")
			if len(cond.GetArguments()) != 2 {
				t.Fatal("wrong arg count")
			}
			if cond.String() != "(created_at >= ? AND created_at <= ?)" {
				t.Fatal("wrong period condition")
			}
		}
	})
}

func TestSearchString_PrepareLikeValue(t *testing.T) {
	t.Run("double_column", func(t *testing.T) {
		s := SearchString("Foo Bar")
		cond := s.PrepareLikeValue("fullname")
		if len(cond.GetArguments()) != 2 {
			t.Fatal("wrong argument count")
		}
		if cond.String() != "(fullname like lower(?) AND fullname like lower(?))" {
			t.Fatal("wrong condition render")
		}
		if v, ok := cond.GetArguments()[0].(string); !ok {
			t.Fatal("must be a string")
		} else {
			if v != "%foo%" {
				t.Fatal("wrong 1st argument")
			}
		}
		if v, ok := cond.GetArguments()[1].(string); !ok {
			t.Fatal("must be a string")
		} else {
			if v != "%bar%" {
				t.Fatal("wrong 1st argument")
			}
		}
	})
}
