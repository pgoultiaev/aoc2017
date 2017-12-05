package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("5.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	slicedString := strings.Split(s, "\n")
	slicedInts := convStringArrayToIntArray(slicedString)

	fmt.Printf("part one: %d\n", solvePartOne(slicedInts))
	fmt.Printf("part two: %d\n", solvePartTwo(convStringArrayToIntArray(slicedString)))

}

func solvePartOne(ia []int) int {
	count1 := 0
	i := 0
	for i < len(ia) {
		offset := ia[i]

		count1++

		ia[i] = ia[i] + 1
		i += offset
		if i > len(ia) {
			break
		}

		//fmt.Printf("%v\n", ia)
	}
	return count1
}

func solvePartTwo(ia []int) int {
	count1 := 0
	i := 0
	for i < len(ia) {
		offset := ia[i]

		count1++

		if offset >= 3 {
			ia[i] = ia[i] - 1
		} else {
			ia[i] = ia[i] + 1
		}

		i += offset
		if i > len(ia) {
			break
		}

		//fmt.Printf("%v\n", ia)
	}
	return count1
}

func convStringArrayToIntArray(sa []string) []int {
	var output []int
	for _, e := range sa {
		i, _ := strconv.Atoi(e)
		output = append(output, i)
	}
	return output
}
