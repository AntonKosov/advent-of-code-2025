package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	aocMath "github.com/AntonKosov/advent-of-code-2025/aoc/math"
	"github.com/AntonKosov/advent-of-code-2025/aoc/transform"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

type Machine struct {
	indicator uint16
	buttons   []uint16
}

func (m Machine) MinPresses() int {
	options := uint16(1 << len(m.buttons))
	minPresses := math.MaxInt
	for option := range options {
		presses := aocMath.CountBits(option)
		if presses >= minPresses {
			continue
		}
		var indicator uint16
		for i, button := range m.buttons {
			if option&(1<<i) != 0 {
				indicator ^= button
			}
		}
		if indicator == m.indicator {
			minPresses = presses
		}
	}

	if minPresses == math.MaxInt {
		panic("solution not found")
	}

	return minPresses
}

func parseMachine(inputLine string) Machine {
	parts := strings.Split(inputLine, " ")
	indicator := parts[0]
	indicator = indicator[1 : len(indicator)-1]
	buttons := parts[1 : len(parts)-1]

	return Machine{
		indicator: parseIndicator(indicator),
		buttons:   parseButtons(len(indicator), buttons),
	}
}

func parseIndicator(str string) uint16 {
	var value uint16
	for i, v := range str {
		if v == '#' {
			value |= 1 << (len(str) - i - 1)
		}
	}

	return value
}

func parseButtons(lightsCount int, values []string) []uint16 {
	buttons := make([]uint16, len(values))
	for i, v := range values {
		nums := transform.StrToInts(v)
		var button uint16
		for _, num := range nums {
			button |= 1 << (lightsCount - num - 1)
		}

		buttons[i] = button
	}

	return buttons
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []Machine {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	machines := make([]Machine, len(lines))
	for i, line := range lines {
		machines[i] = parseMachine(line)
	}

	return machines
}

func process(machines []Machine) int {
	presses := 0
	for _, machine := range machines {
		presses += machine.MinPresses()
	}

	return presses
}
