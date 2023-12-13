package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func transpose(pattern []string, n int) []string {
	if n == 0 {
		return []string{}
	}
	m := len(pattern[0])
	transposed := make([]string, m)
	for i := 0; i < m; i++ {
		var newRow []byte
		for j := 0; j < n; j++ {
			newRow = append(newRow, pattern[j][i])
		}
		transposed[i] = string(newRow)
	}
	return transposed
}

func verifyVerticalSymmetry(pattern []string, n int) (int, bool) {
	transposedPattern := transpose(pattern, n)
	return verifyHorizontalSymmetry(transposedPattern, len(transposedPattern))
}

func verifyHorizontalSymmetry(pattern []string, n int) (int, bool) {
	for i := 0; i < n-1; i++ {
		if pattern[i] == pattern[i+1] {
			symmetric := true
			for j := 1; j <= i && i+j+1 < n; j++ {
				if pattern[i-j] != pattern[i+j+1] {
					symmetric = false
					break
				}
			}
			if symmetric {
				return i, true
			}
		}
	}
	return 0, false
}

func findReflection(pattern []string) int {
	n := len(pattern)
	i, ok := verifyHorizontalSymmetry(pattern, n)
	if ok {
		return (i + 1) * 100
	}
	i, ok = verifyVerticalSymmetry(pattern, n)
	if ok {
		return i + 1
	}
	return 0
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
	total := 0
	pattern := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total += findReflection(pattern)
			pattern = pattern[:0]
		} else {
			pattern = append(pattern, line)
		}
	}
	if len(pattern) > 0 {
		total += findReflection(pattern)
	}
	return total
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
