//
// The RXScript author
// EqualDoubleDef represent += in math.

package ast

import "math/big"

var _ Node = (*EqualDoubleDef)(nil)

type EqualDoubleDef struct {
	A Node
	B Node
}

func (x *EqualDoubleDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *EqualDoubleDef) String() string {
	return "'=='"
}

func (x *EqualDoubleDef) Type() CmdType {
	return EqualDoubleDefAstCmd
}

func checkValueEqual(a, b Node, e RxeHandler) any {
	ra := RawAstNode(a, e)
	rb := RawAstNode(b, e)

	switch av := ra.(type) {
	case *BoolDef, *NullDef: // base type
		switch bv := rb.(type) {
		case *IntDef, *StringDef, *BoolDef, *DoubleDef, *NullDef:
			return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
		}
	case *DoubleDef:
		switch bv := rb.(type) {
		case *IntDef:
			aa := new(big.Rat).SetFloat64(av.Value)
			if new(big.Rat).SetInt64(bv.Value).Cmp(aa) == 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		case *StringDef:
			bb, ok := new(big.Rat).SetString(bv.Value)
			if !ok {
				return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
			}
			if new(big.Rat).SetFloat64(av.Value).Cmp(bb) == 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		case *BoolDef, *DoubleDef, *NullDef:
			return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
		}
	case *StringDef:
		switch bv := rb.(type) {
		case *IntDef:
			aa, ok := new(big.Rat).SetString(av.Value)
			if !ok {
				return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
			}
			if new(big.Rat).SetInt64(bv.Value).Cmp(aa) == 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		case *DoubleDef:
			aa, ok := new(big.Rat).SetString(av.Value)
			if !ok {
				return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
			}
			if new(big.Rat).SetFloat64(bv.Value).Cmp(aa) == 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		case *StringDef:
			aa, ok := new(big.Rat).SetString(av.Value)
			if !ok {
				return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
			}
			bb, ok := new(big.Rat).SetString(bv.Value)
			if !ok {
				return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
			}
			if aa.Cmp(bb) == 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		case *BoolDef, *NullDef:
			return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
		}
	case *IntDef:
		switch bv := rb.(type) {
		case *StringDef:
			bb, ok := new(big.Rat).SetString(bv.Value)
			if !ok {
				return &BoolDef{Value: rawValue(av, e) == rawValue(bv, e)}
			}
			if new(big.Rat).SetInt64(av.Value).Cmp(bb) == 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		case *IntDef, *BoolDef, *DoubleDef, *NullDef:
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
				v := checkValueEqual(av.Value[i], bv.Value[i], e)
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
					v := checkValueEqual(r, v, e)
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

func (x *EqualDoubleDef) Execute(e RxeHandler) any {
	return checkValueEqual(x.A, x.B, e)
}
