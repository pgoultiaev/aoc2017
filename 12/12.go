package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nodes := map[string]map[string]bool{}

	for scanner.Scan() {
		link := scanner.Text()

		i := strings.Index(link, "<->")
		keyString := link[0 : i-1]

		valStrings := link[i+4:]
		valStringsSplit := strings.Split(valStrings, ", ")

		//fmt.Printf("found key: %s, val %+v\n", keyString, valStringsSplit)
		nodesToAdd := append(valStringsSplit, keyString)

		for _, val := range nodesToAdd {
			for _, node := range nodesToAdd {
				child, ok := nodes[val]

				if !ok {
					child = map[string]bool{}
					nodes[val] = child
				}
				nodes[val][node] = true
			}
		}
	}

	set0 := traverse("0", map[string]bool{}, nodes, map[string]bool{})
	independentGroups := getIndependentGroups(nodes)

	fmt.Printf("Part one: %d\n", len(set0))
	fmt.Printf("Part two: %d\n", independentGroups)
}

func getIndependentGroups(nodes map[string]map[string]bool) (groupCount int) {
	nodeList := nodes

	for k := range nodeList {
		set := traverse(k, map[string]bool{}, nodes, map[string]bool{})
		for k := range set {
			delete(nodeList, k)
		}
		groupCount++
	}
	return groupCount
}

func traverse(n string, set0 map[string]bool, nodes map[string]map[string]bool, visited map[string]bool) map[string]bool {
	if len(visited) == len(nodes) {
		return set0
	}
	for k := range nodes[n] {
		set0[k] = true

		if visited[k] != true {
			visited[k] = true
			traverse(k, set0, nodes, visited)
		}
	}
	return set0
}
