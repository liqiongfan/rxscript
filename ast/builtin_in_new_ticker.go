//
// The RXScript author
// built-in functions

package ast

import "time"

type NewTicker struct {
	FuncDef
}

func (p *NewTicker) String() string {
	return "'ticker'"
}

func NewNewTicker() *NewTicker {
	return &NewTicker{}
}

func (p *NewTicker) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: newTicker: no arguments")
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
		return ThrowSystemErrorWithFormat("RXScript: newTicker: invalid duration type '%s'", duration)
	}

	return &TimeTicker{Value: time.NewTicker(dura)}
}
