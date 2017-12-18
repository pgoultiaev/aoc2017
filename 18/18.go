package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//registers := map[string]int{}
	instructions := [][]string{}
	for scanner.Scan() {
		slicedString := strings.Split(scanner.Text(), " ")
		instructions = append(instructions, slicedString)
	}

	println(solve(instructions))
}

func solve(instructions [][]string) (lastNonNilFreq int) {
	registers := map[string]int{}

	i := 0
	for i < len(instructions) {
		instr := instructions[i]
		switch instr[0] {
		case "snd":
			// fmt.Printf("instruction: %+v\n", instr)
			lastNonNilFreq = registers[instr[1]]

		case "set":
			// fmt.Printf("instruction: %+v\n", instr)
			setVal, err := strconv.Atoi(instr[2])
			if err != nil {
				registers[instr[1]] = registers[instr[2]]
			} else {
				registers[instr[1]] = setVal
			}

		case "add":
			// fmt.Printf("instruction: %+v\n", instr)
			addVal, err := strconv.Atoi(instr[2])
			if err != nil {
				registers[instr[1]] += registers[instr[2]]
			} else {
				registers[instr[1]] += addVal
			}

		case "mul":
			// fmt.Printf("instruction: %+v\n", instr)
			mulVal, err := strconv.Atoi(instr[2])
			if err != nil {
				registers[instr[1]] *= registers[instr[2]]
			} else {
				registers[instr[1]] *= mulVal
			}
		case "mod":
			// fmt.Printf("instruction: %+v\n", instr)
			modVal, err := strconv.Atoi(instr[2])
			if err != nil {
				registers[instr[1]] %= registers[instr[2]]
			} else {
				registers[instr[1]] %= modVal
			}
		case "rcv":
			// fmt.Printf("instruction: %+v\n", instr)
			x, ok := registers[instr[1]]
			if lastNonNilFreq != 0 {
				return lastNonNilFreq
			}
			if ok && x != 0 {
				registers[instr[1]] = lastNonNilFreq
			}

		case "jgz":
			// fmt.Printf("instruction: %+v\n", instr)
			x, ok := registers[instr[1]]
			if ok && x > 0 {
				skipVal, _ := strconv.Atoi(instr[2])
				println(skipVal)
				i += skipVal - 1
			}
		}

		i++
		//fmt.Printf("i: %d, registers: %+v\n", i, registers)
	}

	return lastNonNilFreq
}
