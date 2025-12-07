package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

const (
	startCell        = 'S'
	splitterCell     = '^'
	usedSplitterCell = 'x'
)

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) [][]rune {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]

	diagram := make([][]rune, len(lines))
	for i, line := range lines {
		diagram[i] = []rune(line)
	}

	return diagram
}

func process(diagram [][]rune) int {
	return trackBeam(diagram, startPosition(diagram[0]), 0)
}

func trackBeam(diagram [][]rune, x, y int) int {
	if y >= len(diagram) {
		return 0
	}

	switch diagram[y][x] {
	case splitterCell:
		diagram[y][x] = usedSplitterCell
		return 1 + trackBeam(diagram, x-1, y) + trackBeam(diagram, x+1, y)
	case usedSplitterCell:
		return 0
	default:
		return trackBeam(diagram, x, y+1)
	}
}

func startPosition(row []rune) int {
	idx := strings.IndexRune(string(row), startCell)
	if idx < 0 {
		panic("start point not found")
	}

	return idx
}
