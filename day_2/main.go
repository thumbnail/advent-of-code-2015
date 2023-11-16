package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var pattern = regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)

func parseLine(line string) (int, int, int) {
	var match = pattern.FindStringSubmatch(line)
	if len(match) != 4 {
		return 0, 0, 0
	}

	var l, _ = strconv.Atoi(match[1])
	var w, _ = strconv.Atoi(match[2])
	var h, _ = strconv.Atoi(match[3])
	return l, w, h
}

func part1() {
	lines := strings.Split(input, "\n")
	var total = 0
	for _, line := range lines {
		l, w, h := parseLine(line)

		var side1 = (l * w)
		var side2 = (w * h)
		var side3 = (h * l)

		total += 2*side1 + 2*side2 + 2*side3 + min(side1, side2, side3)
	}
	fmt.Println("Answer: ", total)
}

func part2() {
	lines := strings.Split(input, "\n")
	var total = 0
	for _, line := range lines {
		l, w, h := parseLine(line)

		// find smallest sides
		xs := []int{l, w, h}
		slices.Sort(xs)

		total += l*w*h + 2*xs[0] + 2*xs[1]
	}
	fmt.Println("Answer: ", total)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		part1()
	} else {
		part2()
	}
}
