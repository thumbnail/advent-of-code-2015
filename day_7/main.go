package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func resolve(wires map[string]int, identifier string) int {
	val, err := strconv.Atoi(identifier)
	if err == nil {
		return val
	}

	if val, ok := wires[identifier]; ok {
		return val
	}

	for _, line := range strings.Split(input, "\n") {
		if strings.HasSuffix(line, fmt.Sprintf(" -> %s", identifier)) {
			return process(wires, line)[identifier]
		}
	}
	panic("Could not resolve " + identifier)
}

var pattern = regexp.MustCompile(`(?:([a-z]+|\d+) )?(?:(NOT|OR|AND|RSHIFT|LSHIFT) )?([a-z]+|\d+) -> ([a-z]+)`)

func process(wires map[string]int, line string) map[string]int {
	parts := pattern.FindStringSubmatch(line)

	switch parts[2] {
	case "NOT":
		wires[parts[3]] = resolve(wires, parts[3])
		wires[parts[4]] = ^wires[parts[3]]
	case "OR":
		wires[parts[1]] = resolve(wires, parts[1])
		wires[parts[3]] = resolve(wires, parts[3])
		wires[parts[4]] = wires[parts[1]] | wires[parts[3]]
	case "LSHIFT":
		wires[parts[1]] = resolve(wires, parts[1])
		wires[parts[3]] = resolve(wires, parts[3])
		wires[parts[4]] = wires[parts[1]] << wires[parts[3]]
	case "RSHIFT":
		wires[parts[1]] = resolve(wires, parts[1])
		wires[parts[3]] = resolve(wires, parts[3])
		wires[parts[4]] = wires[parts[1]] >> wires[parts[3]]
	case "AND":
		wires[parts[1]] = resolve(wires, parts[1])
		wires[parts[3]] = resolve(wires, parts[3])
		wires[parts[4]] = wires[parts[1]] & wires[parts[3]]
	default:
		wires[parts[3]] = resolve(wires, parts[3])
		wires[parts[4]] = wires[parts[3]]
	}
	return wires
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	var wires = make(map[string]int)
	for _, line := range strings.Split(input, "\n") {
		if string(line) == "" {
			continue
		}

		wires = process(wires, line)
	}

	if part == 1 {
		println("Answer:", wires["a"])
		return
	}

	override := wires["a"]

	// reset wires
	wires = make(map[string]int)
	wires["b"] = override
	for _, line := range strings.Split(input, "\n") {
		if string(line) == "" {
			continue
		}

		wires = process(wires, line)
	}
	println("Answer:", wires["a"])
}
