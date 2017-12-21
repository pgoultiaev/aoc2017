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
		name string
		args args
		want map[string]string
	}{
		{"example", args{"example.txt"}, map[string]string{
			"0001":      "110100000",
			"010001111": "1001000000001001",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInput(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ruleToPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"example", args{"0001"}, [][]int{[]int{0, 0}, []int{0, 1}}},
		{"example", args{"010001111"}, [][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}}},
		{"example", args{"1001000010010110"}, [][]int{[]int{1, 0, 0, 1}, []int{0, 0, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ruleToPattern(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ruleToPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		pattern [][]int
	}
	tests := []struct {
		name         string
		args         args
		wantPatternT [][]int
	}{
		{"example", args{[][]int{
			[]int{0, 1, 0},
			[]int{0, 0, 1},
			[]int{1, 1, 1},
		}}, [][]int{
			[]int{1, 0, 0},
			[]int{1, 0, 1},
			[]int{1, 1, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPatternT := rotate(tt.args.pattern); !reflect.DeepEqual(gotPatternT, tt.wantPatternT) {
				t.Errorf("rotate() = %v, want %v", gotPatternT, tt.wantPatternT)
			}
		})
	}
}

func Test_flip(t *testing.T) {
	type args struct {
		pattern [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"example", args{[][]int{
			[]int{0, 1, 0},
			[]int{0, 0, 1},
			[]int{1, 1, 1},
		}}, [][]int{
			[]int{0, 1, 0},
			[]int{1, 0, 0},
			[]int{1, 1, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flip(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		filename   string
		iterations int
	}
	tests := []struct {
		name     string
		args     args
		wantOnes int
	}{
		{"example", args{"example.txt", 2}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOnes := solve(tt.args.filename, tt.args.iterations); gotOnes != tt.wantOnes {
				t.Errorf("solve() = %v, want %v", gotOnes, tt.wantOnes)
			}
		})
	}
}

func Test_extract(t *testing.T) {
	type args struct {
		grid   [][]int
		row    int
		column int
		size   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"example", args{[][]int{
			[]int{1, 0, 0, 1},
			[]int{0, 0, 0, 0},
			[]int{0, 0, 0, 0},
			[]int{1, 0, 0, 1},
		}, 0, 0, 2}, [][]int{
			[]int{1, 0},
			[]int{0, 0},
		}},
		{"example-2", args{[][]int{
			[]int{1, 0, 0, 1},
			[]int{0, 0, 0, 0},
			[]int{0, 0, 0, 0},
			[]int{1, 0, 0, 1},
		}, 0, 2, 2}, [][]int{
			[]int{0, 1},
			[]int{0, 0},
		}},
		{"example-3", args{[][]int{
			[]int{1, 0, 0, 1},
			[]int{0, 0, 0, 0},
			[]int{0, 0, 0, 0},
			[]int{1, 0, 0, 1},
		}, 2, 2, 2}, [][]int{
			[]int{0, 0},
			[]int{0, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extract(tt.args.grid, tt.args.row, tt.args.column, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combine(t *testing.T) {
	type args struct {
		gridOfGrids [][][]int
		size        int
		count       int
	}
	tests := []struct {
		name     string
		args     args
		wantGrid [][]int
	}{
		{"example", args{[][][]int{
			[][]int{[]int{1, 0}, []int{0, 0}},
			[][]int{[]int{0, 1}, []int{0, 0}},
			[][]int{[]int{0, 0}, []int{1, 0}},
			[][]int{[]int{0, 0}, []int{0, 1}},
		}, 2, 2}, [][]int{
			[]int{1, 0, 0, 1},
			[]int{0, 0, 0, 0},
			[]int{0, 0, 0, 0},
			[]int{1, 0, 0, 1},
		}},
		{"example-3", args{[][][]int{
			[][]int{[]int{1, 1, 0}, []int{1, 0, 0}, []int{0, 0, 0}},
			[][]int{[]int{1, 1, 0}, []int{1, 0, 0}, []int{0, 0, 0}},
			[][]int{[]int{1, 1, 0}, []int{1, 0, 0}, []int{0, 0, 0}},
			[][]int{[]int{1, 1, 0}, []int{1, 0, 0}, []int{0, 0, 0}},
		}, 3, 2}, [][]int{
			[]int{1, 1, 0, 1, 1, 0},
			[]int{1, 0, 0, 1, 0, 0},
			[]int{0, 0, 0, 0, 0, 0},
			[]int{1, 1, 0, 1, 1, 0},
			[]int{1, 0, 0, 1, 0, 0},
			[]int{0, 0, 0, 0, 0, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGrid := combine(tt.args.gridOfGrids, tt.args.size, tt.args.count); !reflect.DeepEqual(gotGrid, tt.wantGrid) {
				t.Errorf("combine() = %v, want %v", gotGrid, tt.wantGrid)
			}
		})
	}
}
