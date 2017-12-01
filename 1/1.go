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

	sum := 0
	var next string
	for i := 0; i < len(slicedString); i++ {
		curr := slicedString[i]

		// if end of list reached, point to first item
		if i == len(slicedString)-1 {
			next = slicedString[0]
		} else {
			next = slicedString[i+1]
		}

		if curr == next {
			digit, _ := strconv.Atoi(curr)
			sum += digit
		}
	}

	fmt.Printf("total sum = %d", sum)
}
