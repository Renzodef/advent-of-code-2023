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

func shoelaceFormula(vertices []Point) float64 {
	n := len(vertices)
	area := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += float64(vertices[i].X*vertices[j].Y - vertices[j].X*vertices[i].Y)
	}
	return 0.5 * area
}

func pickTheorem(vertices []Point) int {
	area := shoelaceFormula(vertices)
	perimeter := 0
	for i := 0; i < len(vertices); i++ {
		j := (i + 1) % len(vertices)
		dx := vertices[j].X - vertices[i].X
		dy := vertices[j].Y - vertices[i].Y
		perimeter += int(math.Sqrt(float64(dx*dx + dy*dy)))
	}
	interiorArea := int(area) - perimeter/2 + 1
	return interiorArea + perimeter
}

func findVertices(instructions []string) []Point {
	var vertices []Point
	x, y := 0, 0
	for _, instruction := range instructions {
		var dir rune
		var steps int
		_, err := fmt.Sscanf(instruction, "%c %d", &dir, &steps)
		if err != nil {
			return nil
		}
		switch dir {
		case 'U':
			y -= steps
		case 'D':
			y += steps
		case 'L':
			x -= steps
		case 'R':
			x += steps
		}
		vertices = append(vertices, Point{X: x, Y: y})
	}
	fmt.Println(vertices)
	return vertices
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
	vertices := findVertices(instructions)
	fmt.Println(vertices)
	return pickTheorem(vertices)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
