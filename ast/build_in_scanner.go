//
// The RXScript author
// built-in functions

package ast

import (
	"bufio"
	"bytes"
)

type BuffScanner struct {
	FuncDef
}

func (p *BuffScanner) String() string {
	return "'newScanner'"
}

func NewBuffScanner() *BuffScanner {
	return &BuffScanner{}
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// ScanLines is a split function for a [Scanner] that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// SplitByString split string used by bufio.Scanner Split func
// any char in chars will be split and return
func SplitByString(chars string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.IndexAny(data, chars); i >= 0 {
			return i + 1, dropCR(data[:i]), nil
		}

		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

func (p *BuffScanner) Execute(rxe RxeHandler) any {
	if len(p.Args) < 2 {
		return ThrowSystemError("RXScript: newScanner need file resource & split chars to read data.")
	}

	f := p.Args[0]
	splitChars := p.Args[1]

	switch of := f.(type) {
	case *OsFile:
		s := bufio.NewScanner(of.Value)
		switch sc := splitChars.(type) {
		case *StringDef:
			s.Split(SplitByString(sc.Value))
		default:
			s.Split(ScanLines)
		}
		return &Scanner{Value: s}
	case *StringDef:
		s := bufio.NewScanner(bytes.NewBufferString(of.Value))
		switch sc := splitChars.(type) {
		case *StringDef:
			s.Split(SplitByString(sc.Value))
		default:
			s.Split(ScanLines)
		}
		return &Scanner{Value: s}
	default:
		return ThrowSystemError("RXScript: newScanner need file resource & split chars to read data.")
	}
}
