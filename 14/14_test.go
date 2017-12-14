package main

import "testing"

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
