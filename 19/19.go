package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Point struct {
	X, Y int
}

func main() {
	file, err := os.Open("19.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := map[Point]string{}
	var origin Point

	rowNum := 1
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")

		for i := range row {
			if row[i] != " " {
				grid[Point{i, rowNum}] = row[i]
				if rowNum == 1 {
					origin = Point{i, 1}
					fmt.Printf("origin: %+v\n", origin)
				}
			}
		}
		rowNum++
	}

	trip, steps := solve(grid, origin)
	println(trip)
	println(steps)
}

func solve(grid map[Point]string, origin Point) (trip string, steps int) {
	trip, steps = traverse(grid, origin, "D", "", 0)
	return trip, steps
}

func traverse(grid map[Point]string, p Point, dir string, t string, steps int) (string, int) {
	path := grid[p]
	if isLetter(path) {
		t += path
	}

	if grid[p] == "" {
		return t, steps
	}

	var next Point
	switch {
	case path == "+":
		next, dir = findNext(grid, p, dir)
	case dir == "D":
		next = Point{p.X, p.Y + 1}
	case dir == "U":
		next = Point{p.X, p.Y - 1}
	case dir == "L":
		next = Point{p.X - 1, p.Y}
	case dir == "R":
		next = Point{p.X + 1, p.Y}
	}

	steps++
	t, steps = traverse(grid, next, dir, t, steps)
	return t, steps
}

func findNext(grid map[Point]string, p Point, dir string) (Point, string) {
	down := Point{p.X, p.Y + 1}
	up := Point{p.X, p.Y - 1}
	left := Point{p.X - 1, p.Y}
	right := Point{p.X + 1, p.Y}

	if dir == "U" || dir == "D" {
		switch {
		case len(grid[left]) > 0:
			return left, "L"
		case len(grid[right]) > 0:
			return right, "R"
		}
	} else {
		switch {
		case len(grid[down]) > 0:
			return down, "D"
		case len(grid[up]) > 0:
			return up, "U"
		}
	}

	return Point{}, ""
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
