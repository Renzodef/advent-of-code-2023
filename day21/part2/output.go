package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

type Point struct {
	X, Y int
}

const Steps = 26501365

var directions = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var grid [][]rune

func calculateFirstTermsArithmeticProgression(numTerms int64, terms ...int64) int64 {
	return terms[0] + numTerms*(terms[1]-terms[0]) + numTerms*(numTerms-1)/2*((terms[2]-terms[1])-(terms[1]-terms[0]))
}

func wrapCoordinates(x, y int, grid [][]rune) (int, int) {
	if y >= len(grid) {
		y %= len(grid)
	} else if y < 0 {
		y = (y%len(grid) + len(grid)) % len(grid)
	}
	if x >= len(grid[y]) {
		x %= len(grid[y])
	} else if x < 0 {
		x = (x%len(grid[y]) + len(grid[y])) % len(grid[y])
	}
	return x, y
}

func bfs(grid [][]rune, start Point, maxSteps int) map[string]int {
	distances := map[string]int{fmt.Sprintf("%d,%d", start.X, start.Y): 0}
	toVisit := [][]int{{start.X, start.Y, maxSteps}}
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]
		for _, direction := range directions {
			nextX, nextY := current[0]+direction[0], current[1]+direction[1]
			wrappedX, wrappedY := wrapCoordinates(nextX, nextY, grid)
			if grid[wrappedY][wrappedX] != '#' {
				coordinateKey := fmt.Sprintf("%d,%d", nextX, nextY)
				if _, visited := distances[coordinateKey]; !visited && current[2]-1 >= 0 {
					toVisit = append(toVisit, []int{nextX, nextY, current[2] - 1})
					distances[coordinateKey] = distances[fmt.Sprintf("%d,%d", current[0], current[1])] + 1
				}
			}
		}
	}
	return distances
}

func countReachablePoints(mapGrid [][]rune, start Point, iteration int) int {
	reachablePoints := 0
	for _, distance := range bfs(mapGrid, start, iteration) {
		if (distance+iteration%2)%2 == 0 {
			reachablePoints++
		}
	}
	return reachablePoints
}

func processFile(filePath string) int64 {
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
	var parameters []int64
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
	for i := 0; i < len(grid)*3; i++ {
		if i%len(grid) == int(math.Floor(float64(len(grid))/float64(2))) {
			count := countReachablePoints(grid, start, i)
			parameters = append(parameters, int64(count))
		}
	}
	return calculateFirstTermsArithmeticProgression(int64(math.Floor(float64(Steps)/float64(len(grid)))), parameters...)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
