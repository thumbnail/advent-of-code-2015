package main

import (
	_ "embed"
	"flag"
	"strings"
)

//go:embed input.txt
var input string

type grid [][]bool

func parseInput(input string) grid {
	var grid grid
	for _, chars := range strings.Split(input, "\n") {
		if string(chars) == "" {
			continue
		}
		var row []bool
		for _, char := range chars {
			row = append(row, char == '#')
		}
		grid = append(grid, row)
	}
	return grid
}

func (g grid) String() string {
	var s string
	for x, row := range g {
		for y, _ := range row {
			if g.getLightState(x, y) {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (g grid) Count() int {
	return strings.Count(g.String(), "#")
}

func (g grid) getLightState(x, y int) bool {
	if x < 0 || y < 0 || x > len(g[0])-1 || y > len(g)-1 {
		return false
	}

	if part2 &&
		(x == 0 && y == 0 ||
			x == 0 && y == len(g)-1 ||
			x == len(g[0])-1 && y == 0 ||
			x == len(g[0])-1 && y == len(g)-1) {
		return true
	}

	return g[y][x]
}

func (g grid) nextState(x, y int) bool {
	var total = 0

	if g.getLightState(x-1, y-1) {
		total++
	}
	if g.getLightState(x, y-1) {
		total++
	}
	if g.getLightState(x+1, y-1) {
		total++
	}
	if g.getLightState(x-1, y) {
		total++
	}
	if g.getLightState(x-1, y+1) {
		total++
	}
	if g.getLightState(x, y+1) {
		total++
	}
	if g.getLightState(x+1, y+1) {
		total++
	}
	if g.getLightState(x+1, y) {
		total++
	}

	if g.getLightState(x, y) {
		return total == 2 || total == 3
	} else {
		return total == 3
	}
}

func (g grid) Step() grid {
	newGrid := make(grid, len(g))
	for y, line := range g {
		for x, _ := range line {
			if newGrid[y] == nil {
				newGrid[y] = make([]bool, len(line))
			}

			newGrid[y][x] = g.nextState(x, y)
		}
	}
	return newGrid
}

// cheating a little with the global flag 🤷🏻
var part2 bool

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	part2 = part == 2

	grid := parseInput(input)
	for i := 0; i < 100; i++ {
		grid = grid.Step()
	}

	println("Answer:", grid.Count())
}
