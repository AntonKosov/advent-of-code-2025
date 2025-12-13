package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	"github.com/AntonKosov/advent-of-code-2025/aoc/math"
	"github.com/AntonKosov/advent-of-code-2025/aoc/transform"
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

func read(reader io.Reader) []math.Vector2[int64] {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	tiles := make([]math.Vector2[int64], len(lines))
	for i, line := range lines {
		nums := transform.StrToInt64s(line)
		tiles[i] = math.NewVector2(nums[0], nums[1])
	}

	return tiles
}

func process(tiles []math.Vector2[int64]) int64 {
	maxArea := int64(0)
	for i, tile1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			tile2 := tiles[j]
			width := math.Abs(tile1.X-tile2.X) + 1
			height := math.Abs(tile1.Y-tile2.Y) + 1
			area := width * height
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}
