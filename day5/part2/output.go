package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SeedRange struct to hold the start and length of each seed range
type SeedRange struct {
	StartSeed   int
	RangeLength int
}

// Transformation struct to hold each mapping's details
type Transformation struct {
	Destination int
	Start       int
	Modifier    int
}

// Global variables to store the data
var seedRanges []SeedRange
var seedToSoilMap []Transformation
var soilToFertilizerMap []Transformation
var fertilizerToWaterMap []Transformation
var waterToLightMap []Transformation
var lightToTemperatureMap []Transformation
var temperatureToHumidityMap []Transformation
var humidityToLocationMap []Transformation

// Function to parse the seed ranges
func parseSeeds(line string) []SeedRange {
	seedStrings := strings.Fields(line)
	var seedRanges []SeedRange
	for i := 0; i < len(seedStrings); i += 2 {
		startSeed, err1 := strconv.Atoi(seedStrings[i])
		rangeLength, err2 := strconv.Atoi(seedStrings[i+1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing seed range:", err1, err2)
			continue
		}
		seedRanges = append(seedRanges, SeedRange{StartSeed: startSeed, RangeLength: rangeLength})
	}
	return seedRanges
}

// Function to parse the map entries
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

// Function to process each stage
func processStage(value int, transformations []Transformation) int {
	for _, t := range transformations {
		if t.Start <= value && value < t.Start+t.Modifier {
			return t.Destination + (value - t.Start)
		}
	}
	return value
}

// Function to process a seed range and find the lowest location number in that range
func getLowestLocationNumberInSeedRange(startSeed, rangeLength int) int {
	lowestLocationNumber := -1

	for i := 0; i < rangeLength; i++ {
		currentValue := startSeed + i
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

// Function to process the file and return the lowest location number
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

	lowestLocationNumber := -1
	var currentMap *[]Transformation

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			seedRanges = parseSeeds(strings.TrimPrefix(line, "seeds: "))
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

	for _, seedRange := range seedRanges {
		lowestLocationNumberInSeedRange := getLowestLocationNumberInSeedRange(seedRange.StartSeed, seedRange.RangeLength)
		if lowestLocationNumber == -1 || lowestLocationNumberInSeedRange < lowestLocationNumber {
			lowestLocationNumber = lowestLocationNumberInSeedRange
		}
	}

	return lowestLocationNumber
}

func main() {
	lowestLocationNumber := processFile("../input.txt")
	fmt.Println(lowestLocationNumber)
}
