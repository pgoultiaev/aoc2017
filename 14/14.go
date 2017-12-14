package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := "hfdlxzhv"

	println(solve(input))
}

func solve(s string) (gridOnes int) {
	i := 0
	for i < 128 {
		sHash := fmt.Sprintf("%s-%d", s, i)
		hash := KnotHashDay10(sHash)
		binVal, err := getBinVal(hash)
		if err != nil {
			log.Fatal(err)
			return
		}

		// fmt.Printf("row hash \n\t%s\n\t%s\n", sHash, binVal)
		gridOnes += strings.Count(binVal, "1")
		i++
	}
	return gridOnes
}

func getBinVal(s string) (string, error) {
	var buffer bytes.Buffer
	sa := strings.Split(s, "")
	for _, char := range sa {
		ui, err := strconv.ParseUint(char, 16, 64)
		if err != nil {
			return "", err
		}

		// base 2, zero padded, 4 characters
		buffer.WriteString(fmt.Sprintf("%04b", ui))
	}
	return buffer.String(), nil
}
