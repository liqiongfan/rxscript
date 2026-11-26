//
// The RXScript author
// built-in functions

package ast

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Md5 struct {
	FuncDef
}

func (p *Md5) String() string {
	return "'md5'"
}

func NewMd5() *Md5 {
	return &Md5{}
}

func (p *Md5) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch a := arg.(type) {
		case *IntDef:
			m := md5.New()
			m.Write([]byte(fmt.Sprintf("%d", a.Value)))
			return &StringDef{Value: hex.EncodeToString(m.Sum(nil))}
		case *StringDef:
			m := md5.New()
			m.Write([]byte(a.Value))
			return &StringDef{Value: hex.EncodeToString(m.Sum(nil))}
		default:
			m := md5.New()
			m.Write([]byte(fmt.Sprintf("%v", a)))
			return &StringDef{Value: hex.EncodeToString(m.Sum(nil))}
		}
	}
	return &NullDef{}
}
