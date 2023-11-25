package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

var pattern = regexp.MustCompile(`(-?\d+)`)

func sum(input string) int {
	result := pattern.FindAllString(input, -1)

	total := 0
	for _, match := range result {
		number, _ := strconv.Atoi(match)
		total += number
	}
	return total
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	answer := sum(input)

	println("Answer:", answer)
}
