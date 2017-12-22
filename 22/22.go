package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Virus struct {
	position Point
}

type Point struct {
	X, Y int
}

type Direction struct {
	dx, dy int
}

var (
	up    = Direction{0, -1}
	down  = Direction{0, 1}
	left  = Direction{-1, 0}
	right = Direction{1, 0}

	// directionsLeft = map[string]Direction{
	// 	"up":    left,
	// 	"down":  right,
	// 	"left":  down,
	// 	"right": up,
	// }

	// directionsRight = map[string]Direction{
	// 	"up":    right,
	// 	"down":  left,
	// 	"left":  up,
	// 	"right": down,
	// }
)

func main() {
	grid, middle := readInput("22.txt")
	println(solve(grid, middle, 10000))
}

func solve(grid map[Point]bool, middle Point, bursts int) (infectBursts int) {
	directions := []Direction{up, right, down, left}
	dirPointer := 0

	virus := Virus{middle}
	i := 0
	//fmt.Printf("start at: %+v\n\n", virus.position)
	for i < bursts {
		if grid[virus.position] {
			dirPointer = (dirPointer + 1) % len(directions)
			grid[virus.position] = false
			// fmt.Printf("turn right, cleaned infected at: %+v\n", virus.position)
		} else {
			if dirPointer == 0 {
				dirPointer = len(directions) - 1
			} else {
				dirPointer = (dirPointer - 1) % len(directions)
			}
			grid[virus.position] = true
			infectBursts++
			// fmt.Printf("turn left, infected clean at: %+v, infectbursts: %d\n", virus.position, infectBursts)
		}
		//fmt.Printf("grid: %+v\n", grid)
		// fmt.Printf("moving: %d, position: %+v\n\n", dirPointer, virus.position)
		virus.move(directions[dirPointer])
		i++
	}
	return infectBursts
}

func (p *Virus) move(direction Direction) {
	p.position.X += direction.dx
	p.position.Y += direction.dy
}

func readInput(filename string) (map[Point]bool, Point) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := map[Point]bool{}

	rowNum := 1
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for i := range line {
			if line[i] == "#" {
				grid[Point{i + 1, rowNum}] = true
			}
		}
		rowNum++
	}

	mid := (rowNum - 1) / 2
	middle := Point{mid, mid}

	return grid, middle
}
