package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// from https://pastebin.com/BMd61PUv
// using input: https://pastebin.com/wGmzZHeq

type Pos struct {
	X int
	Y int
}

func main() {
	buf, err := ioutil.ReadFile("0.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	slicedString := strings.Split(s, ", ")

	origin := Pos{0, 0}
	p := &Pos{0, 0}
	maxDistance := 0
	maxDistancePos := origin
	var a []Pos
	var b []Pos

	// Go through all input
	for _, e := range slicedString {
		var currDist int
		switch e {
		case "A":
			a = append(a, *p)
			//fmt.Printf("set <%s>, at: [%d,%d]\n", e, p.X, p.Y)
			currDist = taxiDistance(*p, origin)
		case "B":
			b = append(b, *p)
			//fmt.Printf("set <%s>, at: [%d,%d]\n", e, p.X, p.Y)
			currDist = taxiDistance(*p, origin)
		case "Start":
			fmt.Printf("Ended at: [%d,%d]\n", p.X, p.Y)
		default:
			traverse(e, p)
		}

		if currDist > maxDistance {
			maxDistance = currDist
			maxDistancePos = *p
		}
	}

	// Question 1: Identify the marker furthest from the origin, as measured by the taxicab distance, and return that distance.
	fmt.Printf("Max distance from origin at [%d,%d], distance = %d\n", maxDistancePos.X, maxDistancePos.Y, maxDistance)

	// Question 2: Consider all pairs of *different* markers (where a pair may consist of any 'A' and any 'B' marker).
	// Identify the pair maximally far apart from one another, as measured by the taxicab distance, and return that distance.
	posA, posB, maxD := maxDistanceOnGrid(a, b)
	fmt.Printf("Max distance from A[%d,%d] to B[%d,%d], distance = %d\n", posA.X, posA.Y, posB.X, posB.Y, maxD)
}

// Just iterate over all A's and check their distance to each B
func maxDistanceOnGrid(a, b []Pos) (Pos, Pos, int) {
	maxDistance := 0
	var maxDistPosA Pos
	var maxDistPosB Pos
	for _, posA := range a {
		for _, posB := range b {
			currDist := taxiDistance(posA, posB)
			if currDist > maxDistance {
				maxDistance = currDist
				maxDistPosA = posA
				maxDistPosB = posB
			}
		}
	}
	return maxDistPosA, maxDistPosB, maxDistance
}

func traverse(direction string, p *Pos) {
	switch direction {
	case "Up":
		p.Y++
		//fmt.Printf("moved <%s>, now standing at:    [%d,%d]\n", direction, p.X, p.Y)
	case "Down":
		p.Y--
		//fmt.Printf("moved <%s>, now standing at:  [%d,%d]\n", direction, p.X, p.Y)
	case "Right":
		p.X++
		//fmt.Printf("moved <%s>, now standing at: [%d,%d]\n", direction, p.X, p.Y)
	case "Left":
		p.X--
		//fmt.Printf("moved <%s>, now standing at:  [%d,%d]\n", direction, p.X, p.Y)
	default:
		fmt.Printf("UNKNOWN DIRECTION <%s>\n", direction)
	}
}

func taxiDistance(pos1, pos2 Pos) int {
	dX := abs(pos1.X - pos2.X)
	dY := abs(pos1.Y - pos2.Y)
	return dX + dY
}

// Why does golang not have a Abs(int) function in the math package?
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
