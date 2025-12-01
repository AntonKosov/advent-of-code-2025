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

func read(reader io.Reader) []Rotation {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	rotations := make([]Rotation, len(lines))
	for i, line := range lines {
		rotations[i] = Rotation{
			direction: rune(line[0]),
			distance:  transform.StrToInt(line[1:]),
		}
	}

	return rotations
}

func process(rotations []Rotation) int {
	const totalNumbers = 100
	count := 0
	position := 50
	for _, rotation := range rotations {
		sign := 1
		if rotation.direction == 'L' {
			sign = -1
		}

		newPosition := position + sign*rotation.distance
		if newPosition*position < 0 || newPosition == 0 {
			count++
		}

		count += math.Abs(newPosition) / totalNumbers
		position = math.Mod(newPosition, totalNumbers)
	}

	return count
}

type Rotation struct {
	direction rune
	distance  int
}
