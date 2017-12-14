package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := "hfdlxzhv"

	grid, count := solve(input)
	println(count)
	println(solve2(grid))
}

// Part two
type Point struct {
	x, y int
}

func solve2(grid [128][]string) (numRegions int) {
	gridMap := fillGridMap(grid)

	for k := range gridMap {
		findRegion(k, gridMap)
		numRegions++
	}
	return numRegions
}

func findRegion(p Point, g map[Point]bool) {
	if !g[p] {
		return
	}

	up := Point{p.x, p.y + 1}
	down := Point{p.x, p.y - 1}
	left := Point{p.x - 1, p.y}
	right := Point{p.x + 1, p.y}

	delete(g, p)

	findRegion(up, g)
	findRegion(down, g)
	findRegion(left, g)
	findRegion(right, g)
}

func fillGridMap(grid [128][]string) map[Point]bool {
	gridMap := map[Point]bool{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "1" {
				gridMap[Point{i, j}] = true
			}
		}
	}
	return gridMap
}

// Part One
func solve(s string) (grid [128][]string, gridOnes int) {
	i := 0
	for i < 128 {
		sHash := fmt.Sprintf("%s-%d", s, i)
		hash := KnotHashDay10(sHash)
		binVal, err := getBinVal(hash)
		if err != nil {
			log.Fatal(err)
			return
		}

		grid[i] = strings.Split(binVal, "")
		gridOnes += strings.Count(binVal, "1")
		i++
	}
	return grid, gridOnes
}

func getBinVal(s string) (string, error) {
	var buffer bytes.Buffer
	sa := strings.Split(s, "")
	for _, char := range sa {
		ui, err := strconv.ParseUint(char, 16, 64)
		if err != nil {
			return "", err
		}

		// base 2, zero padded, 4 characters
		buffer.WriteString(fmt.Sprintf("%04b", ui))
	}
	return buffer.String(), nil
}
