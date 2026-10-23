//
// The RXScript author
// built-in functions

package ast

import "time"

type ResetTicker struct {
	FuncDef
}

func (p *ResetTicker) String() string {
	return "'resetTicker'"
}

func NewResetTicker() *ResetTicker {
	return &ResetTicker{}
}

func (p *ResetTicker) Execute(rxe RxeHandler) any {
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
		return ThrowSystemErrorWithFormat("RXScript: newTicker: invalid duration type '%s'", duration)
	}

	switch v := va.(type) {
	case *TimeTicker:
		v.Value.Reset(dura)
		return v
	default:
		return ThrowSystemErrorWithFormat("RXScript: newTicker: invalid type '%s'", dura)
	}
}
