package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transformation struct {
	Destination int
	Start       int
	Modifier    int
}

var seeds []int
var seedToSoilMap []Transformation
var soilToFertilizerMap []Transformation
var fertilizerToWaterMap []Transformation
var waterToLightMap []Transformation
var lightToTemperatureMap []Transformation
var temperatureToHumidityMap []Transformation
var humidityToLocationMap []Transformation

func processStage(value int, transformations []Transformation) int {
	for _, t := range transformations {
		if t.Start <= value && value < t.Start+t.Modifier {
			return t.Destination + (value - t.Start)
		}
	}
	return value
}

func getLowestLocationNumber() int {
	lowestLocationNumber := -1
	for _, seed := range seeds {
		currentValue := seed
		currentValue = processStage(currentValue, seedToSoilMap)
		currentValue = processStage(currentValue, soilToFertilizerMap)
		currentValue = processStage(currentValue, fertilizerToWaterMap)
		currentValue = processStage(currentValue, waterToLightMap)
		currentValue = processStage(currentValue, lightToTemperatureMap)
		currentValue = processStage(currentValue, temperatureToHumidityMap)
		currentValue = processStage(currentValue, humidityToLocationMap)
		if lowestLocationNumber == -1 || currentValue < lowestLocationNumber {
			lowestLocationNumber = currentValue
		}
	}
	return lowestLocationNumber
}

func parseMapEntry(line string, currentMap *[]Transformation) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		fmt.Println("Invalid map entry:", line)
		return
	}
	destination, err1 := strconv.Atoi(parts[0])
	start, err2 := strconv.Atoi(parts[1])
	modifier, err3 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error parsing map entry:", line)
		return
	}
	transformation := Transformation{
		Destination: destination,
		Start:       start,
		Modifier:    modifier,
	}
	*currentMap = append(*currentMap, transformation)
}

func parseSeeds(line string) []int {
	seedStrings := strings.Fields(line)
	var seeds []int
	for _, s := range seedStrings {
		seed, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error parsing seed:", err)
			continue
		}
		seeds = append(seeds, seed)
	}
	return seeds
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
	var currentMap *[]Transformation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "seeds:") {
			seeds = parseSeeds(strings.TrimPrefix(line, "seeds: "))
			continue
		}
		if strings.HasSuffix(line, "map:") {
			switch line {
			case "seed-to-soil map:":
				currentMap = &seedToSoilMap
			case "soil-to-fertilizer map:":
				currentMap = &soilToFertilizerMap
			case "fertilizer-to-water map:":
				currentMap = &fertilizerToWaterMap
			case "water-to-light map:":
				currentMap = &waterToLightMap
			case "light-to-temperature map:":
				currentMap = &lightToTemperatureMap
			case "temperature-to-humidity map:":
				currentMap = &temperatureToHumidityMap
			case "humidity-to-location map:":
				currentMap = &humidityToLocationMap
			}
		} else if line != "" && currentMap != nil {
			parseMapEntry(line, currentMap)
		}
	}
	return getLowestLocationNumber()
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
