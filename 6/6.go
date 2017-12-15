package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pgoultiaev/aoc2017/util"
)

func main() {
	buf, err := ioutil.ReadFile("6.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	slicedString := strings.Split(s, "\t")
	slicedInts := util.ConvStringArrayToIntArray(slicedString)

	// fmt.Printf("input: %v\n", slicedInts)
	fmt.Printf("part one: %d\n", solve(slicedInts))
	// fmt.Printf("input: %v\n", slicedInts)
	fmt.Printf("part two: %d\n", solve(slicedInts))
}

func solve(input []int) (it int) {
	var mem [][]int
	found := false
	for !found {
		//fmt.Printf("input : %+v before iteration: %d\n", input, it)
		inputStore := make([]int, len(input))
		copy(inputStore, input)
		mem = append(mem, inputStore)

		fm, maxVal := firstMax(input)
		input[fm] = 0

		rest := maxVal
		for i := (fm + 1) % len(input); rest > 0; rest-- {
			input[i] = input[i] + 1
			i = (i + 1) % len(input)
		}
		//fmt.Printf("input : %+v after iteration: %d\n", input, it)

		// fmt.Printf("mem: %v\n", mem)
		for _, e := range mem {
			if util.Equals(input, e) {
				found = true
				return it + 1
			}
		}

		it++
	}
	return it
}

func firstMax(a []int) (index, val int) {
	max := 0
	for i, e := range a {
		if e > max {
			max = e
			index = i
		}
	}
	//fmt.Printf("found max %d at %d\n", max, index)
	return index, max
}
