//
// The RXScript author
// built-in functions

package ast

type ObjProps struct {
	FuncDef
}

func (p *ObjProps) String() string {
	return "'objProps'"
}

func NewObjProps() *ObjProps {
	return &ObjProps{}
}

func (p *ObjProps) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		arg := p.Args[0]
		switch v := arg.(type) {
		case *Class:
			arr := &ArrayDef{}
			for _, vd := range v.def.Body {
				if !vd.IsMethod && vd.IsPublic {
					arr.Value = append(arr.Value, &StringDef{Value: vd.Name})
				}
			}
			return arr
		}
	}
	return &NullDef{}
}
