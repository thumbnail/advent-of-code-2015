package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

var pattern = regexp.MustCompile(`(-?\d+)`)
var total = 0

func sum(input string) int {
	result := pattern.FindAllString(input, -1)

	total := 0
	for _, match := range result {
		number, _ := strconv.Atoi(match)
		total += number
	}
	return total
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	if part == 1 {
		total = sum(input)
	} else {
		total = parseInput(input)
	}

	println("Answer:", total)
}

func parseInput(input string) int {
	// Creating the maps for JSON
	a := []interface{}{}

	// Parsing/Unmarshalling JSON encoding/json
	err := json.Unmarshal([]byte(input), &a)

	if err != nil {
		panic(err)
	}
	return parseArray(a)
}

func parseMap(aMap map[string]interface{}) int {
	for _, val := range aMap {
		if val == "red" {
			return 0
		}
	}

	var total = 0
	for _, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			total += parseMap(val.(map[string]interface{}))
		case []interface{}:
			total += parseArray(val.([]interface{}))
		case float64:
			total += int(concreteVal)
		}
	}
	return total
}

func parseArray(anArray []interface{}) int {
	var total = 0
	for _, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			total += parseMap(val.(map[string]interface{}))
		case []interface{}:
			total += parseArray(val.([]interface{}))
		case float64:
			total += int(concreteVal)
		}
	}
	return total
}
