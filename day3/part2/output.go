// https://adventofcode.com/2023/day/3#part2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Function to find the gears and sum their gear ratiosfunc sumGearRatios(grid []string, partNumbers map[int]map[int]int) int {
func sumGearRatios(partNumbers map[int]map[int]int, gearPositions map[int][]int) int {
	var sum int = 0

	for row, gears := range gearPositions {
		for _, gearPos := range gears {
			adjacentParts := make([]int, 0)

			for start, part := range partNumbers[row] {
				var end int = start + len(strconv.Itoa(part)) - 1
				if start-1 <= gearPos && end+1 >= gearPos {
					adjacentParts = append(adjacentParts, part)
				}
			}

			for start, part := range partNumbers[row-1] {
				var end int = start + len(strconv.Itoa(part)) - 1
				if start-1 <= gearPos && end+1 >= gearPos {
					adjacentParts = append(adjacentParts, part)
				}
			}

			for start, part := range partNumbers[row+1] {
				var end int = start + len(strconv.Itoa(part)) - 1
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

// Function to sum the gear ratios of all the gears
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

// Function to process the file and sum the gear ratios
// A gear is any * symbol that is adjacent to exactly two part numbers
// Its gear ratio is the result of multiplying those two numbers together
func processFile(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

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

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	return sumGearRatios(partNumbers, gearPositions)
}

func main() {
	totalSum := processFile("../input.txt")
	fmt.Println(totalSum)
}
