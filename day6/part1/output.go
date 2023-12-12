package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numberOfWaysToWinRace(time int, distance int) int {
	numberOfWaysToWinTheRace := 0
	for i := 1; i < time; i++ {
		if i*(time-i) > distance {
			numberOfWaysToWinTheRace++
		}
	}
	return numberOfWaysToWinTheRace
}

func processLine(line string) []int {
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))
	for i, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return nil
		}
		numbers[i] = number
	}
	return numbers
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
	if !scanner.Scan() {
		fmt.Println("File contains no lines")
		return 0
	}
	times := processLine(strings.Split(scanner.Text(), ":")[1])
	if times == nil {
		fmt.Println("Error processing time line:", err)
		return 0
	}
	if !scanner.Scan() {
		fmt.Println("File contains less than two lines")
		return 0
	}
	distances := processLine(strings.Split(scanner.Text(), ":")[1])
	if distances == nil {
		fmt.Println("Error processing distance line:", err)
		return 0
	}
	if len(times) != len(distances) {
		fmt.Println("Time and distance lines have different lengths")
		return 0
	}
	if scanner.Scan() {
		fmt.Println("File contains more than two lines")
		return 0
	}
	timeDistanceMap := make(map[int]int)
	for i, time := range times {
		timeDistanceMap[time] = distances[i]
	}
	productOfNumbersOfWaysToWinRaces := 1
	for time, distance := range timeDistanceMap {
		productOfNumbersOfWaysToWinRaces *= numberOfWaysToWinRace(time, distance)
	}
	return productOfNumbersOfWaysToWinRaces
}

func main() {
	result := processFile("../input.txt")
	fmt.Println(result)
}
