package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	"github.com/AntonKosov/advent-of-code-2025/aoc/transform"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func countRequiredUnits(lines []string) int {
	count := 0
	for _, row := range lines {
		for _, v := range row {
			if v == '#' {
				count++
			}
		}
	}

	return count
}

type Region struct {
	width  int
	height int
	counts []int
}

func (r Region) Fit(patterns []int) bool {
	unitsLeft := r.width * r.height
	for i, count := range r.counts {
		unitsLeft -= count * patterns[i]
		if unitsLeft < 0 {
			return false
		}
	}

	return true
}

func parseRegion(line string) Region {
	parts := strings.Split(line, ": ")
	size := transform.StrToInts(parts[0])

	return Region{
		width:  size[0],
		height: size[1],
		counts: transform.StrToInts(parts[1]),
	}
}

func run(reader io.Reader, writer io.Writer) {
	answer := process(read(reader))
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) ([]int, []Region) {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]

	patterns := make([]int, 6)
	for i := range patterns {
		patterns[i] = countRequiredUnits(lines[1:4])
		lines = lines[5:]
	}

	regions := make([]Region, len(lines))
	for i, line := range lines {
		regions[i] = parseRegion(line)
	}

	return patterns, regions
}

func process(patterns []int, regions []Region) int {
	count := 0
	for _, region := range regions {
		if region.Fit(patterns) {
			count++
		}
	}

	return count
}
