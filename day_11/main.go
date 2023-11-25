package main

import (
	"flag"
	"fmt"
	"strings"
)

var input = "hxbxwxba"

/*
Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz.
They cannot skip letters; abd doesn't count.
*/
func condition1(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i]+1 == input[i+1] &&
			input[i+1]+1 == input[i+2] {
			return true
		}
	}
	return false
}

/*
Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
*/
func condition2(input string) bool {
	return !strings.ContainsAny(input, "iol")
}

/*
Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
*/
func condition3(input string) bool {
	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			count++
			i++ // skip next character to prevent overlapping
		}
	}
	return count >= 2
}

func nextPassword(input string) string {
	output := []rune(input)
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 'z' {
			output[i] = 'a'
		} else {
			output[i] = rune(input[i] + 1)
			break
		}
	}
	return string(output)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	newPass := nextPassword(input)
	for !(condition1(newPass) && condition2(newPass) && condition3(newPass)) {
		newPass = nextPassword(newPass)
	}

	if part == 2 {
		newPass = nextPassword(newPass)
		for !(condition1(newPass) && condition2(newPass) && condition3(newPass)) {
			newPass = nextPassword(newPass)
		}
	}

	println("Answer:", newPass)
}
