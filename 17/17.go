package main

func main() {
	input := 348

	partOne, valAfter0 := solve(2017, input)
	println(partOne)
	println(valAfter0)

	partTwo := solve2(50000000, input)
	println(partTwo)
}

func solve2(lastValWritten, input int) (valueAfter0 int) {
	i := 1
	currentPos := 0
	lengthBuffer := 1
	for i <= lastValWritten {
		currentPos = (currentPos + input) % lengthBuffer

		if currentPos == 0 {
			valueAfter0 = i
		}
		lengthBuffer++
		currentPos++
		i++
	}

	return valueAfter0
}

func solve(lastValWritten, input int) (int, int) {
	buffer := []int{0}
	i := 1
	currentPos := 0
	for i <= lastValWritten {
		currentPos = (currentPos + input) % len(buffer)

		if currentPos == len(buffer)-1 {
			buffer = append(buffer, i)
		} else {
			buffer = append(buffer, 0)
			copy(buffer[currentPos+1:], buffer[currentPos:])
			buffer[currentPos+1] = i
		}

		currentPos++
		i++
	}

	return buffer[currentPos+1], buffer[1]
}
