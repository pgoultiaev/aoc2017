package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var dance = []string{
	"a", "b", "c", "d", "e", "f", "g", "h",
	"i", "j", "k", "l", "m", "n", "o", "p",
}

func main() {
	buf, err := ioutil.ReadFile("16.txt")
	if err != nil {
		panic(err)
	}

	s := string(buf)

	splitStrings := strings.Split(s, ",")
	danceOrder := solve(dance, splitStrings)
	println(danceOrder)
}

func solve(dance, dancemoves []string) string {
	for _, move := range dancemoves {
		switch move[0] {
		case 's':
			xprograms, _ := strconv.Atoi(move[1:])
			tail := dance[len(dance)-xprograms:]
			head := dance[:len(dance)-xprograms]
			dance = append(tail, head...)
			//fmt.Printf("SWITCH %s, xprograms: %d, tail:\t%+v\nhead:\t%+v\ndance: %+v", move, xprograms, tail, head, dance)

		case 'x':
			instr := move[1:]
			progs := strings.Split(instr, "/")
			index1, _ := strconv.Atoi(progs[0])
			index2, _ := strconv.Atoi(progs[1])
			prog1 := dance[index1]
			dance[index1] = dance[index2]
			dance[index2] = prog1
			//fmt.Printf("EXCHANGE %s, progs: [%d / %d]\ndance: %+v", instr, index1, index2, dance)

		case 'p':
			instr := move[1:]
			progs := strings.Split(instr, "/")

			var index1, index2 int
			for i := range dance {
				if dance[i] == progs[0] {
					index1 = i
				}
				if dance[i] == progs[1] {
					index2 = i
				}
			}

			prog1 := dance[index1]
			dance[index1] = dance[index2]
			dance[index2] = prog1

			//fmt.Printf("PARTNER %s, progs: [%s / %s]\n, dance: %+v", instr, progs[0], progs[1], dance)
		}
	}
	return strings.Join(dance, "")
}
