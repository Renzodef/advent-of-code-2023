package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isSpecular(matrix [][]byte, pivot int) bool {
	for _, row := range matrix {
		left := row[:pivot]
		right := row[pivot:]
		L := minInt(len(left), len(right))
		for i := 0; i < L; i++ {
			if left[len(left)-1-i] != right[i] {
				return false
			}
		}
	}
	return true
}

func transpose(matrix [][]byte) [][]byte {
	transposed := make([][]byte, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]byte, len(matrix))
		for j := range matrix {
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}

func toggleCharacter(c byte) byte {
	if c == '#' {
		return '.'
	}
	return '#'
}

func findReflection(matrix [][]byte, ignore *int) int {
	for vertical := 0; vertical <= 1; vertical++ {
		if vertical == 1 {
			matrix = transpose(matrix)
		}
		for x := 1; x < len(matrix[0]); x++ {
			if isSpecular(matrix, x) {
				res := x
				if vertical == 1 {
					res *= 100
				}
				if ignore == nil || res != *ignore {
					return res
				}
			}
		}
	}
	return 0
}

func smudge(matrix [][]byte) int {
	oldReflection := findReflection(matrix, nil)
	for y, row := range matrix {
		for x := range row {
			matrix[y][x] = toggleCharacter(matrix[y][x])
			res := findReflection(matrix, &oldReflection)
			matrix[y][x] = toggleCharacter(matrix[y][x])
			if res > 0 {
				return res
			}
		}
	}
	panic("No smudge found")
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
	var matrix [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total += smudge(matrix)
			matrix = [][]byte{}
		} else {
			matrix = append(matrix, []byte(line))
		}
	}
	total += smudge(matrix)
	return total
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
