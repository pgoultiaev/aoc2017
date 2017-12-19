package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		grid   map[Point]string
		origin Point
	}
	tests := []struct {
		name      string
		args      args
		wantTrip  string
		wantSteps int
	}{
		{"simple grid", args{map[Point]string{
			Point{1, 0}: "|",
			Point{1, 1}: "A",
			Point{1, 2}: "+",
			Point{0, 2}: "-"},
			Point{1, 0},
		}, "A", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTrip, gotSteps := solve(tt.args.grid, tt.args.origin)
			if gotTrip != tt.wantTrip {
				t.Errorf("solve() = %v, wantTrip %v", gotTrip, tt.wantTrip)
			}
			if gotSteps != tt.wantSteps {
				t.Errorf("solve() = %d, wantSteps %d", gotSteps, tt.wantSteps)
			}
		})
	}
}
