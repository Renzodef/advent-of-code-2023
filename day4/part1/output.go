package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to process a card and return its points
func processCard(cardString string) int {
	parts := strings.Split(cardString, ":")
	if len(parts) != 2 {
		fmt.Println("Invalid game string format")
		return 0
	}

	var points = 0

	winningNumbersSet := strings.Split(parts[1], "|")[0]
	cardNumbersSet := strings.Split(parts[1], "|")[1]

	winningNumbers := strings.Fields(winningNumbersSet)
	cardNumbers := strings.Fields(cardNumbersSet)
	for _, winningNumberString := range winningNumbers {
		winningNumber, err := strconv.Atoi(winningNumberString)
		if err != nil {
			fmt.Println("Invalid number:", winningNumberString)
			return 0
		}

		for _, cardNumberString := range cardNumbers {
			cardNumber, err := strconv.Atoi(cardNumberString)
			if err != nil {
				fmt.Println("Invalid number:", cardNumberString)
				return 0
			}
			if winningNumber == cardNumber {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

	}

	return points
}

// Function to process the file and sum the points of each card
// Every winning number in a card is worth point
// First 1 point, then doubled each time (2, 4, 8, ...)
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

	var sumOfPoints = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		cardString := scanner.Text()
		sumOfPoints += processCard(cardString)
	}

	return sumOfPoints
}

func main() {
	result := processFile("../input.txt")
	fmt.Println(result)
}
