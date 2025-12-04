package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) [][]rune {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func process(grid [][]rune) int {
	removed := 0
	for {
		c := removeRolls(grid)
		if c == 0 {
			break
		}
		removed += c
	}

	return removed
}

func removeRolls(grid [][]rune) int {
	const roll = '@'
	const empty = '.'
	count := 0
	for y, row := range grid {
		for x, v := range row {
			if v != roll {
				continue
			}

			squareCount := 0 // including the central roll
			for y2 := max(0, y-1); y2 <= min(len(grid)-1, y+1); y2++ {
				for x2 := max(0, x-1); x2 <= min(len(row)-1, x+1); x2++ {
					if grid[y2][x2] == roll {
						squareCount++
					}
				}
			}

			if squareCount <= 4 {
				grid[y][x] = empty
				count++
			}
		}
	}

	return count
}
