package main

import (
	_ "embed"
	"flag"
	"math"
	"strings"

	"github.com/thumbnail/advent-of-code-2015/util"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	buckets := parseInput(input)
	var path []int
	paths := calc(buckets, 0, path, 150)

	answer := 0

	if part == 1 {
		answer = len(paths)
	} else {
		minLen := math.MaxInt32
		for _, path := range paths {
			if minLen > len(path) {
				minLen = len(path)
			}
		}

		for _, path := range paths {
			if len(path) == minLen {
				answer++
			}
		}

	}

	println("Answer:", answer)
}

func calc(buckets []int, start int, path []int, liters int) [][]int {
	if liters == 0 {
		return [][]int{path}
	}

	if liters < 0 {
		return nil
	}

	var paths [][]int
	for i := start; i < len(buckets); i++ {
		paths = append(paths, calc(buckets, i+1, append(path, i), liters-buckets[i])...)
	}

	return paths
}

func parseInput(s string) []int {
	var buckets []int
	for _, line := range strings.Split(s, "\n") {
		if string(line) == "" {
			continue
		}
		buckets = append(buckets, util.ParseInt(line))
	}
	return buckets
}
