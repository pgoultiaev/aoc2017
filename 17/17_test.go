package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		lastValWritten int
		input          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{2017, 3}, 638},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := solve(tt.args.lastValWritten, tt.args.input); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
