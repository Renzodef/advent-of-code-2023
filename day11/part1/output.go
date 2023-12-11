package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Point struct to store X (row) and Y (column)
type Point struct {
	X, Y int
}

// Function to calculate the minimum distance (Manhattan distance) between two points
func manhattanDistance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Function to process the file and return the sum of length of the shortest path between every pair of galaxies
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

	rowsNotToExpand := make([]bool, len(grid))
	colsNotToExpand := make([]bool, len(grid[0]))

	for i, row := range grid {
		for j, val := range row {
			if val == "#" {
				rowsNotToExpand[i] = true
				colsNotToExpand[j] = true
			}
		}
	}

	expandedGrid := make([][]string, 0)
	for i, row := range grid {
		newRow := make([]string, 0)
		for j, val := range row {
			newRow = append(newRow, val)
			if !colsNotToExpand[j] {
				newRow = append(newRow, ".")
			}
		}
		expandedGrid = append(expandedGrid, newRow)
		if !rowsNotToExpand[i] {
			newEmptyRow := make([]string, len(newRow))
			for k := range newEmptyRow {
				newEmptyRow[k] = "."
			}
			expandedGrid = append(expandedGrid, newEmptyRow)
		}
	}

	var hashes []Point
	for i, row := range expandedGrid {
		for j, val := range row {
			if val == "#" {
				hashes = append(hashes, Point{i, j})
			}
		}
	}

	var pairs [][]Point
	for i := 0; i < len(hashes); i++ {
		for j := i + 1; j < len(hashes); j++ {
			pair := []Point{hashes[i], hashes[j]}
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
