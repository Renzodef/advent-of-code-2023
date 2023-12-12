package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func sumGearRatios(partNumbers map[int]map[int]int, gearPositions map[int][]int) int {
	var sum = 0
	for row, gears := range gearPositions {
		for _, gearPos := range gears {
			adjacentParts := make([]int, 0)
			for start, part := range partNumbers[row] {
				var end = start + len(strconv.Itoa(part)) - 1
				if start-1 <= gearPos && end+1 >= gearPos {
					adjacentParts = append(adjacentParts, part)
				}
			}
			for start, part := range partNumbers[row-1] {
				var end = start + len(strconv.Itoa(part)) - 1
				if start-1 <= gearPos && end+1 >= gearPos {
					adjacentParts = append(adjacentParts, part)
				}
			}
			for start, part := range partNumbers[row+1] {
				var end = start + len(strconv.Itoa(part)) - 1
				if start-1 <= gearPos && end+1 >= gearPos {
					adjacentParts = append(adjacentParts, part)
				}
			}
			if len(adjacentParts) == 2 {
				sum += adjacentParts[0] * adjacentParts[1]
			}
		}
	}
	return sum
}

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
	partNumbers := make(map[int]map[int]int)
	gearPositions := make(map[int][]int)
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
				if ch == '*' {
					gearPositions[row] = append(gearPositions[row], i)
				}
				if start != -1 {
					partNumberStr := line[start:i]
					if isAdjacentToSymbols(grid, start, i, row) {
						partNumber, err := strconv.Atoi(partNumberStr)
						if err == nil {
							if partNumbers[row] == nil {
								partNumbers[row] = make(map[int]int)
							}
							partNumbers[row][start] = partNumber
						}
					}
					start = -1
				}
			}
		}
		if start != -1 && isAdjacentToSymbols(grid, start, len(line), row) {
			partNumber, err := strconv.Atoi(line[start:])
			if err == nil {
				if partNumbers[row] == nil {
					partNumbers[row] = make(map[int]int)
				}
				partNumbers[row][start] = partNumber
			}
		}
	}
	return sumGearRatios(partNumbers, gearPositions)
}

func main() {
	result := processFile("../input.txt")
	fmt.Println(result)
}
