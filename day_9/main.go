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

var pattern = regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)

type Route struct {
	path []string
	cost int
}

func calc(possibilities map[string]map[string]int, path []string, cost int) []Route {
	current := path[len(path)-1]

	var routes []Route

	for k, v := range possibilities[current] {
		if !slices.Contains(path, k) {
			routes = append(routes, calc(possibilities, append(path, k), cost+v)...)
		}
	}

	// routes is empty when no possibilities were found, which means we visited every location
	if len(routes) == 0 {
		routes = append(routes, Route{path, cost})
	}

	return routes
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	routes := make(map[string]map[string]int)
	for _, line := range strings.Split(input, "\n") {
		if string(line) == "" {
			continue
		}
		r := pattern.FindStringSubmatch(line)
		from := r[1]
		to := r[2]
		distance, _ := strconv.Atoi(r[3])
		_, found := routes[from]

		if !found {
			routes[from] = make(map[string]int)
		}
		_, found2 := routes[to]
		if !found2 {
			routes[to] = make(map[string]int)
		}

		routes[from][to] = distance
		routes[to][from] = distance
	}

	shortest := Route{}
	longest := Route{}

	var paths []Route
	for k, _ := range routes {
		path := []string{k}
		paths = append(paths, calc(routes, path, 0)...)
	}

	for _, path := range paths {
		if path.cost < shortest.cost || shortest.cost == 0 {
			shortest = path
		}
		if path.cost > longest.cost {
			longest = path
		}
	}

	if part == 1 {
		fmt.Printf("Answer: %v\n", shortest)
	} else {
		fmt.Printf("Answer: %v\n", longest)
	}
}
