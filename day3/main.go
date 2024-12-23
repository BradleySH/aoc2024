package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// read the input.txt file and parse the input and look for only the lines that contain the word "mul"
// then print the lines that contain the word "mul"
func readInputAndCalculate(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	input := string(file)

	// Regex to match do(), don't(), and valid mul(X,Y) instructions
	reInstructions := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),\s*(\d{1,3})\)`)

	// Find all matches in the input
	matches := reInstructions.FindAllStringSubmatch(input, -1)

	var total int
	isEnabled := true // Start with mul instructions enabled

	for _, match := range matches {
		// Check for do() and don't()
		if match[0] == "do()" {
			isEnabled = true
			fmt.Println("do() detected: mul instructions enabled")
		} else if match[0] == "don't()" {
			isEnabled = false
			fmt.Println("don't() detected: mul instructions disabled")
		} else if len(match) == 3 && isEnabled {
			// Process valid mul(X,Y) instructions when enabled
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			product := num1 * num2
			total += product

			// Debugging output for each valid match
			fmt.Printf("Matched mul(%d,%d): %d\n", num1, num2, product)
		}
	}

	return total
}

func main() {
	result := readInputAndCalculate("input.txt")

	fmt.Println("Total Result:", result)
}

// /mul\([a-zA-Z0-9]+,[a-zA-Z0-9]+\)/gm
