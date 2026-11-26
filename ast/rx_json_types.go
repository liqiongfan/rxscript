//
// The RXScript author
//

package ast

import (
	"encoding/json"
	"fmt"
)

type JsonLexer struct {
	buff string

	// script file name
	fileName string

	// re2go need
	cursor, backup int64

	// token value
	value any

	// 1. string 2: int 3: bool 4: null 5: double
	valueType ValueType

	// re2go condition
	condition int

	// token position
	line, column int64

	begin int64

	// result
	result Node
}

func (lex *JsonLexer) Lex(lval *rxeJsonSymType) int {
	token := getJsonToken(lex)
	// fmt.Printf("token: %d, %s, %v\n", token, rxeJsonSymNames[rxeJsonXLAT[token]], lex.value)
	switch token {
	case JSON_INT:
		lval.val = &IntDef{Value: lex.value.(int64)}
	case JSON_STRING:
		lval.val = NewStringDef(lex.value.(string))
	case JSON_BOOL:
		lval.val = &BoolDef{Value: lex.value.(bool)}
	case JSON_DOUBLE:
		lval.val = &DoubleDef{Value: lex.value.(float64)}
	case JSON_NULL:
		lval.val = &NullDef{}
	}
	return token
}

// If yacc or re2go occur mistake, Error will be invoked.
func (lex *JsonLexer) Error(s string) {
	fmt.Printf("RXScript: %s on line %d\n", s, lex.line)
}

// jsonDecode parse json buff into RXScript dictionary
// if error meet, return null
func jsonDecode(buff string) Node {
	lexer := &JsonLexer{
		buff:      buff + "\x00",
		cursor:    0,
		backup:    0,
		value:     nil,
		condition: 0,
		line:      1,
		column:    1,
	}
	if rxeJsonParse(lexer) != 0 {
		return &NullDef{}
	}

	return lexer.result
}

// jsonEncode encode RXScript object into json string
func jsonEncode(v Node, e RxeHandler) *StringDef {
	buf, _ := json.Marshal(v.Any(e))
	return &StringDef{Value: string(buf)}
}
