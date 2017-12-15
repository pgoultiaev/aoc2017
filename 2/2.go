package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/pgoultiaev/aoc2017/util"
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

		slicedInts := util.ConvStringArrayToIntArray(slicedString)
		sort.Ints(slicedInts[:])
		sum1 += (slicedInts[len(slicedInts)-1] - slicedInts[0])
		sum2 += findDivisible(slicedInts)
	}

	fmt.Printf("part one: %d\n", sum1)
	fmt.Printf("part two: %d\n", sum2)
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
