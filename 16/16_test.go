package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		dance      []string
		dancemoves []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example", args{[]string{"a", "b", "c", "d", "e"}, []string{"s1", "x3/4", "pe/b"}}, "baedc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.dance, tt.args.dancemoves); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
