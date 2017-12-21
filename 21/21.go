package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/pgoultiaev/aoc2017/util"
)

func main() {
	println(solve("21.txt", 5))
	println(solve("21.txt", 18))
}

func solve(filename string, iterations int) (ones int) {
	grid := [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
	}

	ruleSet := readInput(filename)
	pattern := [][]int{}

	i := 0
	for i < iterations {
		t := time.Now()

		count := 0
		newGrid := [][][]int{}

		size := 3
		if len(grid)%2 == 0 {
			size = 2
		}

		l := 0
		m := 0
		for l < len(grid) {
			count = 0
			for m < len(grid) {
				pattern = extract(grid, l, m, size)
				count++

				matchedRule := false
				for !matchedRule {
					for k := range ruleSet {
						if reflect.DeepEqual(ruleToPattern(k), pattern) || reflect.DeepEqual(ruleToPattern(k), flip(pattern)) {
							matchedRule = true
							newGrid = append(newGrid, ruleToPattern(ruleSet[k]))
							break
						}
						pattern = rotate(pattern)
					}
				}
				m += size
			}
			m = 0
			l += size
		}

		grid = combine(newGrid, size+1, count)
		//fmt.Printf("grid is now: %+v\n", grid)
		fmt.Printf("round %d took %+v\n", i+1, time.Since(t))
		i++
	}

	return countOnes(grid)
}

func countOnes(grid [][]int) (ones int) {
	for i := range grid {
		for _, k := range grid[i] {
			if k == 1 {
				ones++
			}
		}
	}
	return ones
}

func combine(gridOfGrids [][][]int, size, count int) (grid [][]int) {
	j := 0
	grid = make([][]int, size*count)
	for i := 0; i < (size * count); i++ {
		tmp := []int{}

		l := 0
		for l < count {
			tmp = append(tmp, gridOfGrids[l+j][i%size]...)
			l++
		}

		if (i+1)%size == 0 {
			j += count
		}

		grid[i] = tmp
	}

	// print(grid)
	return grid
}

func extract(grid [][]int, row, column, size int) [][]int {
	pattern := make([][]int, size)
	j := 0
	for j < size {
		pattern[j] = grid[row][column : column+size]
		j++
		row++
	}
	return pattern
}

func flip(pattern [][]int) [][]int {
	// duplicate two dimensional slice...
	patternF := make([][]int, len(pattern))
	for i := range pattern {
		patternF[i] = make([]int, len(pattern[i]))
		copy(patternF[i], pattern[i])
	}

	// flip
	for i := range patternF {
		patternF[i] = util.ReverseOrder(patternF[i])
	}

	//print(patternF)
	return patternF
}

func rotate(pattern [][]int) [][]int {
	// duplicate two dimensional slice...
	patternR := make([][]int, len(pattern))
	for i := range pattern {
		patternR[i] = make([]int, len(pattern[i]))
		copy(patternR[i], pattern[i])
	}

	// rotate
	for j := 0; j < len(patternR); j++ {
		for k := 0; k < len(patternR); k++ {
			patternR[j][k] = pattern[len(patternR)-k-1][j]
		}
	}

	//print(patternR)
	return patternR
}

func ruleToPattern(s string) (pattern [][]int) {
	mod := 2
	if len(s)%3 == 0 {
		mod = 3
	} else if len(s) == 16 {
		mod = 4
	}

	ia := util.ConvStringArrayToIntArray(strings.Split(s, ""))
	i := 0
	for i < len(ia) {
		if (i+1)%mod == 0 {
			pattern = append(pattern, ia[i+1-mod:i+1])
		}
		i++
	}

	//print(pattern)
	return pattern
}

func readInput(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ruleSet := map[string]string{}
	for scanner.Scan() {
		rule := scanner.Text()

		i := strings.Index(rule, "=>")
		r := strings.NewReplacer("/", "", ".", "0", "#", "1")

		fromRule := rule[:i-1]
		fromRuleParsed := r.Replace(fromRule)
		toRule := rule[i+3:]
		toRuleParsed := r.Replace(toRule)

		ruleSet[fromRuleParsed] = toRuleParsed
	}

	return ruleSet
}

func print(p [][]int) {
	for i := range p {
		fmt.Printf("%+v\n", p[i])
	}
}
