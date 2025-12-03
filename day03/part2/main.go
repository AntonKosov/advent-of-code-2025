package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	"github.com/AntonKosov/advent-of-code-2025/aoc/math"
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

func process(banks []JoltageBank) uint64 {
	sum := uint64(0)
	for _, bank := range banks {
		sum += bank.MaxJoltage()
	}

	return sum
}

type JoltageBank []byte

func (b JoltageBank) MaxJoltage() uint64 {
	const totalDigits = 12
	joltage := uint64(0)
	exp := math.Pow(uint64(10), totalDigits)
	maxIdx := 0
	for i := range totalDigits {
		exp /= 10
		idx := maxIdx + b[maxIdx:len(b)-(totalDigits-i-1)].maxBatteryIndex()
		joltage += uint64(b[idx]) * exp
		maxIdx = idx + 1
	}

	return joltage
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
