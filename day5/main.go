package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Move struct {
	Count int
	Start int
	End   int
}

var (
	Input     string
	boxEntry  = regexp.MustCompile(`\[(.)\]`)
	moveEntry = regexp.MustCompile(`move (\d+) from (\d) to (\d)`)
)

func main() {
	Input = util.ReadFile("./input.txt")
	parseInput()
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() (result string) {
	var (
		maxRows = 9
	)
	rows, moves := parseInput()

	for i, row := range rows {
		if len(row) == 0 {
			maxRows = i
			break
		}
	}

	for _, move := range moves {
		for i := 0; i < move.Count; i++ {
			var (
				crate string
			)
			crate, rows[move.Start] = rows[move.Start][len(rows[move.Start])-1], rows[move.Start][:len(rows[move.Start])-1]
			rows[move.End] = append(rows[move.End], crate)
		}
	}
	for _, row := range rows[:maxRows] {
		result = result + row[len(row)-1]
	}

	return
}

func part2() (result string) {
	var (
		maxRows = 9
	)
	rows, moves := parseInput()

	for i, row := range rows {
		if len(row) == 0 {
			maxRows = i
			break
		}
	}

	for _, move := range moves {
		crates := rows[move.Start][len(rows[move.Start])-move.Count:]
		rows[move.Start] = rows[move.Start][:len(rows[move.Start])-move.Count]
		rows[move.End] = append(rows[move.End], crates...)
	}
	for _, row := range rows[:maxRows] {
		result = result + row[len(row)-1]
	}

	return
}

func parseInput() ([][]string, []Move) {
	var (
		inputSplit = strings.Split(Input, "\n")
		rows       = make([][]string, 9)
		foundMoves = false
		moves      = []Move{}
	)

	for _, s := range inputSplit {
		if s == " 1   2   3   4   5   6   7   8   9 " || s == "" {
			if s == "" {
				foundMoves = true
			}
			continue
		}

		if !foundMoves {
			for i := 0; i < len(s); i += 4 {
				testStr := s[i : i+3]
				if boxEntry.MatchString(testStr) {
					match := boxEntry.FindStringSubmatch(testStr)
					rows[i/4] = append(rows[i/4], match[1])
				}
			}
		} else {
			if moveEntry.MatchString(s) {
				var (
					matches = moveEntry.FindStringSubmatch(s)
					count   int
					start   int
					end     int
				)
				count, _ = strconv.Atoi(matches[1])
				start, _ = strconv.Atoi(matches[2])
				end, _ = strconv.Atoi(matches[3])

				moves = append(moves, Move{
					Count: count,
					Start: start - 1,
					End:   end - 1,
				})
			}
		}
	}
	for i, row := range rows {
		newRow := []string{}

		for j := len(row) - 1; j >= 0; j-- {
			newRow = append(newRow, row[j])
		}

		rows[i] = newRow
	}
	return rows, moves
}
