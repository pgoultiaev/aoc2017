package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var checksumPerLine1 []int
	var checksumPerLine2 []int
	for scanner.Scan() {
		slicedString := strings.Split(scanner.Text(), "\t")

		slicedInts := convStringArrayToIntArray(slicedString)
		checksumPerLine1 = append(checksumPerLine1, max(slicedInts)-min(slicedInts))
		checksumPerLine2 = append(checksumPerLine2, findDivisible(slicedInts))
	}

	// Total checksum part one
	total1 := 0
	for _, v := range checksumPerLine1 {
		total1 += v
	}
	fmt.Printf("part one: %d\n", total1)

	// Total checksum part two
	total2 := 0
	for _, v := range checksumPerLine2 {
		total2 += v
	}
	fmt.Printf("part two: %d\n", total2)

}

func max(slice []int) int {
	max := 0
	for _, e := range slice {
		if e > max {
			max = e
		}
	}
	return max
}

func min(slice []int) int {
	min := slice[0]
	if len(slice) > 1 {
		for _, e := range slice {
			if e < min {
				min = e
			}
		}
	}
	return min
}

func convStringArrayToIntArray(sa []string) []int {
	var output []int
	for _, e := range sa {
		i, _ := strconv.Atoi(e)
		output = append(output, i)
	}
	return output
}

func findDivisible(ia []int) int {
	for i, e := range ia {
		for j, o := range ia {
			if j == i {
				continue
			}

			div := float64(float64(e) / float64(o))
			if o > e {
				div = float64(float64(o) / float64(e))
			}
			_, frac := math.Modf(div)
			if frac == 0.0 {
				return int(div)
			}
		}
	}
	return 0
}
