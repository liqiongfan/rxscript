//
// The RXScript author
// ClassDef represent a dictionary in RXScript, very similar to Go's map

package ast

var _ Node = (*DictionaryIndexDef)(nil)

type DictionaryIndexDef struct {
	Values Node
	Key    Node
}

func (x *DictionaryIndexDef) Any(e RxeHandler) any {
	values := x.Values.Any(e)
	switch value := values.(type) {
	case *DictionaryDef:
		key := x.Key.Any(e)
		switch k := key.(type) {
		case *StringDef:
			return value.Values[k.Value]
		}
	}
	return nil
}

func (x *DictionaryIndexDef) Type() CmdType {
	return DictionaryIndexDefAstCmd
}

func (x *DictionaryIndexDef) Execute(e RxeHandler) any {
	values := x.Values.Execute(e)
	switch value := values.(type) {
	case *DictionaryDef:
		key := x.Key.Execute(e)
		switch k := key.(type) {
		case *StringDef:
			return value.Values[k.Value]
		}
	}
	return ThrowSystemErrorWithFormat("RXScript: dict[key] need dictionary, not dict[%s]", x.Key.String())
}

func (x *DictionaryIndexDef) String() string {
	return "dict[" + x.Key.String() + "]"
}
