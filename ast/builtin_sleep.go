//
// The RXScript author
// built-in functions

package ast

import "time"

type Sleep struct {
	FuncDef
}

func (p *Sleep) String() string {
	return "'sleep'"
}

func NewSleep() *Sleep {
	return &Sleep{}
}

func (p *Sleep) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: sleep: no arguments")
	}
	duration := p.Args[0]

	switch d := duration.(type) {
	case *StringDef:
		dura, err := time.ParseDuration(d.Value)
		if err != nil {
			return ThrowSystemError(err)
		}
		time.Sleep(dura)
		return &BoolDef{true}
	case *IntDef:
		time.Sleep(time.Second * time.Duration(d.Value))
		return &BoolDef{true}
	case *DoubleDef:
		time.Sleep(time.Duration(d.Value * float64(time.Millisecond)))
		return &BoolDef{true}
	default:
		return ThrowSystemErrorWithFormat("RXScript: sleep: invalid duration type '%s'", duration)
	}
}
