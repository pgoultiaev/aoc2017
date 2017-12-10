package main

import (
	"bytes"
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
	slicedASCII := convStringArrayToByteArray(strings.Split(s, ""))
	slicedInts := convStringArrayToIntArray(strings.Split(s, ","))

	// Part one : 19591
	partOne, _, _ := partOne(makeRange(0, 255), slicedInts, 0, 0)
	fmt.Printf("slicedInts: %d\n", partOne[0]*partOne[1])

	// Part two
	fmt.Printf("knot hash %s", partTwo(makeRange(0, 255), slicedASCII))
}

func partTwo(a []int, ba []byte) (hash string) {
	suffix := []byte{17, 31, 73, 47, 23}
	ba = append(ba, suffix...)
	ia := convByteArrayToIntArray(ba)

	rounds := 0
	currPos := 0
	skipSize := 0
	for rounds < 64 {
		a, currPos, skipSize = partOne(a, ia, currPos, skipSize)
		rounds++
	}

	// dense hash
	xordVals := []int{}
	current := 0
	for i, val := range a {
		current ^= val

		if (i+1)%16 == 0 {
			xordVals = append(xordVals, current)
			current = 0
		}
	}
	//fmt.Printf("xordVals length: %d\n\t%+v\n", len(xordVals), xordVals)

	// knot hash
	hash = knotHash(xordVals)
	return hash
}

func knotHash(ia []int) string {
	var buffer bytes.Buffer
	for i := range ia {
		buffer.WriteString(strconv.FormatInt(int64(ia[i]), 16))
	}
	return buffer.String()
}

func partOne(a []int, ia []int, cp, ss int) (partOne []int, currPos, skipSize int) {
	//fmt.Printf("range: %+v, len %d\n", a, len(a))
	currPos = cp
	skipSize = ss
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
	return a, currPos, skipSize
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

func convStringArrayToByteArray(s []string) []byte {
	var output []byte
	for _, e := range s {
		i := []byte(e)
		output = append(output, i...)
	}
	return output
}

func convByteArrayToIntArray(ba []byte) []int {
	var output []int
	for _, e := range ba {
		output = append(output, int(e))
	}
	return output
}
