//
// The RXScript author
// built-in functions

package ast

type Range struct {
	FuncDef
}

func (p *Range) String() string {
	return "'series'"
}

func NewRange() *Range {
	return &Range{}
}

func (p *Range) Execute(rxe RxeHandler) any {
	// range(0, 10, 2)
	if len(p.Args) < 2 {
		return ThrowSystemErrorWithFormat("RXScript: series: args(start, end, step)")
	}

	var (
		start, end, step int64
	)

	s := p.Args[0]

	switch v := s.(type) {
	case *IntDef:
		start = v.Value
	default:
		return ThrowSystemErrorWithFormat("RXScript: series: argument one need to be an integer, not '%s'", v)
	}

	e := p.Args[1]
	switch v := e.(type) {
	case *IntDef:
		end = v.Value
	default:
		return ThrowSystemErrorWithFormat("RXScript: series: argument one need to be an integer, not '%s'", v)
	}

	step = 1
	if len(p.Args) > 2 {
		p3 := p.Args[2]
		switch a3 := p3.(type) {
		case *IntDef:
			step = a3.Value
		default:
			return ThrowSystemErrorWithFormat("RXScript: series: step needs to be an integer")
		}
	}

	result := &ArrayDef{make([]Node, 0, 32)}
	for i := start; ; /*i <= end*/ i += step {
		if step > 0 {
			if i >= end {
				break
			}
		} else {
			if i <= end {
				break
			}
		}
		result.Value = append(result.Value, &IntDef{Value: i})
	}

	return result
}
