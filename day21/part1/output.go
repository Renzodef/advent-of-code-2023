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

const Steps = 64

var grid [][]rune

func countEvenDistances(distancesMatrix [][]int) int {
	count := 0
	for _, row := range distancesMatrix {
		for _, dist := range row {
			if dist != -1 && dist <= Steps && dist%2 == 0 {
				count++
			}
		}
	}
	return count
}

func bfs(start Point) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	distancesMatrix := make([][]int, rows)
	for i := range distancesMatrix {
		distancesMatrix[i] = make([]int, cols)
		for j := range distancesMatrix[i] {
			distancesMatrix[i][j] = -1
		}
	}
	queue := []Point{start}
	distancesMatrix[start.Y][start.X] = 0
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nx, ny := p.X+dx[i], p.Y+dy[i]
			if nx >= 0 && ny >= 0 && nx < cols && ny < rows && grid[ny][nx] != '#' && distancesMatrix[ny][nx] == -1 {
				distancesMatrix[ny][nx] = distancesMatrix[p.Y][p.X] + 1
				queue = append(queue, Point{nx, ny})
			}
		}
	}
	return distancesMatrix
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
	start := Point{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		for x, char := range row {
			if char == 'S' {
				start = Point{x, y}
			}
		}
		grid = append(grid, row)
		y++
	}
	distanceMatrix := make([][]int, y)
	if start.X != -1 && start.Y != -1 {
		distanceMatrix = bfs(start)
	}
	return countEvenDistances(distanceMatrix)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
