package main

func main() {
	input := 348

	println(solve(2017, input))
}

func solve(lastValWritten, input int) int {
	buffer := []int{0}
	i := 1
	currentPos := 0
	for i <= 2017 {
		currentPos = (currentPos + input) % len(buffer)
		//fmt.Printf("currentPosition before: %d\n", currentPos)

		if currentPos == len(buffer)-1 {
			buffer = append(buffer, i)
		} else {
			buffer = append(buffer, 0)
			copy(buffer[currentPos+1:], buffer[currentPos:])
			buffer[currentPos+1] = i
		}
		currentPos++

		//fmt.Printf("buffer: %+v\n", buffer)
		//fmt.Printf("currentPosition after: %d\n", currentPos)
		i++
	}

	return buffer[currentPos+1]
}
