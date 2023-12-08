package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to process a line and return the concatenation of all its numbers
func processLine(line string) int {
	parts := strings.Fields(line)
	concatenated := ""

	for _, part := range parts {
		_, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return 0
		}
		concatenated += part
	}

	result, err := strconv.Atoi(concatenated)
	if err != nil {
		fmt.Println("Error converting concatenated string to int:", err)
		return 0
	}

	return result
}

// Function to sum the number of ways to win the race
func numberOfWaysToWinRace(time int, distance int) int {
	numberOfWaysToWinTheRace := 0
	for i := 1; i < time; i++ {
		if i*(time-i) > distance {
			numberOfWaysToWinTheRace++
		}
	}
	return numberOfWaysToWinTheRace
}

// Function to process the file and return the number of ways to win the race
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

	if !scanner.Scan() {
		fmt.Println("File contains no lines")
		return 0
	}

	time := processLine(strings.Split(scanner.Text(), ":")[1])
	if time == 0 {
		fmt.Println("Error processing time line:", err)
		return 0
	}

	if !scanner.Scan() {
		fmt.Println("File contains less than two lines")
		return 0
	}

	distance := processLine(strings.Split(scanner.Text(), ":")[1])
	if distance == 0 {
		fmt.Println("Error processing distance line:", err)
		return 0
	}

	if scanner.Scan() {
		fmt.Println("File contains more than two lines")
		return 0
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	return numberOfWaysToWinRace(time, distance)
}

func main() {
	productOfNumbersOfWaysToWinRaces := processFile("../input.txt")
	fmt.Println(productOfNumbersOfWaysToWinRaces)
}
