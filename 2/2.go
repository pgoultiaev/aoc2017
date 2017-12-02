package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		slicedString := strings.Split(scanner.Text(), "\t")

		slicedInts := convStringArrayToIntArray(slicedString)
		sort.Ints(slicedInts[:])
		sum1 += (slicedInts[len(slicedInts)-1] - slicedInts[0])
		sum2 += findDivisible(slicedInts)
	}

	fmt.Printf("part one: %d\n", sum1)
	fmt.Printf("part two: %d\n", sum2)
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
			if j != i && e%o == 0 {
				return e / o
			}
		}
	}
	return 0
}
