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
			isExist := make(map[string]struct{}, 32)
			// class owner props
			for _, vd := range v.def.Body {
				if !vd.IsMethod && vd.IsPublic {
					isExist[vd.Name] = struct{}{}
					arr.Value = append(arr.Value, &StringDef{Value: vd.Name})
				}
			}
			// dynamic setting props
			for name, _ := range v.variables {
				if _, ok := isExist[name]; !ok {
					arr.Value = append(arr.Value, &StringDef{Value: name})
				}
			}
			return arr
		}
	}
	return &NullDef{}
}
