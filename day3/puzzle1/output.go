package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Function to check if a number is adjacent to any symbol except for dots
func isAdjacentToSymbols(grid []string, start int, end int, row int) bool {
	// Check characters immediately before and after the number in the same row
	if start > 0 && grid[row][start-1] != '.' && grid[row][start-1] != ' ' {
		return true
	}
	if end < len(grid[row]) && grid[row][end] != '.' && grid[row][end] != ' ' {
		return true
	}

	// Check characters in the row above and below the number, if applicable
	if row > 0 {
		for i := start - 1; i <= end && i < len(grid[row-1]); i++ {
			if i >= 0 && grid[row-1][i] != '.' && grid[row-1][i] != ' ' {
				return true
			}
		}
	}
	if row < len(grid)-1 {
		for i := start - 1; i <= end && i < len(grid[row+1]); i++ {
			if i >= 0 && grid[row+1][i] != '.' && grid[row+1][i] != ' ' {
				return true
			}
		}
	}

	return false
}

// Function to process the file and sum the part numbers in each line
// A part number is a number in the line not adjacent to any symbol (just adjacent to dots even diagonally)
func processFile(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	var sumPartNumbers int = 0

	for row, line := range grid {
		start := -1
		for i, ch := range line {
			if unicode.IsDigit(ch) {
				if start == -1 {
					start = i
				}
			} else {
				if start != -1 {
					partNumberStr := line[start:i]
					if isAdjacentToSymbols(grid, start, i, row) {
						partNumber, err := strconv.Atoi(partNumberStr)
						if err == nil {
							sumPartNumbers += partNumber
						}
					}
					start = -1
				}
			}
		}
		if start != -1 && isAdjacentToSymbols(grid, start, len(line), row) {
			partNumber, err := strconv.Atoi(line[start:])
			if err == nil {
				sumPartNumbers += partNumber
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	return sumPartNumbers
}

func main() {
	// Call the function with the file path and print the result
	totalSum := processFile("../input.txt")
	fmt.Println(totalSum)
}
