package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pgoultiaev/aoc2017/util"
)

func main() {
	buf, err := ioutil.ReadFile("10.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)
	slicedASCII := util.ConvStringArrayToByteArray(strings.Split(s, ""))
	slicedInts := util.ConvStringArrayToIntArray(strings.Split(s, ","))

	// Part one : 19591
	partOne, _, _ := partOne(util.MakeRange(0, 255), slicedInts, 0, 0)
	fmt.Printf("slicedInts: %d\n", partOne[0]*partOne[1])

	// Part two
	fmt.Printf("knot hash %s\n", partTwo(util.MakeRange(0, 255), slicedASCII))
}

func partTwo(a []int, ba []byte) (hash string) {
	suffix := []byte{17, 31, 73, 47, 23}
	ba = append(ba, suffix...)
	ia := util.ConvByteArrayToIntArray(ba)

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
	hexString := ""
	for _, element := range ia {
		hexString += fmt.Sprintf("%.02x", element)
	}
	return hexString
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

			toSort = util.ReverseOrder(toSort)
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
