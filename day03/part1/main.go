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

func read(reader io.Reader) []JoltageBank {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	banks := make([]JoltageBank, len(lines))
	for i, line := range lines {
		jb := make(JoltageBank, len(line))
		for j, c := range line {
			jb[j] = byte(c - '0')
		}
		banks[i] = jb
	}

	return banks
}

func process(banks []JoltageBank) int {
	sum := 0
	for _, bank := range banks {
		sum += bank.MaxJoltage()
	}

	return sum
}

type JoltageBank []byte

func (b JoltageBank) MaxJoltage() int {
	firstIdx := b[:len(b)-1].maxBatteryIndex()
	secondIdx := firstIdx + 1 + b[firstIdx+1:].maxBatteryIndex()

	return int(b[firstIdx])*10 + int(b[secondIdx])
}

func (b JoltageBank) maxBatteryIndex() int {
	idx := 0
	for i, v := range b {
		if b[idx] < v {
			idx = i
		}
	}

	return idx
}
