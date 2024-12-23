package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if len(values) != 2 {
			return nil, nil, fmt.Errorf("invalid line: %s", line)
		}

		leftValue, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid left value: %s", values[0])
		}
		rightValue, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid right value: %s", values[1])
		}
		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading input: %w", err)
	}

	return left, right, nil
}

func similarityScore(left []int, right []int) int {
	// count the number of times each value appears in the right column
	rightCount := make(map[int]int)
	for _, num := range right {
		rightCount[num]++
	}

	score := 0
	for _, num := range left {
		score += num * rightCount[num]
	}
	return score
}

func main() {
	filename := "input.txt"

	left, right, err := readInput(filename)
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		return
	}

	score := similarityScore(left, right)
	fmt.Printf("Similarity score: %d\n", score)
}
