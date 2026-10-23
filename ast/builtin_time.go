//
// The RXScript author
// built-in functions

package ast

import (
	"reflect"
	"time"
)

type Time struct {
	FuncDef
}

func (p *Time) String() string {
	return "'time'"
}

func NewTime() *Time {
	return &Time{}
}

func (p *Time) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		switch p.Args[0].(type) {
		case *StringDef:
		default:
			return ThrowSystemErrorWithFormat("RXScript: arguments must be string, not: '%T'", p.Args[0])
		}
		v, err := time.Parse(time.DateTime, rawValue(p.Args[0], rxe).(string))
		if err != nil {
			return ThrowSystemError(err)
		}
		return &SystemResource{obj: reflect.ValueOf(v)}
	}
	return &SystemResource{obj: reflect.ValueOf(time.Now())}
}
