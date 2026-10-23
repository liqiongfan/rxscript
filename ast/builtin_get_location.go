//
// The RXScript author
// built-in functions

package ast

import (
	"reflect"
	"time"
)

type Location struct {
	FuncDef
}

func (p *Location) String() string {
	return "'getLocation'"
}

func NewLocation() *Location {
	return &Location{}
}

func (p *Location) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		switch p.Args[0].(type) {
		case *StringDef:
		default:
			return ThrowSystemErrorWithFormat("RXScript: arguments must be string, not: '%T'", p.Args[0])
		}
		location, err := time.LoadLocation(rawValue(p.Args[0], rxe).(string))
		if err != nil {
			return ThrowSystemError(err)
		}
		return &SystemResource{obj: reflect.ValueOf(location)}
	}
	return ThrowSystemError("RXScript: getLocation argument empty.")
}
