package main

import (
	_ "embed"
	"flag"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	if part == 1 {
		// cheeky way to calculate the answer, for the next part we'll actually need to parse the input
		var up = strings.Count(input, "(")
		var down = strings.Count(input, ")")

		println("Answer:", up-down)
	} else {
		var floor = 0
		for i, c := range input {
			if string(c) == "(" {
				floor++
			} else {
				floor--
			}

			if floor == -1 {
				println("Answer:", i+1)
				return
			}
		}
	}
}
