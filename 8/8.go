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

var register = map[string]int{}

func main() {
	file, err := os.Open("8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	largestEvar := 0
	for scanner.Scan() {
		instruction := scanner.Text()

		i := strings.Index(instruction, "if")
		instrString := instruction[0 : i-1]
		instrStrings := strings.Split(instrString, " ")

		ruleString := instruction[i+3:]
		ruleStrings := strings.Split(ruleString, " ")

		if evalRule(ruleStrings) {
			if l := getLargest(); l > largestEvar {
				largestEvar = l
			}
			runInstruction(instrStrings)
		}
	}

	fmt.Printf("largest number after instructions: %d\n", getLargest())
	fmt.Printf("largest number ever: %d", largestEvar)
}

func getLargest() int {
	if len(register) == 0 {
		return 0
	}
	var values []int
	for _, v := range register {
		values = append(values, v)
	}
	sort.Ints(values)
	return values[len(values)-1]
}

func evalRule(s []string) bool {
	regVal := register[s[0]]
	val, _ := strconv.Atoi(s[2])

	switch s[1] {
	case ">=":
		if regVal >= val {
			return true
		}
	case "<=":
		if regVal <= val {
			return true
		}
	case ">":
		if regVal > val {
			return true
		}
	case "<":
		if regVal < val {
			return true
		}
	case "!=":
		if regVal != val {
			return true
		}
	case "==":
		if regVal == val {
			return true
		}
	}
	return false
}

func runInstruction(s []string) {
	val, _ := strconv.Atoi(s[2])
	if s[1] == "inc" {
		register[s[0]] += val
	} else {
		register[s[0]] -= val
	}
}
