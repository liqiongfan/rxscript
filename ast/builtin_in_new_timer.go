//
// The RXScript author
// built-in functions

package ast

import "time"

type NewTimer struct {
	FuncDef
}

func (p *NewTimer) String() string {
	return "'timer'"
}

func NewNewTimer() *NewTimer {
	return &NewTimer{}
}

func (p *NewTimer) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: newTimer: no arguments")
	}
	duration := p.Args[0]

	var err error
	var dura time.Duration

	switch d := duration.(type) {
	case *StringDef:
		dura, err = time.ParseDuration(d.Value)
		if err != nil {
			return ThrowSystemError(err)
		}
	case *IntDef:
		dura = time.Duration(d.Value) * time.Second
	case *DoubleDef:
		dura = time.Duration(d.Value * float64(time.Millisecond))
	default:
		return ThrowSystemErrorWithFormat("RXScript: newTimer: invalid duration type '%s'", duration)
	}

	return &TimeTimer{Value: time.NewTimer(dura)}
}
