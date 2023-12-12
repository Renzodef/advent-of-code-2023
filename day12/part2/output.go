package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var cache = make(map[string]int)

func generateCombinations(line string, groups []int) int {
	key := line
	for _, group := range groups {
		key += strconv.Itoa(group) + ","
	}
	if v, ok := cache[key]; ok {
		return v
	}
	if len(line) == 0 {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if strings.HasPrefix(line, "?") {
		return generateCombinations(strings.Replace(line, "?", ".", 1), groups) +
			generateCombinations(strings.Replace(line, "?", "#", 1), groups)
	}
	if strings.HasPrefix(line, ".") {
		res := generateCombinations(strings.TrimPrefix(line, "."), groups)
		cache[key] = res
		return res
	}
	if strings.HasPrefix(line, "#") {
		if len(groups) == 0 {
			cache[key] = 0
			return 0
		}
		if len(line) < groups[0] {
			cache[key] = 0
			return 0
		}
		if strings.Contains(line[0:groups[0]], ".") {
			cache[key] = 0
			return 0
		}
		if len(groups) > 1 {
			if len(line) < groups[0]+1 || string(line[groups[0]]) == "#" {
				cache[key] = 0
				return 0
			}
			res := generateCombinations(line[groups[0]+1:], groups[1:])
			cache[key] = res
			return res
		} else {
			res := generateCombinations(line[groups[0]:], groups[1:])
			cache[key] = res
			return res
		}
	}
	return 0
}

func countValidArrangements(line string) int {
	parts := strings.Fields(line)
	input := parts[0]
	groupParts := strings.Split(parts[1], ",")
	groups := make([]int, len(groupParts))
	for i, gp := range groupParts {
		group, err := strconv.Atoi(gp)
		if err != nil {
			fmt.Println("Error converting group to integer:", err)
			return 0
		}
		groups[i] = group
	}
	return generateCombinations(input, groups)
}

func unfoldRow(line string) string {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		fmt.Println("Invalid line format")
		return ""
	}
	springConditions := strings.Repeat(parts[0]+"?", 4) + parts[0]
	groupParts := strings.Split(parts[1], ",")
	unfoldedGroups := strings.Join(groupParts, ",") + ","
	unfoldedGroups = strings.Repeat(unfoldedGroups, 4) + strings.Join(groupParts, ",")
	return springConditions + " " + unfoldedGroups
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
	sumOfArrangements := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		unfoldedRow := unfoldRow(scanner.Text())
		sumOfArrangements += countValidArrangements(unfoldedRow)
	}
	return sumOfArrangements
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
