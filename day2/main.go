package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		ints := make([]int, len(numbers))
		for i, n := range numbers {
			ints[i], _ = strconv.Atoi(n)
		}
		result = append(result, ints)
	}

	return result
}

func isSafeLine(line []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]

		// Check if the difference is out of range
		if diff > 3 || diff < -3 {
			return false
		}
		// if two values are the same it is unsafe
		if line[i] == line[i+1] {
			return false
		}
		// Check if the sequence is not strictly increasing or decreasing
		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		}
	}
	// The line must be either strictly increasing or strictly decreasing
	return increasing || decreasing
}

func canBeMadeSafe(line []int) bool {
	for i := 0; i < len(line); i++ {
		newLine := append([]int{}, line[:i]...)
		newLine = append(newLine, line[i+1:]...)

		if isSafeLine(newLine) {
			return true
		}
	}
	return false
}

func main() {
	filename := "input.txt"
	lines := readInput(filename)

	safeLines := 0
	for _, line := range lines {
		if isSafeLine(line) || canBeMadeSafe(line) {
			safeLines++
		}
	}

	fmt.Println("Number of safe lines:", safeLines)
}
