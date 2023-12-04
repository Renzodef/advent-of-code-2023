// https://adventofcode.com/2023/day/1#part1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Function to concatenate the first and last digits of a line
func sumFirstLastDigit(line string) int {
	firstDigit, lastDigit := -1, -1

	for _, char := range line {
		if unicode.IsDigit(char) {
			firstDigit, _ = strconv.Atoi(string(char))
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			lastDigit, _ = strconv.Atoi(string(line[i]))
			break
		}
	}

	if firstDigit == -1 || lastDigit == -1 {
		return 0
	}

	concatenated := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
	result, err := strconv.Atoi(concatenated)
	if err != nil {
		fmt.Println("Error converting concatenated string to int:", err)
		return 0
	}

	return result
}

// Function to process the file and sum the concatented first and last digits of each line
func processFile(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	var count int = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		count += sumFirstLastDigit(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	return count
}

func main() {
	totalSum := processFile("../input.txt")
	fmt.Println(totalSum)
}
