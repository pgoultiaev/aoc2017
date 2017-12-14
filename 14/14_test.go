package main

import (
	"testing"
)

func Test_getBinVal(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"0", "0", "0000"},
		{"1", "1", "0001"},
		{"e", "e", "1110"},
		{"f", "f", "1111"},
		{"a0c2017", "a0c2017", "1010000011000010000000010111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getBinVal(tt.s); got != tt.want {
				t.Errorf("getBinVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		grid [128][]string
	}
	tests := []struct {
		name           string
		grid           [128][]string
		wantNumRegions int
	}{
		{name: "example", grid: [128][]string{[]string{"1", "1", "0"}, []string{"0", "1", "0"}, []string{"0", "0", "1"}}, wantNumRegions: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumRegions := solve2(tt.grid); gotNumRegions != tt.wantNumRegions {
				t.Errorf("solve2() = %v, want %v", gotNumRegions, tt.wantNumRegions)
			}
		})
	}
}
