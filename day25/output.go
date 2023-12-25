package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type NODE struct {
	name     string
	edges    map[*EDGE]*NODE
	traveled bool
}

func newNode(name string) *NODE { return &NODE{name: name, edges: make(map[*EDGE]*NODE)} }

type EDGE struct {
	traveled bool
}

type Graph struct {
	nodes map[string]*NODE
	edges []*EDGE
}

func (g *Graph) resetNodes() {
	for _, n := range g.nodes {
		n.traveled = false
	}
}

func (g *Graph) resetEdges() {
	for _, e := range g.edges {
		e.traveled = false
	}
}

func (g *Graph) removeShortestPath(source, dest *NODE) bool {
	type QueueItem struct {
		edge     *EDGE
		node     *NODE
		previous *QueueItem
	}
	queue := make([]*QueueItem, 0, len(g.nodes))
	queue = append(queue, &QueueItem{node: source})
	found := false
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.node == dest {
			for itr := current; itr.edge != nil; itr = itr.previous {
				itr.edge.traveled = true
			}
			found = true
			break
		}
		for e, n := range current.node.edges {
			if e.traveled || n.traveled {
				continue
			}
			n.traveled = true
			queue = append(queue, &QueueItem{e, n, current})
		}
	}
	g.resetNodes()
	return found
}

func (g *Graph) cutPaths(source, dest *NODE, pathNum int) bool {
	complete := true
	for i := 0; i < pathNum; i++ {
		if !g.removeShortestPath(source, dest) {
			complete = false
			break
		}
	}
	g.resetEdges()
	return complete
}

func (g *Graph) split(cuts int) ([]*NODE, []*NODE) {
	var g1 []*NODE
	var g2 []*NODE
	var source *NODE
	for _, n := range g.nodes {
		source = n
		break
	}
	g1 = append(g1, source)
	for _, dest := range g.nodes {
		if source == dest {
			continue
		}
		if g.cutPaths(source, dest, cuts+1) {
			g1 = append(g1, dest)
		} else {
			g2 = append(g2, dest)
		}
	}
	return g1, g2
}

const cutCount = 3

func MultipleOfTwoGroups(input []string) *Graph {
	nodes := make(map[string]*NODE)
	for _, line := range input {
		name, _, _ := strings.Cut(line, ": ")
		nodes[name] = newNode(name)
	}
	edges := make([]*EDGE, 0)
	for _, line := range input {
		sourceName, destNames, _ := strings.Cut(line, ": ")
		source := nodes[sourceName]
		for _, destName := range strings.Split(destNames, " ") {
			dest, ok := nodes[destName]
			if !ok {
				dest = newNode(destName)
				nodes[destName] = dest
			}
			newEdge := &EDGE{}
			edges = append(edges, newEdge)
			source.edges[newEdge] = dest
			dest.edges[newEdge] = source
		}
	}
	return &Graph{nodes, edges}
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
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	groups := MultipleOfTwoGroups(input)
	firstGroup, secondGroup := groups.split(cutCount)
	return len(firstGroup) * len(secondGroup)
}

func main() {
	startTime := time.Now()
	result := processFile("input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
