//
// The RXScript author
// built-in functions

package ast

import (
	"strings"
)

type ToLower struct {
	FuncDef
}

func (p *ToLower) String() string {
	return "'toLower'"
}

func NewToLower() *ToLower {
	return &ToLower{}
}

func (p *ToLower) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: toLower: need at least 1 argument.")
	}

	v := p.Args[0]
	switch vv := v.(type) {
	case *StringDef:
		return &StringDef{Value: strings.ToLower(vv.Value)}
	}

	return ThrowSystemErrorWithFormat("RXScript: toLower: invalid argument, expect string not '%s'", v)
}
