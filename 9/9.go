package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumAllGroups := 0
	for scanner.Scan() {
		t := scanner.Text()

		sumAllGroups += findSumGroups(t)
	}

	fmt.Printf("Sum of all groups %d", sumAllGroups)
}

// in: string
// out: sum of groups in string
func findSumGroups(s string) int {
	//fmt.Printf("input: %s\n", s)
	st := make(stack, 0)

	sumGroups := 0
	garbage := false
	neglectNext := false
	level := 1
	levelWhenUpping := level
	mem := []Paren{}

	for i, rune := range s {
		char := string(rune)

		if neglectNext == true {
			//fmt.Print("neglecting this char\n")
			neglectNext = false
			continue
		}

		switch {
		case char == "!":
			neglectNext = true
		case char == ">":
			garbage = false
		case char == "{" && !garbage:
			level--
			st = st.push(Paren{char, i, -1})
		case char == "}" && !garbage:

			peekVal, err := st.peek()
			if err != nil {
				continue
			} else if peekVal.char == "{" {
				level++
				var p Paren
				st, p, _ = st.pop()

				if level != levelWhenUpping {
					levelWhenUpping = level
					upOne(mem, p.index)
				}
				mem = append(mem, Paren{"", i, 1})

				//fmt.Printf("found group: \n\tmem: %+v\n\tstack: %+v\n\tlevel: %d\n", mem, st, level)
				if len(st) == 0 {
					sumGroups += sum(mem)
					//fmt.Printf("set away mem, sumGroups is %d\n", sumGroups)
					mem = []Paren{}
				}
			}
		case char == "<" && !garbage:
			garbage = true
		}
		i++
	}
	return sumGroups + sum(mem)
}

func upOne(pa []Paren, index int) {
	for i := range pa {
		if pa[i].index >= index {
			pa[i].value = pa[i].value + 1
		}
	}
}

func sum(pa []Paren) (sum int) {
	for i := range pa {
		sum += pa[i].value
	}
	return sum
}

// Barebones stack implementation
type Paren struct {
	char  string
	index int
	value int
}
type stack []Paren

func (s stack) push(p Paren) stack {
	return append(s, p)
}

func (s stack) pop() (stack, Paren, error) {
	l := len(s)
	if l < 1 {
		return nil, Paren{}, errors.New("stack is empty")
	}
	return s[:l-1], s[l-1], nil
}

func (s stack) peek() (Paren, error) {
	l := len(s)
	if l < 1 {
		return Paren{}, errors.New("stack is empty")
	}
	return s[l-1], nil
}
