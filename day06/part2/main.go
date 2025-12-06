package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	"github.com/AntonKosov/advent-of-code-2025/aoc/transform"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

type Problem struct {
	numbers   []uint64
	operation rune
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []Problem {
	lines := input.Lines(reader)
	lines = transpose(lines[:len(lines)-1])

	var problems []Problem
	var problem *Problem
	for _, line := range lines {
		if problem == nil {
			problem = &Problem{operation: rune(line[len(line)-1])}
		}

		nums := transform.StrToUint64s(line)
		if len(nums) == 0 {
			problems = append(problems, *problem)
			problem = nil
			continue
		}

		problem.numbers = append(problem.numbers, nums[0])
	}

	return append(problems, *problem)
}

func process(problems []Problem) uint64 {
	var sum uint64
	for _, problem := range problems {
		nums := problem.numbers
		solution := nums[0]
		for j := 1; j < len(nums); j++ {
			switch value := nums[j]; problem.operation {
			case '*':
				solution *= value
			case '+':
				solution += value
			}
		}

		sum += solution
	}

	return sum
}

func transpose(lines []string) []string {
	transposed := make([][]rune, len(lines[0]))
	for _, line := range lines {
		for j, v := range line {
			transposed[j] = append(transposed[j], v)
		}
	}

	result := make([]string, len(transposed))
	for i, line := range transposed {
		result[i] = string(line)
	}

	return result
}
