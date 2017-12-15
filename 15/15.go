package main

import (
	"strconv"
)

func main() {
	// from input
	a := int64(516)
	b := int64(190)

	factorA := int64(16807)
	factorB := int64(48271)

	modAB := int64(2147483647)

	finalCount := solve(a, b, factorA, factorB, modAB)
	println(finalCount)

	finalCount2 := solve2(a, b, factorA, factorB, modAB)
	println(finalCount2)
}

// Part two
func solve2(a, b, factorA, factorB, modAB int64) (finalCount int) {
	i := 0
	for i < 5000000 {
		a, b = generatePartTwo(a, b, factorA, factorB, modAB)
		finalCount += compareLast16bits(a, b)

		i++
	}
	return finalCount
}

func generatePartTwo(a, b, factorA, factorB, modAB int64) (int64, int64) {
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
func solve(a, b, factorA, factorB, modAB int64) (finalCount int) {
	i := 0
	for i < 40000000 {
		a = (a * factorA) % modAB
		b = (b * factorB) % modAB

		finalCount += compareLast16bits(a, b)

		i++
	}
	return finalCount
}

func compareLast16bits(a, b int64) int {
	stringA := strconv.FormatInt(a, 2)
	stringB := strconv.FormatInt(b, 2)

	lA := len(stringA)
	lB := len(stringB)

	if lA > 16 {
		stringA = stringA[lA-16:]
	} else {
		stringA = stringA[:]
	}

	if lB > 16 {
		stringB = stringB[lB-16:]
	} else {
		stringB = stringB[:]
	}

	//fmt.Printf("a last 16 bits: %s\nb last 16 bits: %s\n", stringA, stringB)
	if stringA == stringB {
		return 1
	}

	return 0
}
