package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func split(toSplit, where [2]int) [][2]int {
	ts, te := toSplit[0], toSplit[1]
	ws, we := where[0], where[1]
	before := [2]int{ts, minInt(ws, te)}
	intersection := [2]int{maxInt(ws, ts), minInt(we, te)}
	after := [2]int{maxInt(we, ts), te}
	var valid [][2]int
	if before[0] < before[1] {
		valid = append(valid, before)
	}
	if intersection[0] < intersection[1] {
		valid = append(valid, intersection)
	}
	if after[0] < after[1] {
		valid = append(valid, after)
	}
	return valid
}

func shred(seeds [][2]int, operations map[[2]int]int) [][2]int {
	for operation := range operations {
		var done [][2]int
		for _, seed := range seeds {
			done = append(done, split(seed, operation)...)
		}
		seeds = done
	}
	return seeds
}

func solve(steps []map[[2]int]int, seeds [][2]int) [][2]int {
	for _, step := range steps {
		seeds = shred(seeds, step)
		for i, seed := range seeds {
			for operation, offset := range step {
				if operation[0] <= seed[0] && seed[1] <= operation[1] {
					seeds[i] = [2]int{seed[0] + offset, seed[1] + offset}
				}
			}
		}
	}
	return seeds
}

func minRange(ranges [][2]int) int {
	minVal := ranges[0][0]
	for _, r := range ranges {
		minVal = minInt(minVal, r[0])
	}
	return minVal
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
	var seeds []int
	var steps []map[[2]int]int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seedsLine := strings.TrimPrefix(scanner.Text(), "seeds: ")
	for _, n := range strings.Fields(seedsLine) {
		seed, _ := strconv.Atoi(n)
		seeds = append(seeds, seed)
	}
	seeds1 := make([][2]int, 0, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		seeds1 = append(seeds1, [2]int{seeds[i], seeds[i] + seeds[i+1]})
	}
	var currentStep map[[2]int]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "map") {
			if currentStep != nil {
				steps = append(steps, currentStep)
			}
			currentStep = make(map[[2]int]int)
		} else if len(line) > 0 {
			parts := strings.Fields(line)
			destination, _ := strconv.Atoi(parts[0])
			start, _ := strconv.Atoi(parts[1])
			modifier, _ := strconv.Atoi(parts[2])
			currentStep[[2]int{start, start + modifier}] = destination - start
		}
	}
	steps = append(steps, currentStep)
	return minRange(solve(steps, seeds1))
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
