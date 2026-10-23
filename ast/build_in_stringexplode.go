//
// The RXScript author
// built-in functions

package ast

import "strings"

type Explode struct {
	FuncDef
}

func (p *Explode) String() string {
	return "'explode'"
}

func NewExplode() *Explode {
	return &Explode{}
}

func (p *Explode) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch a1 := arg.(type) {
		case *StringDef:
			if len(p.Args) > 1 {
				switch a2 := p.Args[1].(type) {
				case *StringDef:
					arr := strings.Split(a1.Value, a2.Value)
					ret := &ArrayDef{Value: make([]Node, 0, len(arr))}
					for _, val := range arr {
						ret.Value = append(ret.Value, &StringDef{Value: val})
					}
					return ret
				default:
					return ThrowSystemErrorWithFormat("RXScript: explode second argument must be a string, not '%s'", a2)
				}
			}
			return &ArrayDef{Value: []Node{a1}}
		}
	}
	return &NullDef{}
}
