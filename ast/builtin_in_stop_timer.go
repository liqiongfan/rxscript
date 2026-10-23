//
// The RXScript author
// built-in functions

package ast

type StopTimer struct {
	FuncDef
}

func (p *StopTimer) String() string {
	return "'stopTimer'"
}

func NewStopTimer() *StopTimer {
	return &StopTimer{}
}

func (p *StopTimer) Execute(rxe RxeHandler) any {
	if len(p.Args) < 1 {
		return ThrowSystemError("RXScript: need one arguments")
	}
	va := p.Args[0]

	switch v := va.(type) {
	case *TimeTimer:
		v.Value.Stop()
		return v
	default:
		return ThrowSystemErrorWithFormat("RXScript: stopTimer: invalid type '%s'", v)
	}
}
