package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X, Y, Z int
}

type Brick struct {
	Start, End Point
}

func countNumberOfBricksThatCanBeSafelyDeleted(bricks []Brick) int {
	zeroBrick := Brick{}
	tmp := make([]Brick, len(bricks))
	numberOfBricksThatCanBeSafelyDeleted := 0
	for i := range bricks {
		copy(tmp, bricks)
		tmp[i] = zeroBrick
		if fall(tmp, true) {
			numberOfBricksThatCanBeSafelyDeleted++
		}
	}
	return numberOfBricksThatCanBeSafelyDeleted
}

func fall(bricks []Brick, checkIfBrickCanBeSafelyDeleted bool) bool {
	for i := range bricks {
		a := &bricks[i]
		for a.Start.Z > 1 {
			for j := i - 1; j >= 0; j-- {
				b := &bricks[j]
				if a.End.X >= b.Start.X &&
					a.Start.X <= b.End.X &&
					a.End.Y >= b.Start.Y &&
					a.Start.Y <= b.End.Y &&
					(a.End.Z-1) >= b.Start.Z &&
					(a.Start.Z-1) <= b.End.Z {
					goto nextBrick
				}
			}
			if checkIfBrickCanBeSafelyDeleted {
				return false
			}
			a.Start.Z--
			a.End.Z--
		}
	nextBrick:
	}
	return true
}

func parsePoint(s string) Point {
	coords := strings.Split(s, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	return Point{x, y, z}
}

func parseBrick(line string) Brick {
	parts := strings.Split(line, "~")
	start := parsePoint(parts[0])
	end := parsePoint(parts[1])
	return Brick{start, end}
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
	var bricks []Brick
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		brick := parseBrick(line)
		bricks = append(bricks, brick)
	}
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].Start.Z < bricks[j].Start.Z
	})
	fall(bricks, false)
	return countNumberOfBricksThatCanBeSafelyDeleted(bricks)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
