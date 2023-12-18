package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"time"
)

type Block struct {
	HeatLoss      int
	X, Y          int
	Direction     int
	StraightMoves int
	Index         int
}

type PriorityQueue []*Block

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].HeatLoss < (*pq)[j].HeatLoss
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Block)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].Index = i
	(*pq)[j].Index = j
}

var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var grid [][]int

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func findMinHeatLoss() int {
	rows, cols := len(grid), len(grid[0])
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Block{HeatLoss: 0, X: 0, Y: 0, Direction: -1, StraightMoves: 0})
	visited := make(map[[4]int]bool)
	for pq.Len() > 0 {
		block := heap.Pop(&pq).(*Block)
		x, y, direction, straightMoves := block.X, block.Y, block.Direction, block.StraightMoves
		if x == rows-1 && y == cols-1 {
			if straightMoves >= 4 {
				return block.HeatLoss
			}
			continue
		}
		if visited[[4]int{x, y, direction, straightMoves}] {
			continue
		}
		visited[[4]int{x, y, direction, straightMoves}] = true
		for i, d := range directions {
			nx, ny := x+d[0], y+d[1]
			if isValid(nx, ny, rows, cols) && (direction+2)%4 != i {
				newStraightMoves := straightMoves
				if direction == -1 || direction == i {
					newStraightMoves++
					if newStraightMoves <= 10 {
						newHeatLoss := block.HeatLoss + grid[nx][ny]
						heap.Push(&pq, &Block{HeatLoss: newHeatLoss, X: nx, Y: ny, Direction: i, StraightMoves: newStraightMoves})
					}
				} else {
					if straightMoves >= 4 {
						newHeatLoss := block.HeatLoss + grid[nx][ny]
						heap.Push(&pq, &Block{HeatLoss: newHeatLoss, X: nx, Y: ny, Direction: i, StraightMoves: 1})
					}
				}
			}
		}
	}
	panic("No solution found")
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
		var row []int
		for _, char := range scanner.Text() {
			row = append(row, int(char-'0'))
		}
		grid = append(grid, row)
	}
	return findMinHeatLoss()
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
