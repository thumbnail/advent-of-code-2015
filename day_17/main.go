package main

import (
	_ "embed"
	"flag"
	"slices"
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

	buckets := parseInput(input)
	var path []int
	println("Answer:", len(calc(buckets, path, 150)))
}

func calc(buckets []int, path []int, liters int) (paths [][]int) {
	if liters == 0 {
		return [][]int{path}
	}

	for i, bucket := range buckets {
		if bucket <= liters && !slices.Contains(path, i) {
			path = append(path, i)
			paths = append(paths, calc(buckets, path, liters-bucket)...)
		}
	}

	return
}

func parseInput(s string) []int {
	var buckets []int
	for _, line := range strings.Split(s, "\n") {
		if string(line) == "" {
			continue
		}
		buckets = append(buckets, toInt(line))
	}
	return buckets
}

func toInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
