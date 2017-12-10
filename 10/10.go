package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("10.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	slicedString := strings.Split(s, ",")
	slicedInts := convStringArrayToIntArray(slicedString)

	fmt.Printf("slicedInts: %d, should be 12\n", partOne([]int{3, 4, 1, 5}, 4))
	fmt.Printf("slicedInts: %d\n", partOne(slicedInts, 255))
}

func partOne(ia []int, max int) int {
	a := makeRange(0, max)

	//fmt.Printf("range: %+v, len %d\n", a, len(a))

	skipSize := 0
	currPos := 0
	for i, num := range ia {
		if num > 1 {
			// fmt.Printf("currentPos: %d, currentLength %d\n", currPos, num)
			toSort := []int{}
			if currPos+num-1 < len(a) {
				// fmt.Printf("selecting [%d:%d]\n", currPos, currPos+num-1)
				toSort = a[currPos : currPos+num]
			} else {
				// fmt.Printf("wrapping [%d:] and [:%d]\n", currPos, (currPos+num)%len(a))
				toSort = append(a[currPos:], a[:((currPos+num)%len(a))]...)
			}

			toSort = reverseOrder(toSort)
			// fmt.Printf("reversed toSort: %+v\n", toSort)

			for j := range toSort {
				a[(currPos+j)%len(a)] = toSort[j]
				j++
			}
		}

		//fmt.Printf("state: %+v\n\n", a)
		currPos = (currPos + num + skipSize) % len(a)
		skipSize++
		i++
	}
	return a[0] * a[1]
}

func reverseOrder(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func convStringArrayToIntArray(sa []string) []int {
	var output []int
	for _, e := range sa {
		i, _ := strconv.Atoi(e)
		output = append(output, i)
	}
	return output
}
