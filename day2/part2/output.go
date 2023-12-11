package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to process a game and return its power
func processGame(gameString string) int {
	var minimumValueForBlueCubes = 0
	var minimumValueForGreenCubes = 0
	var minimumValueForRedCubes = 0

	parts := strings.Split(gameString, ":")
	if len(parts) != 2 {
		fmt.Println("Invalid game string format")
		return 0
	}

	sets := strings.Split(parts[1], ";")

	for _, set := range sets {
		setParts := strings.Split(set, ",")
		for _, setPart := range setParts {
			setPart = strings.TrimSpace(setPart)
			if setPart == "" {
				continue
			}

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

			switch color {
			case "blue":
				if number > minimumValueForBlueCubes {
					minimumValueForBlueCubes = number
				}
			case "green":
				if number > minimumValueForGreenCubes {
					minimumValueForGreenCubes = number
				}
			case "red":
				if number > minimumValueForRedCubes {
					minimumValueForRedCubes = number
				}
			default:
				fmt.Println("Invalid color:", color)
				return 0
			}
		}
	}

	return minimumValueForBlueCubes * minimumValueForGreenCubes * minimumValueForRedCubes
}

// Function to process the file and sum the power of each game
// The power is the product of the minimum number of cubes for each color needed to resolve the game
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

	var sumOfPowers = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gameString := scanner.Text()
		sumOfPowers += processGame(gameString)
	}

	return sumOfPowers
}

func main() {
	result := processFile("../input.txt")
	fmt.Println(result)
}
