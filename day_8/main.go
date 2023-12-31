package main

import (
	_ "embed"
	"flag"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	var codeChars, memChars int

	if part == 1 {
		for _, line := range strings.Split(input, "\n") {
			codeChars += len(line)
			x, _ := strconv.Unquote(line)
			memChars += len(x)
		}
		println("Answer:", codeChars-memChars)
	} else {
		for _, line := range strings.Split(input, "\n") {
			if line == "" {
				continue
			}

			codeChars += len(line)
			x := strconv.Quote(line)
			memChars += len(x)
		}
		println("Answer:", memChars-codeChars)
	}

}
