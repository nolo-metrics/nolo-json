package nolo

import (
//	"reflect"
	"testing"
)

type parseTest struct {
	name  string
	input string
	items string
}

var parseTests = []parseTest{
	{"empty", "", "[]"},
	{"single",
		"guage 99\n",
		"[{\"identifier\":\"guage\", \"value\":99}]",
	},
	{"signed value",
		"guage +99\n",
		"[{\"identifier\":\"guage\", \"value\":99}]",
	},
	{"negative signed value",
		"guage -99\n",
		"[{\"identifier\":\"guage\", \"value\":-99}]",
	},
	{"floating point value",
		"guage 99.9\n",
		"[{\"identifier\":\"guage\", \"value\":99.9}]",
	},
	{"single with options",
		"guage 99 key=value\n",
		"[{\"identifier\":\"guage\", \"value\":99, \"meta\":{\"key\":\"value\"}}]",
	},
	{"single with numeric options",
		"guage 99 key=216\n",
		"[{\"identifier\":\"guage\", \"value\":99, \"meta\":{\"key\":\"216\"}}]",
	},
	{"multiple",
		"guage 999\ncounter 5\n",
		"[{\"identifier\":\"guage\", \"value\":99}], [{\"identifier\":\"counter\", \"value\":5}]",
	},
	{"extra spaces",
		" guage  99  key=value  foo=bar \n",
		"[{\"identifier\":\"guage\", \"value\":99, \"meta\":{\"key\":\"value\", \"foo\":\"bar\"}]",
	},
	{"single with signed numeric options",
		"guage 99 key=+216\n",
		"[{\"identifier\":\"guage\", \"value\":99, \"meta\":{\"key\":\"+216\"}]",
	},
	{"single with quoted options",
		"guage 99 key=\"value\"\n",
		"[{\"identifier\":\"guage\", \"value\":99, \"meta\":{\"key\":\"value\"}]",
	},
	{"single with quoted options including an escaped quote",
		"guage 99 key=\"va\\\"lue\"\n",
		"[{\"identifier\":\"guage\", \"value\":99, \"meta\":{\"key\":\"va\\\"lue\"}]",
	},
	{"single with quoted options including an escaped newline",
		"guage 99 key=\"va\\nlue\"\n",
		"[{\"identifier\":\"guage\", \"value\":99, \"meta\":{\"key\":\"va\\\"lue\"}]",
	},

	// Failures
	//	{"fails for value m", "guage m\n", []item{
	//		{itemIdentifier, "guage"},
	//		{itemError, "bad number syntax: \"m\""},
	//	}},
	//	{"fails for value .", "guage .\n", []item{
	//		{itemIdentifier, "guage"},
	//		{itemError, "bad number syntax: \".\""},
	//	}},
	//	{"fails for missing value", "guage\n", []item{
	//		{itemError, "unexpected character after identifier: \"guage\""},
	//	}},
	//	{"fails for eof after identifier", "guage", []item{
	//		{itemError, "unexpected character after identifier: \"guage\""},
	//	}},
	//	{"fail for quoted string with newline", "guage 99 key=\"va\n", []item{
	//		{itemIdentifier, "guage"},
	//		{itemValue, "99"},
	//		{itemOptionIdentifier, "key"},
	//		{itemError, "unterminated quoted string: \"\\\"va\\n\""},
	//	}},
	//	{"fail for quoted string with eof", "guage 99 key=\"va", []item{
	//		{itemIdentifier, "guage"},
	//		{itemValue, "99"},
	//		{itemOptionIdentifier, "key"},
	//		{itemError, "unterminated quoted string: \"\\\"va\""},
	//	}},
}

// collect gathers the emitted items into a slice.
func collectParseTests(t *parseTest) (items []item) {
//	l := Lex(t.name, t.input)
//	for {
//		item := l.nextItem()
//		items = append(items, item)
//		if item.typ == itemEOF || item.typ == itemError {
//			break
//		}
//	}
	return
}

func TestParse(t *testing.T) {
//	for _, test := range lexTests {
//		items := collect(&test)
//		if !reflect.DeepEqual(items, test.items) {
//			t.Errorf("%s: got\n\t%v\nexpected\n\t%v", test.name, items, test.items)
//		}
//	}
}
