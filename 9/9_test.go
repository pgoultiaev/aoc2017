package main

import "testing"

func Test_findSumGroups(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"unbalanced a", "{{}{}", 2},
		{"unbalanced b", "{}{}}", 2},
		{"unbalanced c", "{{{}{}}", 5},
		{"1", "{}", 1},
		{"2", "{{{}}}", 6},
		{"3", "{{},{}}", 5},
		{"4a", "{}{{}}", 4},
		{"4", "{{{},{},{{}}}}", 16},
		{"5", "{<a>,<a>,<a>,<a>}", 1},
		{"6", "{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"7", "{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"8", "{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumGroups(tt.text); got != tt.want {
				t.Errorf("findSumGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
