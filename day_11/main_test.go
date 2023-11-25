package main

import (
	"testing"
)

func TestCondition1(t *testing.T) {
	r := condition1("hijklmmn")
	want := true
	if r != want {
		t.Fatalf(`condition1("hijklmmn") = %v, want match for %v`, r, want)
	}

	r = condition1("aabc")
	want = true
	if r != want {
		t.Fatalf(`condition1("aabc") = %v, want match for %v`, r, want)
	}

	r = condition1("abbceffg")
	want = false
	if r != want {
		t.Fatalf(`condition1("abbceffg") = %v, want match for %v`, r, want)
	}
}

func TestCondition2(t *testing.T) {
	r := condition2("hijklmmn")
	want := false
	if r != want {
		t.Fatalf(`condition2("hijklmmn") = %v, want match for %v`, r, want)
	}

	r = condition2("abcdefg")
	want = true
	if r != want {
		t.Fatalf(`condition2("abcdefg") = %v, want match for %v`, r, want)
	}
}

func TestCondition3(t *testing.T) {
	r := condition3("abbceffg")
	want := true
	if r != want {
		t.Fatalf(`condition3("abbceffg") = %v, want match for %v`, r, want)
	}

	r = condition3("abbcegjk")
	want = false
	if r != want {
		t.Fatalf(`condition3("abbcegjk") = %v, want match for %v`, r, want)
	}
}

func TestNextPassword(t *testing.T) {
	r := nextPassword("aaaa")
	want := "aaab"
	if r != want {
		t.Fatalf(`nextPassword("aaaa") = %v, want match for %v`, r, want)
	}

	r = nextPassword("aaaz")
	want = "aaba"
	if r != want {
		t.Fatalf(`nextPassword("aaaz") = %v, want match for %v`, r, want)
	}

	// rollover doesn't add a new character. I guess that's fine
	r = nextPassword("zzzz")
	want = "aaaa"
	if r != want {
		t.Fatalf(`nextPassword("zzzz") = %v, want match for %v`, r, want)
	}
}
