package main

import (
	"advent/util"
	"fmt"
)

var (
	Input string
)

func main() {
	Input = util.ReadFile("./input.txt")
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() (sum int) {
	for i := 0; i < len(Input); i++ {
		chars := map[rune]struct{}{}

		for j := 0; j < 4 && i+j < len(Input); j++ {
			chars[rune(Input[i+j])] = struct{}{}
		}
		if len(chars) == 4 {
			return i + 4
		}
	}
	return
}

func part2() (sum int) {
	for i := 0; i < len(Input); i++ {
		chars := map[rune]struct{}{}

		for j := 0; j < 14 && i+j < len(Input); j++ {
			chars[rune(Input[i+j])] = struct{}{}
		}
		if len(chars) == 14 {
			return i + 14
		}
	}
	return
}
