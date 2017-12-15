package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pgoultiaev/aoc2017/util"
)

func main() {
	buf, err := ioutil.ReadFile("5.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	slicedString := strings.Split(s, "\n")
	slicedInts := util.ConvStringArrayToIntArray(slicedString)

	si2 := make([]int, len(slicedInts))
	copy(si2, slicedInts)

	//fmt.Printf("part one: %d\n")
	//fmt.Printf("part two: %d\n")

	messages := make(chan int)
	go func() {
		messages <- solvePartOne(slicedInts)
	}()
	go func() {
		messages <- solvePartTwo(si2)
	}()
	for i := 0; i < 2; i++ {
		fmt.Println(<-messages)
	}
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

func solvePartTwo(ia []int) (count1 int) {
	i := 0
	l := len(ia)
	for i < l {
		offset := ia[i]

		count1++

		if offset >= 3 {
			ia[i] = offset - 1
		} else {
			ia[i] = offset + 1
		}

		i += offset
		if i > l {
			break
		}
		//fmt.Printf("%v\n", ia)
	}
	return count1
}
