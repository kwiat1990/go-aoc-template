package utils

import (
	"strconv"
	"strings"
)

// ReadLines splits input by newlines and returns non-empty lines
func ReadLines(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var result []string
	for _, line := range lines {
		if line != "" {
			result = append(result, line)
		}
	}
	return result
}

// ParseInt converts string to int, panics on error (for AoC convenience)
func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// Sum returns the sum of all integers in the slice
func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// Abs returns absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
