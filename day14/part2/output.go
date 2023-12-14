package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func stateToString(matrix [][]byte) string {
	var builder strings.Builder
	for _, row := range matrix {
		builder.Write(row)
		builder.WriteByte('\n')
	}
	return builder.String()
}

func moveWest(matrix [][]byte) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 'O' {
				k := j
				for k > 0 && matrix[i][k-1] != '#' && matrix[i][k-1] != 'O' {
					k--
				}
				if k != j {
					matrix[i][k] = matrix[i][j]
					matrix[i][j] = '.'
				}
			}
		}
	}
}

func moveSouth(matrix [][]byte) {
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := range matrix[i] {
			if matrix[i][j] == 'O' {
				k := i
				for k < len(matrix)-1 && matrix[k+1][j] != '#' && matrix[k+1][j] != 'O' {
					k++
				}
				if k != i {
					matrix[k][j] = matrix[i][j]
					matrix[i][j] = '.'
				}
			}
		}
	}
}

func moveEast(matrix [][]byte) {
	for i := range matrix {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if matrix[i][j] == 'O' {
				k := j
				for k < len(matrix[i])-1 && matrix[i][k+1] != '#' && matrix[i][k+1] != 'O' {
					k++
				}
				if k != j {
					matrix[i][k] = matrix[i][j]
					matrix[i][j] = '.'
				}
			}
		}
	}
}

func moveNorth(matrix [][]byte) {
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
			}
		}
	}
}

func calculateTotalLoad(matrix [][]byte, cycles int) int {
	seenStates := make(map[string]int)
	cycleLength := 0
	cycleStart := 0
	for i := 1; i <= cycles; i++ {
		moveNorth(matrix)
		moveWest(matrix)
		moveSouth(matrix)
		moveEast(matrix)
		currentState := stateToString(matrix)
		if start, found := seenStates[currentState]; found {
			cycleLength = i - start
			cycleStart = start
			break
		}
		seenStates[currentState] = i
	}
	if cycleLength > 0 {
		remainingCycles := (cycles - cycleStart) % cycleLength
		for i := 0; i < remainingCycles; i++ {
			moveNorth(matrix)
			moveWest(matrix)
			moveSouth(matrix)
			moveEast(matrix)
		}
	}
	total := 0
	for x := range matrix {
		for y := range matrix[x] {
			if matrix[x][y] == 'O' {
				total += len(matrix) - x
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
	return calculateTotalLoad(matrix, 1000000000)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
