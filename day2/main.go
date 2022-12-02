package main

import (
	"advent/util"
	"fmt"
	"strings"
)

var (
	Input        string
	ScoresByMove = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
		"A": 1,
		"B": 2,
		"C": 3,
	}
	EquivalentsByMove = map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	LosersByMove = map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
		"A": "C",
		"B": "A",
		"C": "B",
	}
)

func main() {
	Input = util.ReadFile("./input.txt")
	fmt.Println(part1())
	fmt.Println(part2())
}

/*
	0 if you lost
	3 if the round was a draw
	6 if you won
*/
func part1() (score int) {
	for _, str := range strings.Split(Input, "\n") {
		moves := strings.Split(str, " ")
		score += getScoreForPart1(moves[0], moves[1])
	}
	return
}

/*
	X means you need to lose
	Y means you need to end the round in a draw
	Z means you need to win
*/
func part2() (score int) {
	for _, str := range strings.Split(Input, "\n") {
		moves := strings.Split(str, " ")
		score += getScoreForPart2(moves[0], moves[1])
	}
	return
}

func getScoreForPart1(oppMove, myMove string) (score int) {
	score = ScoresByMove[myMove]

	if loser := LosersByMove[myMove]; loser == oppMove {
		score += 6
	} else if equivalent := EquivalentsByMove[myMove]; equivalent == oppMove {
		score += 3
	}
	return
}

/*
	Rock:			A
	Paper:		B
	Scissors:	C

	Got -> Want	== IntGot -> IntWant
	A 	-> B 		== 0 			-> 1
	B 	-> C 		== 1 			-> 2
	C 	-> A 		== 2 			-> 0

	A => 0 + 1 % 3 = 1 => B
	B => 1 + 1 % 3 = 2 => C
	C => 2 + 1 % 3 = 0 => A
*/
func getScoreForPart2(oppMove, strategy string) (score int) {
	switch strategy {
	case "X":
		return ScoresByMove[LosersByMove[oppMove]]
	case "Y":
		return 3 + ScoresByMove[oppMove]
	case "Z":
		return 6 + ScoresByMove[string((((oppMove[0]-65)+1)%3)+65)] // magic
	}
	return
}
