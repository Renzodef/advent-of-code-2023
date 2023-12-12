package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func b2i(b bool) int {
	if b {
		return 1
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
	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}
	numRows := len(grid)
	numCols := len(grid[0])
	rowHasHash := make([]bool, numRows)
	colHasHash := make([]bool, numCols)
	hashesStartingPosition := make([]Point, 0)
	for i, row := range grid {
		for j, val := range row {
			if val == "#" {
				hashesStartingPosition = append(hashesStartingPosition, Point{i, j})
				rowHasHash[i] = true
				colHasHash[j] = true
			}
		}
	}
	distortionFactor := 1000000
	rowShift := make([]int, numRows)
	colShift := make([]int, numCols)
	for i := 1; i < numRows; i++ {
		rowShift[i] = rowShift[i-1] + (distortionFactor-1)*b2i(!rowHasHash[i-1])
	}
	for j := 1; j < numCols; j++ {
		colShift[j] = colShift[j-1] + (distortionFactor-1)*b2i(!colHasHash[j-1])
	}
	var hashesEndingPosition []Point
	for _, hash := range hashesStartingPosition {
		newX := hash.X + rowShift[hash.X]
		newY := hash.Y + colShift[hash.Y]
		hashesEndingPosition = append(hashesEndingPosition, Point{newX, newY})
	}
	var pairs [][]Point
	for i := 0; i < len(hashesEndingPosition); i++ {
		for j := i + 1; j < len(hashesEndingPosition); j++ {
			pair := []Point{hashesEndingPosition[i], hashesEndingPosition[j]}
			pairs = append(pairs, pair)
		}
	}
	sumOfShortestPath := 0
	for _, pair := range pairs {
		sumOfShortestPath += manhattanDistance(pair[0], pair[1])
	}
	return sumOfShortestPath
}

func main() {
	result := processFile("../input.txt")
	fmt.Println(result)
}
