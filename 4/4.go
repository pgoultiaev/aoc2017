package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pgoultiaev/aoc2017/util"
)

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum1 := 0
	sum2 := 0
	i := 0
	for scanner.Scan() {
		slicedString := strings.Split(scanner.Text(), " ")

		valid1 := true
		valid2 := true
		m := map[string]string{}
		for j, word := range slicedString {
			if m[word] != "" {
				// fmt.Printf("found duplicate word %s at line %d\n", word, i)
				valid1 = false
				break
			}
			m[word] = word

			for k, compareWord := range slicedString {
				if util.IsPermutation(word, compareWord) && k != j {
					//fmt.Printf("found permutation [%s,%s] at line %d\n", word, compareWord, i)
					valid2 = false
				}
			}
		}

		if valid1 == true {
			sum1++
		}
		if valid2 == true {
			sum2++
		}
		i++
	}

	fmt.Printf("part one: %d\n", sum1)
	fmt.Printf("part two: %d\n", sum2)
}
