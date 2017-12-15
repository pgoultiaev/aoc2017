package main

import (
	"testing"
)

func Test_compareLast16bits(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1092455", 1092455, 430625591, 0},
		{"1181022009", 1181022009, 1233683848, 0},
		{"245556042", 245556042, 1431495498, 1},
		{"1744312007", 1744312007, 137874439, 0},
		{"1352636452", 1352636452, 285222916, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareLast16bits(tt.a, tt.b); got != tt.want {
				t.Errorf("compareLast16bits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		a       int
		b       int
		factorA int
		factorB int
		modAB   int
	}
	tests := []struct {
		name           string
		args           args
		wantFinalCount int
	}{
		{"example", args{65, 8921, 16807, 48271, 2147483647}, 588},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFinalCount := solve(tt.args.a, tt.args.b, tt.args.factorA, tt.args.factorB, tt.args.modAB); gotFinalCount != tt.wantFinalCount {
				t.Errorf("solve() = %v, want %v", gotFinalCount, tt.wantFinalCount)
			}
		})
	}
}

func Test_generatePartTwo(t *testing.T) {
	type args struct {
		a       int
		b       int
		factorA int
		factorB int
		modAB   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"example1", args{65, 8921, 16807, 48271, 2147483647}, 1352636452, 1233683848},
		{"example2", args{1352636452, 1233683848, 16807, 48271, 2147483647}, 1992081072, 862516352},
		{"example3", args{1992081072, 862516352, 16807, 48271, 2147483647}, 530830436, 1159784568},
		{"example4", args{530830436, 1159784568, 16807, 48271, 2147483647}, 1980017072, 1616057672},
		{"example5", args{1980017072, 1616057672, 16807, 48271, 2147483647}, 740335192, 412269392},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generatePartTwo(tt.args.a, tt.args.b, tt.args.factorA, tt.args.factorB, tt.args.modAB)
			if got != tt.want {
				t.Errorf("generatePartTwo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("generatePartTwo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		a       int
		b       int
		factorA int
		factorB int
		modAB   int
	}
	tests := []struct {
		name           string
		args           args
		wantFinalCount int
	}{
		{"example", args{65, 8921, 16807, 48271, 2147483647}, 309},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFinalCount := solve2(tt.args.a, tt.args.b, tt.args.factorA, tt.args.factorB, tt.args.modAB); gotFinalCount != tt.wantFinalCount {
				t.Errorf("solve2() = %v, want %v", gotFinalCount, tt.wantFinalCount)
			}
		})
	}
}
