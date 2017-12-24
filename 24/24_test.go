package main

import (
	"reflect"
	"testing"
)

func Test_readInput(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name           string
		args           args
		wantComponents []Component
	}{
		{"example", args{"example.txt"}, []Component{
			Component{0, 2, false},
			Component{2, 2, false},
			Component{2, 3, false},
			Component{3, 4, false},
			Component{3, 5, false},
			Component{0, 1, false},
			Component{10, 1, false},
			Component{9, 10, false},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotComponents := readInput(tt.args.filename); !reflect.DeepEqual(gotComponents, tt.wantComponents) {
				t.Errorf("readInput() = %v, want %v", gotComponents, tt.wantComponents)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{"example.txt"}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.filename); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{"example.txt"}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve2(tt.args.filename); got != tt.want {
				t.Errorf("solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}
