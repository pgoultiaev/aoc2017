package main

import (
	"fmt"
	"io/ioutil"
)

// from http://adventofcode.com/2017/day/1

func main() {
	buf, err := ioutil.ReadFile("1.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	l := len(s)

	fmt.Printf("Part one total sum = %d\n", solve(s, 1))
	fmt.Printf("Part two total sum = %d\n", solve(s, l/2))
}

func solve(input string, offset int) int {
	sum := 0
	for i := range input {
		curr := input[i]
		next := input[(i+offset)%len(input)]

		if curr == next {
			sum += int(curr - '0')
		}
	}
	return sum
}
