package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Cookie struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseInput(input string) []Ingredient {
	l := make([]Ingredient, 0)
	for _, line := range strings.Split(input, "\n") {
		if string(line) == "" {
			continue
		}

		var name string
		var capacity, durability, flavor, texture, calories int
		_, err := fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)

		if err != nil {
			panic(err)
		}

		l = append(l, Ingredient{
			name[:len(name)-1],
			capacity,
			durability,
			flavor,
			texture,
			calories,
		})
	}
	return l
}

func bake(ingredients map[Ingredient]int) Cookie {
	cookie := Cookie{}
	for ingredient, teaspoons := range ingredients {
		cookie.flavor += ingredient.flavor * teaspoons
		cookie.durability += ingredient.durability * teaspoons
		cookie.capacity += ingredient.capacity * teaspoons
		cookie.texture += ingredient.texture * teaspoons
		cookie.calories += ingredient.calories * teaspoons
	}
	return cookie
}

func (c Cookie) score() int {
	if c.capacity < 0 || c.durability < 0 || c.flavor < 0 || c.texture < 0 {
		return 0
	}
	return c.capacity * c.durability * c.flavor * c.texture
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	ingredients := parseInput(input)

	winner := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			for k := 0; k < 100-i-j; k++ {
				var l = 100 - i - j - k
				recipe := map[Ingredient]int{
					ingredients[0]: i,
					ingredients[1]: j,
					ingredients[2]: k,
					ingredients[3]: l,
				}

				cookie := bake(recipe)

				if part == 2 {
					if bake(recipe).calories != 500 {
						continue
					}
				}

				score := cookie.score()
				if winner < score {
					winner = score
				}
			}
		}
	}

	println("Answer:", winner)
}
