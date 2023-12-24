package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Hailstone struct {
	Position [2]int
	Velocity [2]int
}

const (
	MinXYIntersection = 200000000000000
	MaxXYIntersection = 400000000000000
)

func checkIntersection(h1, h2 Hailstone) bool {
	p0, v0 := h1.Position, h1.Velocity
	p1, v1 := h2.Position, h2.Velocity
	if v0[0] == 0 || v1[0] == 0 {
		return false
	}
	if v0[1]*v1[0] == v0[0]*v1[1] {
		return false
	}
	x := ((float64(p1[1]) - float64(p0[1])) + float64(p0[0])*(float64(v0[1])/float64(v0[0])) - float64(p1[0])*(float64(v1[1])/float64(v1[0]))) / ((float64(v0[1]) / float64(v0[0])) - (float64(v1[1]) / float64(v1[0])))
	t0 := (x - float64(p0[0])) / float64(v0[0])
	if t0 < 0 {
		return false
	}
	t1 := (x - float64(p1[0])) / float64(v1[0])
	if t1 < 0 {
		return false
	}
	y := float64(p0[1]) + t0*float64(v0[1])
	return (x >= MinXYIntersection && x <= MaxXYIntersection) && (y >= MinXYIntersection && y <= MaxXYIntersection)
}

func parseAndTrim(s string) int {
	trimmed := strings.TrimSpace(s)
	val, err := strconv.Atoi(trimmed)
	if err != nil {
		fmt.Println("Error parsing int:", err)
	}
	return val
}

func parseLine(line string) Hailstone {
	parts := strings.Split(line, " @ ")
	posStr := strings.Split(parts[0], ",")
	velStr := strings.Split(parts[1], ",")
	return Hailstone{
		Position: [2]int{parseAndTrim(posStr[0]), parseAndTrim(posStr[1])},
		Velocity: [2]int{parseAndTrim(velStr[0]), parseAndTrim(velStr[1])},
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
	var hailstones []Hailstone
	for scanner.Scan() {
		line := scanner.Text()
		hailstone := parseLine(line)
		hailstones = append(hailstones, hailstone)
	}
	count := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if checkIntersection(hailstones[i], hailstones[j]) {
				count++
			}
		}
	}
	return count
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
