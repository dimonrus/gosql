package gosql

import "strings"

// On conflict query part
type conflict struct {
	// object of conflict
	object string
	// action on conflict
	action string
	// set of changes
	set expression
	// condition
	where Condition
	// constraint
	constraint string
}

// String conflict expression
func (c *conflict) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("ON CONFLICT")
	if c.object != "" {
		b.WriteString(" (" + c.object + ")")
	}
	if c.constraint != "" {
		b.WriteString(" ON CONSTRAINT " + c.constraint)
	}
	if c.set.Len() > 0 {
		if c.action != "" {
			b.WriteString(" DO " + c.action)
		}
		b.WriteString(" SET " + c.set.String(", "))
		if !c.where.IsEmpty() {
			b.WriteString(" WHERE " + c.where.String())
		}
	} else {
		if !c.where.IsEmpty() {
			b.WriteString(" WHERE " + c.where.String())
		}
		if c.action != "" {
			b.WriteString(" DO " + c.action)
		}
	}
	return b.String()
}

// GetArguments get all arguments
func (c *conflict) GetArguments() []any {
	return append(c.set.Params(), c.where.GetArguments()...)
}

// IsEmpty Is conflict empty
func (c *conflict) IsEmpty() bool {
	return c.object == "" && c.action == "" && c.set.Len() == 0 && c.where.IsEmpty() && c.constraint == ""
}

// Object of conflict
func (c *conflict) Object(object string) *conflict {
	c.object = object
	return c
}

// ResetObject reset
func (c *conflict) ResetObject() *conflict {
	c.object = ""
	return c
}

// Action of conflict
func (c *conflict) Action(action string) *conflict {
	c.action = action
	return c
}

// ResetAction reset action
func (c *conflict) ResetAction() *conflict {
	c.action = ""
	return c
}

// Set of expressions on conflict
func (c *conflict) Set(expr ...string) *conflict {
	c.set.AddExpressions(expr...)
	return c
}

// Add expression on conflict
func (c *conflict) Add(expr string, args ...any) *conflict {
	c.set.Add(expr, args...)
	return c
}

// AddArguments add expression arguments
func (c *conflict) AddArguments(args ...any) *conflict {
	c.set.AddParams(args...)
	return c
}

// ResetSet of expressions on conflict
func (c *conflict) ResetSet() *conflict {
	c.set.Reset()
	return c
}

// Where get condition
func (c *conflict) Where() *Condition {
	return &c.where
}

// Constraint set constraint
func (c *conflict) Constraint(constraint string) *conflict {
	c.constraint = constraint
	return c
}

// ResetConstraint reset constraint
func (c *conflict) ResetConstraint() *conflict {
	c.constraint = ""
	return c
}

// NewConflict conflict constructor
func NewConflict() *conflict {
	return &conflict{
		where: Condition{operator: ConditionOperatorAnd},
	}
}
