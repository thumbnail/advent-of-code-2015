package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"flag"
	"fmt"
	"math"
	"strings"
)

var input = "yzbqklnj"

func hash(s string, i int) string {
	var hash = md5.Sum([]byte(fmt.Sprintf("%s%d", s, i)))
	return hex.EncodeToString(hash[:])
}

func findHash(prefix string, seed string, data chan int, result chan int) {
	for i := range data {
		h := hash(seed, i)

		if strings.HasPrefix(h, prefix) {
			result <- i
			break
		}
	}
}

func task(prefix string, seed string) int {
	channel := make(chan int, math.MaxInt32)
	defer close(channel)

	result := make(chan int, 1)

	for i := 0; i < 10; i++ {
		go findHash(prefix, seed, channel, result)
	}

	for i := 0; i < math.MaxInt32; i++ {
		select {
		case r := <-result:
			return r
		default:
			channel <- i
		}
	}

	panic("No solution")
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	println("Running part", part)

	var prefix = "00000"
	if part == 2 {
		prefix = "000000"
	}

	println("Answer:", task(prefix, input))
}
