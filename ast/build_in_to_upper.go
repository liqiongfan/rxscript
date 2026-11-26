//
// The RXScript author
// built-in functions

package ast

import (
	"strings"
)

type ToUpper struct {
	FuncDef
}

func (p *ToUpper) String() string {
	return "'toUpper'"
}

func NewToUpper() *ToUpper {
	return &ToUpper{}
}

func (p *ToUpper) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: toUpper: need at least 1 argument.")
	}

	v := p.Args[0]
	switch vv := v.(type) {
	case *StringDef:
		return &StringDef{Value: strings.ToUpper(vv.Value)}
	}

	return ThrowSystemErrorWithFormat("RXScript: toUpper: invalid argument, expect string not '%s'", v)
}
