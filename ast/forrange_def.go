//
// The RXScript author
// ForRangeDef represent for x, x = range x {}.

package ast

var _ Node = (*ForRangeDef)(nil)

type ForRangeDef struct {
	NameA string
	NameB string
	Value Node
	Stmts *FunctionBodyDef
}

func (x *ForRangeDef) Any(handler RxeHandler) any {
	return nil
}

func (x *ForRangeDef) String() string {
	return "'for range'"
}

// Type ForRangeDefAstCmd
func (x *ForRangeDef) Type() CmdType {
	return ForRangeDefAstCmd
}

// Execute return the raw value of the array list
// user should use `CheckSystemError` to check value or system error occurred.
func (x *ForRangeDef) Execute(e RxeHandler) any {
	switch xv := x.Value.(type) {
	case *ArrayDef, *DictionaryDef, *Class, *StringDef:
	default:
		x.Value = xv.Execute(e).(Node)
	}
	defer e.DeleteArgs(x.NameA, x.NameB)

	switch v := x.Value.(type) {
	case *DictionaryDef:
	breakDictFor:
		for nameA, nameB := range v.Values {
			e.SetArgs(x.NameA, &StringDef{Value: nameA})
			e.SetArgs(x.NameB, nameB)
			switch n := x.Stmts.Execute(e).(type) {
			case *ReturnDef:
				return n
			case *BreakDef:
				break breakDictFor
			}
		}
	case *ArrayDef:
	breakArrayDefFor:
		for nameA, nameB := range v.Value {
			e.SetArgs(x.NameA, &IntDef{Value: int64(nameA)})
			e.SetArgs(x.NameB, nameB)
			switch n := x.Stmts.Execute(e).(type) {
			case *ReturnDef:
				return n
			case *BreakDef:
				break breakArrayDefFor
			}
		}
	case *StringDef:
	breakStringDefFor:
		for index, byteData := range v.Value {
			e.SetArgs(x.NameA, &IntDef{Value: int64(index)})
			e.SetArgs(x.NameB, &StringDef{Value: string(byteData)})
			switch n := x.Stmts.Execute(e).(type) {
			case *ReturnDef:
				return n
			case *BreakDef:
				break breakStringDefFor
			}
		}
	case *Class:
	breakClassFor:
		for nameA, nameB := range v.variables {
			e.SetArgs(x.NameA, &StringDef{Value: nameA})
			e.SetArgs(x.NameB, nameB)
			switch n := x.Stmts.Execute(e).(type) {
			case *ReturnDef:
				return n
			case *BreakDef:
				break breakClassFor
			}
		}
	default:
		return ThrowSystemErrorWithFormat("RXScript: for range only support array, string and object, not: '%s'", x.Value)
	}

	return &NullDef{}
}
