package main

import "testing"

func Test_knotHash(t *testing.T) {
	tests := []struct {
		name string
		ia   []int
		want string
	}{
		{name: "255", ia: []int{255}, want: "ff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knotHash(tt.ia); got != tt.want {
				t.Errorf("knotHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
