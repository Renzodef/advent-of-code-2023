package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const TypeButton = 'b'
const TypeBroadcaster = 'a'
const TypeConjunction = '&'
const TypeFlipFlop = '%'

type Module struct {
	Type         rune
	Destinations []string
	On           bool
	InputStates  map[string]bool
}

type Pulse struct {
	Source      string
	Destination string
	IsHighPulse bool
}

func (p *Pulse) handlePulse(modules map[string]Module) []Pulse {
	statusOut := false
	module := modules[p.Destination]
	switch module.Type {
	case TypeFlipFlop:
		if p.IsHighPulse {
			return nil
		} else {
			module.On = !module.On
			statusOut = module.On
		}
		break
	case TypeConjunction:
		module.InputStates[p.Source] = p.IsHighPulse
		statusOut = false
		for _, m := range module.InputStates {
			if !m {
				statusOut = true
				break
			}
		}
		break
	case TypeBroadcaster:
		statusOut = p.IsHighPulse
		break
	}
	out := make([]Pulse, len(module.Destinations))
	for i, d := range module.Destinations {
		out[i] = Pulse{
			Source:      p.Destination,
			Destination: d,
			IsHighPulse: statusOut,
		}
	}
	modules[p.Destination] = module
	return out
}

func processPulses(modules map[string]Module) int {
	lowPulses := 0
	highPulses := 0
	pulses := make([]Pulse, 0)
	for i := 0; i < 1000; i++ {
		pulses = pulses[:0]
		pulses = append(pulses, Pulse{
			Source:      "button",
			Destination: "broadcaster",
			IsHighPulse: false,
		})
		for len(pulses) > 0 {
			currentPulse := pulses[0]
			pulses = pulses[1:]
			if currentPulse.IsHighPulse {
				highPulses += 1
			} else {
				lowPulses += 1
			}
			pulses = append(pulses, currentPulse.handlePulse(modules)...)
		}
	}
	return lowPulses * highPulses
}

func parseModules(input []byte) map[string]Module {
	modules := map[string]Module{
		"button": {
			Type:         TypeButton,
			Destinations: []string{"broadcaster"},
		},
	}
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		m := Module{}
		pos := strings.Index(line, " ")
		var name string
		if line[0] == TypeFlipFlop || line[0] == TypeConjunction {
			m.Type = rune(line[0])
			name = line[1:pos]
		} else {
			m.Type = TypeBroadcaster
			name = line[:pos]
		}
		line = line[strings.Index(line, " -> ")+4:]
		m.Destinations = strings.Split(line, ", ")
		m.InputStates = make(map[string]bool)
		modules[name] = m
	}
	for k, input := range modules {
		for _, t := range input.Destinations {
			dest := modules[t]
			if dest.Type == TypeConjunction {
				dest.InputStates[k] = false
			}
		}
	}
	return modules
}

func processFile(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	var inputBuilder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputBuilder.WriteString(scanner.Text() + "\n")
	}
	modules := parseModules([]byte(inputBuilder.String()))
	result := processPulses(modules)
	return result
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
