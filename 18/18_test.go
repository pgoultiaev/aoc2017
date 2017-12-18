package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name               string
		instructions       [][]string
		wantLastNonNilFreq int
	}{
		{"example", [][]string{
			[]string{"set", "a", "1"},
			[]string{"add", "a", "2"},
			[]string{"mul", "a", "a"},
			[]string{"mod", "a", "5"},
			[]string{"snd", "a"},
			[]string{"set", "a", "0"},
			[]string{"rcv", "a"},
			[]string{"jgz", "a", "-1"},
			[]string{"set", "a", "1"},
			[]string{"jgz", "a", "-2"}},
			4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLastNonNilFreq := solve(tt.instructions); gotLastNonNilFreq != tt.wantLastNonNilFreq {
				t.Errorf("solve() = %v, want %v", gotLastNonNilFreq, tt.wantLastNonNilFreq)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	tests := []struct {
		name         string
		instructions [][]string
		want         int
	}{
		{"example", [][]string{
			[]string{"snd", "1"},
			[]string{"snd", "2"},
			[]string{"snd", "p"},
			[]string{"rcv", "a"},
			[]string{"rcv", "b"},
			[]string{"rcv", "c"},
			[]string{"rcv", "d"}},
			3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := partTwo(tt.instructions); got != tt.want {
				t.Errorf("solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}
