//
// The RXScript author
// built-in functions

package ast

import "time"

type ResetTimer struct {
	FuncDef
}

func (p *ResetTimer) String() string {
	return "'resetTimer'"
}

func NewResetTimer() *ResetTimer {
	return &ResetTimer{}
}

func (p *ResetTimer) Execute(rxe RxeHandler) any {
	if len(p.Args) < 2 {
		return ThrowSystemError("RXScript: need two arguments")
	}
	va := p.Args[0]
	duration := p.Args[1]

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

	switch v := va.(type) {
	case *TimeTimer:
		v.Value.Reset(dura)
		return v
	default:
		return ThrowSystemErrorWithFormat("RXScript: newTimer: invalid type '%s'", dura)
	}
}
