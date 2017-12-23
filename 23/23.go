package main

import (
	"bufio"
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

	mulInvoked := solve(instructions)
	println(mulInvoked)

	println(solve2())
}

func solve2() int {
	h := 0
	b := (79 * 100) + 100000
	c := b + 17000 // sub b -17 1000 times

	g := b - c
	for g != 0 {
		f0 := false

		for d := 2; (d * d) < b; d++ {
			if (b % d) == 0 {
				f0 = true
				break
			}
		}
		if f0 {
			h++
		}
		g = b - c
		b += 17
	}

	return h
}

// Part one
func solve(instructions [][]string) (mulInvoked int) {
	registers := map[string]int{}
	i := 0
	for i < len(instructions) {
		instr := instructions[i]
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
