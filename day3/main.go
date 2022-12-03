package main

import (
	"advent/util"
	"fmt"
	"strings"
	"unicode"
)

var (
	Input = util.ReadFile("./input.txt")
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

/*
	a == 1
	A == 27

*/
func part1() (sum int) {
	var (
		inputSplit = strings.Split(Input, "\n")
	)

	for _, s := range inputSplit {
		var (
			set        = map[rune]struct{}{}
			firstHalf  = s[:len(s)/2]
			secondHalf = s[len(s)/2:]
		)

		for _, char := range firstHalf {
			set[char] = struct{}{}
		}

		for _, char := range secondHalf {
			// char is already in set
			if _, ok := set[char]; ok {
				sum += getScore(char)
				break
			}
		}
	}
	return
}

func part2() (sum int) {
	var (
		inputSplit = strings.Split(Input, "\n")
	)

	for i := 0; i < len(inputSplit); i += 3 {
		var (
			first  = map[rune]struct{}{}
			second = map[rune]struct{}{}
			third  = map[rune]struct{}{}
			badge  = map[rune]int{}
		)

		// generate sets
		for _, char := range inputSplit[i] {
			first[char] = struct{}{}
		}

		for _, char := range inputSplit[i+1] {
			second[char] = struct{}{}
		}

		for _, char := range inputSplit[i+2] {
			third[char] = struct{}{}
		}

		for char := range first {
			badge[char] = 1
		}

		for char := range second {
			if count, ok := badge[char]; ok && count != 2 {
				badge[char] = 2
			}
		}

		for char := range third {
			if count, ok := badge[char]; ok && count == 2 && count != 3 {
				badge[char] = 3
			}
		}
		for char, val := range badge {
			if val == 3 {
				sum += getScore(char)
			}
		}

	}
	return
}

func getScore(char rune) int {
	// A => 65 - 38 = 27
	// a => 97 - 96 = 1
	if unicode.IsUpper(char) {
		return int(char) - 38
	}
	return int(char) - 96
}
