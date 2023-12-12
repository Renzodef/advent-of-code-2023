package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

func isAdjacentToSymbols(grid []string, start int, end int, row int) bool {
	if start > 0 && grid[row][start-1] != '.' && grid[row][start-1] != ' ' {
		return true
	}
	if end < len(grid[row]) && grid[row][end] != '.' && grid[row][end] != ' ' {
		return true
	}
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

func processFile(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			return
		}
	}(file)
	scanner := bufio.NewScanner(file)
	var grid []string
	var sumPartNumbers = 0
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
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
	return sumPartNumbers
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
