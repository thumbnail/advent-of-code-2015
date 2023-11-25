package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	r := sum(`{"a":2,"b":4}`)
	want := 6
	if r != want {
		t.Fatalf(`sum("{"a":2,"b":4}") = %v, want match for %v`, r, want)
	}

	r = sum(`[1,2,3]`)
	want = 6
	if r != want {
		t.Fatalf(`sum("[1,2,3]") = %v, want match for %v`, r, want)
	}

	r = sum(`[[[3]]]`)
	want = 3
	if r != want {
		t.Fatalf(`sum("[[[3]]]") = %v, want match for %v`, r, want)
	}

	r = sum(`{"a":{"b":4},"c":-1}`)
	want = 3
	if r != want {
		t.Fatalf(`sum("{"a":{"b":4},"c":-1}") = %v, want match for %v`, r, want)
	}

	r = sum(`{"a":[-1,1]}`)
	want = 0
	if r != want {
		t.Fatalf(`sum("{"a":[-1,1]}") = %v, want match for %v`, r, want)
	}

	r = sum(`[-1,{"a":1}]`)
	want = 0
	if r != want {
		t.Fatalf(`sum("[-1,{"a":1}]") = %v, want match for %v`, r, want)
	}
	r = sum(`{}`)
	want = 0
	if r != want {
		t.Fatalf(`sum("{}") = %v, want match for %v`, r, want)
	}

	r = sum(`[]`)
	want = 0
	if r != want {
		t.Fatalf(`sum("[]") = %v, want match for %v`, r, want)
	}

}
