package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	"github.com/AntonKosov/advent-of-code-2025/aoc/transform"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

type Worksheet struct {
	lines      [][]uint64
	operations []string
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) Worksheet {
	lines := input.Lines(reader)
	operationsLine := lines[len(lines)-2]
	dataLines := lines[:len(lines)-2]

	var worksheet Worksheet
	for _, line := range dataLines {
		worksheet.lines = append(worksheet.lines, transform.StrToUint64s(line))
	}
	worksheet.operations = regexp.MustCompile(`\S`).FindAllString(operationsLine, -1)

	return worksheet
}

func process(worksheet Worksheet) uint64 {
	var sum uint64
	for i, op := range worksheet.operations {
		columnValue := worksheet.lines[0][i]
		for j := 1; j < len(worksheet.lines); j++ {
			value := worksheet.lines[j][i]
			switch op {
			case "*":
				columnValue *= value
			case "+":
				columnValue += value
			}
		}

		sum += columnValue
	}

	return sum
}
