//
// The RXScript author
// TriEqualDef represent += in math.

package ast

var _ Node = (*TriEqualDef)(nil)

type TriEqualDef struct {
	A Node
	B Node
}

func (x *TriEqualDef) Any(handler RxeHandler) any {
	return CheckSystemError(
		x.Execute(handler),
		handler,
	).(Node).Any(handler)
}

func (x *TriEqualDef) String() string {
	return "'==='"
}

func (x *TriEqualDef) Type() CmdType {
	return TriEqualDefAstCmd
}

func checkNodeEqualWithType(a, b Node, e RxeHandler) any {
	ra := RawAstNode(a, e)
	rb := RawAstNode(b, e)

	switch av := ra.(type) {
	case *IntDef, *StringDef, *DoubleDef, *BoolDef, *NullDef: // base type
		switch bv := rb.(type) {
		case *IntDef, *StringDef, *BoolDef, *DoubleDef, *NullDef:
			return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
		}
	case *Class: // class object
		switch bv := rb.(type) {
		case *Class:
			return &BoolDef{Value: av == bv}
		}
	case *ArrayDef: // array
		switch bv := rb.(type) {
		case *ArrayDef:
			if len(av.Value) != len(bv.Value) {
				return &BoolDef{Value: false}
			}
			for i, _ := range av.Value {
				v := checkNodeEqualWithType(av.Value[i], bv.Value[i], e)
				switch vv := v.(type) {
				case *BoolDef:
					if !vv.Value {
						return vv
					}
				case *SystemErrorDef:
					return vv
				}
			}
			return &BoolDef{Value: true}
		}
	case *DictionaryDef: // dictionary
		switch bv := rb.(type) {
		case *DictionaryDef:
			if len(av.Values) != len(bv.Values) {
				return &BoolDef{Value: false}
			}
			for k, v := range av.Values {
				if r, ok := bv.Values[k]; !ok || r != v {
					v := checkNodeEqualWithType(r, v, e)
					switch vv := v.(type) {
					case *BoolDef:
						if !vv.Value {
							return vv
						}
					case *SystemErrorDef:
						return vv
					}
				}
			}
			return &BoolDef{Value: true}
		}
	}

	// other type return false
	return &BoolDef{Value: false}
}

func (x *TriEqualDef) Execute(e RxeHandler) any {
	return checkNodeEqualWithType(x.A, x.B, e)
}
