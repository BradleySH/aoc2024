package main

import (
	"bufio"
	"fmt"
	"os"
)

// read the input file and return how many times XMAS appears.  It can be horizontal, vertical,diagonal, written backwards, or even overlapping other words.  any irreleant chaacter should be replaced with a .

func readInputAndReplace(input string) string {
	// read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return 

}

func main() {
	fmt.Println("Hello, World!")
}
