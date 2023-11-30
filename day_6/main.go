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

var pattern = regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)

func parseLine(line string) (int, int, int, int, string) {
	match := pattern.FindStringSubmatch(line)
	return util.ParseInt(match[2]),
		util.ParseInt(match[3]),
		util.ParseInt(match[4]),
		util.ParseInt(match[5]),
		match[1]
}

func step1(x int, y int, action string, lights map[[2]int]int) map[[2]int]int {
	switch action {
	case "turn on":
		lights[[2]int{x, y}] = 1
	case "turn off":
		lights[[2]int{x, y}] = 0
	case "toggle":
		lights[[2]int{x, y}] = 1 - lights[[2]int{x, y}]
	}
	return lights
}

func step2(x int, y int, action string, lights map[[2]int]int) map[[2]int]int {
	switch action {
	case "turn on":
		lights[[2]int{x, y}] += 1
	case "turn off":
		lights[[2]int{x, y}] = max(0, lights[[2]int{x, y}]-1)
	case "toggle":
		lights[[2]int{x, y}] += 2
	}
	return lights
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	var lights = make(map[[2]int]int)

	for _, line := range strings.Split(input, "\n") {
		if string(line) == "" {
			continue
		}

		x1, y1, x2, y2, action := parseLine(string(line))
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if part == 1 {
					lights = step1(x, y, action, lights)
				} else {
					lights = step2(x, y, action, lights)
				}
			}
		}
	}

	var total = 0
	for _, light := range lights {
		total += light
	}

	println("Answer:", total)
}
