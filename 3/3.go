package main

import "fmt"

type Player struct {
	currPosition     Point
	index            int
	visited          map[Point]int
	endSumValue      int
	endSumValueFound bool
}

type Point struct {
	X, Y int
}

func main() {
	taxiDistance, sum := solve(277678)
	fmt.Printf("solution for 277678: taxiDistance: %d, sum: %d\n", taxiDistance, sum)
}

func solve(end int) (int, int) {
	if end == 1 {
		return 0, 1
	}

	origin := Point{0, 0}
	p := &Player{currPosition: origin, index: 1, visited: map[Point]int{Point{0, 0}: 1}}
	sqSize := 0

	for p.index < end {
		p.traverse("Up", sqSize-1, end)
		p.traverse("Left", sqSize, end)
		p.traverse("Down", sqSize, end)
		p.traverse("Right", sqSize+1, end)
		sqSize += 2
	}

	return taxiDistance(origin, p.currPosition), p.endSumValue
}

func (p *Player) traverse(direction string, dist, end int) {
	if dist > 0 {
		distTravelled := 0
		for distTravelled < dist && p.index < end {
			switch direction {
			case "Up":
				p.currPosition.Y++
			case "Down":
				p.currPosition.Y--
			case "Right":
				p.currPosition.X++
			case "Left":
				p.currPosition.X--
			}
			p.index++
			sum := p.sum()

			if sum > end && !p.endSumValueFound {
				p.endSumValue = sum
				p.endSumValueFound = true
			}
			p.visited[p.currPosition] = sum
			distTravelled++
		}
	}
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
	dX := abs(p1.X - p2.X)
	dY := abs(p1.Y - p2.Y)
	return dX + dY
}

// Why does golang not have a Abs(int) function in the math package?
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
