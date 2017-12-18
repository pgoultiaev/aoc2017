package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

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
	partTwo(instructions)
}

func partTwo(instructions [][]string) (sends0 int, sends1 int) {
	ch0 := make(chan int, 100)
	ch1 := make(chan int, 100)
	register0 := map[string]int{"p": 0}
	register1 := map[string]int{"p": 1}

	waitGroup.Add(2)
	go func() {
		sends0 = solve2(instructions, ch0, ch1, register0, 0)
	}()
	go func() {
		sends1 = solve2(instructions, ch1, ch0, register1, 1)
	}()
	waitGroup.Wait()

	return sends0, sends1
}

// Part two
func solve2(instructions [][]string, chRcv, chSend chan int, registers map[string]int, id int) (timesSent int) {
	i := 0
	for i < len(instructions) {
		instr := instructions[i]
		x := instr[1]

		switch instr[0] {
		case "snd":
			sndVal, err := strconv.Atoi(x)
			if err != nil {
				chSend <- registers[x]
			} else {
				chSend <- sndVal
			}
			timesSent++
		case "rcv":
			rcvVal, timeout := readChannelWithTimeout(chRcv)
			if timeout {
				fmt.Printf("id: %d, times sent before deadlock: %d\n", id, timesSent)
				waitGroup.Done()
				return timesSent
			}
			registers[x] = rcvVal
		case "set":
			registers[x] = getInstructionValue(instr[2], registers)
		case "add":
			registers[x] += getInstructionValue(instr[2], registers)
		case "mul":
			registers[x] *= getInstructionValue(instr[2], registers)
		case "mod":
			registers[x] %= getInstructionValue(instr[2], registers)
		case "jgz":
			if getInstructionValue(x, registers) > 0 {
				i += getInstructionValue(instr[2], registers) - 1
			}
		}
		i++
	}

	fmt.Printf("id: %d, times sent before deadlock: %d\n", id, timesSent)
	waitGroup.Done()
	return timesSent
}

func readChannelWithTimeout(c chan int) (int, bool) {
	var rcvVal int
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()
	select {
	case rcvVal = <-c:
		return rcvVal, false
	case <-timeout:
		return rcvVal, true
	}
}

// Part one
func solve(instructions [][]string) (lastNonNilFreq int) {
	registers := map[string]int{}
	i := 0
	for i < len(instructions) {
		instr := instructions[i]
		x := instr[1]

		switch instr[0] {
		case "snd":
			lastNonNilFreq = registers[x]
		case "rcv":
			valX, _ := registers[x]
			if lastNonNilFreq != 0 {
				return lastNonNilFreq
			}
			if valX != 0 {
				registers[x] = lastNonNilFreq
			}
		case "set":
			registers[x] = getInstructionValue(instr[2], registers)
		case "add":
			registers[x] += getInstructionValue(instr[2], registers)
		case "mul":
			registers[x] *= getInstructionValue(instr[2], registers)
		case "mod":
			registers[x] %= getInstructionValue(instr[2], registers)
		case "jgz":
			if getInstructionValue(x, registers) > 0 {
				i += getInstructionValue(instr[2], registers) - 1
			}
		}
		i++
	}

	return lastNonNilFreq
}

func getInstructionValue(y string, registers map[string]int) int {
	setVal, err := strconv.Atoi(y)
	if err != nil {
		return registers[y]
	}
	return setVal
}
