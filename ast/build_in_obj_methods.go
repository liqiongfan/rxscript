//
// The RXScript author
// built-in functions

package ast

type ObjMethods struct {
	FuncDef
}

func (p *ObjMethods) String() string {
	return "'objMethods'"
}

func NewObjMethods() *ObjMethods {
	return &ObjMethods{}
}

func (p *ObjMethods) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		arg := p.Args[0]
		switch v := arg.(type) {
		case *Class:
			arr := &ArrayDef{}
			for _, vd := range v.def.Body {
				if vd.IsMethod && vd.IsPublic {
					arr.Value = append(arr.Value, &StringDef{Value: vd.Name})
				}
			}
			return arr
		}
	}
	return &NullDef{}
}
