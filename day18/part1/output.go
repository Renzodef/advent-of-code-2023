package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X, Y int
}

func countHashes(grid [][]rune) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == '#' {
				count++
			}
		}
	}
	return count
}

func floodFill(grid [][]rune, x, y int) {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) || grid[y][x] != '.' {
		return
	}
	grid[y][x] = ' '
	floodFill(grid, x-1, y)
	floodFill(grid, x+1, y)
	floodFill(grid, x, y-1)
	floodFill(grid, x, y+1)
}

func move(direction string, steps int, position Point) Point {
	switch direction {
	case "U":
		position.Y -= steps
	case "D":
		position.Y += steps
	case "L":
		position.X -= steps
	case "R":
		position.X += steps
	}
	return position
}

func createGrid(instructions []string) [][]rune {
	position := Point{0, 0}
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for _, instruction := range instructions {
		parts := strings.Fields(instruction)
		if len(parts) < 2 {
			continue
		}
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		position = move(direction, steps, position)

		if position.X < minX {
			minX = position.X
		}
		if position.X > maxX {
			maxX = position.X
		}
		if position.Y < minY {
			minY = position.Y
		}
		if position.Y > maxY {
			maxY = position.Y
		}
	}
	gridWidth := maxX - minX + 1
	gridHeight := maxY - minY + 1
	grid := make([][]rune, gridHeight)
	for i := range grid {
		grid[i] = make([]rune, gridWidth)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	position = Point{-minX, -minY} // Adjust the origin based on the min values
	for _, instruction := range instructions {
		parts := strings.Fields(instruction)
		if len(parts) < 2 {
			continue
		}
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		for s := 0; s < steps; s++ {
			switch direction {
			case "U":
				position.Y--
			case "D":
				position.Y++
			case "L":
				position.X--
			case "R":
				position.X++
			}
			grid[position.Y][position.X] = '#'
		}
	}
	return grid
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
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	grid := createGrid(instructions)
	for x := 0; x < len(grid[0]); x++ {
		floodFill(grid, x, 0)
		floodFill(grid, x, len(grid)-1)
	}
	for y := 0; y < len(grid); y++ {
		floodFill(grid, 0, y)
		floodFill(grid, len(grid[0])-1, y)
	}
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == '.' {
				grid[y][x] = '#'
			}
		}
	}
	hashCount := countHashes(grid)
	return hashCount
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
