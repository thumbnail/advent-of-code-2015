package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	h := hash("Input", 1)
	want := "0214fc5959f10ec946469e065afa7ae0"
	if h != want {
		t.Fatalf(`hash("Input", 1) = %s, want match for %s`, h, want)
	}
}

func TestTask(t *testing.T) {
	result := task("00000", "abcdef")
	want := 609043
	if result != want {
		t.Fatalf(`task("00000", "abcdef") = %d, want match for %d`, result, want)
	}
}

func TestTask2(t *testing.T) {
	result := task("00000", "pqrstuv")
	want := 1048970
	if result != want {
		t.Fatalf(`task("00000", "pqrstuv") = %d, want match for %d`, result, want)
	}
}

func TestTask3(t *testing.T) {
	result := task("000000", "abcdef")
	want := 6742839
	if result != want {
		t.Fatalf(`task("000000", "abcdef") = %d, want match for %d`, result, want)
	}
}
