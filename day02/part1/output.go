package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var cubesContainedInsideBag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func processGame(gameString string) int {
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
			if currentCount, ok := cubesContainedInsideBag[color]; !ok || number > currentCount {
				return 0
			}
		}
	}
	gameIdString := strings.Fields(parts[0])[1]
	id, err := strconv.Atoi(gameIdString)
	if err != nil {
		fmt.Println("Invalid number:", id)
		return 0
	}
	return id
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
	var sumOfIds = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameString := scanner.Text()
		sumOfIds += processGame(gameString)
	}
	return sumOfIds
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
