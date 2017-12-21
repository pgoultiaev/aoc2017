package main

import "testing"

func Test_findSumGroups(t *testing.T) {
	tests := []struct {
		name        string
		text        string
		wantSum     int
		wantGarbage int
	}{
		{"unbalanced a", "{{}{}", 2, 0},
		{"unbalanced b", "{}{}}", 2, 0},
		{"unbalanced c", "{{{}{}}", 5, 0},
		{"1", "{}", 1, 0},
		{"2", "{{{}}}", 6, 0},
		{"3", "{{},{}}", 5, 0},
		{"4a", "{}{{}}", 4, 0},
		{"4", "{{{},{},{{}}}}", 16, 0},
		{"5", "{<a>,<a>,<a>,<a>}", 1, 4},
		{"6", "{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
		{"7", "{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
		{"8", "{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gCount := findSumGroups(tt.text)
			if gotSum != tt.wantSum {
				t.Errorf("findSumGroups() = %v, want %v", gotSum, tt.wantSum)
			}
			if gCount != tt.wantGarbage {
				t.Errorf("findSumGroups() = %v, want %v", gCount, tt.wantGarbage)
			}
		})
	}
}
