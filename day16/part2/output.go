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

var grid [][]rune

func energizeLocal(grid [][]rune, point Point, dx int, dy int, localEnergizedTiles map[Point]bool, localVisited map[Point]string) {
	if point.X < 0 || point.X >= len(grid[0]) || point.Y < 0 || point.Y >= len(grid) {
		return
	}
	direction := fmt.Sprintf("%d,%d", dx, dy)
	if dir, exists := localVisited[point]; exists && dir == direction {
		return
	}
	localVisited[point] = direction
	current := grid[point.Y][point.X]
	if current == '.' {
		localEnergizedTiles[point] = true
		energizeLocal(grid, Point{X: point.X + dx, Y: point.Y + dy}, dx, dy, localEnergizedTiles, localVisited)
	} else if current == '|' {
		if dx != 0 {
			localEnergizedTiles[point] = true
			energizeLocal(grid, Point{X: point.X, Y: point.Y - 1}, 0, -1, localEnergizedTiles, localVisited)
			energizeLocal(grid, Point{X: point.X, Y: point.Y + 1}, 0, 1, localEnergizedTiles, localVisited)
		} else {
			localEnergizedTiles[point] = true
			energizeLocal(grid, Point{X: point.X + dx, Y: point.Y + dy}, dx, dy, localEnergizedTiles, localVisited)
		}
	} else if current == '-' {
		if dy != 0 {
			localEnergizedTiles[point] = true
			energizeLocal(grid, Point{X: point.X - 1, Y: point.Y}, -1, 0, localEnergizedTiles, localVisited)
			energizeLocal(grid, Point{X: point.X + 1, Y: point.Y}, 1, 0, localEnergizedTiles, localVisited)
		} else {
			localEnergizedTiles[point] = true
			energizeLocal(grid, Point{X: point.X + dx, Y: point.Y + dy}, dx, dy, localEnergizedTiles, localVisited)
		}
	} else if current == '/' {
		localEnergizedTiles[point] = true
		if dx == 1 && dy == 0 {
			energizeLocal(grid, Point{X: point.X, Y: point.Y - 1}, 0, -1, localEnergizedTiles, localVisited)
		} else if dx == -1 && dy == 0 {
			energizeLocal(grid, Point{X: point.X, Y: point.Y + 1}, 0, 1, localEnergizedTiles, localVisited)
		} else if dx == 0 && dy == 1 {
			energizeLocal(grid, Point{X: point.X - 1, Y: point.Y}, -1, 0, localEnergizedTiles, localVisited)
		} else if dx == 0 && dy == -1 {
			energizeLocal(grid, Point{X: point.X + 1, Y: point.Y}, 1, 0, localEnergizedTiles, localVisited)
		}
	} else if current == '\\' {
		localEnergizedTiles[point] = true
		if dx == 1 && dy == 0 {
			energizeLocal(grid, Point{X: point.X, Y: point.Y + 1}, 0, 1, localEnergizedTiles, localVisited)
		} else if dx == -1 && dy == 0 {
			energizeLocal(grid, Point{X: point.X, Y: point.Y - 1}, 0, -1, localEnergizedTiles, localVisited)
		} else if dx == 0 && dy == 1 {
			energizeLocal(grid, Point{X: point.X + 1, Y: point.Y}, 1, 0, localEnergizedTiles, localVisited)
		} else if dx == 0 && dy == -1 {
			energizeLocal(grid, Point{X: point.X - 1, Y: point.Y}, -1, 0, localEnergizedTiles, localVisited)
		}
	}
}

func calculateMaxNumberOfEnergizedTiles() int {
	maxLen := 0
	width, height := len(grid[0]), len(grid)
	results := make(chan int)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == 0 || y == height-1 || x == 0 || x == width-1 {
				go func(x, y int) {
					var dx, dy int
					if y == 0 {
						dx, dy = 0, 1
					} else if y == height-1 {
						dx, dy = 0, -1
					} else if x == 0 {
						dx, dy = 1, 0
					} else {
						dx, dy = -1, 0
					}
					localEnergizedTiles := make(map[Point]bool)
					localVisited := make(map[Point]string)
					energizeLocal(grid, Point{X: x, Y: y}, dx, dy, localEnergizedTiles, localVisited)
					results <- len(localEnergizedTiles)
				}(x, y)
			}
		}
	}
	for i := 0; i < (width*2 + height*2 - 4); i++ {
		lenEnergized := <-results
		if lenEnergized > maxLen {
			maxLen = lenEnergized
		}
	}
	close(results)
	return maxLen
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
	return calculateMaxNumberOfEnergizedTiles()
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
