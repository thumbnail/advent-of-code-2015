package main

import (
	"reflect"
	"testing"
)

func TestParseGraph(t *testing.T) {
	input := `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`
	r := parseGraph(input, false)
	want := map[string]map[string]int{
		"Alice": {
			"Bob":   137,
			"Carol": -141,
			"David": 44,
		},
		"Bob": {
			"Alice": 137,
			"Carol": 53,
			"David": -70,
		},
		"Carol": {
			"Alice": -141,
			"Bob":   53,
			"David": 96,
		},
		"David": {
			"Alice": 44,
			"Bob":   -70,
			"Carol": 96,
		},
	}

	if !reflect.DeepEqual(r, want) {
		t.Fatalf(`parseGraph(input, false) = %v, want match for %v`, r, want)
	}

	r = parseGraph(input, true)
	want = map[string]map[string]int{
		"Alice": {
			"Bob":   137,
			"Carol": -141,
			"David": 44,
			"me":    0,
		},
		"Bob": {
			"Alice": 137,
			"Carol": 53,
			"David": -70,
			"me":    0,
		},
		"Carol": {
			"Alice": -141,
			"Bob":   53,
			"David": 96,
			"me":    0,
		},
		"David": {
			"Alice": 44,
			"Bob":   -70,
			"Carol": 96,
			"me":    0,
		},
	}

	if !reflect.DeepEqual(r, want) {
		t.Fatalf(`parseGraph(input, true) = %v, want match for %v`, r, want)
	}
}
