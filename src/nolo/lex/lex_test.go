package nolo

import (
	"reflect"
	"testing"
)

type lexTest struct {
	name  string
	input string
	items []item
}

var (
	tEOF = item{itemEOF, ""}
	raw  = "`" + `abc\n\t\" ` + "`"
)

var lexTests = []lexTest{
	{"empty", "", []item{tEOF}},
	{"single", "guage 99\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		tEOF,
	}},
	{"signed value", "guage +99\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "+99"},
		tEOF,
	}},
	{"negative signed value", "guage -99\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "-99"},
		tEOF,
	}},
	{"floating point value", "guage 99.9\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99.9"},
		tEOF,
	}},

	{"single with options", "guage 99 key=value\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemOptionValue, "value"},
		tEOF,
	}},
	{"single with numeric options", "guage 99 key=216\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemOptionValue, "216"},
		tEOF,
	}},
	{"multiple", "guage 999\ncounter 5\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "999"},
		{itemIdentifier, "counter"},
		{itemValue, "5"},
		tEOF,
	}},

	{"extra spaces", " guage  99  key=value  foo=bar \n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemOptionValue, "value"},
		{itemOptionIdentifier, "foo"},
		{itemOptionValue, "bar"},
		tEOF,
	}},

	{"single with signed numeric options", "guage 99 key=+216\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemOptionValue, "+216"},
		tEOF,
	}},
	{"single with quoted options", "guage 99 key=\"value\"\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemOptionValue, "\"value\""},
		tEOF,
	}},
	{"single with quoted options including an escaped quote", "guage 99 key=\"va\\\"lue\"\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemOptionValue, "\"va\\\"lue\""},
		tEOF,
	}},
	{"single with quoted options including an escaped newline", "guage 99 key=\"va\\nlue\"\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemOptionValue, "\"va\\nlue\""},
		tEOF,
	}},

	// Failures
	{"fails for value m", "guage m\n", []item{
		{itemIdentifier, "guage"},
		{itemError, "bad number syntax: \"m\""},
	}},
	{"fails for value .", "guage .\n", []item{
		{itemIdentifier, "guage"},
		{itemError, "bad number syntax: \".\""},
	}},
	{"fails for missing value", "guage\n", []item{
		{itemError, "unexpected character after identifier: \"guage\""},
	}},
	{"fails for eof after identifier", "guage", []item{
		{itemError, "unexpected character after identifier: \"guage\""},
	}},
	{"fail for quoted string with newline", "guage 99 key=\"va\n", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemError, "unterminated quoted string: \"\\\"va\\n\""},
	}},
	{"fail for quoted string with eof", "guage 99 key=\"va", []item{
		{itemIdentifier, "guage"},
		{itemValue, "99"},
		{itemOptionIdentifier, "key"},
		{itemError, "unterminated quoted string: \"\\\"va\""},
	}},
}

// collect gathers the emitted items into a slice.
func collect(t *lexTest) (items []item) {
	l := Lex(t.name, t.input)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.typ == itemEOF || item.typ == itemError {
			break
		}
	}
	return
}

func TestLex(t *testing.T) {
	for _, test := range lexTests {
		items := collect(&test)
		if !reflect.DeepEqual(items, test.items) {
			t.Errorf("%s: got\n\t%v\nexpected\n\t%v", test.name, items, test.items)
		}
	}
}
