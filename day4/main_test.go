package main

import (
	"testing"
)

func init() {
	Input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name    string
		wantSum int
	}{
		{
			name:    "1",
			wantSum: 2,
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
			name:    "2",
			wantSum: 4,
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
