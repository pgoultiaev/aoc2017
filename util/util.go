package util

import (
	"sort"
	"strconv"
	"strings"
)

// Abs returns the absolute value of an int
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IsPermutation returns true if two strings are permutations of each other
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

// Equals returns true if two int arrays have exactly the same elements in the same order
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

// ReverseOrder Reverses order of an int array
func ReverseOrder(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// MakeRange Creates an int array with ints starting with min and ending with max
func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// ConvStringArrayToIntArray converts a string array to an int array
func ConvStringArrayToIntArray(sa []string) []int {
	var output []int
	for _, e := range sa {
		i, _ := strconv.Atoi(e)
		output = append(output, i)
	}
	return output
}

// ConvStringArrayToByteArray converts a string array to a byte array
func ConvStringArrayToByteArray(s []string) []byte {
	var output []byte
	for _, e := range s {
		i := []byte(e)
		output = append(output, i...)
	}
	return output
}

// ConvByteArrayToIntArray converts a byte array to an int array
func ConvByteArrayToIntArray(ba []byte) []int {
	var output []int
	for _, e := range ba {
		output = append(output, int(e))
	}
	return output
}

// Stripchars removes all characters in str that match any char in chr
func Stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
