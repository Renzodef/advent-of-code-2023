// https://adventofcode.com/2023/day/4#part2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to process a card and return the count of its matching numbers
func processCard(cardString string) int {
	parts := strings.Split(cardString, ":")
	if len(parts) != 2 {
		fmt.Println("Invalid game string format")
		return 0
	}

	var countOfMatchingNumbers = 0

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
				countOfMatchingNumbers += 1
			}
		}

	}

	return countOfMatchingNumbers
}

// Function to process the file and the number of scratchcards processed
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

	var numberOfScratchcardsProcessed int

	scanner := bufio.NewScanner(file)

	var cardStrings []string

	for scanner.Scan() {
		cardStrings = append(cardStrings, scanner.Text())
	}

	scratchcardsProcessedArray := make([]int, len(cardStrings))

	for index, cardString := range cardStrings {
		countOfMatchingNumbers := processCard(cardString)

		scratchcardsProcessedArray[index]++

		if countOfMatchingNumbers > 0 {
			for i := index + 1; i < len(cardStrings) && i <= index+countOfMatchingNumbers; i++ {
				scratchcardsProcessedArray[i] += scratchcardsProcessedArray[index]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	for _, count := range scratchcardsProcessedArray {
		numberOfScratchcardsProcessed += count
	}

	return numberOfScratchcardsProcessed
}

func main() {
	numberOfScratchcardsProcessed := processFile("../input.txt")
	fmt.Println(numberOfScratchcardsProcessed)
}
