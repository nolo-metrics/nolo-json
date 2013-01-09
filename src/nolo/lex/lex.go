package nolo

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return fmt.Sprintf("%s: %s", i.typ, i.val)
	case len(i.val) > 10:
		return fmt.Sprintf("%s: %.10q...", i.typ, i.val)
	}
	return fmt.Sprintf("%s: %q", i.typ, i.val)
}

type itemType int

const (
	itemError            itemType = iota // error occurred; value is text of error
	itemIdentifier                       // metric identifier
	itemValue                            // metric value
	itemOptionIdentifier                 // option identifier
	itemEqual                            // equals sign, seperates option identifier from value
	itemOptionValue                      // option value
	itemEOF
)

// Make the types prettyprint.
var itemName = map[itemType]string{
	itemError:            "error",
	itemIdentifier:       "identifier",
	itemValue:            "value",
	itemOptionIdentifier: "option identifier",
	itemEqual:            "=",
	itemOptionValue:      "option value",
	itemEOF:              "EOF",
}

func (i itemType) String() string {
	s := itemName[i]
	if s == "" {
		return fmt.Sprintf("item%d", int(i))
	}
	return s
}

const eof = -1

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

type termFn func(*lexer) bool

// lexer holds the state of the scanner.
type lexer struct {
	name  string    // the name of the input; used only for error reports.
	input string    // the string being scanned.
	state stateFn   // the next lexing function to enter.
	pos   int       // current position in the input.
	start int       // start position of this item.
	width int       // width of last rune read from input.
	items chan item // channel of scanned items.
}

// next returns the next rune in the input.
func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// error returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, fmt.Sprintf(format, args...)}
	return nil
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() item {
	for {
		select {
		case item := <-l.items:
			return item
		default:
			l.state = l.state(l)
		}
	}
	panic("not reached")
}

// lex creates a new scanner for the input string.
func Lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		state: lexText,
		items: make(chan item, 2), // Two items sufficient.
	}
	return l
}

// lexEOL scans the end of line
func lexEOL(l *lexer) stateFn {
	l.ignore() // drop EOL token, transition info is sufficient
	return lexText
}

// lexEOF scans the end of line
func lexEOF(l *lexer) stateFn {
	l.emit(itemEOF)
	return nil
}

// lexText scans the elements for a single row
func lexText(l *lexer) stateFn {
	switch r := l.next(); {
	case r == eof:
		if l.pos > l.start {
			// TODO: write test case
			return l.errorf("Unexpected character: %#U", r)
		}
		return lexEOF
	case r == '\n':
		return lexEOL
	case isAlphaNumeric(r):
		l.backup()
		return lexIdentifier
	case isSpace(r):
		l.ignore()
		return lexText
	default:
		// TODO: write test case
		return l.errorf("unrecognized character in text: %#U", r)
	}
	panic("not reached")
}

// lexIdentifier scans an alphanumeric
func lexIdentifier(l *lexer) stateFn {
	if !l.scanIdentifierUntil(atIdentifierTerminator) {
		return l.errorf("unexpected character after identifier: %q",
			l.input[l.start:l.pos])
	}
	l.emit(itemIdentifier)
	return lexValue
}

// lexValue scans a decimal number, allowing for optional sign and point
func lexValue(l *lexer) stateFn {
	// ignore spaces
	for {
		if isSpace(l.next()) {
			l.ignore()
		} else {
			l.backup()
			break
		}
	}

	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	l.emit(itemValue)
	return lexTail
}

// lexTail scans the optional pairs at the end of a line
func lexTail(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof:
			if l.pos > l.start {
				// TODO: write test case
				return l.errorf("Unexpected EOF after option identifier: %q", l.input[l.start:l.pos])
			}
			return lexEOF
		case r == '\n':
			return lexEOL
		case isSpace(r):
			l.ignore()
		case isAlphaNumeric(r):
			l.backup()
			return lexOptionIdentifier
		default:
			l.backup()
			// TODO: write test case
			return l.errorf("unexpected character before option identifier: %q", l.input[l.start:l.pos])
		}
	}
	panic("not reached")
}

// lexOptionIdentifier scans an alphanumeric up to a '='
func lexOptionIdentifier(l *lexer) stateFn {
	if !l.scanIdentifierUntil(atOptionIdentifierTerminator) {
		// TODO: write test case
		return l.errorf("unexpected character in option identifier: %q", l.input[l.start:l.pos])
	}
	l.emit(itemOptionIdentifier)
	l.next()
	l.ignore() // drop '=' as a token, transitioning to the right lexer is enough
	return lexOptionValue
}

// lexOptionValue scans an alphanumeric, number or quoted string
func lexOptionValue(l *lexer) stateFn {
	r := l.peek()
	switch {
	case l.scanNumber():
		l.emit(itemOptionValue)
	case r == '"':
		return lexQuote
	case l.scanIdentifierUntil(atOptionValueTerminator):
		l.emit(itemOptionValue)
	default:
		// TODO: write test case
		return l.errorf("unexpected character in option value: %#U", r)
	}
	return lexTail
}

// lexQuote scans a quoted string.
func lexQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string: %q", l.input[l.start:l.pos])
		case '"':
			break Loop
		}
	}
	l.emit(itemOptionValue)
	return lexTail
}

// scanIdentifierUntil scans an alphanumeric up to a terminator
func (l *lexer) scanIdentifierUntil(terminator termFn) bool {
	for {
		r := l.next()
		// absorb alphanumerics
		if !isAlphaNumeric(r) {
			l.backup()
			if !terminator(l) {
				return false
			}
			return true
		}
	}
	panic("not reached")
}

// scanNumber scans a decimal with optional sign and point 
func (l *lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	digits := "0123456789"
	// requires at least one digit
	if !l.accept(digits) {
		l.next()
		return false
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	// Next thing mustn't be alphanumeric.
	if isAlphaNumeric(l.peek()) {
		l.next()
		return false
	}
	return true
}

// atIdentifierTerminator reports whether the input is at
// valid termination character to appear after an identifier.
func atIdentifierTerminator(l *lexer) bool {
	r := l.peek()
	switch {
	case isSpace(r):
		return true
	}
	return false
}

// atOptionIdentifierTerminator reports whether the input is at
// valid termination character to appear after an option identifier.
func atOptionIdentifierTerminator(l *lexer) bool {
	r := l.peek()
	switch {
	case r == '=':
		return true
	}
	return false
}

// atOptionValueTerminator reports whether the input is at
// valid termination character to appear after an option value
func atOptionValueTerminator(l *lexer) bool {
	r := l.peek()
	switch {
	case isSpace(r):
		return true
	case r == '\n':
		return true
	case r == eof:
		return true
	}
	return false
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	switch r {
	case ' ', '\t', '\r':
		return true
	}
	return false
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
