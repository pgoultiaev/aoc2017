package main

import "testing"

var (
	testA = stateInstruction{1, 0, 1, -1, "B", "B"}
	testB = stateInstruction{1, 1, -1, 1, "A", "A"}

	testStates = map[string]stateInstruction{"A": testA, "B": testB}
)

func Test_solve(t *testing.T) {
	type args struct {
		diagnoseAfter int
		states        map[string]stateInstruction
	}
	tests := []struct {
		name         string
		args         args
		wantChecksum int
	}{
		{"example", args{6, testStates}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChecksum := solve(tt.args.diagnoseAfter, tt.args.states); gotChecksum != tt.wantChecksum {
				t.Errorf("solve() = %v, want %v", gotChecksum, tt.wantChecksum)
			}
		})
	}
}
