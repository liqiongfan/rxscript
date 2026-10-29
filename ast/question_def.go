//
// The RXScript author
// QuestionDef represent ? :

package ast

var _ Node = (*QuestionDef)(nil)

// QuestionDef Ternary Operator like  xx ? a : b
// if xx is true, 'a' is returned
// otherwise 'b' is returned
type QuestionDef struct {
	// condition
	Condition Node

	// A is returned when condition is true otherwise return B
	A, B Node
}

func (q *QuestionDef) Any(handler RxeHandler) any {
	return CheckSystemError(q.Execute(handler), handler).(Node).Any(handler)
}

func (q *QuestionDef) String() string {
	return "object(?:)"
}

// Execute the ternary operator
func (q *QuestionDef) Execute(e RxeHandler) any {
	if convertToBool(q.Condition, e) {
		return q.A
	}
	return q.B
}

// Type is ast.QuestionAstCmd
func (q *QuestionDef) Type() CmdType {
	return QuestionAstCmd
}
