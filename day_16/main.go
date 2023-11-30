package main

import (
	_ "embed"
	"flag"
	"regexp"
	"strings"

	"github.com/thumbnail/advent-of-code-2015/util"
)

//go:embed input.txt
var input string

var ticker = `children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`

var pattern = regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

func mismatch(ticker map[string]int, properties map[string]int, name string, fn func(a int, b int) bool) bool {
	v, ok := properties[name]
	if !ok {
		return false
	}

	return !fn(v, ticker[name])
}

func matcher(name string, part1 bool) func(a int, b int) bool {
	if part1 {
		return equals
	}

	switch name {
	case "cats":
	case "trees":
		return greaterThan
	case "pomeranians":
	case "goldfish":
		return lessThan
	}
	return equals
}

func matches(ticker map[string]int, properties map[string]int, part1 bool) bool {
	for name, _ := range ticker {
		matcher := matcher(name, part1)
		if mismatch(ticker, properties, name, matcher) {
			return false
		}
	}
	return true
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	ticker := parseTicker(ticker)
	sues := parseInput(input)

	giver := 0
	for sue, properties := range sues {
		if matches(ticker, properties, part == 1) {
			giver = sue
		}
	}

	println("Answer:", giver)
}

func parseInput(t string) map[int]map[string]int {
	sues := make(map[int]map[string]int)
	for _, line := range strings.Split(t, "\n") {
		if string(line) == "" {
			continue
		}

		matches := pattern.FindStringSubmatch(line)
		sue := util.ParseInt(matches[1])
		_, found := sues[sue]
		if !found {
			sues[sue] = make(map[string]int)
		}
		sues[sue][matches[2]] = util.ParseInt(matches[3])
		sues[sue][matches[4]] = util.ParseInt(matches[5])
		sues[sue][matches[6]] = util.ParseInt(matches[7])
	}
	return sues
}

func parseTicker(s string) map[string]int {
	ticker := make(map[string]int)
	for _, line := range strings.Split(s, "\n") {
		r := strings.Split(line, ": ")
		ticker[r[0]] = util.ParseInt(r[1])
	}
	return ticker
}

var equals = func(a, b int) bool {
	return a == b
}

var greaterThan = func(a, b int) bool {
	return a > b
}

var lessThan = func(a, b int) bool {
	return a < b
}
