package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Point struct {
	X, Y, Z int
}

type Player struct {
	position Point
}

type Direction struct {
	dx, dy int
}

var (
	n      = Direction{0, 1}
	nw     = Direction{-1, 1}
	ne     = Direction{1, 0}
	s      = Direction{0, -1}
	sw     = Direction{-1, 0}
	se     = Direction{1, -1}
	origin = Point{0, 0, 0}

	distances = []int{}

	directions = map[string]Direction{
		"n":  n,
		"nw": nw,
		"ne": ne,
		"s":  s,
		"sw": sw,
		"se": se,
	}
)

func main() {
	buf, err := ioutil.ReadFile("11.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	slicedString := strings.Split(s, ",")

	solve([]string{"se", "se", "se"})
	solve(slicedString)
}

func solve(sa []string) {
	p := Player{position: origin}
	for _, dir := range sa {
		p.move(directions[dir])
		distances = append(distances, hexGridDistance(origin, p.position))
	}
	fmt.Printf("Part one : last coordinate [%d,%d,%d], hexgridDistance: %d\n", p.position.X, p.position.Y, p.position.Z, hexGridDistance(origin, p.position))
	fmt.Printf("Part two : largest hexgridDistance: %d\n", max(distances))
}

func max(ia []int) int {
	sort.Ints(ia[:])
	return ia[len(ia)-1]
}

func (p *Player) move(direction Direction) {
	p.position.X += direction.dx
	p.position.Y += direction.dy
	p.position.Z = 0 - p.position.X - p.position.Y
}

func hexGridDistance(p1, p2 Point) int {
	dX := abs(p1.X - p2.X)
	dY := abs(p1.Y - p2.Y)
	dZ := abs(p1.Z - p2.Z)
	return (dX + dY + dZ) / 2
}

// Why does golang not have a Abs(int) function in the math package?
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
