package gosql

import (
	"fmt"
	"testing"
	"time"
)

func TestConditions_AddExpression(t *testing.T) {
	c := NewSqlCondition(ConditionOperatorAnd)
	c.AddExpression("created_at BETWEEN ? AND ?", time.Now(), time.Now().Add(time.Minute))
	c.AddExpression("is_active IS TRUE")
	c.AddExpression("middle = ?", "center")

	fmt.Println(c.String())
	fmt.Println(c.GetArguments())

	// (((condition = ? OR lazy IS ?) AND (easy IS ? AND size = ?)) OR (price = ?) OR (created_at BETWEEN ? AND ? AND is_active IS TRUE AND middle = ?))

	c1 := NewSqlCondition(ConditionOperatorAnd)
	c1.AddExpression("easy IS ?", true)
	c1.AddExpression("size = ?", 20)

	c11 := NewSqlCondition(ConditionOperatorOr)
	c11.AddExpression("condition = ?", "empty")
	c11.AddExpression("lazy IS ?", true)

	c2 := NewSqlCondition(ConditionOperatorOr)
	c2.AddExpression("price = ?", 200.23)

	//c.Merge(ConditionOperatorOr, c1.Merge(ConditionOperatorAnd, c11), c2)

	c.Merge(ConditionOperatorOr, c1)
	c.Merge(ConditionOperatorOr, c2)
	c.Merge(ConditionOperatorOr, nil)

	fmt.Println(c.String())
	fmt.Println(c.GetArguments())
}

func TestCondition_Replace(t *testing.T) {
	c := NewSqlCondition(ConditionOperatorAnd)
	c.AddExpression("created_at BETWEEN ? AND ?", time.Now(), time.Now().Add(time.Minute))
	c.AddExpression("is_active IS TRUE")
	c.AddExpression("middle = ?", "center")
	c2 := NewSqlCondition(ConditionOperatorAnd)
	c2.AddExpression("is_active IS NOT TRUE")
	c.Replace(c2)
	t.Log(c.String())
	if c.String() != "(is_active IS NOT TRUE)" {
		t.Fatal("wrong replace")
	}
}

func BenchmarkCondition_AddExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewSqlCondition(ConditionOperatorAnd)
		c.AddExpression("created_at BETWEEN ? AND ?", time.Now(), time.Now().Add(time.Minute))
		c.AddExpression("is_active IS TRUE")
		c.AddExpression("middle = ?", "center")
		c.AddExpression("left = ?", "left")
		c.AddExpression("top = ?", "top")

		c1 := NewSqlCondition(ConditionOperatorAnd)
		c1.AddExpression("created_at BETWEEN ? AND ?", time.Now(), time.Now().Add(time.Minute))
		c1.AddExpression("is_active IS TRUE")
		c1.AddExpression("middle = ?", "center")

		c.Merge(ConditionOperatorOr, c1)
	}
	b.ReportAllocs()
}
