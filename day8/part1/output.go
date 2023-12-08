package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to calculate the number of steps required to go from AAA to ZZZ given the pattern
func processPattern(nodeMap map[string][2]string, pattern string) int {
	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		for _, runeValue := range pattern {
			direction := string(runeValue)
			if direction == "L" {
				currentNode = nodeMap[currentNode][0]
			} else if direction == "R" {
				currentNode = nodeMap[currentNode][1]
			}

			steps++

			if currentNode == "ZZZ" {
				break
			}
		}
	}

	return steps
}

// Function to process the file and calculate the number of steps required to go from AAA to ZZZ given the pattern
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

	return processPattern(nodeMap, pattern)
}

func main() {
	numberOfSteps := processFile("../input.txt")
	fmt.Println(numberOfSteps)
}
