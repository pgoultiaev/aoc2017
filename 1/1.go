package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// from http://adventofcode.com/2017/day/1

func main() {
	buf, err := ioutil.ReadFile("1.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	slicedString := strings.Split(s, "")
	l := len(slicedString)

	sum := 0
	var next string
	for i := 0; i < l; i++ {
		curr := slicedString[i]

		// if end of list reached, point to first item
		if i == l-1 {
			next = slicedString[0]
		} else {
			next = slicedString[i+1]
		}

		if curr == next {
			digit, _ := strconv.Atoi(curr)
			sum += digit
		}
	}

	fmt.Printf("Part one total sum = %d\n", sum)

	halfway := l / 2
	sumHalfway := 0
	var nextHalfway string
	for i := 0; i < l; i++ {
		curr := slicedString[i]

		// if end of list reached, point to first item
		if i == l-1 {
			nextHalfway = slicedString[halfway-1]
		} else if i+l/2 >= l {
			n := i + halfway
			nextHalfway = slicedString[n-l]
		} else {
			nextHalfway = slicedString[i+halfway]
		}

		if curr == nextHalfway {
			digit, _ := strconv.Atoi(curr)
			sumHalfway += digit
		}
	}

	fmt.Printf("Part two total sum = %d\n", sumHalfway)
}
