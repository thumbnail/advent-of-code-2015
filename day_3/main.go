package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed input.txt
var input string

func step(c string, pos [2]int, houses map[[2]int]int) ([2]int, map[[2]int]int) {
	houses[pos]++
	switch c {
	case "^":
		pos[1]++
	case "v":
		pos[1]--
	case ">":
		pos[0]++
	case "<":
		pos[0]--
	}
	houses[pos]++

	return pos, houses
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var santaPos = [2]int{0, 0}
	var robotPos = [2]int{0, 0}
	var houses = make(map[[2]int]int)
	for i, c := range input {
		if i%2 == 0 || part == 1 {
			santaPos, houses = step(string(c), santaPos, houses)
		} else {
			robotPos, houses = step(string(c), robotPos, houses)
		}
	}

	var total = 0
	for _, house := range houses {
		if house >= 1 {
			total++
		}
	}

	fmt.Println("Answer: ", total)
}
