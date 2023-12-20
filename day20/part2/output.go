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

func (p *Pulse) handle(modules map[string]Module) []Pulse {
	statusOut := false
	m := modules[p.Destination]
	switch m.Type {
	case TypeFlipFlop:
		if p.IsHighPulse {
			return nil
		} else {
			m.On = !m.On
			statusOut = m.On
		}
		break
	case TypeConjunction:
		m.InputStates[p.Source] = p.IsHighPulse
		statusOut = false
		for _, v := range m.InputStates {
			if !v {
				statusOut = true
				break
			}
		}
		break
	case TypeBroadcaster:
		statusOut = p.IsHighPulse
		break
	}
	out := make([]Pulse, len(m.Destinations))
	for i, t := range m.Destinations {
		out[i] = Pulse{
			Source:      p.Destination,
			Destination: t,
			IsHighPulse: statusOut,
		}
	}
	modules[p.Destination] = m
	return out
}

func findInputs(modules map[string]Module, dest string) []string {
	inputs := make([]string, 0)
	for k, m := range modules {
		for _, t := range m.Destinations {
			if t == dest {
				inputs = append(inputs, k)
			}
		}
	}
	return inputs
}

func processPulses(modules map[string]Module) int {
	rxFeed := findInputs(modules, "rx")[0]
	inputs := findInputs(modules, rxFeed)
	factors := make(map[string]int)
	pulses := make([]Pulse, 0)
	for i := 1; len(factors) != len(inputs); i++ {
		pulses = pulses[:0]
		pulses = append(pulses, Pulse{
			Source:      "button",
			Destination: "broadcaster",
			IsHighPulse: false,
		})
		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]
			pulses = append(pulses, pulse.handle(modules)...)
			for _, k := range inputs {
				_, ok := factors[k]
				if !ok && modules[rxFeed].InputStates[k] {
					factors[k] = i
				}
			}
		}
	}
	product := 1
	for _, v := range factors {
		product *= v
	}
	return product
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
