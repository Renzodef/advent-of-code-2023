package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Function to check if an array is all zeros
func allZeros(array []int) bool {
	return reflect.DeepEqual(array, make([]int, len(array)))
}

// Function to calculate the array of differences of an array
func calculateArrayOfDifferences(array []int) []int {
	var differences []int

	for i := 1; i < len(array); i++ {
		diff := array[i] - array[i-1]
		differences = append(differences, diff)
	}

	return differences
}

// Function to calculate the next value of a line
func calculateNextValue(line string) int {
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
	var lastValuesArray []int
	lastValuesArray = append(lastValuesArray, differences[len(differences)-1])

	for !allZeros(differences) {
		differences = calculateArrayOfDifferences(differences)
		lastValue := differences[len(differences)-1]
		lastValuesArray = append(lastValuesArray, lastValue)
	}

	var sumOfLastValues = 0

	for _, lastValue := range lastValuesArray {
		sumOfLastValues += lastValue
	}

	return sumOfLastValues
}

// Function to process the file and sum the next values of each line
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

	sumOfNextValues := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		sumOfNextValues += calculateNextValue(line)
	}

	return sumOfNextValues
}

func main() {
	sumOfNextValues := processFile("../input.txt")
	fmt.Println(sumOfNextValues)
}
