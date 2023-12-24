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

var (
	grid  [][]rune
	paths [][]Point
)

func isValid(p Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(grid) && p.X < len(grid[p.Y]) && grid[p.Y][p.X] != '#'
}

func dfs(current Point, path []Point, start Point, end Point) {
	if !isValid(current) {
		return
	}
	if current == end {
		paths = append(paths, append(path, current))
		return
	}
	original := grid[current.Y][current.X]
	grid[current.Y][current.X] = '#'
	if current != start {
		path = append(path, current)
	}
	switch original {
	case '.', '>', '<', '^', 'v':
		if original == '.' || original == '>' {
			dfs(Point{current.X + 1, current.Y}, path, start, end)
		}
		if original == '.' || original == '<' {
			dfs(Point{current.X - 1, current.Y}, path, start, end)
		}
		if original == '.' || original == 'v' {
			dfs(Point{current.X, current.Y + 1}, path, start, end)
		}
		if original == '.' || original == '^' {
			dfs(Point{current.X, current.Y - 1}, path, start, end)
		}
	}
	grid[current.Y][current.X] = original
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
	start, end := Point{Y: 0, X: -1}, Point{Y: len(grid) - 1, X: -1}
	for x, cell := range grid[start.Y] {
		if cell == '.' {
			start.X = x
			break
		}
	}
	for x, cell := range grid[end.Y] {
		if cell == '.' {
			end.X = x
			break
		}
	}
	if start.X == -1 || end.X == -1 {
		fmt.Println("Start or end point not found")
		return 0
	}
	dfs(start, []Point{}, start, end)
	maxPathLength := 0
	for _, path := range paths {
		if len(path) > maxPathLength {
			maxPathLength = len(path)
		}
	}
	return maxPathLength
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
