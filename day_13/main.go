package main

import (
	_ "embed"
	"flag"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseGraph(input string, included bool) map[string]map[string]int {
	pattern := regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).`)

	graph := make(map[string]map[string]int)
	for _, line := range strings.Split(input, "\n") {
		if string(line) == "" {
			continue
		}

		matches := pattern.FindStringSubmatch(line)
		_, found := graph[matches[1]]
		if !found {
			graph[matches[1]] = make(map[string]int)
			if included {
				graph[matches[1]]["me"] = 0
			}
		}

		_, found2 := graph[matches[4]]
		if !found2 {
			graph[matches[4]] = make(map[string]int)
			if included {
				graph[matches[4]]["me"] = 0
			}
		}

		distance, _ := strconv.Atoi(matches[3])
		if matches[2] == "lose" {
			distance = -distance
		}
		graph[matches[1]][matches[4]] += distance
		graph[matches[4]][matches[1]] += distance
	}

	return graph
}

type Seating struct {
	arrangement []string
	happiness   int
}

func calc(guests map[string]map[string]int, arrangement []string, happiness int) []Seating {
	current := arrangement[len(arrangement)-1]

	var options []Seating

	for k, v := range guests[current] {
		if !slices.Contains(arrangement, k) {
			options = append(options, calc(guests, append(arrangement, k), happiness+v)...)
		}
	}

	// options is empty when no guests were found, which means we visited every location
	if len(options) == 0 {
		arrangement = append(arrangement, arrangement[0])
		happiness += guests[current][arrangement[0]]
		options = append(options, Seating{arrangement, happiness})
	}

	return options
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	graph := parseGraph(input, part == 2)

	longest := Seating{}

	var paths []Seating
	for k, _ := range graph {
		path := []string{k}
		paths = append(paths, calc(graph, path, 0)...)
	}

	for _, path := range paths {
		if path.happiness > longest.happiness {
			longest = path
		}
	}

	println("Answer:", longest.happiness)
}
