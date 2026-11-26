//
// The RXScript author
// ForNormalDef represent for {}.

package ast

var _ Node = (*ForNormalDef)(nil)

type ForNormalDef struct {
	Name      string
	Value     Node
	Condition Node
	Step      Node
	Stmts     *FunctionBodyDef
}

func (x *ForNormalDef) Any(handler RxeHandler) any {
	return nil
}

func (x *ForNormalDef) String() string {
	return "object(for)"
}

// Type ForNormalDefAstCmd
func (x *ForNormalDef) Type() CmdType {
	return ForNormalDefAstCmd
}

// Execute return the raw value of the array list
// user should use `CheckSystemError` to check value or system error occurred.
func (x *ForNormalDef) Execute(e RxeHandler) any {
	e.SetArgs(x.Name, x.Value)
	defer e.DeleteArgs(x.Name)
breakFor:
	for {
		cond := x.Condition.Execute(e)
		if !convertToBool(cond, e) {
			break
		}
		switch n := x.Stmts.Execute(e).(type) {
		case *ContinueDef:
			continue
		case *ReturnDef:
			return n
		case *BreakDef:
			break breakFor
		}

		x.Step.Execute(e)
	}
	return &NullDef{}
}
