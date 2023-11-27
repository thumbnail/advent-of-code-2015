package main

import (
	"flag"
	"fmt"
	"strings"
)

var input = "3113322113"

func lookAndSay(input string) string {
	var count = 0
	var previous rune = 0
	var output strings.Builder
	for _, char := range input {
		if char == previous {
			count++
			continue
		}

		// not the start condition, not the same character anymore. reset loop
		if previous != 0 {
			output.WriteString(fmt.Sprintf("%d%s", count, string(previous)))
		}
		count = 1
		previous = char
	}

	// last character(s) need to be added
	output.WriteString(fmt.Sprintf("%d%s", count, string(previous)))
	return output.String()
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	cycles := 40
	if part == 2 {
		cycles = 50
	}

	output := input
	for i := 0; i < cycles; i++ {
		output = lookAndSay(output)
	}

	println("Answer:", len(output))
}
