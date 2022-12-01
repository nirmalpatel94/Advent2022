package main

import (
	"testing"
)

func init() {
	Input = `
	1000
	2000
	3000

	4000

	5000
	6000

	7000
	8000
	9000

	10000
`
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "pt 1",
			want: 24000,
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
		name string
		want int
	}{
		{
			name: "pt 2",
			want: 45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
