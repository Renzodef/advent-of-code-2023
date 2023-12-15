package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func calculateHash(input string) int {
	currentValue := 0
	for _, char := range input {
		asciiCode := int(char)
		currentValue += asciiCode
		currentValue *= 17
		currentValue = currentValue % 256
	}
	return currentValue
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
	sumOfHashes := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ",")
		for _, linePart := range lineParts {
			sumOfHashes += calculateHash(linePart)
		}
	}
	return sumOfHashes
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
