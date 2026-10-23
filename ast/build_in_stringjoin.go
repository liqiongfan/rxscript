//
// The RXScript author
// built-in functions

package ast

import (
	"fmt"
	"strings"
)

type Join struct {
	FuncDef
}

func (p *Join) String() string {
	return "'join'"
}

func NewJoin() *Join {
	return &Join{}
}

func (p *Join) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch a1 := arg.(type) {
		case *ArrayDef:
			if len(p.Args) > 1 {
				switch a2 := p.Args[1].(type) {
				case *StringDef:
					nodes := make([]string, 0, len(a1.Value))
					for _, node := range a1.Value {
						nodes = append(nodes, fmt.Sprintf("%v", rawValue(node, rxe)))
					}
					return &StringDef{Value: strings.Join(nodes, a2.Value)}
				default:
					return ThrowSystemErrorWithFormat("RXScript: join second argument must be a string, not '%s'", a2)
				}
			}
			return &StringDef{}
		default:
			return ThrowSystemErrorWithFormat("RXScript: join argument must be a array, not '%s'", a1)
		}
	}
	return &StringDef{}
}
