package gosql

import (
	"testing"
)

func TestUpdate_String(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("kind = ?", "Drama")
		u.
			Table("films").
			Set("kind = ?", "Dramatic").
			Condition(*cond)
		t.Log(u.String())
		if u.String() != "UPDATE films SET kind = ? WHERE (kind = ?);" || len(append(u.GetValues(), cond.GetArguments()...)) != 2 {
			t.Fatal("wrong simple")
		}
	})

	t.Run("complex_set", func(t *testing.T) {
		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("city = ?", "San Francisco")
		cond.AddExpression("date = ?", "2003-07-03")
		u.
			Table("weather").
			Set("temp_lo = temp_lo+1").
			Set("temp_hi = temp_lo+15").
			Set("prcp = DEFAULT").
			Condition(*cond)
		t.Log(u.String())
		if u.String() != "UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT WHERE (city = ? AND date = ?);" || len(append(u.GetValues(), cond.GetArguments()...)) != 2 {
			t.Fatal("wrong complex_set")
		}
	})

	t.Run("complex_set_with_returning", func(t *testing.T) {
		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.
			AddExpression("city = ?", "San Francisco").
			AddExpression("date = ?", "2003-07-03")
		u.
			Table("weather").
			Set("temp_lo = temp_lo+1").
			Set("temp_hi = temp_lo+15").
			Set("prcp = DEFAULT").
			AddReturning("temp_lo", "temp_hi", "prcp").
			Condition(*cond)
		t.Log(u.String())
		if u.String() != "UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT WHERE (city = ? AND date = ?) RETURNING temp_lo, temp_hi, prcp;" || len(append(u.GetValues(), cond.GetArguments()...)) != 2 {
			t.Fatal("wrong complex_set_with_returning")
		}
	})

	t.Run("complex_set_1", func(t *testing.T) {
		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("city = ?", "San Francisco")
		cond.AddExpression("date = ?", "2003-07-03")
		u.
			Table("weather").
			Set("(temp_lo, temp_hi, prcp) = (temp_lo+1, temp_lo+15, DEFAULT)").
			Condition(*cond)
		t.Log(u.String())
		if u.String() != "UPDATE weather SET (temp_lo, temp_hi, prcp) = (temp_lo+1, temp_lo+15, DEFAULT) WHERE (city = ? AND date = ?);" || len(append(u.GetValues(), cond.GetArguments()...)) != 2 {
			t.Fatal("wrong complex_set_1")
		}
	})

	t.Run("complex_set_2", func(t *testing.T) {
		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("accounts.name = ?", "Acme Corporation")
		cond.AddExpression("employees.id = accounts.sales_person")
		u.
			Table("employees").
			From("accounts").
			Set("sales_count = sales_count + 1").
			Condition(*cond)
		t.Log(u.String())
		if u.String() != "UPDATE employees SET sales_count = sales_count + 1 FROM accounts WHERE (accounts.name = ? AND employees.id = accounts.sales_person);" || len(append(u.GetValues(), cond.GetArguments()...)) != 1 {
			t.Fatal("wrong complex_set_2")
		}
	})

	t.Run("complex_where", func(t *testing.T) {
		sub := NewSelect()
		sub.From("accounts")
		sub.Columns("sales_person")
		sub.Where().AddExpression("name = ?", "Acme Corporation")
		sub.SubQuery = true

		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("id = "+sub.String(), sub.GetArguments()...)
		u.
			Table("employees").
			Set("sales_count = sales_count + 1").
			Condition(*cond)
		t.Log(u.String())
		if u.String() != "UPDATE employees SET sales_count = sales_count + 1 WHERE (id = (SELECT sales_person FROM accounts WHERE (name = ?)));" || len(append(u.GetValues(), cond.GetArguments()...)) != 1 {
			t.Fatal("wrong complex_where")
		}
	})

	t.Run("complex_where_1", func(t *testing.T) {
		sub := NewSelect()
		sub.From("salesmen")
		sub.Columns("first_name", "last_name")
		sub.Where().AddExpression("salesmen.id = accounts.sales_id")
		sub.SubQuery = true

		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("id = "+sub.String(), sub.GetArguments()...)
		u.
			Table("accounts").
			Set("(contact_first_name, contact_last_name) = " + sub.String())
		t.Log(u.String())
		if u.String() != "UPDATE accounts SET (contact_first_name, contact_last_name) = (SELECT first_name, last_name FROM salesmen WHERE (salesmen.id = accounts.sales_id));" || len(append(u.GetValues(), cond.GetArguments()...)) != 0 {
			t.Fatal("wrong complex_where_1")
		}
	})

	t.Run("from", func(t *testing.T) {
		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("salesmen.id = accounts.sales_id")
		u.
			Table("accounts").
			Set("contact_first_name = first_name").
			Set("contact_last_name = last_name").
			From("salesmen").
			Condition(*cond)
		t.Log(u.String())
		if u.String() != "UPDATE accounts SET contact_first_name = first_name, contact_last_name = last_name FROM salesmen WHERE (salesmen.id = accounts.sales_id);" || len(append(u.GetValues(), u.WithValues()...)) != 0 {
			t.Fatal("wrong from")
		}
	})

	t.Run("agg", func(t *testing.T) {
		sub := NewSelect()
		sub.From("data d")
		sub.Columns("sum(x)", "sum(y)", "avg(x)", "avg(y)")
		sub.Where().AddExpression("d.group_id = s.group_id")
		sub.SubQuery = true

		u := NewUpdate()
		cond := NewSqlCondition(ConditionOperatorAnd)
		cond.AddExpression("salesmen.id = accounts.sales_id")
		u.
			Table("summary s").
			Set("(sum_x, sum_y, avg_x, avg_y) = " + sub.String())
		t.Log(u.String())
		if u.String() != "UPDATE summary s SET (sum_x, sum_y, avg_x, avg_y) = (SELECT sum(x), sum(y), avg(x), avg(y) FROM data d WHERE (d.group_id = s.group_id));" || len(append(u.GetValues(), u.WithValues()...)) != 0 {
			t.Fatal("wrong agg")
		}
	})
}
