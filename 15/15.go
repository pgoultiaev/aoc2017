package main

func main() {
	// from input
	a := 516
	b := 190

	factorA := 16807
	factorB := 48271

	modAB := 2147483647

	finalCount := solve(a, b, factorA, factorB, modAB)
	println(finalCount)

	finalCount2 := solve2(a, b, factorA, factorB, modAB)
	println(finalCount2)
}

// Part two
func solve2(a, b, factorA, factorB, modAB int) (finalCount int) {
	i := 0
	for i < 5000000 {
		a, b = generatePartTwo(a, b, factorA, factorB, modAB)
		finalCount += compareLast16bits(a, b)

		i++
	}
	return finalCount
}

func generatePartTwo(a, b, factorA, factorB, modAB int) (int, int) {
	a = (a * factorA) % modAB
	b = (b * factorB) % modAB

	for a%4 != 0 {
		a = (a * factorA) % modAB
	}
	for b%8 != 0 {
		b = (b * factorB) % modAB
	}
	return a, b
}

// Part one
func solve(a, b, factorA, factorB, modAB int) (finalCount int) {
	i := 0
	for i < 40000000 {
		a = (a * factorA) % modAB
		b = (b * factorB) % modAB

		finalCount += compareLast16bits(a, b)

		i++
	}
	return finalCount
}

func compareLast16bits(a, b int) int {
	bitMask := 1<<16 - 1
	if a&bitMask == b&bitMask {
		return 1
	}

	return 0
}
