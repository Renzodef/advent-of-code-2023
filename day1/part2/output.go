package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Variable to map string digits to their numeric equivalents
var stringToDigitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// Function to check if a string contains any of the digit words as substring
func stringToDigit(s string) (int, bool) {
	for word := range stringToDigitMap {
		if strings.Contains(s, word) {
			return stringToDigitMap[word], true
		}
	}
	return 0, false
}

// Function to concatenate the first and last digits or string digits of a line
func sumFirstLastDigit(line string) int {
	firstDigit, lastDigit := -1, -1

	for i := 0; i < len(line); i++ {
		substring := line[:i+1]
		if digit, ok := stringToDigit(substring); ok {
			firstDigit = digit
			break
		} else if unicode.IsDigit(rune(line[i])) {
			firstDigit, _ = strconv.Atoi(string(line[i]))
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		substring := line[i:]
		if digit, ok := stringToDigit(substring); ok {
			lastDigit = digit
			break
		} else if unicode.IsDigit(rune(line[i])) {
			lastDigit, _ = strconv.Atoi(string(line[i]))
			break
		}
	}

	if firstDigit == -1 || lastDigit == -1 {
		return 0
	}

	concatenated := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
	result, err := strconv.Atoi(concatenated)
	if err != nil {
		fmt.Println("Error converting concatenated string to int:", err)
		return 0
	}

	return result
}

// Function to process the file and sum the concatenated first and last digits or string digits of each line
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

	var totalSumOfFirstLastDigits = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		totalSumOfFirstLastDigits += sumFirstLastDigit(line)
	}

	return totalSumOfFirstLastDigits
}

func main() {
	totalSumOfFirstLastDigits := processFile("../input.txt")
	fmt.Println(totalSumOfFirstLastDigits)
}
