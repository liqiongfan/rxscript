package ast

const (
	ScriptSuffix = ".rxt"
)

var (
	FloatPrecisionWidth = 32
)

type ValueType int

const (
	String ValueType = iota + 1
	Int
	Bool
	Null
	Double
	Variable
	GetAttribute
	FuncCall
	Array
	FetchArray
)

type RxeLexer struct {
	// script contents
	buff string

	// script file name
	fileName string

	// re2go need
	cursor, backup int64

	// token value
	value interface{}

	// 1. string 2: int 3: bool 4: null 5: double
	valueType ValueType

	// re2go condition
	condition int

	// token position
	line, column int64

	// result
	result []any
}

const (
	yycstd     = 0
	yycpath    = 1
	yyccall    = 2
	yyccomment = 3
)
