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
	return p.X >= 0 && p.Y >= 0 && p.X < len(grid) && p.Y < len(grid[p.X]) && grid[p.X][p.Y] != '#'
}

func dfs(current Point, path []Point, start Point, end Point) {
	if !isValid(current) {
		return
	}
	if current == end {
		paths = append(paths, append(path, current))
		return
	}
	original := grid[current.X][current.Y]
	grid[current.X][current.Y] = '#'
	if current != start {
		path = append(path, current)
	}
	switch original {
	case '.', '>', '<', '^', 'v':
		if original == '.' || original == 'v' {
			dfs(Point{current.X + 1, current.Y}, path, start, end)
		}
		if original == '.' || original == '^' {
			dfs(Point{current.X - 1, current.Y}, path, start, end)
		}
		if original == '.' || original == '>' {
			dfs(Point{current.X, current.Y + 1}, path, start, end)
		}
		if original == '.' || original == '<' {
			dfs(Point{current.X, current.Y - 1}, path, start, end)
		}
	}
	grid[current.X][current.Y] = original
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
	start, end := Point{X: 0, Y: -1}, Point{X: len(grid) - 1, Y: -1}
	for y, cell := range grid[start.X] {
		if cell == '.' {
			start.Y = y
			break
		}
	}
	for y, cell := range grid[end.X] {
		if cell == '.' {
			end.Y = y
			break
		}
	}
	if start.Y == -1 || end.Y == -1 {
		fmt.Println("Start or end point not found")
		return 0
	}
	dfs(start, []Point{}, start, end)
	for _, path := range paths {
		fmt.Println(len(path))
	}
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
