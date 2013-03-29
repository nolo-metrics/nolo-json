package nolo

import (
	"reflect"
	"testing"
)

type ToMapTest struct {
	name  string
	input Plugin
	expected map[string][]map[string]string
}

var ToMapTests = []ToMapTest{
	{"empty",
		Plugin{"empty", []Metric{}}, 
		map[string][]map[string]string { "empty": []map[string]string {}}},
	{"single", 
		Plugin{"single", []Metric{ Metric { "guage", "99", map[string]string {}}}}, 
		map[string][]map[string]string {
			"single": []map[string]string {
				map[string]string {
					"identifier": "guage",
					"value": "99" }}}},
	{"single with options", 
		Plugin{"single-with-options",
			[]Metric{ Metric { "guage", "99",
					map[string]string { "key": "value" }}}}, 
		map[string][]map[string]string {
			"single-with-options": []map[string]string {
				map[string]string {
					"identifier": "guage",
					"value": "99",
					"key": "value" }}}},
	{"multiple", 
		Plugin{"multiple",
			[]Metric{
				Metric { "guage", "99", map[string]string {}},
				Metric { "counter", "5", map[string]string {}}}}, 
		map[string][]map[string]string {
			"multiple": []map[string]string {
				map[string]string {
					"identifier": "guage",
					"value": "99" },
				map[string]string {
					"identifier": "counter",
					"value": "5" }}}},
}

func TestToMap(t *testing.T) {
	for _, test := range ToMapTests {
		actual := test.input.ToMap()
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s: got\n\t%v\nexpected\n\t%v", test.name, actual, test.expected)
		}
	}
}

func TestPluginEquiv(t *testing.T) {
	first := Plugin{"multiple", []Metric{
			Metric { "guage", "99", map[string]string {}},
			Metric { "counter", "5", map[string]string {}}}} 
	second := Plugin{"multiple", []Metric{
			Metric { "guage", "99", map[string]string {}},
			Metric { "counter", "5", map[string]string {}}}} 
	if !reflect.DeepEqual(first, second) {
		t.Errorf("Plugin Equivalance: got\n\t%v\nexpected\n\t%v", first, second)
	}
}

func TestMetricEquiv(t *testing.T) {
	first := Metric { "guage", "99", map[string]string {}}
	second := Metric { "guage", "99", map[string]string {}}
	if !reflect.DeepEqual(first, second) {
		t.Errorf("Metric Equivalance: got\n\t%v\nexpected\n\t%v", first, second)
	}
}