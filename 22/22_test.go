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
		name       string
		args       args
		want       map[Point]string
		wantMiddle Point
	}{
		{"example", args{"example.txt"}, map[Point]string{
			Point{1, 1}: "#",
			Point{3, 1}: "#",
			Point{2, 2}: "#",
			Point{1, 3}: "#",
			Point{3, 3}: "#",
		}, Point{2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGrid, gotMiddle := readInput(tt.args.filename)
			if !reflect.DeepEqual(gotGrid, tt.want) {
				t.Errorf("readInput() = %v, want %v", gotGrid, tt.want)
			}
			if !reflect.DeepEqual(gotMiddle, tt.wantMiddle) {
				t.Errorf("readInput() = %v, wantMiddle %v", gotMiddle, tt.wantMiddle)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		grid   map[Point]string
		middle Point
		bursts int
	}
	tests := []struct {
		name             string
		args             args
		wantInfectBursts int
	}{
		{"example", args{map[Point]string{
			Point{3, 1}: "#",
			Point{1, 2}: "#"},
			Point{2, 2}, 10000},
			5587},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInfectBursts := solve(tt.args.grid, tt.args.middle, tt.args.bursts); gotInfectBursts != tt.wantInfectBursts {
				t.Errorf("solve() = %v, want %v", gotInfectBursts, tt.wantInfectBursts)
			}
		})
	}
}
func Test_solve2(t *testing.T) {
	type args struct {
		grid   map[Point]string
		middle Point
		bursts int
	}
	tests := []struct {
		name             string
		args             args
		wantInfectBursts int
	}{
		{"example", args{map[Point]string{
			Point{3, 1}: "#",
			Point{1, 2}: "#"},
			Point{2, 2}, 100},
			26},
		{"example-long", args{map[Point]string{
			Point{3, 1}: "#",
			Point{1, 2}: "#"},
			Point{2, 2}, 10000000},
			2511944},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInfectBursts := solve2(tt.args.grid, tt.args.middle, tt.args.bursts); gotInfectBursts != tt.wantInfectBursts {
				t.Errorf("solve2() = %v, want %v", gotInfectBursts, tt.wantInfectBursts)
			}
		})
	}
}
