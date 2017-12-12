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

	set0 := nodes["0"]
	for key := range nodes["0"] {
		traverse(key, set0, nodes, map[string]bool{})
	}

	fmt.Printf("Part one: %d\n", len(set0))
}

func traverse(n string, set0 map[string]bool, nodes map[string]map[string]bool, visited map[string]bool) {
	if len(visited) == len(nodes) {
		return
	}
	for k := range nodes[n] {
		set0[k] = true

		if visited[k] != true {
			visited[k] = true
			traverse(k, set0, nodes, visited)
		}
	}
}
