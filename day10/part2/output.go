package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

var directions = map[rune]Point{
	'W': {X: 0, Y: -1},
	'S': {X: 1, Y: 0},
	'E': {X: 0, Y: 1},
	'N': {X: -1, Y: 0},
}

var pipeDirections = map[rune]map[rune]rune{
	'|': {'S': 'S', 'N': 'N'},
	'-': {'W': 'W', 'E': 'E'},
	'L': {'S': 'E', 'W': 'N'},
	'J': {'S': 'W', 'E': 'N'},
	'7': {'N': 'W', 'E': 'S'},
	'F': {'N': 'E', 'W': 'S'},
}

func shoelace(vertices []Point) float64 {
	area := 0.0
	j := len(vertices) - 1
	for i := 0; i < len(vertices); i++ {
		area += float64(vertices[j].X+vertices[i].X) * float64(vertices[j].Y-vertices[i].Y)
		j = i
	}
	return math.Abs(area) / 2
}

func countTilesInLoop(vertices []Point, boundaryCount int) int {
	area := shoelace(vertices)
	return int(area) + 1 - boundaryCount/2
}

func isBend(pipe rune) bool {
	return pipe == 'L' || pipe == 'J' || pipe == 'F' || pipe == '7'
}

func getNextDirection(currentPipe rune, direction rune) (rune, bool) {
	if nextDir, ok := pipeDirections[currentPipe][direction]; ok {
		return nextDir, true
	}
	return ' ', false
}

func isValidStep(step Point, grid [][]rune) bool {
	return step.X >= 0 && step.Y >= 0 && step.X < len(grid) && step.Y < len(grid[step.X]) && grid[step.X][step.Y] != '.'
}

func moveInDirection(direction rune, startPos Point) Point {
	offset := directions[direction]
	return Point{offset.X + startPos.X, offset.Y + startPos.Y}
}

func countPoints(grid [][]rune, startPos Point) int {
	var vertices []Point
	boundaryCount := 0
	for direction := range directions {
		nextStep := moveInDirection(direction, startPos)
		vertices = append(vertices, startPos)
		if !isValidStep(nextStep, grid) {
			continue
		}
		currentPipe := grid[nextStep.X][nextStep.Y]
		vertices = append(vertices, nextStep)
		boundaryCount++
		nextDirection, ok := getNextDirection(currentPipe, direction)
		for ok {
			direction = nextDirection
			nextStep = moveInDirection(direction, nextStep)
			if !isValidStep(nextStep, grid) {
				break
			}
			currentPipe = grid[nextStep.X][nextStep.Y]
			if isBend(currentPipe) {
				vertices = append(vertices, nextStep)
			}
			boundaryCount++
			nextDirection, ok = getNextDirection(currentPipe, direction)
			if nextStep == startPos && currentPipe == 'S' {
				return countTilesInLoop(vertices, boundaryCount)
			}
		}
	}
	fmt.Println("No loop found.")
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
	scanner := bufio.NewScanner(file)
	var grid [][]rune
	var startPos Point
	foundStart := false
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		if idx := strings.IndexRune(line, 'S'); idx != -1 {
			startPos = Point{X: len(grid) - 1, Y: idx}
			foundStart = true
		}
	}
	if !foundStart {
		fmt.Println("Start point 'S' not found in grid.")
		return 0
	}
	return countPoints(grid, startPos)
}

func main() {
	result := processFile("../input.txt")
	fmt.Println(result)
}
