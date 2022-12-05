package main

import (
	"advent/util"
	"fmt"
	"strconv"
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

func part1() (sum int) {
	var (
		inputSplit = strings.Split(Input, "\n")
	)

	// 2-8,3-7
	for _, str := range inputSplit {
		assignments := strings.Split(str, ",")
		elf0 := strings.Split(assignments[0], "-")
		elf1 := strings.Split(assignments[1], "-")
		e0 := getRange(elf0)
		e1 := getRange(elf1)

		if contains(e0, e1) || contains(e1, e0) {
			sum++
		}
	}
	return
}

func part2() (sum int) {
	var (
		inputSplit = strings.Split(Input, "\n")
	)

	// 2-8,3-7
	for _, str := range inputSplit {
		assignments := strings.Split(str, ",")
		elf0 := strings.Split(assignments[0], "-")
		elf1 := strings.Split(assignments[1], "-")
		e0 := getRange(elf0)
		e1 := getRange(elf1)

		if overlaps(e0, e1) || overlaps(e1, e0) {
			sum++
		}
	}
	return
}

// check if r2 is within r1
func contains(r1, r2 []int) bool {
	min := r1[0]
	max := r1[1]

	if r2[0] >= min && r2[1] <= max {
		return true
	}
	return false
}

/*
	2,8				3,7
	2,6				4,8
*/
func overlaps(r1, r2 []int) bool {
	min := r1[0]
	max := r1[1]

	if (r2[0] >= min && r2[0] <= max) || (r2[1] >= min && r2[1] <= max) {
		return true
	}
	return false
}

func getRange(strs []string) (vals []int) {
	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		vals = append(vals, i)
	}
	return
}
