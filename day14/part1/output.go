package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func calculateTotalLoad(matrix [][]byte) int {
	total := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'O' {
				k := i
				for k > 0 && matrix[k-1][j] != '#' && matrix[k-1][j] != 'O' {
					k--
				}
				if k != i {
					matrix[k][j] = matrix[i][j]
					matrix[i][j] = '.'
				}
				total += len(matrix) - k
			}
		}
	}
	return total
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
	var matrix [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []byte(line))
	}
	return calculateTotalLoad(matrix)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
