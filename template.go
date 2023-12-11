package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to process the file
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

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return 0
}

func main() {
	result := processFile("../input.txt")
	fmt.Println(result)
}
