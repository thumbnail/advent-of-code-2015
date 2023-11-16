package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// It contains at least three vowels (aeiou only),
func condition1(s string) bool {
	return (strings.Count(s, "a") +
		strings.Count(s, "e") +
		strings.Count(s, "i") +
		strings.Count(s, "o") +
		strings.Count(s, "u")) >= 3
}

func condition2(s string) bool {
	seqs := [26]string{
		"aa",
		"bb",
		"cc",
		"dd",
		"ee",
		"ff",
		"gg",
		"hh",
		"ii",
		"jj",
		"kk",
		"ll",
		"mm",
		"nn",
		"oo",
		"pp",
		"qq",
		"rr",
		"ss",
		"tt",
		"uu",
		"vv",
		"ww",
		"xx",
		"yy",
		"zz",
	}
	for _, seq := range seqs {
		if strings.Contains(s, string(seq)) {
			return true
		}
	}
	return false
}

func condition3(s string) bool {
	seqs := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
	for _, seq := range seqs {
		if strings.Contains(s, string(seq)) {
			return false
		}
	}
	return true
}

func condition4(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		for j := i + 2; j < len(s)-1; j++ {
			if s[i] == s[j] && s[i+1] == s[j+1] {
				return true
			}
		}
	}
	return false
}

func condition5(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var total = 0
	for _, line := range strings.Split(input, "\n") {
		if part == 1 {
			if condition1(line) && condition2(line) && condition3(line) {
				total++
			}
		} else {
			if condition4(line) && condition5(line) {
				total++
			}
		}
	}

	fmt.Println("Answer: ", total)

}
