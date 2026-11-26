%{
//
// The RXScript author
// RXScript's json grammer analyzer
// use goyacc to generate: goyacc -p rxeJson -o rx_json_yacc.go rx_json_yacc.y
//

package ast
%}

%union {
  val Node
}

%type <val>
  start object value array
  members member elements element

%token <val>
  JSON_ERROR "symbol"
  JSON_LEFT_BRACE "'{'"
  JSON_RIGHT_BRACE "'}'"
  JSON_LEFT_SQUARE_BRACKET "'['"
  JSON_RIGHT_SQUARE_BRACKET "']'"
  JSON_COLON "':'"
  JSON_COMMA "','"
  JSON_NULL "'null'"
  JSON_STRING "'string'"
  JSON_INT "'int'"
  JSON_BOOL "'bool'"
  JSON_DOUBLE "'double'"

%% /* Rules */


start:
  value
  { yylex.(*JsonLexer).result = $1; return 0; }
;

object:
  JSON_LEFT_BRACE members JSON_RIGHT_BRACE
  { $$ = $2 }
;

members:
  /* empty */
  { $$ = &DictionaryDef{} }
| member
  { $$ = $1 }
;

member:
  JSON_STRING JSON_COLON value
  { $$ = &DictionaryDef{Values: map[string]Node{$1.(*StringDef).Value: $3.(Node)}} }
| member JSON_COMMA JSON_STRING JSON_COLON value
  {
    $1.(*DictionaryDef).Values[$3.(*StringDef).Value] = $5.(Node)
    $$ = $1;
  }
;

array:
  JSON_LEFT_SQUARE_BRACKET elements JSON_RIGHT_SQUARE_BRACKET
  { $$ = $2 }
;

elements:
  /*empty*/
  { $$ = &ArrayDef{} }
| element
  { $$ = $1 }
;

element:
  value
  { $$ = &ArrayDef{Value: []Node{$1.(Node)}} }
| element JSON_COMMA value
  { $1.(*ArrayDef).Value = append($1.(*ArrayDef).Value, $3.(Node)); $$ = $1; }
;

value:
  object
  { $$ = $1; }
| array
  { $$ = $1; }
| JSON_NULL
  { $$ = $1 }
| JSON_STRING
  { $$ = $1 }
| JSON_INT
  { $$ = $1 }
| JSON_BOOL
  { $$ = $1 }
| JSON_DOUBLE
  { $$ = $1 }
;


%%
