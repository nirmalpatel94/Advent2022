package main

import (
	"advent/util"
	"fmt"
	"strings"
)

var (
	Input string
)

func main() {
	Input = util.ReadFile("./input.txt")
	fmt.Println(part1())
	fmt.Println(part2())
}

// find elf with most calories
func part1() int {
	var (
		elfCalories = getElfCalories()
		maxCount    int
	)

	for _, cal := range elfCalories {
		if cal > maxCount {
			maxCount = cal
		}
	}
	return maxCount
}

// find sum for top 3 elves
func part2() int {
	var (
		first       = 0
		second      = 0
		third       = 0
		elfCalories = getElfCalories()
	)

	for _, cal := range elfCalories {
		switch true {
		case cal >= first:
			third = second
			second = first
			first = cal
		case cal >= second:
			third = second
			second = cal
		case cal > third:
			third = cal
		}
	}
	return first + second + third
}

func getElfCalories() (elfCalories []int) {
	var (
		currCount int
	)

	for _, l := range strings.Split(Input, "\n") {
		if l == "" {
			elfCalories = append(elfCalories, currCount)
			currCount = 0
			continue
		}
		currCount += util.ToInt(strings.TrimSpace(l))
	}
	return
}
