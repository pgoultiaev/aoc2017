package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var firewall = map[int]int{}

func main() {
	file, err := os.Open("13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		link := scanner.Text()

		splitStrings := strings.Split(link, ": ")
		column, _ := strconv.Atoi(splitStrings[0])
		depth, _ := strconv.Atoi(splitStrings[1])

		firewall[column] = depth
	}

	threat := solve(firewall)
	fmt.Printf("Part one, threat: %d\n", threat)

	delay := 0
	for {
		caught := solve2(firewall, delay)
		if !caught {
			fmt.Printf("did not get caught with delay %d\n", delay)
			break
		}
		delay++
	}
}

func solve2(fw map[int]int, delay int) bool {
	for k, v := range fw {
		if (k+delay)%((v-1)*2) == 0 {
			return true
		}
	}
	return false
}

// Part one
func solve(fw map[int]int) (threat int) {
	for k, v := range fw {
		if k%((v-1)*2) == 0 {
			threat += k * v
		}
	}
	return threat
}
