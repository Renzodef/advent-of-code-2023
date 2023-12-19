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

func processWorkflows(workflows map[string][]Rule, ratings []map[string]int) int {
	totalScore := 0
	for _, rating := range ratings {
		currentWorkflow := "in"
		for {
			if currentWorkflow == "A" {
				for _, v := range rating {
					totalScore += v
				}
				break
			} else if currentWorkflow == "R" {
				break
			}
			for _, rule := range workflows[currentWorkflow] {
				if rule.Condition == "" {
					currentWorkflow = rule.Action
					break
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
				if strings.Contains(rule.Condition, ">") && rating[category] > threshold {
					currentWorkflow = rule.Action
					break
				} else if strings.Contains(rule.Condition, "<") && rating[category] < threshold {
					currentWorkflow = rule.Action
					break
				}
			}
		}
	}
	return totalScore
}

func parseRating(line string) map[string]int {
	trimmedLine := strings.Trim(strings.Trim(line, " "), "{}")
	parts := strings.Split(trimmedLine, ",")
	rating := make(map[string]int)
	for _, part := range parts {
		part1 := strings.Split(part, "=")
		var err error
		rating[part1[0]], err = strconv.Atoi(part1[1])
		if err != nil {
			fmt.Println("Error parsing rating:", err)
			return nil
		}
	}
	return rating
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
	var workflows = make(map[string][]Rule)
	var ratings []map[string]int
	workflowsFinished := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			workflowsFinished = true
		} else {
			if !workflowsFinished {
				name, rules := parseWorkflow(line)
				workflows[name] = rules
			} else {
				ratings = append(ratings, parseRating(line))
			}
		}
	}
	return processWorkflows(workflows, ratings)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
