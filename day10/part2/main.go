package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	aocMath "github.com/AntonKosov/advent-of-code-2025/aoc/math"
	"github.com/AntonKosov/advent-of-code-2025/aoc/pointer"
	"github.com/AntonKosov/advent-of-code-2025/aoc/transform"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

type Joltages [10]int16

func (joltages *Joltages) PressButton(button []int) bool {
	for _, idx := range button {
		if joltages[idx] == 0 {
			return false
		}
	}

	for _, idx := range button {
		joltages[idx]--
	}

	return true
}

func (joltages *Joltages) OddLevels() bool {
	for _, j := range joltages {
		if j%2 == 1 {
			return true
		}
	}

	return false
}

type Machine struct {
	joltages Joltages
	buttons  [][]int
}

func (m Machine) MinPresses() int {
	cache := map[Joltages]*int{{}: pointer.Ref(0)}
	presses := findMinPresses(m.buttons, m.joltages, cache)
	if presses == nil {
		panic(fmt.Sprintf("solution not found: %v", m.buttons))
	}

	return *presses
}

func findMinPresses(buttons [][]int, joltages Joltages, cache map[Joltages]*int) (result *int) {
	buttonsMasksCount := uint16(1 << len(buttons))
	var minPresses *int
nextMask:
	for buttonsMask := range buttonsMasksCount {
		joltages := joltages
		for i, button := range buttons {
			if buttonsMask&(1<<i) != 0 {
				if !joltages.PressButton(button) {
					continue nextMask
				}
			}
		}

		if joltages.OddLevels() {
			continue
		}

		if c := splitInHalf(buttons, joltages, cache); c != nil {
			presses := *c + aocMath.CountBits(buttonsMask)
			if minPresses == nil {
				minPresses = &presses
			} else {
				*minPresses = min(*minPresses, presses)
			}
		}
	}

	return minPresses
}

func splitInHalf(buttons [][]int, joltages Joltages, cache map[Joltages]*int) (result *int) {
	if v, ok := cache[joltages]; ok {
		return v
	}

	defer func() { cache[joltages] = result }()

	halfJoltages := joltages
	for i := range halfJoltages {
		halfJoltages[i] /= 2
	}

	result = findMinPresses(buttons, halfJoltages, cache)
	if result == nil {
		return nil
	}

	return pointer.Ref(*result * 2)
}

func parseMachine(inputLine string) Machine {
	parts := strings.Split(inputLine, " ")
	joltage := parts[len(parts)-1]
	buttons := parts[1 : len(parts)-1]

	var joltages Joltages
	for i, j := range transform.StrToInts(joltage) {
		joltages[i] = int16(j)
	}

	return Machine{
		joltages: joltages,
		buttons:  parseButtons(buttons),
	}
}

func parseButtons(values []string) [][]int {
	buttons := make([][]int, len(values))
	for i, v := range values {
		buttons[i] = transform.StrToInts(v)
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
	minPresses := 0
	for _, machine := range machines {
		minPresses += machine.MinPresses()
	}

	return minPresses
}
