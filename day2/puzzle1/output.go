package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to process a game string and return the id of the game if it can be resolved
func processGame(gameString string, cubesContainedInsideBag map[string]int) int {
	// Split the game string by ':'
	parts := strings.Split(gameString, ":")
	if len(parts) != 2 {
		fmt.Println("Invalid game string format")
		return 0
	}

	// Further split the second part by ';'
	sets := strings.Split(parts[1], ";")

	for _, set := range sets {
		// Split each set by ','
		setParts := strings.Split(set, ",")
		for _, setPart := range setParts {
			// Remove leading and trailing spaces
			setPart = strings.TrimSpace(setPart)
			if setPart == "" {
				continue
			}

			// Split each part into number and color dividing the string by ' '
			splitPart := strings.Fields(setPart)
			if len(splitPart) != 2 {
				fmt.Println("Invalid format:", setPart)
				return 0
			}

			number, err := strconv.Atoi(splitPart[0])
			if err != nil {
				fmt.Println("Invalid number:", splitPart[0])
				return 0
			}

			color := splitPart[1]

			// Check if the requested color number exceeds the count in the map
			if currentCount, ok := cubesContainedInsideBag[color]; !ok || number > currentCount {
				return 0
			}
		}
	}

	// Extract the game ID from the first part of the string divided by ' '
	gameIdString := strings.Fields(parts[0])[1]
	id, err := strconv.Atoi(gameIdString)
	if err != nil {
		fmt.Println("Invalid number:", id)
		return 0
	}
	return id
}

// Function to process the file and sum the ids of each game that can be resolved
func processFile(filePath string, cubesContainedInsideBag map[string]int) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	var sumOfIds int = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gameString := scanner.Text()
		sumOfIds += processGame(gameString, cubesContainedInsideBag)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	return sumOfIds
}

func main() {
	// Define the cubes contained inside the bag
	var cubesContainedInsideBag = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// Call the function with the file path and cubesContainedInsideBag and print the result
	sumOfIds := processFile("../input.txt", cubesContainedInsideBag)
	fmt.Println(sumOfIds)
}
