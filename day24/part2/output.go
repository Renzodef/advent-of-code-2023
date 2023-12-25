package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tuple struct {
	PX, PY, PZ int
	VX, VY, VZ int
}

func getSingleElementFromSet(set map[int]bool) int {
	for k := range set {
		return k
	}
	return 0
}

func setIntersection(set1, set2 map[int]bool) map[int]bool {
	intersection := newSet()
	for v := range set1 {
		if set2[v] {
			intersection[v] = true
		}
	}
	return intersection
}

func addToSet(set map[int]bool, value int) {
	set[value] = true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func newSet() map[int]bool {
	return make(map[int]bool)
}

func processTuples(inputList []Tuple) int {
	var potentialXSet, potentialYSet, potentialZSet map[int]bool
	for i := 0; i < len(inputList); i++ {
		for j := i + 1; j < len(inputList); j++ {
			A, B := inputList[i], inputList[j]
			if A.VX == B.VX && abs(A.VX) > 100 {
				newXSet := newSet()
				difference := B.PX - A.PX
				for v := -1000; v <= 1000; v++ {
					if v != A.VX && difference%(v-A.VX) == 0 {
						addToSet(newXSet, v)
					}
				}
				if potentialXSet == nil {
					potentialXSet = newXSet
				} else {
					potentialXSet = setIntersection(potentialXSet, newXSet)
				}
			}
			if A.VY == B.VY && abs(A.VY) > 100 {
				newYSet := newSet()
				difference := B.PY - A.PY
				for v := -1000; v <= 1000; v++ {
					if v != A.VY && difference%(v-A.VY) == 0 {
						addToSet(newYSet, v)
					}
				}
				if potentialYSet == nil {
					potentialYSet = newYSet
				} else {
					potentialYSet = setIntersection(potentialYSet, newYSet)
				}
			}
			if A.VZ == B.VZ && abs(A.VZ) > 100 {
				newZSet := newSet()
				difference := B.PZ - A.PZ
				for v := -1000; v <= 1000; v++ {
					if v != A.VZ && difference%(v-A.VZ) == 0 {
						addToSet(newZSet, v)
					}
				}
				if potentialZSet == nil {
					potentialZSet = newZSet
				} else {
					potentialZSet = setIntersection(potentialZSet, newZSet)
				}
			}
		}
	}
	RVX, RVY, RVZ := getSingleElementFromSet(potentialXSet), getSingleElementFromSet(potentialYSet), getSingleElementFromSet(potentialZSet)
	A := inputList[0]
	B := inputList[1]
	MA := float64(A.VY-RVY) / float64(A.VX-RVX)
	MB := float64(B.VY-RVY) / float64(B.VX-RVX)
	CA := float64(A.PY) - (MA * float64(A.PX))
	CB := float64(B.PY) - (MB * float64(B.PX))
	XPos := int((CB - CA) / (MA - MB))
	YPos := int(MA*float64(XPos) + CA)
	Time := (XPos - A.PX) / (A.VX - RVX)
	ZPos := A.PZ + (A.VZ-RVZ)*Time
	return XPos + YPos + ZPos
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
	var inputList []Tuple
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " @ ")
		positions := strings.Split(parts[0], ", ")
		velocities := strings.Split(parts[1], ", ")
		pX, _ := strconv.Atoi(positions[0])
		pY, _ := strconv.Atoi(positions[1])
		pZ, _ := strconv.Atoi(positions[2])
		vX, _ := strconv.Atoi(velocities[0])
		vY, _ := strconv.Atoi(velocities[1])
		vZ, _ := strconv.Atoi(velocities[2])
		inputList = append(inputList, Tuple{PX: pX, PY: pY, PZ: pZ, VX: vX, VY: vY, VZ: vZ})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	return processTuples(inputList)
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
