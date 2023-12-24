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

func valid(point Point, grid [][]rune) bool {
	return point.X >= 0 && point.Y >= 0 && point.X < len(grid[0]) && point.Y < len(grid) && grid[point.Y][point.X] != '#'
}

func getNeighbours(point Point) []Point {
	return []Point{
		{X: point.X + 1, Y: point.Y},
		{X: point.X - 1, Y: point.Y},
		{X: point.X, Y: point.Y + 1},
		{X: point.X, Y: point.Y - 1},
	}
}

func contains(slice []Point, point Point) bool {
	for _, item := range slice {
		if item == point {
			return true
		}
	}
	return false
}

func getNext(path []Point, grid [][]rune) [][]Point {
	head := path[len(path)-1]
	var paths [][]Point
	for _, next := range getNeighbours(head) {
		if valid(next, grid) && !contains(path, next) {
			nextPath := make([]Point, len(path))
			copy(nextPath, path)
			paths = append(paths, append(nextPath, next))
		}
	}
	return paths
}

func search(start Point, end Point) int {
	graph := make(map[Point]map[Point]int)
	graph[start] = make(map[Point]int)
	queue := [][]Point{{start}}
	for len(queue) > 0 {
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		next := getNext(curr, grid)
		vertex := curr[len(curr)-1]
		if ((vertex == start || vertex == end) && len(curr) > 1) || len(next) > 1 {
			_, visited := graph[vertex]
			if !visited {
				graph[vertex] = make(map[Point]int)
				queue = append(queue, getNext([]Point{vertex}, grid)...)
			}
			graph[curr[0]][vertex] = len(curr) - 1
		} else {
			queue = append(queue, next...)
		}
	}
	maxDistance := 0
	queue = [][]Point{{start}}
	for len(queue) > 0 {
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		head := curr[len(curr)-1]
		dist, last := graph[head][end]
		if last {
			for i := 0; i < len(curr)-1; i++ {
				dist += graph[curr[i]][curr[i+1]]
			}
			if dist > maxDistance {
				maxDistance = dist
			}
		} else {
			for next := range graph[head] {
				if !contains(curr, next) {
					nextPath := make([]Point, len(curr))
					copy(nextPath, curr)
					queue = append(queue, append(nextPath, next))
				}
			}
		}
	}
	return maxDistance
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
	return search(start, end)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
