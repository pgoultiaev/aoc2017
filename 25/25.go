package main

type stateInstruction struct {
	writeVal0, writeVal1         int
	cursorChange0, cursorChange1 int
	nextState0, nextState1       string
}

var (
	A = stateInstruction{1, 0, 1, -1, "B", "D"}
	B = stateInstruction{1, 0, 1, 1, "C", "F"}
	C = stateInstruction{1, 1, -1, -1, "C", "A"}
	D = stateInstruction{0, 1, -1, 1, "E", "A"}
	E = stateInstruction{1, 0, -1, 1, "A", "B"}
	F = stateInstruction{0, 0, 1, 1, "C", "E"}

	states = map[string]stateInstruction{"A": A, "B": B, "C": C, "D": D, "E": E, "F": F}
)

func main() {
	println(solve(12302209, states))
}

func solve(diagnoseAfter int, states map[string]stateInstruction) (checksum int) {
	t := map[int]int{}
	currentState := "A"

	cursor := 0
	i := 0
	for i < diagnoseAfter {
		cursorChange, nextState, writeVal := evalState(currentState, t[cursor], states)
		t[cursor] = writeVal
		cursor += cursorChange
		currentState = nextState
		i++
	}

	for k := range t {
		if t[k] == 1 {
			checksum++
		}
	}

	return checksum
}

func evalState(state string, cursorVal int, states map[string]stateInstruction) (cursorChange int, nextState string, writeVal int) {
	s := states[state]
	if cursorVal == 0 {
		return s.cursorChange0, s.nextState0, s.writeVal0
	}
	return s.cursorChange1, s.nextState1, s.writeVal1
}
