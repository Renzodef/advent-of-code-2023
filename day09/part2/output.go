package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func calculateArrayOfDifferences(array []int) []int {
	var differences []int
	for i := 1; i < len(array); i++ {
		diff := array[i] - array[i-1]
		differences = append(differences, diff)
	}
	return differences
}

func allZeros(array []int) bool {
	return reflect.DeepEqual(array, make([]int, len(array)))
}

func calculatePreviousValue(line string) int {
	parts := strings.Fields(line)
	var integers []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error converting '%s' to an integer: %v\n", part, err)
			continue
		}
		integers = append(integers, num)
	}
	var differences = integers
	var previousValuesArray []int
	previousValuesArray = append(previousValuesArray, differences[0])
	for !allZeros(differences) {
		differences = calculateArrayOfDifferences(differences)
		previousValuesArray = append(previousValuesArray, differences[0])
	}
	for i := len(previousValuesArray) - 2; i >= 0; i-- {
		previousValuesArray[i] = previousValuesArray[i] - previousValuesArray[i+1]
	}
	return previousValuesArray[0]
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
		}
	}(file)
	sumOfPreviousValues := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sumOfPreviousValues += calculatePreviousValue(line)
	}
	return sumOfPreviousValues
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
