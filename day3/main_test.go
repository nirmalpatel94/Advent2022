package main

import (
	"testing"
)

func init() {
	Input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name    string
		wantSum int
	}{
		{
			name:    "part1",
			wantSum: 157,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := part1(); gotSum != tt.wantSum {
				t.Errorf("part1() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name    string
		wantSum int
	}{
		{
			name:    "part2",
			wantSum: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := part2(); gotSum != tt.wantSum {
				t.Errorf("part2() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
