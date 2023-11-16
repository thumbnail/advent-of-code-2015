package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"flag"
	"fmt"
	"strings"
)

var input = "yzbqklnj"

func hash(s string, i int) string {
	var hash = md5.Sum([]byte(fmt.Sprintf("%s%d", s, i)))
	return hex.EncodeToString(hash[:])
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var prefix = "00000"
	if part == 2 {
		prefix = "000000"
	}

	var h = ""
	var i = 0
	for !strings.HasPrefix(h, prefix) {
		i++
		h = hash(input, i)
		if i%100000 == 0 {
			fmt.Println("attempt", i)
		}
	}
	fmt.Println("Answer: ", i)
}
