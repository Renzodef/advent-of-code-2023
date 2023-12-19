package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Rule struct {
	Condition string
	Action    string
}

func cloneMap(original map[string][2]int) map[string][2]int {
	c := make(map[string][2]int)
	for k, v := range original {
		c[k] = v
	}
	return c
}

func countAcceptedDistinctCombinations(workflows map[string][]Rule, currentWorkflow string, values map[string][2]int) int {
	if currentWorkflow == "A" {
		product := 1
		for _, v := range values {
			product *= v[1] - v[0] + 1
		}
		return product
	} else if currentWorkflow == "R" {
		return 0
	}
	total := 0
	for _, rule := range workflows[currentWorkflow] {
		if rule.Condition == "" {
			total += countAcceptedDistinctCombinations(workflows, rule.Action, values)
			continue
		}
		parts := strings.Split(rule.Condition, ">")
		if len(parts) == 1 {
			parts = strings.Split(rule.Condition, "<")
		}
		category := parts[0]
		threshold, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing threshold:", err)
			return 0
		}
		v := values[category]
		var tv, fv [2]int
		if strings.Contains(rule.Condition, ">") {
			tv = [2]int{threshold + 1, v[1]}
			fv = [2]int{v[0], threshold}
		} else if strings.Contains(rule.Condition, "<") {
			tv = [2]int{v[0], threshold - 1}
			fv = [2]int{threshold, v[1]}
		}
		if tv[0] <= tv[1] {
			v2 := cloneMap(values)
			v2[category] = tv
			total += countAcceptedDistinctCombinations(workflows, rule.Action, v2)
		}
		if fv[0] > fv[1] {
			break
		}
		values[category] = fv
	}
	return total
}

func parseRule(ruleStr string) Rule {
	parts := strings.Split(ruleStr, ":")
	if len(parts) == 1 {
		return Rule{Condition: "", Action: parts[0]}
	}
	return Rule{Condition: parts[0], Action: parts[1]}
}

func parseWorkflow(line string) (string, []Rule) {
	parts := strings.SplitN(line, "{", 2)
	name := strings.TrimSpace(parts[0])
	ruleStr := strings.Split(strings.Trim(parts[1], "}"), ",")
	var rules []Rule
	for _, ruleStr := range ruleStr {
		rules = append(rules, parseRule(ruleStr))
	}
	return name, rules
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
	var workflows = make(map[string][]Rule)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		name, rules := parseWorkflow(line)
		workflows[name] = rules
	}
	values := map[string][2]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}
	return countAcceptedDistinctCombinations(workflows, "in", values)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
