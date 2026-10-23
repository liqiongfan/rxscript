// DO NOT EDIT.
//line "rx_engine_language.re":1
//
// The RXScript author
//

package ast

import (
	"strconv"
)

func getToken(lex *RxeLexer) int {

	//line "rx_engine_language_scanner.go":16
	var t1 int64
	var t2 int64
	//line "rx_engine_language.re":12

	//line "rx_engine_language_scanner.go":22
	var yyt1 int64
	//line "rx_engine_language.re":13

std:

	//line "rx_engine_language_scanner.go":29
	{
		var yych byte
		yyaccept := 0
		switch lex.condition {
		case yycstd:
			goto yyc_std
		case yycpath:
			goto yyc_path
		case yyccall:
			goto yyc_call
		case yyccomment:
			goto yyc_comment
		default:
			panic("internal lexer error")
		}
		/* *********************************** */
	yyc_std:
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			goto yy1
		case 0x07, 0x08, '\t':
			fallthrough
		case '\r':
			fallthrough
		case ' ':
			yyt1 = lex.cursor
			goto yy4
		case '\n':
			yyt1 = lex.cursor
			goto yy6
		case '!':
			goto yy8
		case '"':
			goto yy9
		case '&':
			goto yy10
		case '(':
			goto yy12
		case ')':
			goto yy13
		case '*':
			goto yy14
		case '+':
			goto yy15
		case ',':
			goto yy17
		case '-':
			goto yy18
		case '.':
			yyt1 = lex.cursor
			goto yy20
		case '/':
			yyt1 = lex.cursor
			goto yy22
		case '0':
			yyt1 = lex.cursor
			goto yy24
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			yyt1 = lex.cursor
			goto yy26
		case ':':
			goto yy27
		case '<':
			goto yy28
		case '=':
			goto yy30
		case '>':
			goto yy32
		case '?':
			goto yy34
		case '@':
			goto yy35
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'h':
			fallthrough
		case 'j', 'k', 'l', 'm':
			fallthrough
		case 'o':
			fallthrough
		case 'q':
			fallthrough
		case 's':
			fallthrough
		case 'u', 'v', 'w', 'x', 'y', 'z':
			yyt1 = lex.cursor
			goto yy36
		case '[':
			goto yy39
		case ']':
			goto yy40
		case '^':
			goto yy41
		case 'a':
			yyt1 = lex.cursor
			goto yy42
		case 'b':
			yyt1 = lex.cursor
			goto yy43
		case 'c':
			yyt1 = lex.cursor
			goto yy44
		case 'd':
			yyt1 = lex.cursor
			goto yy45
		case 'e':
			yyt1 = lex.cursor
			goto yy46
		case 'f':
			yyt1 = lex.cursor
			goto yy47
		case 'g':
			yyt1 = lex.cursor
			goto yy48
		case 'i':
			yyt1 = lex.cursor
			goto yy49
		case 'n':
			yyt1 = lex.cursor
			goto yy50
		case 'p':
			yyt1 = lex.cursor
			goto yy51
		case 'r':
			yyt1 = lex.cursor
			goto yy52
		case 't':
			yyt1 = lex.cursor
			goto yy53
		case '{':
			goto yy54
		case '|':
			goto yy55
		case '}':
			goto yy57
		default:
			goto yy2
		}
	yy1:
		lex.cursor++
		//line "rx_engine_language.re":28
		{
			return 0
		}
		//line "rx_engine_language_scanner.go":175
	yy2:
		lex.cursor++
	yy3:
		//line "rx_engine_language.re":454
		{
			println(lex.condition)
			return ERROR
		}
		//line "rx_engine_language_scanner.go":181
	yy4:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x07, 0x08, '\t':
			fallthrough
		case '\r':
			fallthrough
		case ' ':
			goto yy4
		default:
			goto yy5
		}
	yy5:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":395
		{
			lex.column += t2 - t1
			goto std
		}
		//line "rx_engine_language_scanner.go":203
	yy6:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '\n':
			goto yy6
		default:
			goto yy7
		}
	yy7:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":400
		{
			lex.column = 1
			lex.line += t2 - t1
			goto std
		}
		//line "rx_engine_language_scanner.go":222
	yy8:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '=':
			goto yy58
		default:
			goto yy3
		}
	yy9:
		yyaccept = 0
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			fallthrough
		case '\n':
			goto yy3
		case '"':
			yyt1 = lex.cursor
			goto yy61
		default:
			yyt1 = lex.cursor
			goto yy59
		}
	yy10:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '&':
			goto yy62
		default:
			goto yy11
		}
	yy11:
		//line "rx_engine_language.re":267
		{
			lex.condition = yycstd
			lex.column += 1
			return BITAND
		}
		//line "rx_engine_language_scanner.go":265
	yy12:
		lex.cursor++
		//line "rx_engine_language.re":99
		{
			lex.condition = yycstd
			lex.column += 1
			return LEFT_ROUND_BRACKET
		}
		//line "rx_engine_language_scanner.go":274
	yy13:
		lex.cursor++
		//line "rx_engine_language.re":105
		{
			lex.condition = yyccall
			lex.column += 1
			return RIGHT_ROUND_BRACKET
		}
		//line "rx_engine_language_scanner.go":283
	yy14:
		lex.cursor++
		//line "rx_engine_language.re":377
		{
			lex.condition = yycstd
			lex.column += 1
			return MULTIPLY
		}
		//line "rx_engine_language_scanner.go":292
	yy15:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '+':
			goto yy63
		case '=':
			goto yy64
		default:
			goto yy16
		}
	yy16:
		//line "rx_engine_language.re":365
		{
			lex.condition = yycstd
			lex.column += 1
			return PLUS
		}
		//line "rx_engine_language_scanner.go":311
	yy17:
		lex.cursor++
		//line "rx_engine_language.re":123
		{
			lex.condition = yycstd
			lex.column += 1
			return COMMA
		}
		//line "rx_engine_language_scanner.go":320
	yy18:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '-':
			goto yy65
		case '=':
			goto yy66
		default:
			goto yy19
		}
	yy19:
		//line "rx_engine_language.re":371
		{
			lex.condition = yycstd
			lex.column += 1
			return MINUS
		}
		//line "rx_engine_language_scanner.go":339
	yy20:
		yyaccept = 1
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case '.':
			goto yy67
		default:
			goto yy69
		}
	yy21:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":357
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseFloat(lex.buff[t1:t2], 64)
			lex.valueType = Double
			lex.column += t2 - t1
			return DOUBLE
		}
		//line "rx_engine_language_scanner.go":362
	yy22:
		yyaccept = 2
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case '*':
			goto yy70
		case '/':
			goto yy71
		default:
			goto yy23
		}
	yy23:
		//line "rx_engine_language.re":383
		{
			lex.condition = yycstd
			lex.column += 1
			return DIVIDE
		}
		//line "rx_engine_language_scanner.go":383
	yy24:
		yyaccept = 3
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case '.':
			goto yy68
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'E':
			fallthrough
		case 'e':
			goto yy73
		case 'B':
			fallthrough
		case 'b':
			goto yy76
		case 'O':
			fallthrough
		case 'o':
			goto yy78
		case 'X':
			fallthrough
		case 'x':
			goto yy79
		default:
			goto yy25
		}
	yy25:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":291
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 10, 64)
			lex.column += t2 - t1
			lex.valueType = Int
			return INT
		}
		//line "rx_engine_language_scanner.go":424
	yy26:
		yyaccept = 3
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case '.':
			goto yy68
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			goto yy26
		case 'E':
			fallthrough
		case 'e':
			goto yy77
		default:
			goto yy25
		}
	yy27:
		lex.cursor++
		//line "rx_engine_language.re":219
		{
			lex.condition = yycstd
			lex.column += 1
			return COLON
		}
		//line "rx_engine_language_scanner.go":450
	yy28:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '-':
			goto yy80
		case '<':
			goto yy81
		case '=':
			goto yy82
		default:
			goto yy29
		}
	yy29:
		//line "rx_engine_language.re":243
		{
			lex.condition = yycstd
			lex.column += 1
			return LT
		}
		//line "rx_engine_language_scanner.go":471
	yy30:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '=':
			goto yy83
		default:
			goto yy31
		}
	yy31:
		//line "rx_engine_language.re":153
		{
			lex.condition = yycstd
			lex.column += 1
			return EQUAL
		}
		//line "rx_engine_language_scanner.go":488
	yy32:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '=':
			goto yy84
		case '>':
			goto yy85
		default:
			goto yy33
		}
	yy33:
		//line "rx_engine_language.re":231
		{
			lex.condition = yycstd
			lex.column += 1
			return GT
		}
		//line "rx_engine_language_scanner.go":507
	yy34:
		lex.cursor++
		//line "rx_engine_language.re":207
		{
			lex.condition = yycstd
			lex.column += 1
			return QUESTION
		}
		//line "rx_engine_language_scanner.go":516
	yy35:
		lex.cursor++
		//line "rx_engine_language.re":129
		{
			lex.condition = yycstd
			lex.column += 2
			return AT
		}
		//line "rx_engine_language_scanner.go":525
	yy36:
		lex.cursor++
		yych = lex.buff[lex.cursor]
	yy37:
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy38
		}
	yy38:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":434
		{
			lex.value = lex.buff[t1:t2]
			lex.column += t2 - t1
			lex.condition = yyccall
			return NAME
		}
		//line "rx_engine_language_scanner.go":552
	yy39:
		lex.cursor++
		//line "rx_engine_language.re":111
		{
			lex.condition = yycstd
			lex.column += 1
			return LEFT_SQUARE_BRACKET
		}
		//line "rx_engine_language_scanner.go":561
	yy40:
		lex.cursor++
		//line "rx_engine_language.re":117
		{
			lex.condition = yycstd
			lex.column += 1
			return RIGHT_SQUARE_BRACKET
		}
		//line "rx_engine_language_scanner.go":570
	yy41:
		lex.cursor++
		//line "rx_engine_language.re":279
		{
			lex.condition = yycstd
			lex.column += 1
			return BITEXP
		}
		//line "rx_engine_language_scanner.go":579
	yy42:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy86
		default:
			goto yy37
		}
	yy43:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy88
		default:
			goto yy37
		}
	yy44:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy89
		default:
			goto yy37
		}
	yy45:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy90
		default:
			goto yy37
		}
	yy46:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy91
		default:
			goto yy37
		}
	yy47:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy92
		case 'o':
			goto yy93
		default:
			goto yy37
		}
	yy48:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'o':
			goto yy94
		default:
			goto yy37
		}
	yy49:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'f':
			goto yy96
		case 'm':
			goto yy98
		default:
			goto yy37
		}
	yy50:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'u':
			goto yy99
		default:
			goto yy37
		}
	yy51:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy100
		case 'u':
			goto yy101
		default:
			goto yy37
		}
	yy52:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy102
		case 'e':
			goto yy103
		default:
			goto yy37
		}
	yy53:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy104
		default:
			goto yy37
		}
	yy54:
		lex.cursor++
		//line "rx_engine_language.re":75
		{
			lex.condition = yycstd
			lex.column += 1
			return LEFT_BRACES
		}
		//line "rx_engine_language_scanner.go":704
	yy55:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '|':
			goto yy105
		default:
			goto yy56
		}
	yy56:
		//line "rx_engine_language.re":273
		{
			lex.condition = yycstd
			lex.column += 1
			return BITOR
		}
		//line "rx_engine_language_scanner.go":721
	yy57:
		lex.cursor++
		//line "rx_engine_language.re":81
		{
			lex.condition = yycstd
			lex.column += 1
			return RIGHT_BRACES
		}
		//line "rx_engine_language_scanner.go":730
	yy58:
		lex.cursor++
		//line "rx_engine_language.re":285
		{
			lex.condition = yycstd
			lex.column += 2
			return NE
		}
		//line "rx_engine_language_scanner.go":739
	yy59:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			fallthrough
		case '\n':
			goto yy60
		case '"':
			goto yy61
		default:
			goto yy59
		}
	yy60:
		lex.cursor = lex.backup
		switch yyaccept {
		case 0:
			goto yy3
		case 1:
			goto yy21
		case 2:
			goto yy23
		case 3:
			goto yy25
		default:
			goto yy74
		}
	yy61:
		lex.cursor++
		t1 = yyt1
		t2 = lex.cursor
		t2 += -1
		//line "rx_engine_language.re":61
		{
			lex.condition = yycstd
			lex.value = lex.buff[t1:t2]
			lex.valueType = String
			lex.column += t2 - t1 + 2
			return STRING
		}
		//line "rx_engine_language_scanner.go":780
	yy62:
		lex.cursor++
		//line "rx_engine_language.re":225
		{
			lex.condition = yycstd
			lex.column += 2
			return AND
		}
		//line "rx_engine_language_scanner.go":789
	yy63:
		lex.cursor++
		//line "rx_engine_language.re":171
		{
			lex.condition = yycstd
			lex.column += 2
			return PLUSDOUBLE
		}
		//line "rx_engine_language_scanner.go":798
	yy64:
		lex.cursor++
		//line "rx_engine_language.re":159
		{
			lex.condition = yycstd
			lex.column += 2
			return PLUSEQUAL
		}
		//line "rx_engine_language_scanner.go":807
	yy65:
		lex.cursor++
		//line "rx_engine_language.re":201
		{
			lex.condition = yycstd
			lex.column += 2
			return MINUSDOUBLE
		}
		//line "rx_engine_language_scanner.go":816
	yy66:
		lex.cursor++
		//line "rx_engine_language.re":177
		{
			lex.condition = yycstd
			lex.column += 2
			return MINUSQUAL
		}
		//line "rx_engine_language_scanner.go":825
	yy67:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '.':
			goto yy106
		default:
			goto yy60
		}
	yy68:
		lex.cursor++
		yych = lex.buff[lex.cursor]
	yy69:
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			goto yy68
		default:
			goto yy21
		}
	yy70:
		lex.cursor++
		t1 = lex.cursor
		t1 += -2
		t2 = lex.cursor
		//line "rx_engine_language.re":406
		{
			lex.column += t2 - t1
			lex.condition = yyccomment
			goto std
		}
		//line "rx_engine_language_scanner.go":856
	yy71:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '\n':
			goto yy60
		default:
			goto yy108
		}
	yy72:
		yyaccept = 4
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
	yy73:
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7':
			goto yy72
		case '8', '9':
			goto yy75
		case 'E':
			fallthrough
		case 'e':
			goto yy77
		default:
			goto yy74
		}
	yy74:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":307
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
			lex.column += t2 - t1
			lex.valueType = Int
			return INT
		}
		//line "rx_engine_language_scanner.go":895
	yy75:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			goto yy75
		case 'E':
			fallthrough
		case 'e':
			goto yy77
		default:
			goto yy60
		}
	yy76:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1':
			goto yy109
		default:
			goto yy60
		}
	yy77:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '-':
			goto yy111
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			goto yy68
		default:
			goto yy60
		}
	yy78:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7':
			goto yy112
		default:
			goto yy60
		}
	yy79:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f':
			goto yy113
		default:
			goto yy60
		}
	yy80:
		lex.cursor++
		//line "rx_engine_language.re":147
		{
			lex.condition = yycstd
			lex.column += 2
			return CHANNELTO
		}
		//line "rx_engine_language_scanner.go":959
	yy81:
		lex.cursor++
		//line "rx_engine_language.re":255
		{
			lex.condition = yycstd
			lex.column += 2
			return LEFTSHIFT
		}
		//line "rx_engine_language_scanner.go":968
	yy82:
		lex.cursor++
		//line "rx_engine_language.re":249
		{
			lex.condition = yycstd
			lex.column += 2
			return LTE
		}
		//line "rx_engine_language_scanner.go":977
	yy83:
		lex.cursor++
		//line "rx_engine_language.re":165
		{
			lex.condition = yycstd
			lex.column += 2
			return EQUALDOUBLE
		}
		//line "rx_engine_language_scanner.go":986
	yy84:
		lex.cursor++
		//line "rx_engine_language.re":237
		{
			lex.condition = yycstd
			lex.column += 2
			return GTE
		}
		//line "rx_engine_language_scanner.go":995
	yy85:
		lex.cursor++
		//line "rx_engine_language.re":261
		{
			lex.condition = yycstd
			lex.column += 2
			return RIGHTSHIFT
		}
		//line "rx_engine_language_scanner.go":1004
	yy86:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy87
		}
	yy87:
		//line "rx_engine_language.re":345
		{
			lex.condition = yycstd
			lex.column += 2
			return AS
		}
		//line "rx_engine_language_scanner.go":1027
	yy88:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy115
		default:
			goto yy37
		}
	yy89:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy116
		default:
			goto yy37
		}
	yy90:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'f':
			goto yy117
		default:
			goto yy37
		}
	yy91:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy118
		default:
			goto yy37
		}
	yy92:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy119
		default:
			goto yy37
		}
	yy93:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy120
		default:
			goto yy37
		}
	yy94:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy95
		}
	yy95:
		//line "rx_engine_language.re":135
		{
			lex.condition = yycstd
			lex.column += 2
			return GO
		}
		//line "rx_engine_language_scanner.go":1104
	yy96:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy97
		}
	yy97:
		//line "rx_engine_language.re":339
		{
			lex.condition = yycstd
			lex.column += 2
			return IF
		}
		//line "rx_engine_language_scanner.go":1127
	yy98:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'p':
			goto yy122
		default:
			goto yy37
		}
	yy99:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy123
		default:
			goto yy37
		}
	yy100:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'i':
			goto yy124
		default:
			goto yy37
		}
	yy101:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'b':
			goto yy125
		default:
			goto yy37
		}
	yy102:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'n':
			goto yy126
		default:
			goto yy37
		}
	yy103:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 't':
			goto yy127
		default:
			goto yy37
		}
	yy104:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'u':
			goto yy128
		default:
			goto yy37
		}
	yy105:
		lex.cursor++
		//line "rx_engine_language.re":213
		{
			lex.condition = yycstd
			lex.column += 2
			return OR
		}
		//line "rx_engine_language_scanner.go":1199
	yy106:
		lex.cursor++
		//line "rx_engine_language.re":447
		{
			lex.column += 3
			lex.value = `...`
			lex.condition = yycstd
			return OMITMORE
		}
		//line "rx_engine_language_scanner.go":1209
	yy107:
		lex.cursor++
		yych = lex.buff[lex.cursor]
	yy108:
		switch yych {
		case 0x00:
			goto yy60
		case '\n':
			goto yy129
		default:
			goto yy107
		}
	yy109:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1':
			goto yy109
		default:
			goto yy110
		}
	yy110:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":299
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
			lex.column += t2 - t1
			lex.valueType = Int
			return INT
		}
		//line "rx_engine_language_scanner.go":1242
	yy111:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			goto yy68
		default:
			goto yy60
		}
	yy112:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7':
			goto yy112
		default:
			goto yy74
		}
	yy113:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f':
			goto yy113
		default:
			goto yy114
		}
	yy114:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":315
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
			lex.column += t2 - t1
			lex.valueType = Int
			return INT
		}
		//line "rx_engine_language_scanner.go":1285
	yy115:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy130
		default:
			goto yy37
		}
	yy116:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy131
		default:
			goto yy37
		}
	yy117:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy132
		default:
			goto yy37
		}
	yy118:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy133
		default:
			goto yy37
		}
	yy119:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy128
		default:
			goto yy37
		}
	yy120:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy121
		}
	yy121:
		//line "rx_engine_language.re":183
		{
			lex.condition = yycstd
			lex.column += 3
			return FOR
		}
		//line "rx_engine_language_scanner.go":1353
	yy122:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'o':
			goto yy135
		default:
			goto yy37
		}
	yy123:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy136
		default:
			goto yy37
		}
	yy124:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'v':
			goto yy138
		default:
			goto yy37
		}
	yy125:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy139
		default:
			goto yy37
		}
	yy126:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'g':
			goto yy140
		default:
			goto yy37
		}
	yy127:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'u':
			goto yy141
		default:
			goto yy37
		}
	yy128:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy142
		default:
			goto yy37
		}
	yy129:
		lex.cursor++
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":412
		{
			lex.column += 1
			lex.line += 1
			goto std
		}
		//line "rx_engine_language_scanner.go":1427
	yy130:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'k':
			goto yy144
		default:
			goto yy37
		}
	yy131:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy146
		default:
			goto yy37
		}
	yy132:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy148
		default:
			goto yy37
		}
	yy133:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy134
		}
	yy134:
		//line "rx_engine_language.re":351
		{
			lex.condition = yycstd
			lex.column += 4
			return ELSE
		}
		//line "rx_engine_language_scanner.go":1477
	yy135:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy150
		default:
			goto yy37
		}
	yy136:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy137
		}
	yy137:
		t1 = lex.cursor
		t1 += -4
		t2 = lex.cursor
		//line "rx_engine_language.re":331
		{
			lex.condition = yycstd
			lex.value = lex.buff[t1:t2]
			lex.valueType = Null
			lex.column += t2 - t1
			return NULL
		}
		//line "rx_engine_language_scanner.go":1514
	yy138:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy151
		default:
			goto yy37
		}
	yy139:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'i':
			goto yy152
		default:
			goto yy37
		}
	yy140:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy153
		default:
			goto yy37
		}
	yy141:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy155
		default:
			goto yy37
		}
	yy142:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy143
		}
	yy143:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":323
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseBool(lex.buff[t1:t2])
			lex.valueType = Bool
			lex.column += t2 - t1
			return BOOL
		}
		//line "rx_engine_language_scanner.go":1577
	yy144:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy145
		}
	yy145:
		//line "rx_engine_language.re":189
		{
			lex.condition = yycstd
			lex.column += 5
			return BREAK
		}
		//line "rx_engine_language_scanner.go":1600
	yy146:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy147
		}
	yy147:
		//line "rx_engine_language.re":69
		{
			lex.condition = yycstd
			lex.column += 5
			return CLASS
		}
		//line "rx_engine_language_scanner.go":1623
	yy148:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy149
		}
	yy149:
		//line "rx_engine_language.re":141
		{
			lex.condition = yycstd
			lex.column += 5
			return DEFER
		}
		//line "rx_engine_language_scanner.go":1646
	yy150:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 't':
			goto yy156
		default:
			goto yy37
		}
	yy151:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 't':
			goto yy158
		default:
			goto yy37
		}
	yy152:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'c':
			goto yy159
		default:
			goto yy37
		}
	yy153:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy154
		}
	yy154:
		//line "rx_engine_language.re":195
		{
			lex.condition = yycstd
			lex.column += 5
			return RANGE
		}
		//line "rx_engine_language_scanner.go":1696
	yy155:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'n':
			goto yy161
		default:
			goto yy37
		}
	yy156:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy157
		}
	yy157:
		//line "rx_engine_language.re":30
		{
			lex.column += 6
			lex.condition = yycpath
			return IMPORT
		}
		//line "rx_engine_language_scanner.go":1728
	yy158:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy163
		default:
			goto yy37
		}
	yy159:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy160
		}
	yy160:
		//line "rx_engine_language.re":93
		{
			lex.condition = yycstd
			lex.column += 6
			return PUBLIC
		}
		//line "rx_engine_language_scanner.go":1760
	yy161:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy162
		}
	yy162:
		//line "rx_engine_language.re":389
		{
			lex.condition = yycstd
			lex.column += 6
			return RETURN
		}
		//line "rx_engine_language_scanner.go":1783
	yy163:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy36
		default:
			goto yy164
		}
	yy164:
		//line "rx_engine_language.re":87
		{
			lex.condition = yycstd
			lex.column += 7
			return PRIVATE
		}
		//line "rx_engine_language_scanner.go":1806
		/* *********************************** */
	yyc_path:
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			goto yy166
		case 0x07, 0x08, '\t':
			fallthrough
		case '\r':
			fallthrough
		case ' ':
			goto yy169
		case '\n':
			yyt1 = lex.cursor
			goto yy170
		case '"':
			goto yy172
		case '.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			yyt1 = lex.cursor
			goto yy173
		default:
			goto yy167
		}
	yy166:
		lex.cursor++
		//line "rx_engine_language.re":28
		{
			return 0
		}
		//line "rx_engine_language_scanner.go":1840
	yy167:
		lex.cursor++
	yy168:
		//line "rx_engine_language.re":454
		{
			println(lex.condition)
			return ERROR
		}
		//line "rx_engine_language_scanner.go":1846
	yy169:
		lex.cursor++
		t1 = lex.cursor
		t1 += -1
		t2 = lex.cursor
		//line "rx_engine_language.re":36
		{
			lex.column += t2 - t1
			goto std
		}
		//line "rx_engine_language_scanner.go":1857
	yy170:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '\n':
			goto yy170
		default:
			goto yy171
		}
	yy171:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":41
		{
			lex.line += t2 - t1
			lex.column = 1
			goto std
		}
		//line "rx_engine_language_scanner.go":1876
	yy172:
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case '.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			yyt1 = lex.cursor
			goto yy175
		default:
			goto yy168
		}
	yy173:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy173
		default:
			goto yy174
		}
	yy174:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":54
		{
			lex.condition = yycstd
			lex.value = lex.buff[t1:t2]
			lex.column += t2 - t1
			return SYSTEMPATH
		}
		//line "rx_engine_language_scanner.go":1919
	yy175:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '"':
			goto yy177
		case '.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy175
		default:
			goto yy176
		}
	yy176:
		lex.cursor = lex.backup
		goto yy168
	yy177:
		lex.cursor++
		t1 = yyt1
		t2 = lex.cursor
		t2 += -1
		//line "rx_engine_language.re":47
		{
			lex.condition = yycstd
			lex.value = lex.buff[t1:t2]
			lex.column += t2 - t1
			return PATH
		}
		//line "rx_engine_language_scanner.go":1952
		/* *********************************** */
	yyc_call:
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			goto yy179
		case 0x07, 0x08, '\t':
			fallthrough
		case '\r':
			fallthrough
		case ' ':
			yyt1 = lex.cursor
			goto yy182
		case '\n':
			yyt1 = lex.cursor
			goto yy184
		case '!':
			goto yy186
		case '"':
			goto yy187
		case '&':
			goto yy188
		case '(':
			goto yy190
		case ')':
			goto yy191
		case '*':
			goto yy192
		case '+':
			goto yy193
		case ',':
			goto yy195
		case '-':
			goto yy196
		case '.':
			goto yy198
		case '/':
			yyt1 = lex.cursor
			goto yy200
		case '0':
			yyt1 = lex.cursor
			goto yy202
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			yyt1 = lex.cursor
			goto yy204
		case ':':
			goto yy205
		case '<':
			goto yy206
		case '=':
			goto yy208
		case '>':
			goto yy210
		case '?':
			goto yy212
		case '@':
			goto yy213
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'h':
			fallthrough
		case 'j', 'k', 'l', 'm':
			fallthrough
		case 'o':
			fallthrough
		case 'q':
			fallthrough
		case 's':
			fallthrough
		case 'u', 'v', 'w', 'x', 'y', 'z':
			yyt1 = lex.cursor
			goto yy214
		case '[':
			goto yy217
		case ']':
			goto yy218
		case '^':
			goto yy219
		case 'a':
			yyt1 = lex.cursor
			goto yy220
		case 'b':
			yyt1 = lex.cursor
			goto yy221
		case 'c':
			yyt1 = lex.cursor
			goto yy222
		case 'd':
			yyt1 = lex.cursor
			goto yy223
		case 'e':
			yyt1 = lex.cursor
			goto yy224
		case 'f':
			yyt1 = lex.cursor
			goto yy225
		case 'g':
			yyt1 = lex.cursor
			goto yy226
		case 'i':
			yyt1 = lex.cursor
			goto yy227
		case 'n':
			yyt1 = lex.cursor
			goto yy228
		case 'p':
			yyt1 = lex.cursor
			goto yy229
		case 'r':
			yyt1 = lex.cursor
			goto yy230
		case 't':
			yyt1 = lex.cursor
			goto yy231
		case '{':
			goto yy232
		case '|':
			goto yy233
		case '}':
			goto yy235
		default:
			goto yy180
		}
	yy179:
		lex.cursor++
		//line "rx_engine_language.re":28
		{
			return 0
		}
		//line "rx_engine_language_scanner.go":2082
	yy180:
		lex.cursor++
	yy181:
		//line "rx_engine_language.re":454
		{
			println(lex.condition)
			return ERROR
		}
		//line "rx_engine_language_scanner.go":2088
	yy182:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x07, 0x08, '\t':
			fallthrough
		case '\r':
			fallthrough
		case ' ':
			goto yy182
		default:
			goto yy183
		}
	yy183:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":395
		{
			lex.column += t2 - t1
			goto std
		}
		//line "rx_engine_language_scanner.go":2110
	yy184:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '\n':
			goto yy184
		default:
			goto yy185
		}
	yy185:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":400
		{
			lex.column = 1
			lex.line += t2 - t1
			goto std
		}
		//line "rx_engine_language_scanner.go":2129
	yy186:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '=':
			goto yy236
		default:
			goto yy181
		}
	yy187:
		yyaccept = 0
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			fallthrough
		case '\n':
			goto yy181
		case '"':
			yyt1 = lex.cursor
			goto yy239
		default:
			yyt1 = lex.cursor
			goto yy237
		}
	yy188:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '&':
			goto yy240
		default:
			goto yy189
		}
	yy189:
		//line "rx_engine_language.re":267
		{
			lex.condition = yycstd
			lex.column += 1
			return BITAND
		}
		//line "rx_engine_language_scanner.go":2172
	yy190:
		lex.cursor++
		//line "rx_engine_language.re":99
		{
			lex.condition = yycstd
			lex.column += 1
			return LEFT_ROUND_BRACKET
		}
		//line "rx_engine_language_scanner.go":2181
	yy191:
		lex.cursor++
		//line "rx_engine_language.re":105
		{
			lex.condition = yyccall
			lex.column += 1
			return RIGHT_ROUND_BRACKET
		}
		//line "rx_engine_language_scanner.go":2190
	yy192:
		lex.cursor++
		//line "rx_engine_language.re":377
		{
			lex.condition = yycstd
			lex.column += 1
			return MULTIPLY
		}
		//line "rx_engine_language_scanner.go":2199
	yy193:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '+':
			goto yy241
		case '=':
			goto yy242
		default:
			goto yy194
		}
	yy194:
		//line "rx_engine_language.re":365
		{
			lex.condition = yycstd
			lex.column += 1
			return PLUS
		}
		//line "rx_engine_language_scanner.go":2218
	yy195:
		lex.cursor++
		//line "rx_engine_language.re":123
		{
			lex.condition = yycstd
			lex.column += 1
			return COMMA
		}
		//line "rx_engine_language_scanner.go":2227
	yy196:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '-':
			goto yy243
		case '=':
			goto yy244
		default:
			goto yy197
		}
	yy197:
		//line "rx_engine_language.re":371
		{
			lex.condition = yycstd
			lex.column += 1
			return MINUS
		}
		//line "rx_engine_language_scanner.go":2246
	yy198:
		yyaccept = 1
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case '.':
			goto yy245
		default:
			goto yy199
		}
	yy199:
		//line "rx_engine_language.re":441
		{
			lex.column += 1
			lex.condition = yycstd
			return CALL
		}
		//line "rx_engine_language_scanner.go":2265
	yy200:
		yyaccept = 2
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case '*':
			goto yy246
		case '/':
			goto yy247
		default:
			goto yy201
		}
	yy201:
		//line "rx_engine_language.re":383
		{
			lex.condition = yycstd
			lex.column += 1
			return DIVIDE
		}
		//line "rx_engine_language_scanner.go":2286
	yy202:
		yyaccept = 3
		lex.cursor++
		lex.backup = lex.cursor
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7':
			goto yy248
		case 'B':
			fallthrough
		case 'b':
			goto yy250
		case 'O':
			fallthrough
		case 'o':
			goto yy251
		case 'X':
			fallthrough
		case 'x':
			goto yy252
		default:
			goto yy203
		}
	yy203:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":291
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 10, 64)
			lex.column += t2 - t1
			lex.valueType = Int
			return INT
		}
		//line "rx_engine_language_scanner.go":2321
	yy204:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			goto yy204
		default:
			goto yy203
		}
	yy205:
		lex.cursor++
		//line "rx_engine_language.re":219
		{
			lex.condition = yycstd
			lex.column += 1
			return COLON
		}
		//line "rx_engine_language_scanner.go":2339
	yy206:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '-':
			goto yy253
		case '<':
			goto yy254
		case '=':
			goto yy255
		default:
			goto yy207
		}
	yy207:
		//line "rx_engine_language.re":243
		{
			lex.condition = yycstd
			lex.column += 1
			return LT
		}
		//line "rx_engine_language_scanner.go":2360
	yy208:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '=':
			goto yy256
		default:
			goto yy209
		}
	yy209:
		//line "rx_engine_language.re":153
		{
			lex.condition = yycstd
			lex.column += 1
			return EQUAL
		}
		//line "rx_engine_language_scanner.go":2377
	yy210:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '=':
			goto yy257
		case '>':
			goto yy258
		default:
			goto yy211
		}
	yy211:
		//line "rx_engine_language.re":231
		{
			lex.condition = yycstd
			lex.column += 1
			return GT
		}
		//line "rx_engine_language_scanner.go":2396
	yy212:
		lex.cursor++
		//line "rx_engine_language.re":207
		{
			lex.condition = yycstd
			lex.column += 1
			return QUESTION
		}
		//line "rx_engine_language_scanner.go":2405
	yy213:
		lex.cursor++
		//line "rx_engine_language.re":129
		{
			lex.condition = yycstd
			lex.column += 2
			return AT
		}
		//line "rx_engine_language_scanner.go":2414
	yy214:
		lex.cursor++
		yych = lex.buff[lex.cursor]
	yy215:
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy216
		}
	yy216:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":434
		{
			lex.value = lex.buff[t1:t2]
			lex.column += t2 - t1
			lex.condition = yyccall
			return NAME
		}
		//line "rx_engine_language_scanner.go":2441
	yy217:
		lex.cursor++
		//line "rx_engine_language.re":111
		{
			lex.condition = yycstd
			lex.column += 1
			return LEFT_SQUARE_BRACKET
		}
		//line "rx_engine_language_scanner.go":2450
	yy218:
		lex.cursor++
		//line "rx_engine_language.re":117
		{
			lex.condition = yycstd
			lex.column += 1
			return RIGHT_SQUARE_BRACKET
		}
		//line "rx_engine_language_scanner.go":2459
	yy219:
		lex.cursor++
		//line "rx_engine_language.re":279
		{
			lex.condition = yycstd
			lex.column += 1
			return BITEXP
		}
		//line "rx_engine_language_scanner.go":2468
	yy220:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy259
		default:
			goto yy215
		}
	yy221:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy261
		default:
			goto yy215
		}
	yy222:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy262
		default:
			goto yy215
		}
	yy223:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy263
		default:
			goto yy215
		}
	yy224:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy264
		default:
			goto yy215
		}
	yy225:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy265
		case 'o':
			goto yy266
		default:
			goto yy215
		}
	yy226:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'o':
			goto yy267
		default:
			goto yy215
		}
	yy227:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'f':
			goto yy269
		default:
			goto yy215
		}
	yy228:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'u':
			goto yy271
		default:
			goto yy215
		}
	yy229:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy272
		case 'u':
			goto yy273
		default:
			goto yy215
		}
	yy230:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy274
		case 'e':
			goto yy275
		default:
			goto yy215
		}
	yy231:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy276
		default:
			goto yy215
		}
	yy232:
		lex.cursor++
		//line "rx_engine_language.re":75
		{
			lex.condition = yycstd
			lex.column += 1
			return LEFT_BRACES
		}
		//line "rx_engine_language_scanner.go":2591
	yy233:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '|':
			goto yy277
		default:
			goto yy234
		}
	yy234:
		//line "rx_engine_language.re":273
		{
			lex.condition = yycstd
			lex.column += 1
			return BITOR
		}
		//line "rx_engine_language_scanner.go":2608
	yy235:
		lex.cursor++
		//line "rx_engine_language.re":81
		{
			lex.condition = yycstd
			lex.column += 1
			return RIGHT_BRACES
		}
		//line "rx_engine_language_scanner.go":2617
	yy236:
		lex.cursor++
		//line "rx_engine_language.re":285
		{
			lex.condition = yycstd
			lex.column += 2
			return NE
		}
		//line "rx_engine_language_scanner.go":2626
	yy237:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			fallthrough
		case '\n':
			goto yy238
		case '"':
			goto yy239
		default:
			goto yy237
		}
	yy238:
		lex.cursor = lex.backup
		switch yyaccept {
		case 0:
			goto yy181
		case 1:
			goto yy199
		case 2:
			goto yy201
		default:
			goto yy203
		}
	yy239:
		lex.cursor++
		t1 = yyt1
		t2 = lex.cursor
		t2 += -1
		//line "rx_engine_language.re":61
		{
			lex.condition = yycstd
			lex.value = lex.buff[t1:t2]
			lex.valueType = String
			lex.column += t2 - t1 + 2
			return STRING
		}
		//line "rx_engine_language_scanner.go":2665
	yy240:
		lex.cursor++
		//line "rx_engine_language.re":225
		{
			lex.condition = yycstd
			lex.column += 2
			return AND
		}
		//line "rx_engine_language_scanner.go":2674
	yy241:
		lex.cursor++
		//line "rx_engine_language.re":171
		{
			lex.condition = yycstd
			lex.column += 2
			return PLUSDOUBLE
		}
		//line "rx_engine_language_scanner.go":2683
	yy242:
		lex.cursor++
		//line "rx_engine_language.re":159
		{
			lex.condition = yycstd
			lex.column += 2
			return PLUSEQUAL
		}
		//line "rx_engine_language_scanner.go":2692
	yy243:
		lex.cursor++
		//line "rx_engine_language.re":201
		{
			lex.condition = yycstd
			lex.column += 2
			return MINUSDOUBLE
		}
		//line "rx_engine_language_scanner.go":2701
	yy244:
		lex.cursor++
		//line "rx_engine_language.re":177
		{
			lex.condition = yycstd
			lex.column += 2
			return MINUSQUAL
		}
		//line "rx_engine_language_scanner.go":2710
	yy245:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '.':
			goto yy278
		default:
			goto yy238
		}
	yy246:
		lex.cursor++
		t1 = lex.cursor
		t1 += -2
		t2 = lex.cursor
		//line "rx_engine_language.re":406
		{
			lex.column += t2 - t1
			lex.condition = yyccomment
			goto std
		}
		//line "rx_engine_language_scanner.go":2731
	yy247:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '\n':
			goto yy238
		default:
			goto yy280
		}
	yy248:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7':
			goto yy248
		default:
			goto yy249
		}
	yy249:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":307
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
			lex.column += t2 - t1
			lex.valueType = Int
			return INT
		}
		//line "rx_engine_language_scanner.go":2761
	yy250:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1':
			goto yy281
		default:
			goto yy238
		}
	yy251:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7':
			goto yy248
		default:
			goto yy238
		}
	yy252:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f':
			goto yy283
		default:
			goto yy238
		}
	yy253:
		lex.cursor++
		//line "rx_engine_language.re":147
		{
			lex.condition = yycstd
			lex.column += 2
			return CHANNELTO
		}
		//line "rx_engine_language_scanner.go":2801
	yy254:
		lex.cursor++
		//line "rx_engine_language.re":255
		{
			lex.condition = yycstd
			lex.column += 2
			return LEFTSHIFT
		}
		//line "rx_engine_language_scanner.go":2810
	yy255:
		lex.cursor++
		//line "rx_engine_language.re":249
		{
			lex.condition = yycstd
			lex.column += 2
			return LTE
		}
		//line "rx_engine_language_scanner.go":2819
	yy256:
		lex.cursor++
		//line "rx_engine_language.re":165
		{
			lex.condition = yycstd
			lex.column += 2
			return EQUALDOUBLE
		}
		//line "rx_engine_language_scanner.go":2828
	yy257:
		lex.cursor++
		//line "rx_engine_language.re":237
		{
			lex.condition = yycstd
			lex.column += 2
			return GTE
		}
		//line "rx_engine_language_scanner.go":2837
	yy258:
		lex.cursor++
		//line "rx_engine_language.re":261
		{
			lex.condition = yycstd
			lex.column += 2
			return RIGHTSHIFT
		}
		//line "rx_engine_language_scanner.go":2846
	yy259:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy260
		}
	yy260:
		//line "rx_engine_language.re":345
		{
			lex.condition = yycstd
			lex.column += 2
			return AS
		}
		//line "rx_engine_language_scanner.go":2869
	yy261:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy285
		default:
			goto yy215
		}
	yy262:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy286
		default:
			goto yy215
		}
	yy263:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'f':
			goto yy287
		default:
			goto yy215
		}
	yy264:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy288
		default:
			goto yy215
		}
	yy265:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy289
		default:
			goto yy215
		}
	yy266:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy290
		default:
			goto yy215
		}
	yy267:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy268
		}
	yy268:
		//line "rx_engine_language.re":135
		{
			lex.condition = yycstd
			lex.column += 2
			return GO
		}
		//line "rx_engine_language_scanner.go":2946
	yy269:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy270
		}
	yy270:
		//line "rx_engine_language.re":339
		{
			lex.condition = yycstd
			lex.column += 2
			return IF
		}
		//line "rx_engine_language_scanner.go":2969
	yy271:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy292
		default:
			goto yy215
		}
	yy272:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'i':
			goto yy293
		default:
			goto yy215
		}
	yy273:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'b':
			goto yy294
		default:
			goto yy215
		}
	yy274:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'n':
			goto yy295
		default:
			goto yy215
		}
	yy275:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 't':
			goto yy296
		default:
			goto yy215
		}
	yy276:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'u':
			goto yy297
		default:
			goto yy215
		}
	yy277:
		lex.cursor++
		//line "rx_engine_language.re":213
		{
			lex.condition = yycstd
			lex.column += 2
			return OR
		}
		//line "rx_engine_language_scanner.go":3032
	yy278:
		lex.cursor++
		//line "rx_engine_language.re":447
		{
			lex.column += 3
			lex.value = `...`
			lex.condition = yycstd
			return OMITMORE
		}
		//line "rx_engine_language_scanner.go":3042
	yy279:
		lex.cursor++
		yych = lex.buff[lex.cursor]
	yy280:
		switch yych {
		case 0x00:
			goto yy238
		case '\n':
			goto yy298
		default:
			goto yy279
		}
	yy281:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1':
			goto yy281
		default:
			goto yy282
		}
	yy282:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":299
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
			lex.column += t2 - t1
			lex.valueType = Int
			return INT
		}
		//line "rx_engine_language_scanner.go":3075
	yy283:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f':
			goto yy283
		default:
			goto yy284
		}
	yy284:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":315
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseInt(lex.buff[t1:t2], 0, 64)
			lex.column += t2 - t1
			lex.valueType = Int
			return INT
		}
		//line "rx_engine_language_scanner.go":3100
	yy285:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy299
		default:
			goto yy215
		}
	yy286:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy300
		default:
			goto yy215
		}
	yy287:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy301
		default:
			goto yy215
		}
	yy288:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy302
		default:
			goto yy215
		}
	yy289:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy297
		default:
			goto yy215
		}
	yy290:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy291
		}
	yy291:
		//line "rx_engine_language.re":183
		{
			lex.condition = yycstd
			lex.column += 3
			return FOR
		}
		//line "rx_engine_language_scanner.go":3168
	yy292:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy304
		default:
			goto yy215
		}
	yy293:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'v':
			goto yy306
		default:
			goto yy215
		}
	yy294:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'l':
			goto yy307
		default:
			goto yy215
		}
	yy295:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'g':
			goto yy308
		default:
			goto yy215
		}
	yy296:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'u':
			goto yy309
		default:
			goto yy215
		}
	yy297:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy310
		default:
			goto yy215
		}
	yy298:
		lex.cursor++
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":412
		{
			lex.column += 1
			lex.line += 1
			goto std
		}
		//line "rx_engine_language_scanner.go":3233
	yy299:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'k':
			goto yy312
		default:
			goto yy215
		}
	yy300:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 's':
			goto yy314
		default:
			goto yy215
		}
	yy301:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy316
		default:
			goto yy215
		}
	yy302:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy303
		}
	yy303:
		//line "rx_engine_language.re":351
		{
			lex.condition = yycstd
			lex.column += 4
			return ELSE
		}
		//line "rx_engine_language_scanner.go":3283
	yy304:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy305
		}
	yy305:
		t1 = lex.cursor
		t1 += -4
		t2 = lex.cursor
		//line "rx_engine_language.re":331
		{
			lex.condition = yycstd
			lex.value = lex.buff[t1:t2]
			lex.valueType = Null
			lex.column += t2 - t1
			return NULL
		}
		//line "rx_engine_language_scanner.go":3311
	yy306:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'a':
			goto yy318
		default:
			goto yy215
		}
	yy307:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'i':
			goto yy319
		default:
			goto yy215
		}
	yy308:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy320
		default:
			goto yy215
		}
	yy309:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'r':
			goto yy322
		default:
			goto yy215
		}
	yy310:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy311
		}
	yy311:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":323
		{
			lex.condition = yycstd
			lex.value, _ = strconv.ParseBool(lex.buff[t1:t2])
			lex.valueType = Bool
			lex.column += t2 - t1
			return BOOL
		}
		//line "rx_engine_language_scanner.go":3374
	yy312:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy313
		}
	yy313:
		//line "rx_engine_language.re":189
		{
			lex.condition = yycstd
			lex.column += 5
			return BREAK
		}
		//line "rx_engine_language_scanner.go":3397
	yy314:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy315
		}
	yy315:
		//line "rx_engine_language.re":69
		{
			lex.condition = yycstd
			lex.column += 5
			return CLASS
		}
		//line "rx_engine_language_scanner.go":3420
	yy316:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy317
		}
	yy317:
		//line "rx_engine_language.re":141
		{
			lex.condition = yycstd
			lex.column += 5
			return DEFER
		}
		//line "rx_engine_language_scanner.go":3443
	yy318:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 't':
			goto yy323
		default:
			goto yy215
		}
	yy319:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'c':
			goto yy324
		default:
			goto yy215
		}
	yy320:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy321
		}
	yy321:
		//line "rx_engine_language.re":195
		{
			lex.condition = yycstd
			lex.column += 5
			return RANGE
		}
		//line "rx_engine_language_scanner.go":3484
	yy322:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'n':
			goto yy326
		default:
			goto yy215
		}
	yy323:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 'e':
			goto yy328
		default:
			goto yy215
		}
	yy324:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy325
		}
	yy325:
		//line "rx_engine_language.re":93
		{
			lex.condition = yycstd
			lex.column += 6
			return PUBLIC
		}
		//line "rx_engine_language_scanner.go":3525
	yy326:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy327
		}
	yy327:
		//line "rx_engine_language.re":389
		{
			lex.condition = yycstd
			lex.column += 6
			return RETURN
		}
		//line "rx_engine_language_scanner.go":3548
	yy328:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case '_':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			goto yy214
		default:
			goto yy329
		}
	yy329:
		//line "rx_engine_language.re":87
		{
			lex.condition = yycstd
			lex.column += 7
			return PRIVATE
		}
		//line "rx_engine_language_scanner.go":3571
		/* *********************************** */
	yyc_comment:
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			goto yy331
		case '\n':
			yyt1 = lex.cursor
			goto yy334
		case '*':
			yyt1 = lex.cursor
			goto yy336
		default:
			yyt1 = lex.cursor
			goto yy332
		}
	yy331:
		lex.cursor++
		//line "rx_engine_language.re":28
		{
			return 0
		}
		//line "rx_engine_language_scanner.go":3592
	yy332:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case 0x00:
			fallthrough
		case '\n':
			fallthrough
		case '*':
			goto yy333
		default:
			goto yy332
		}
	yy333:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":424
		{
			lex.column += t2 - t1
			goto std
		}
		//line "rx_engine_language_scanner.go":3614
	yy334:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '\n':
			goto yy334
		default:
			goto yy335
		}
	yy335:
		t1 = yyt1
		t2 = lex.cursor
		//line "rx_engine_language.re":418
		{
			lex.line += t2 - t1
			lex.column = 1
			goto std
		}
		//line "rx_engine_language_scanner.go":3633
	yy336:
		lex.cursor++
		yych = lex.buff[lex.cursor]
		switch yych {
		case '/':
			goto yy339
		default:
			goto yy338
		}
	yy337:
		lex.cursor++
		yych = lex.buff[lex.cursor]
	yy338:
		switch yych {
		case 0x00:
			fallthrough
		case '\n':
			fallthrough
		case '/':
			goto yy333
		default:
			goto yy337
		}
	yy339:
		lex.cursor++
		//line "rx_engine_language.re":429
		{
			lex.condition = yycstd
			goto std
		}
		//line "rx_engine_language_scanner.go":3664
	}
	//line "rx_engine_language.re":456

}
