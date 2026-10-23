//
// The RXScript author
// ClassDef represent a dictionary in RXScript, very similar to Go's map

package ast

import (
	"fmt"
)

var _ Node = (*DictionaryDef)(nil)

type DictionaryDef struct {
	Values map[string]Node
}

func (x *DictionaryDef) Any(handler RxeHandler) any {
	ret := make(map[string]any, len(x.Values))
	for k, v := range x.Values {
		ret[k] = v.Any(handler)
	}
	return ret
}

func (x *DictionaryDef) Type() CmdType {
	return MinusDefAstCmd
}

func (x *DictionaryDef) Execute(e RxeHandler) any {
	for k, v := range x.Values {
		x.Values[k] = v.Execute(e).(Node)
	}
	return x
}

func (x *DictionaryDef) String() string {
	r := "dict{"
	for k, v := range x.Values {
		r += fmt.Sprintf(`"%s": %s,`, k, v)
	}
	if len(x.Values) > 0 {
		return r[0:len(r)-1] + "}"
	}
	return r + "}"
}
