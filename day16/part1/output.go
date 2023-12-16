package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Point struct {
	X, Y int
}

var energizedTiles = make(map[Point]bool)
var grid [][]rune
var visited = make(map[Point]string)

func energize(grid [][]rune, point Point, dx int, dy int) {
	if point.X < 0 || point.X >= len(grid[0]) || point.Y < 0 || point.Y >= len(grid) {
		return
	}
	direction := fmt.Sprintf("%d,%d", dx, dy)
	if dir, exists := visited[point]; exists && dir == direction {
		return
	}
	visited[point] = direction
	current := grid[point.Y][point.X]
	if current == '.' {
		energizedTiles[point] = true
		energize(grid, Point{X: point.X + dx, Y: point.Y + dy}, dx, dy)
	} else if current == '|' {
		if dx != 0 {
			energizedTiles[point] = true
			energize(grid, Point{X: point.X, Y: point.Y - 1}, 0, -1)
			energize(grid, Point{X: point.X, Y: point.Y + 1}, 0, 1)
		} else {
			energizedTiles[point] = true
			energize(grid, Point{X: point.X + dx, Y: point.Y + dy}, dx, dy)
		}
	} else if current == '-' {
		if dy != 0 {
			energizedTiles[point] = true
			energize(grid, Point{X: point.X - 1, Y: point.Y}, -1, 0)
			energize(grid, Point{X: point.X + 1, Y: point.Y}, 1, 0)
		} else {
			energizedTiles[point] = true
			energize(grid, Point{X: point.X + dx, Y: point.Y + dy}, dx, dy)
		}
	} else if current == '/' {
		energizedTiles[point] = true
		if dx == 1 && dy == 0 {
			energize(grid, Point{X: point.X, Y: point.Y - 1}, 0, -1)
		} else if dx == -1 && dy == 0 {
			energize(grid, Point{X: point.X, Y: point.Y + 1}, 0, 1)
		} else if dx == 0 && dy == 1 {
			energize(grid, Point{X: point.X - 1, Y: point.Y}, -1, 0)
		} else if dx == 0 && dy == -1 {
			energize(grid, Point{X: point.X + 1, Y: point.Y}, 1, 0)
		}
	} else if current == '\\' {
		energizedTiles[point] = true
		if dx == 1 && dy == 0 {
			energize(grid, Point{X: point.X, Y: point.Y + 1}, 0, 1)
		} else if dx == -1 && dy == 0 {
			energize(grid, Point{X: point.X, Y: point.Y - 1}, 0, -1)
		} else if dx == 0 && dy == 1 {
			energize(grid, Point{X: point.X + 1, Y: point.Y}, 1, 0)
		} else if dx == 0 && dy == -1 {
			energize(grid, Point{X: point.X - 1, Y: point.Y}, -1, 0)
		}
	}
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
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	energize(grid, Point{X: 0, Y: 0}, 1, 0)
	return len(energizedTiles)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
