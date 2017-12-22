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
		want       map[Point]bool
		wantMiddle Point
	}{
		{"example", args{"example.txt"}, map[Point]bool{
			Point{1, 1}: true,
			Point{3, 1}: true,
			Point{2, 2}: true,
			Point{1, 3}: true,
			Point{3, 3}: true,
		}, Point{1, 1}},
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
		grid   map[Point]bool
		middle Point
		bursts int
	}
	tests := []struct {
		name             string
		args             args
		wantInfectBursts int
	}{
		{"example", args{map[Point]bool{
			Point{2, 0}: true,
			Point{0, 1}: true},
			Point{1, 1}, 10000},
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
