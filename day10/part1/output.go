package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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

func findMaximumDistanceInLoop(grid [][]rune, startPos Point) int {
	for direction := range directions {
		nextStep := moveInDirection(direction, startPos)
		if !isValidStep(nextStep, grid) {
			continue
		}
		currentPipe := grid[nextStep.X][nextStep.Y]
		stepCounter := 1
		nextDirection, ok := getNextDirection(currentPipe, direction)
		for ok {
			direction = nextDirection
			nextStep = moveInDirection(direction, nextStep)
			if !isValidStep(nextStep, grid) {
				break
			}
			currentPipe = grid[nextStep.X][nextStep.Y]
			stepCounter++
			nextDirection, ok = getNextDirection(currentPipe, direction)
		}
		if currentPipe == 'S' {
			return stepCounter / 2
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
	return findMaximumDistanceInLoop(grid, startPos)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
