package gosql

import "testing"

func TestNewDetailedExpression(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		de := NewDetailedExpression()
		de.Expression().Add("id", "name", "some_int")
		de.SetDetail("keyword")
		t.Log(de.String())
		if de.String() != "(id, name, some_int) keyword" {
			t.Fatal("normal wrong detailed expression")
		}
	})
	t.Run("detail_first", func(t *testing.T) {
		de := NewDetailedExpression()
		de.Expression().Add("id", "name", "some_int")
		de.SetDetail("keyword")
		de.SetRightAlign(true)
		t.Log(de.String())
		if de.String() != "keyword (id, name, some_int)" {
			t.Fatal("detail_first wrong detailed expression")
		}
	})
	t.Run("empty", func(t *testing.T) {
		de := NewDetailedExpression()
		t.Log(de.String())
		if de.String() != "" {
			t.Fatal("empty wrong detailed expression")
		}
	})
	t.Run("only_expression", func(t *testing.T) {
		de := NewDetailedExpression()
		de.Expression().Add("id")
		t.Log(de.String())
		if de.String() != "(id)" {
			t.Fatal("only_expression wrong detailed expression")
		}
	})
	t.Run("only_detail", func(t *testing.T) {
		de := NewDetailedExpression()
		de.SetDetail("xor")
		t.Log(de.String())
		if de.String() != "" {
			t.Fatal("only_expression wrong detailed expression")
		}
	})
	t.Run("detail_methods", func(t *testing.T) {
		de := NewDetailedExpression()
		de.SetDetail("xor")
		if de.GetDetail() != "xor" {
			t.Fatal("wrong get detail")
		}
		de.ResetDetail()
		t.Log(de.String())
		if de.String() != "" {
			t.Fatal("detail_methods wrong detailed expression")
		}
	})
}
