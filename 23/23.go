package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("23.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := [][]string{}
	for scanner.Scan() {
		slicedString := strings.Split(scanner.Text(), " ")
		instructions = append(instructions, slicedString)
	}

	println(solve(instructions))
}

// Part one
func solve(instructions [][]string) (mulInvoked int) {
	registers := map[string]int{}
	i := 0
	for i < len(instructions) {
		instr := instructions[i]
		fmt.Printf("instr: %+v\n", instr)
		x := instr[1]

		switch instr[0] {
		case "set":
			registers[x] = getInstructionValue(instr[2], registers)
		case "sub":
			registers[x] -= getInstructionValue(instr[2], registers)
		case "mul":
			registers[x] *= getInstructionValue(instr[2], registers)
			mulInvoked++
		case "jnz":
			if getInstructionValue(x, registers) != 0 {
				i += getInstructionValue(instr[2], registers) - 1
			}
		}
		i++
	}

	return mulInvoked
}

func getInstructionValue(y string, registers map[string]int) int {
	setVal, err := strconv.Atoi(y)
	if err != nil {
		return registers[y]
	}
	return setVal
}
