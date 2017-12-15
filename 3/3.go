package main

import "fmt"
import "github.com/pgoultiaev/aoc2017/util"

type Player struct {
	currPosition     Point
	index            int
	visited          map[Point]int
	endSumValue      int
	endSumValueFound bool
	direction        Direction
	distTravelled    int
}

type Point struct {
	X, Y int
}

type Direction struct {
	dx, dy int
}

var (
	Right  = Direction{1, 0}
	Up     = Direction{0, 1}
	Left   = Direction{-1, 0}
	Down   = Direction{0, -1}
	origin = Point{0, 0}
)

var nextDirection = map[Direction]Direction{
	Right: Up,
	Up:    Left,
	Left:  Down,
	Down:  Right,
}

func main() {
	taxiDistance, sum := solve(277678)
	fmt.Printf("solution for 277678: taxiDistance: %d, sum: %d\n", taxiDistance, sum)
}

func solve(end int) (int, int) {
	if end == 1 {
		return 0, 1
	}

	p := &Player{currPosition: origin, index: 1, visited: map[Point]int{Point{0, 0}: 1}, direction: Right, distTravelled: 0}
	sqSize := 2

	for p.index < end {
		sum := p.sum()
		if sum > end && !p.endSumValueFound {
			p.endSumValue = sum
			p.endSumValueFound = true
		}
		p.visited[p.currPosition] = sum

		if p.distTravelled == sqSize/2 {
			p.direction = nextDirection[p.direction]
		} else if p.distTravelled == sqSize {
			p.distTravelled = 0
			sqSize += 2
			p.direction = nextDirection[p.direction]
		}
		p.traverse(p.direction)
		p.distTravelled++
		p.index++
	}

	return taxiDistance(origin, p.currPosition), p.endSumValue
}

func (p *Player) traverse(direction Direction) {
	p.currPosition.X += direction.dx
	p.currPosition.Y += direction.dy
}

func (p *Player) sum() (sum int) {
	sum += p.visited[Point{p.currPosition.X + 1, p.currPosition.Y}]
	sum += p.visited[Point{p.currPosition.X + 1, p.currPosition.Y + 1}]
	sum += p.visited[Point{p.currPosition.X, p.currPosition.Y + 1}]
	sum += p.visited[Point{p.currPosition.X - 1, p.currPosition.Y + 1}]
	sum += p.visited[Point{p.currPosition.X - 1, p.currPosition.Y}]
	sum += p.visited[Point{p.currPosition.X - 1, p.currPosition.Y - 1}]
	sum += p.visited[Point{p.currPosition.X, p.currPosition.Y - 1}]
	sum += p.visited[Point{p.currPosition.X + 1, p.currPosition.Y - 1}]

	if sum == 0 {
		sum++
	}

	return sum
}

func taxiDistance(p1, p2 Point) int {
	dX := util.Abs(p1.X - p2.X)
	dY := util.Abs(p1.Y - p2.Y)
	return dX + dY
}
