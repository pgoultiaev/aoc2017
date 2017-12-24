package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/pgoultiaev/aoc2017/util"
)

type Component struct {
	port1, port2 int
	inUse        bool
}

var maxBridgeStrength int
var longestBridgeStrength int
var maxBridgeLength int

func main() {
	println(solve("24.txt"))
	println(solve2("24.txt"))
}

// Part two
func solve2(filename string) int {
	components := readInput(filename)
	return buildBridge2(components, 0, 0, 0)
}

func buildBridge2(components []Component, bridgeStrength, bridgeLength, currentPort int) int {
	if bridgeLength > maxBridgeLength || bridgeLength == maxBridgeLength && bridgeStrength > longestBridgeStrength {
		maxBridgeLength = bridgeLength
		longestBridgeStrength = bridgeStrength
	}
	for i, component := range components {
		if component.inUse {
			continue
		}
		if component.port1 == currentPort {
			components[i].inUse = true
			buildBridge2(components, bridgeStrength+component.port1+component.port2, bridgeLength+1, component.port2)
			components[i].inUse = false
		} else if component.port2 == currentPort {
			components[i].inUse = true
			buildBridge2(components, bridgeStrength+component.port1+component.port2, bridgeLength+1, component.port1)
			components[i].inUse = false
		}
	}

	return longestBridgeStrength
}

// Part one
func solve(filename string) int {
	components := readInput(filename)
	return buildBridge(components, 0, 0)
}

func buildBridge(components []Component, bridgeStrength, currentPort int) int {
	if bridgeStrength > maxBridgeStrength {
		maxBridgeStrength = bridgeStrength
	}
	for i, component := range components {
		if component.inUse {
			continue
		}
		if component.port1 == currentPort {
			components[i].inUse = true
			buildBridge(components, bridgeStrength+component.port1+component.port2, component.port2)
			components[i].inUse = false
		} else if component.port2 == currentPort {
			components[i].inUse = true
			buildBridge(components, bridgeStrength+component.port1+component.port2, component.port1)
			components[i].inUse = false
		}
	}

	return maxBridgeStrength
}

func readInput(filename string) []Component {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	components := []Component{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "/")
		lineInts := util.ConvStringArrayToIntArray(line)

		c := Component{lineInts[0], lineInts[1], false}
		components = append(components, c)
	}

	return components
}
