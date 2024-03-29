package gosql

import (
	"fmt"
	"testing"
)

func TestQB_String(t *testing.T) {
	mr := NewSelect()
	mr.From("mv_right")
	mr.Columns().Add("id", "contract_id", "object_id")
	mr.Where().AddExpression("object_id = ?", "84f3ba22-5b7f-4967-80e2-451a123deff6")
	mr.AddOrder("terrirtory_name")
	mr.SetPagination(10, 0)

	c := NewSelect()
	c.From("mv_contracts")
	c.Columns().Add("id", "contract_name")
	c.Where().AddExpression("contract_sum > ?", 23.45)
	c.SetPagination(5, 0)

	qb := NewSelect()
	qb.With().
		Add("mv_right_items", mr).
		Add("mv_contracts_items", c)
	qb.From("mv_object mo")
	qb.Columns().Add("mo.id", "mo.title", "mo.rightholder_ids", "mr.id", "mr.contract_id")
	qb.Relate("JOIN mv_right_items AS mr ON mr.object_id = mo.id")
	qb.Relate("LEFT JOIN mv_contracts_items AS ci ON ci.id = mr.contract_id")
	qb.Where().AddExpression("mr.object_id IS NOT NULL")

	if qb.String() != "WITH mv_right_items AS (SELECT id, contract_id, object_id FROM mv_right WHERE (object_id = ?) ORDER BY terrirtory_name LIMIT 10 OFFSET 0),mv_contracts_items AS (SELECT id, contract_name FROM mv_contracts WHERE (contract_sum > ?) LIMIT 5 OFFSET 0) SELECT mo.id, mo.title, mo.rightholder_ids, mr.id, mr.contract_id FROM mv_object mo JOIN mv_right_items AS mr ON mr.object_id = mo.id LEFT JOIN mv_contracts_items AS ci ON ci.id = mr.contract_id WHERE (mr.object_id IS NOT NULL)" {
		t.Fatal("wrong query")
	}
	fmt.Println(qb.String())
}

func TestWith_Recursive(t *testing.T) {
	// WITH RECURSIVE employee_recursive(distance, employee_name, manager_name) AS (
	//    SELECT 1, employee_name, manager_name
	//    FROM employee
	//    WHERE manager_name = 'Mary'
	//  UNION ALL
	//    SELECT er.distance + 1, e.employee_name, e.manager_name
	//    FROM employee_recursive er, employee e
	//    WHERE er.employee_name = e.manager_name
	//  )
	//SELECT distance, employee_name FROM employee_recursive;

	employee := NewSelect().From("employee")
	employee.Columns().Add("1", "employee_name", "manager_name")
	employee.Where().AddExpression("manager_name = ?", "Mary")

	reqEmployee := NewSelect().From("employee_recursive er", "employee e")
	reqEmployee.Columns().Add("er.distance + 1", "e.employee_name", "e.manager_name")
	reqEmployee.Where().AddExpression("er.employee_name = e.manager_name")
	employee.Union(reqEmployee)

	s := NewSelect().From("employee_recursive")
	s.Columns().Add("distance", "employee_name")
	s.With().Recursive().Add("employee_recursive(distance, employee_name, manager_name)", employee)

	t.Log(s.String())
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dimonrus/gosql
// BenchmarkSelect
// BenchmarkSelect-12    	 4600735	       253.6 ns/op	     296 B/op	      10 allocs/op
func BenchmarkSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		employee := NewSelect().From("employee")
		employee.Columns().Add("1", "employee_name", "manager_name")
		employee.Where().AddExpression("manager_name = ?", "Mary")
		employee.Where().AddExpression("some_other = ?", 1000)
		employee.Where().AddExpression("some_other = ?", "1000")
	}
	b.ReportAllocs()
}

func TestCondition_IsEmpty(t *testing.T) {
	var c *Condition
	if !c.IsEmpty() {
		t.Fatal("wrong")
	}
	condFuncEmpty := func(c *Condition) {
		if !c.IsEmpty() {
			t.Fatal("wrong")
		}
	}
	condFuncNotEmpty := func(c *Condition) {
		if c.IsEmpty() {
			t.Fatal("wrong")
		}
	}
	condFuncEmpty(c)
	c = NewSqlCondition(ConditionOperatorAnd)
	condFuncEmpty(c)
	c.AddExpression("some = ?", 1)
	condFuncNotEmpty(c)
}

func TestQB_Union(t *testing.T) {
	q := NewSelect()
	q.Columns().Add("*")
	q.From("some_table")

	u1 := NewSelect()
	u1.Columns().Add("*")
	u1.From("some_table_union_1")

	u2 := NewSelect()
	u2.Columns().Add("*")
	u2.From("some_table_union_2")

	q.Union(u1)
	q.Union(u2)

	fmt.Println(q.String())
}

func TestQB_Intersect(t *testing.T) {
	m := NewSelect()
	m.Columns().Add("*")
	m.From("main_table")

	q := NewSelect()
	q.Columns().Add("*")
	q.From("some_table")

	u1 := NewSelect()
	u1.Columns().Add("*")
	u1.From("some_table_union_1")

	u2 := NewSelect()
	u2.Columns().Add("*")
	u2.From("some_table_union_2")

	u1.Intersect(u2)
	u1.SubQuery = true
	q.Union(u1)

	m.With().Add("some", q)
	fmt.Println(m.String())
}

func TestQB_Except(t *testing.T) {
	q := NewSelect()
	q.Columns().Add("*")
	q.From("some_table")

	u1 := NewSelect()
	u1.Columns().Add("*")
	u1.From("some_table_union_1")

	u2 := NewSelect()
	u2.Columns().Add("*")
	u2.From("some_table_except_2")

	q.Union(u1)
	q.Except(u2)
	fmt.Println(q.String())
}
