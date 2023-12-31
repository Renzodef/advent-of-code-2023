package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func gcdTwo(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcmTwo(a, b int) int {
	return a * b / gcdTwo(a, b)
}

func lcm(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = lcmTwo(result, number)
	}
	return result
}

func processPattern(nodeMap map[string][2]string, pattern string) []int {
	var steps []int
	for node := range nodeMap {
		if strings.HasSuffix(node, "A") {
			currentNode := node
			stepCount := 0
			for !strings.HasSuffix(currentNode, "Z") {
				for _, runeValue := range pattern {
					direction := string(runeValue)
					if direction == "L" {
						currentNode = nodeMap[currentNode][0]
					} else if direction == "R" {
						currentNode = nodeMap[currentNode][1]
					}
					stepCount++
					if strings.HasSuffix(currentNode, "Z") {
						break
					}
				}
			}
			steps = append(steps, stepCount)
		}
	}
	return steps
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
	scanner.Scan()
	pattern := scanner.Text()
	nodeMap := make(map[string][2]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) == 2 {
			cleanedPart := strings.Trim(parts[1], "()")
			nodes := strings.Split(cleanedPart, ", ")
			if len(nodes) == 2 {
				nodeMap[parts[0]] = [2]string{nodes[0], nodes[1]}
			}
		}
	}
	steps := processPattern(nodeMap, pattern)
	return lcm(steps)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
