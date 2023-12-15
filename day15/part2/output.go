package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type OrderedMap struct {
	Keys   []string
	Values map[string]int
}

var boxMap = make(map[int]*OrderedMap)

func calculateSumOfFocusingPowers() int {
	sumOfFocusingPowers := 0
	for boxStep, orderedMap := range boxMap {
		for slotPosition, key := range orderedMap.Keys {
			focalLength := orderedMap.Values[key]
			focusingPower := (boxStep + 1) * (slotPosition + 1) * focalLength
			sumOfFocusingPowers += focusingPower
		}
	}
	return sumOfFocusingPowers
}

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

func processLine(input string) {
	if strings.HasSuffix(input, "-") {
		trimmedInput := strings.TrimSuffix(input, "-")
		boxStep := calculateHash(trimmedInput)
		if orderedMap, exists := boxMap[boxStep]; exists {
			for i, k := range orderedMap.Keys {
				if k == trimmedInput {
					delete(orderedMap.Values, k)
					orderedMap.Keys = append(orderedMap.Keys[:i], orderedMap.Keys[i+1:]...)
					break
				}
			}
			if len(orderedMap.Values) == 0 {
				delete(boxMap, boxStep)
			}
		}
	} else {
		inputParts := strings.Split(input, "=")
		elementName := inputParts[0]
		elementValue, err := strconv.Atoi(inputParts[1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		boxStep := calculateHash(elementName)
		if _, exists := boxMap[boxStep]; !exists {
			boxMap[boxStep] = &OrderedMap{
				Keys:   make([]string, 0),
				Values: make(map[string]int),
			}
		}
		if _, keyExists := boxMap[boxStep].Values[elementName]; !keyExists {
			boxMap[boxStep].Keys = append(boxMap[boxStep].Keys, elementName)
		}
		boxMap[boxStep].Values[elementName] = elementValue
	}
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ",")
		for _, linePart := range lineParts {
			processLine(linePart)
		}
	}
	return calculateSumOfFocusingPowers()
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
