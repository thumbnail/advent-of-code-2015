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

var pattern = regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

type Reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
	state    *State
}

type State struct {
	distance      int
	points        int
	restRemaining int
	stamina       int
}

func parseInput(input string) []Reindeer {
	l := make([]Reindeer, 0)
	for _, line := range strings.Split(input, "\n") {
		if string(line) == "" {
			continue
		}
		r := pattern.FindStringSubmatch(line)

		duration := util.ParseInt(r[3])
		l = append(l, Reindeer{r[1], util.ParseInt(r[2]), duration, util.ParseInt(r[4]), &State{0, 0, 0, duration}})
	}
	return l
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	deer := parseInput(input)
	for tick := 0; tick < 2503; tick++ {
		advance(&deer)
	}

	if part == 1 {
		var winner = deer[0]
		for _, deer := range deer {
			if deer.state.distance > winner.state.distance {
				winner = deer
			}
		}
		println("Answer:", winner.state.distance)
	} else {
		var winner = deer[0]
		for _, deer := range deer {
			if deer.state.points > winner.state.points {
				winner = deer
			}
		}
		println("Answer:", winner.state.points)

	}
}

func advance(reindeer *[]Reindeer) {

	var leaderDistance = 0
	var leaders []string

	for _, deer := range *reindeer {
		state := deer.state
		if state.stamina > 0 {
			state.distance += deer.speed
			state.restRemaining = deer.rest
			state.stamina--
		} else {
			state.restRemaining--
			if state.restRemaining == 0 {
				state.stamina = deer.duration
			}
		}

		if state.distance > leaderDistance {
			leaderDistance = state.distance
			leaders = []string{deer.name}
		} else {
			if state.distance == leaderDistance {
				leaders = append(leaders, deer.name)
			}
		}
	}

	for _, deer := range *reindeer {
		for _, leader := range leaders {
			if deer.name == leader {
				deer.state.points++
			}
		}
	}
}
