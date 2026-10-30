//
// The RXScript author
//

package ast

import (
	"fmt"
	"log"
	"strconv"
)

type CmdType int

const (
	AndDefAstCmd CmdType = iota + 1
	ArrayDefAstCmd
	ArrayIndexDefAstCmd
	BoolDefAstCmd
	ClassAstCmd
	ClassBodyDefAstCmd
	ClassDefAstCmd
	DivideDefAstCmd
	ModDefAstCmd
	DoubleDefAstCmd
	ElseDefAstCmd
	ElseIfDefAstCmd
	FuncCallDefAstCmd
	FunctionBodyDefAstCmd
	FunctionDefAstCmd
	GetAttributeDefAstCmd
	GtDefAstCmd
	GteDefAstCmd
	IfStmtDefAstCmd
	IntDefAstCmd
	LtDefAstCmd
	LteDefAstCmd
	MinusDefAstCmd
	MultiplyDefAstCmd
	NeDefAstCmd
	NotDefAstCmd
	NegValueDefAstCmd
	NullDefAstCmd
	OrDefAstCmd
	PlusDefAstCmd
	ReturnDefAstCmd
	StringDefAstCmd
	VariableAssignDefAstCmd
	VariableDefAstCmd
	SystemErrorDefAstCmd
	ImportAstCmd
	QuestionAstCmd
	PlusEqualDefAstCmd
	PlusDoubleDefAstCmd
	MinusEqualDefAstCmd
	MinusDoubleDefAstCmd
	ForRangeDefAstCmd
	EqualDoubleDefAstCmd
	TriEqualDefAstCmd
	SystemResourceAstCmd
	GetNodeAttributeDefAstCmd
	GoDefAstCmd
	ChannelDefAstCmd
	ReadFromChannelDefAstCmd
	WriteToChannelDefAstCmd
	WaitDefAstCmd
	AtomicDefAstCmd
	BreakDefAstCmd
	ForDefAstCmd
	SplitArrayDefAstCmd
	LeftShiftDefAstCmd
	RightShiftDefAstCmd
	BitAndDefAstCmd
	BitOrDefAstCmd
	BitExpDefAstCmd
	GoSystemResourceDefAstCmd
)

type NodeStore map[string]Node

func NewNodeStore() NodeStore {
	return make(NodeStore, 32)
}

func (ans NodeStore) Dump() NodeStore {
	newNodeStore := NewNodeStore()
	for name, node := range ans {
		newNodeStore[name] = node
	}
	return newNodeStore
}

func (ans NodeStore) Get(name string) Node {
	return ans[name]
}

func (ans NodeStore) Set(name string, value Node) {
	ans[name] = value
}

func (ans NodeStore) Delete(names ...string) {
	for _, name := range names {
		delete(ans, name)
	}
}

type RxeHandler interface {
	GetVariables() NodeStore
	GetCallArgs() NodeStore
	GetFromArgs(string) (any, bool)
	SetArgs(string, any)
	GetFromVariables(string) (any, bool)
	GetFromFuncDefs(string) (any, bool)
	GetFromClassDefs(string) (any, bool)
	GetFromPackageDefs(string) (any, bool)
	SetCallingObject(*Class)
	PushArgStack(NodeStore)
	PopArgStack()
	GetCallingObject() *Class
	DumpRxe() RxeHandler
	GoroutineChan() chan struct{}
	DeleteArgs(names ...string)
	GetCurrentClassDef() *ClassDef
	SetCurrentClassDef(*ClassDef)
	PushClassDef(def *ClassDef)
	PopClassDef() *ClassDef
}

type Node interface {
	Execute(RxeHandler) any
	Type() CmdType
	String() string
	Any(handler RxeHandler) any
}

var _ Node = (*SystemErrorDef)(nil)

func CheckSystemError(v any, e RxeHandler) any {
	switch vv := v.(type) {
	case *SystemErrorDef:
		return vv.Execute(e)
	case SystemErrorDef:
		return vv.Execute(e)
	default:
		return v
	}
}

// ThrowSystemError Throw system error
func ThrowSystemError(v any) *SystemErrorDef {
	return &SystemErrorDef{Value: v}
}

// ThrowSystemErrorWithFormat Throw system error
func ThrowSystemErrorWithFormat(format string, args ...any) *SystemErrorDef {
	return &SystemErrorDef{Value: fmt.Sprintf(format, args...)}
}

type SystemErrorDef struct {
	Value any
}

func (s *SystemErrorDef) Any(handler RxeHandler) any {
	return s.Value
}

func (s *SystemErrorDef) Execute(e RxeHandler) any {
	log.Fatal(s.Value)
	return nil
}

func (s *SystemErrorDef) Type() CmdType {
	return SystemErrorDefAstCmd
}

func (s *SystemErrorDef) String() string {
	return fmt.Sprintf(`%s`, s.Value)
}

type FuncDef struct {
	Args []any
}

// WithArgs parse the input args for built-in functions
// if argument is b... then b will explode into several arguments
func (b *FuncDef) WithArgs(node any, e RxeHandler, args ...any) {
	for _, arg := range args {
		switch arg := arg.(type) {
		case *SplitArrayDef:
			explodeArgs := CheckSystemError(arg.Execute(e), e)
			for _, val := range explodeArgs.([]Node) {
				b.Args = append(b.Args, RawAstNode(val, e))
			}
		default:
			b.Args = append(b.Args, RawAstNode(arg, e))
		}
	}
}

func (b *FuncDef) Type() CmdType {
	return FunctionDefAstCmd
}

// RawAstNode fetch the raw value in the RXScript
// used in built-in functions
func RawAstNode(astNode any, e RxeHandler) any {
	for {
		switch v := astNode.(type) {
		case *ReturnDef:
			astNode = v.Value
		case *IntDef, *StringDef, *BoolDef, *DoubleDef, *NullDef:
			return v
		case *ArrayDef, *DictionaryDef:
			return v
		case *ChannelDef, *WaitDef, *HttpRequest, *OsFile, *Scanner, *TimeTicker, *TimeTimer:
			return v
		case *ClassBodyDef, *FunctionDef, *SplitArrayDef, *AtomicIntDef:
			return v
		case *SystemResource:
			return v.Execute(e)
		case *Class:
			return v
		case Node:
			astNode = v.Execute(e)
		default:
			panic(fmt.Sprintf("unknown type %T", astNode))
		}
	}
}

// rawValue fetch the original value in the RXScript
// used in RXScript engine && built-in functions.
func rawValue(qv any, e RxeHandler) any {
	for {
		switch v := qv.(type) {
		case *IntDef:
			return v.Value
		case *StringDef:
			return v.Value
		case *BoolDef:
			return v.Value
		case *DoubleDef:
			return v.Value
		case *NullDef:
			return nil
		case *ArrayDef:
			return v.Value
		case *AtomicIntDef:
			return v.Value.Load()
		case *ChannelDef:
			return v.Value
		case *DictionaryDef:
			return v.Values
		case *ReturnDef:
			qv = v.Value
		case *SplitArrayDef:
			return v
		case *SystemResource:
			return v.obj.Interface()
		case *HttpRequest:
			return v.Value
		case *OsFile:
			return v.Value
		case *Scanner:
			return v.Value
		case *TimeTicker:
			return v.Value
		case *TimeTimer:
			return v.Value
		case *Class:
			return v.variables
		case Node:
			qv = v.Execute(e)
		default:
			return qv
		}
	}
}

// rawValueToPtr fetch the original value in the RXScript
// used in RXScript engine && built-in functions.
func rawValueToPtr(qv any, e RxeHandler) any {
	for {
		switch v := qv.(type) {
		case *IntDef:
			return &v.Value
		case *StringDef:
			return &v.Value
		case *BoolDef:
			return &v.Value
		case *DoubleDef:
			return &v.Value
		case *NullDef:
			return nil
		case Node:
			qv = v.Execute(e)
		default:
			return &qv
		}
	}
}

// toNodeValue turn Go's base type into RXScript Node type
// support int, string, bool, null, boolean, []any, map[string]any
// other types return null
func toNodeValue(qv any, e RxeHandler) any {
	switch v := qv.(type) {
	case int64:
		return &IntDef{Value: v}
	case int:
		return &IntDef{Value: int64(v)}
	case int8:
		return &IntDef{Value: int64(v)}
	case int16:
		return &IntDef{Value: int64(v)}
	case int32:
		return &IntDef{Value: int64(v)}
	case uint64:
		// if value is uint64, turn into string
		return &StringDef{Value: strconv.FormatUint(v, 10)}
	case uint:
		return &IntDef{Value: int64(v)}
	case uint8:
		return &IntDef{Value: int64(v)}
	case uint16:
		return &IntDef{Value: int64(v)}
	case uint32:
		return &IntDef{Value: int64(v)}
	case string:
		return &StringDef{Value: v}
	case bool:
		return &BoolDef{Value: v}
	case float64:
		return &DoubleDef{Value: v}
	case float32:
		return &DoubleDef{Value: float64(v)}
	case []any:
		ad := &ArrayDef{Value: make([]Node, 0, 32)}
		for i, _ := range v {
			ad.Value = append(ad.Value, toNodeValue(v[i], e).(Node))
		}
		return ad
	case map[string]any:
		vd := &DictionaryDef{Values: make(map[string]Node, 32)}
		for key, value := range v {
			vd.Values[key] = toNodeValue(value, e).(Node)
		}
		return vd
	}
	return &NullDef{}
}

// NodeStringValue fetch the string value of the node in RXScript
func NodeStringValue(astNode any, e RxeHandler) any {
	for {
		switch v := astNode.(type) {
		case *IntDef:
			return v.String()
		case *StringDef:
			return v.String()
		case *BoolDef:
			return v.String()
		case *DoubleDef:
			return v.String()
		case *ArrayDef:
			return v.String()
		case *NullDef:
			return v.String()
		case *ReturnDef:
			astNode = v.Value
		case *SplitArrayDef:
			return v.Value.String()
		case *AtomicIntDef:
			return v.String()
		case *DictionaryDef:
			return v.String()
		case *ChannelDef:
			return v.String()
		case *SystemResource:
			return v.String()
		case *TimeTicker:
			return v.String()
		case *TimeTimer:
			return v.String()
		case *HttpRequest:
			return v.String()
		case *OsFile:
			return v.String()
		case *Scanner:
			return v.String()
		case *Class:
			return v.String()
		case Node:
			astNode = v.Execute(e)
		default:
			return astNode
		}
	}
}
