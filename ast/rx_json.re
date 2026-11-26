//
// The RXScript author
// RXScript's json lexer
// use re2go generate: re2go -c -o rx_json_scanner.go rx_json.re
//

package ast

import (
  "os"

  "strconv"
)

func getJsonToken(lex *JsonLexer) int {
    /*!svars:re2c format='var @@ int64\n';*/
    /*!stags:re2c format='var @@ int64\n';*/

    std:
    %{
      re2c:yyfill:enable=0;
      re2c:YYCTYPE="byte";
      re2c:YYPEEK="lex.buff[lex.cursor]";
      re2c:YYGETCONDITION="lex.condition";
      re2c:YYSTAGP="@@{tag}=lex.cursor";
      re2c:YYSKIP="lex.cursor++";
      re2c:YYBACKUP="lex.backup=lex.cursor";
      re2c:YYRESTORE="lex.cursor=lex.backup";
      re2c:YYSHIFTSTAG="@@{tag}+=@@{shift}";
      re2c:tags=1;

      <*> "\x00" {
        if lex.condition != yycstd {
          lex.Error("Unexpected end")
          os.Exit(1)
        }
        return 0
      }

      <std> "{"  { return JSON_LEFT_BRACE }

      <std> "}"  { return JSON_RIGHT_BRACE }

      <std> "["  { return JSON_LEFT_SQUARE_BRACKET }

      <std> "]"  { return JSON_RIGHT_SQUARE_BRACKET }

      <std> ":"  { return JSON_COLON }

      <std> ","  { return JSON_COMMA }

      <std> "null"  { return JSON_NULL }

      <std> ["] {
        lex.condition = yycstring
        lex.begin = lex.cursor
        lex.column++
        goto std
      }

      <string> @t1 [^\x00\n"]+ @t2 {
        lex.column+= t2 - t1
        goto std
      }

      <string> @t1 [\n]+ @t2 {
        lex.line+= t2 - t1
        lex.column=1
        goto std
      }

      <string> ["] {
        if lex.buff[lex.cursor-2] == '\\' {
          goto std
        }
        lex.value = lex.buff[lex.begin:lex.cursor-1]
        lex.valueType = String
        lex.column += lex.cursor-lex.begin+2
        lex.condition=yycstd
        return JSON_STRING
      }

      <std> @t1 ("0" | [1-9][0-9]*) @t2 {
        lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 10, 64)
        lex.column += t2-t1
        return JSON_INT
      }

      <std> @t1 '0b'[0-1]+ @t2 {
        lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
        lex.column += t2-t1
        return JSON_INT
      }

      <std> @t1 ('0'|'0o')[0-7]+ @t2 {
        lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
        lex.column += t2-t1
        return JSON_INT
      }

      <std> @t1 '0x'[0-9a-fA-F]+ @t2 {
        lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
        lex.column += t2-t1
        return JSON_INT
      }

      <std> @t1 ("true"|"false") @t2 {
        lex.value, _ = strconv.ParseBool(lex.buff[t1:t2])
        lex.column += t2-t1
        return JSON_BOOL
      }

      <std> @t1 (("0."|"."|[1-9][0-9]*".")[0-9]*|[0-9]+'e'"-"?[0-9]+) @t2 {
        lex.value, _ = strconv.ParseFloat(lex.buff[t1:t2], 64)
        lex.column += t2-t1
        return JSON_DOUBLE
      }

      <std> @t1 [\n]+ @t2 {
        lex.line += t2 - t1
        lex.column = 1
        goto std
      }

      <std> @t1 [ \a\b\r\t] @t2 {
        lex.column += t2-t1
        goto std
      }

      <*>* { return JSON_ERROR; }
    %}
}
