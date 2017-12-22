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
)

func main() {
	grid, middle := readInput("22.txt")
	grid2, _ := readInput("22.txt")
	println(solve(grid, middle, 10000))
	println(solve2(grid2, middle, 10000000))
}

// Part two
func solve2(grid map[Point]string, middle Point, bursts int) int {
	directions := []Direction{up, right, down, left}
	dirPointer := 0

	infectBursts := 0
	virus := Virus{middle}
	i := 0

	for i < bursts {
		switch grid[virus.position] {
		case "":
			if dirPointer == 0 {
				dirPointer = len(directions) - 1
			} else {
				dirPointer = (dirPointer - 1) % len(directions)
			}
			grid[virus.position] = "W"
		case "W":
			grid[virus.position] = "#"
			infectBursts++
		case "#":
			dirPointer = (dirPointer + 1) % len(directions)
			grid[virus.position] = "F"
		case "F":
			dirPointer = (dirPointer + 2) % len(directions)
			grid[virus.position] = ""
		}

		virus.move(directions[dirPointer])
		i++
	}

	return infectBursts
}

// Part one
func solve(grid map[Point]string, middle Point, bursts int) (infectBursts int) {
	directions := []Direction{up, right, down, left}
	dirPointer := 0

	virus := Virus{middle}
	i := 0
	for i < bursts {
		if grid[virus.position] == "#" {
			dirPointer = (dirPointer + 1) % len(directions)
			grid[virus.position] = ""
		} else {
			if dirPointer == 0 {
				dirPointer = len(directions) - 1
			} else {
				dirPointer = (dirPointer - 1) % len(directions)
			}
			grid[virus.position] = "#"
			infectBursts++
		}
		virus.move(directions[dirPointer])
		i++
	}
	return infectBursts
}

func (p *Virus) move(direction Direction) {
	p.position.X += direction.dx
	p.position.Y += direction.dy
}

func readInput(filename string) (map[Point]string, Point) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := map[Point]string{}
	rowNum := 1
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for i := range line {
			if line[i] == "#" {
				grid[Point{i + 1, rowNum}] = "#"
			}
		}
		rowNum++
	}

	mid := rowNum / 2
	middle := Point{mid, mid}

	return grid, middle
}
