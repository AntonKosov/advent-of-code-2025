package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	"github.com/AntonKosov/advent-of-code-2025/aoc/math"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

const (
	startCell    = 'S'
	splitterCell = '^'
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

func process(diagram [][]rune) uint64 {
	startPos := startPosition(diagram[0])
	memo := map[math.Vector2[int]]uint64{}

	return trackBeam(diagram, memo, startPos, 0)
}

func trackBeam(diagram [][]rune, memo map[math.Vector2[int]]uint64, x, y int) uint64 {
	if y >= len(diagram) {
		return 1
	}

	pos := math.NewVector2(x, y)
	if timelines, ok := memo[pos]; ok {
		return timelines
	}

	if diagram[y][x] == splitterCell {
		timelines := trackBeam(diagram, memo, x-1, y) + trackBeam(diagram, memo, x+1, y)
		memo[pos] = timelines
		return timelines
	}

	return trackBeam(diagram, memo, x, y+1)
}

func startPosition(row []rune) int {
	idx := strings.IndexRune(string(row), startCell)
	if idx < 0 {
		panic("start point not found")
	}

	return idx
}
