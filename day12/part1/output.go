package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func replaceAtIndex(in string, i int, r rune) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func calculateRunLengths(arrangement string) []int {
	var runLengths []int
	count := 0
	for _, ch := range arrangement {
		if ch == '#' {
			count++
		} else if count > 0 {
			runLengths = append(runLengths, count)
			count = 0
		}
	}
	if count > 0 {
		runLengths = append(runLengths, count)
	}
	return runLengths
}

func isValidArrangement(arrangement string, groups []int) bool {
	runLengths := calculateRunLengths(arrangement)
	if len(runLengths) != len(groups) {
		return false
	}
	for i, length := range runLengths {
		if length != groups[i] {
			return false
		}
	}
	return true
}

func countValidArrangements(index int, arrangement string, unknownIndices []int, groups []int) int {
	if index == len(unknownIndices) {
		if isValidArrangement(arrangement, groups) {
			return 1
		}
		return 0
	}

	count := 0
	i := unknownIndices[index]
	count += countValidArrangements(index+1, replaceAtIndex(arrangement, i, '.'), unknownIndices, groups)
	count += countValidArrangements(index+1, replaceAtIndex(arrangement, i, '#'), unknownIndices, groups)

	return count
}

func findUnknownIndices(input string) []int {
	var indices []int
	for i, ch := range input {
		if ch == '?' {
			indices = append(indices, i)
		}
	}
	return indices
}

func countArrangements(input string, groups []int) int {
	unknownIndices := findUnknownIndices(input)
	return countValidArrangements(0, input, unknownIndices, groups)
}

func countArrangementsOfLine(line string) int {
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
	return countArrangements(input, groups)
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
		sumOfArrangements += countArrangementsOfLine(scanner.Text())
	}
	return sumOfArrangements
}

func main() {
	result := processFile("../input.txt")
	fmt.Println(result)
}
