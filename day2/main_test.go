package main

import (
	"testing"
)

func init() {
	Input =
		`A Y
B X
C Z`
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "part1",
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name      string
		wantScore int
	}{
		{
			name:      "part2",
			wantScore: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := part2(); gotScore != tt.wantScore {
				t.Errorf("part2() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
