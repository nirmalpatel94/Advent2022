package main

import "testing"

func init() {
	Input = `    [D]    
[N] [C]    
[Z] [M] [P]
	1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "1",
			want: "CMZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := part1(); gotSum != tt.want {
				t.Errorf("part1() = %v, want %v", gotSum, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "2",
			want: "MCD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := part2(); gotSum != tt.want {
				t.Errorf("part2() = %v, want %v", gotSum, tt.want)
			}
		})
	}
}
