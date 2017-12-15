package util

import (
	"sort"
	"strconv"
	"strings"
)

func ConvStringArrayToIntArray(sa []string) []int {
	var output []int
	for _, e := range sa {
		i, _ := strconv.Atoi(e)
		output = append(output, i)
	}
	return output
}

// Why does golang not have a Abs(int) function in the math package?
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1split := strings.Split(s1, "")
	sort.Strings(s1split)
	s2split := strings.Split(s2, "")
	sort.Strings(s2split)

	for i := range s1split {
		if s1split[i] != s2split[i] {
			return false
		}
	}
	return true
}

func Equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	//fmt.Printf("equals: %+v == %+v\n", a, b)
	return true
}

func ReverseOrder(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func ConvStringArrayToByteArray(s []string) []byte {
	var output []byte
	for _, e := range s {
		i := []byte(e)
		output = append(output, i...)
	}
	return output
}

func ConvByteArrayToIntArray(ba []byte) []int {
	var output []int
	for _, e := range ba {
		output = append(output, int(e))
	}
	return output
}
