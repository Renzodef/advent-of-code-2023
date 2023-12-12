package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func canBeValid(combination string, groups []int) bool {
	count, groupIndex := 0, 0
	for _, c := range combination {
		if c == '#' {
			count++
			if groupIndex < len(groups) && count > groups[groupIndex] {
				return false
			}
		} else if c == '.' {
			if count > 0 {
				if groupIndex >= len(groups) || count != groups[groupIndex] {
					return false
				}
				groupIndex++
				count = 0
			}
		}
	}
	return true
}

func isValid(combination string, groups []int) bool {
	index, count := 0, 0
	for _, c := range combination + "." {
		if c == '#' {
			count++
		} else {
			if count > 0 {
				if index >= len(groups) || count != groups[index] {
					return false
				}
				index++
				count = 0
			}
		}
	}
	return index == len(groups)
}

func generateCombinations(line string, pos int, current string, groups []int) int {
	if pos == len(line) {
		if isValid(current, groups) {
			return 1
		}
		return 0
	}
	if !canBeValid(current, groups) {
		return 0
	}
	count := 0
	switch line[pos] {
	case '.':
		count += generateCombinations(line, pos+1, current+".", groups)
	case '#':
		count += generateCombinations(line, pos+1, current+"#", groups)
	case '?':
		count += generateCombinations(line, pos+1, current+".", groups)
		count += generateCombinations(line, pos+1, current+"#", groups)
	}
	return count
}

func countValidArrangements(line string) int {
	parts := strings.Fields(line)
	input := parts[0]
	groupParts := strings.Split(parts[1], ",")
	groups := make([]int, len(groupParts))
	for i, gp := range groupParts {
		group, err := strconv.Atoi(gp)
		if err != nil {
			fmt.Println("Error converting group to integer:", err)
			return 0
		}
		groups[i] = group
	}
	return generateCombinations(input, 0, "", groups)
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
	sumOfArrangements := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sumOfArrangements += countValidArrangements(scanner.Text())
	}
	return sumOfArrangements
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
