package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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
				if isPermutation(word, compareWord) && k != j {
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

func isPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1split := strings.Split(s1, "")
	sort.Strings(s1split)
	s2split := strings.Split(s2, "")
	sort.Strings(s2split)

	for i := range s1split {
		if s1split[i] != s2split[i] {
			return false
		}
	}
	return true
}
