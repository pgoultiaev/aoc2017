package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var listOfNodes = map[string]Node{}
var weightRgx = regexp.MustCompile(`\((.*?)\)`)

type Node struct {
	name      string
	weight    int
	count     int
	children  []string
	sumWeight int
}

func main() {
	readInput("7.txt")

	var bottomNodeName string
	for k, v := range listOfNodes {
		if v.count == 1 {
			bottomNodeName = k
			break
		}
	}
	root := listOfNodes[bottomNodeName]
	fmt.Printf("Bottom node name: %s\n", bottomNodeName) //bsfpjtc
	fmt.Println("Root weight:", getSumWeight(root))
	findUnbalancedNodes(root, 0)

}

func printListOfNodes() {
	for _, v := range listOfNodes {
		fmt.Printf("%+v\n", v)
	}
}

func getSumWeight(n Node) int {
	weight := n.weight

	for i := 0; i < len(n.children); i++ {
		weight += getSumWeight(listOfNodes[n.children[i]])
	}

	n.sumWeight = weight
	listOfNodes[n.name] = n

	return weight
}

func findUnbalancedNodes(n Node, diff int) {
	if len(n.children) > 2 {
		childWeights := make(map[int]int)

		for i := range n.children {
			childNode := listOfNodes[n.children[i]]
			childWeights[childNode.sumWeight]++
		}

		fmt.Printf("childWeights: %+v\n", childWeights)

		if len(childWeights) == 1 {
			fmt.Printf("Node %s with weight: %d is unbalanced with diff: %d, should be |%d|\n", n.name, n.weight, diff, n.weight-diff)
		} else {
			for value, count := range childWeights {
				if count == 1 {
					for j := range n.children {
						childNode := listOfNodes[n.children[j]]
						if childNode.sumWeight == value {
							fmt.Printf("Node %s is unbalanced with sumWeight %d and weight %d\n", childNode.name, childNode.sumWeight, childNode.weight)
							diff := getDiff(childWeights, value)
							findUnbalancedNodes(childNode, diff)
						}
					}
				}
			}
		}
	}
}

func getDiff(m map[int]int, v int) int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	if keys[0] == v {
		return v - keys[1]
	}
	return v - keys[0]
}

func findName(s string) string {
	i := strings.Index(s, "(")
	return s[0 : i-1]
}

func findChildren(s string) (na []string) {
	i := strings.Index(s, ">")
	if i == -1 {
		return
	}

	names := s[i+2:]
	namesArray := strings.Split(names, ", ")

	for _, e := range namesArray {
		n := listOfNodes[e]
		n.count++
		listOfNodes[e] = n

		na = append(na, e)
	}
	return na
}

func readInput(f string) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		weight, _ := strconv.Atoi(weightRgx.FindStringSubmatch(text)[1])
		name := findName(text)
		children := findChildren(text)

		n := Node{
			name:     name,
			weight:   weight,
			children: children,
			count:    listOfNodes[name].count + 1,
		}
		listOfNodes[name] = n
	}
}
