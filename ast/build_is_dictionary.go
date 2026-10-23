//
// The RXScript author
// built-in functions

package ast

type IsDictionary struct {
	FuncDef
}

func (p *IsDictionary) String() string {
	return "'isDictionary'"
}

func NewIsDictionary() *IsDictionary {
	return &IsDictionary{}
}

func (p *IsDictionary) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *DictionaryDef:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
