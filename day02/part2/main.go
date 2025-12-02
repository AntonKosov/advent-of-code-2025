package main

import (
	"fmt"
	"io"
	"os"
	"strings"

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

func read(reader io.Reader) []Range {
	line := input.Lines(reader)[0]
	ur := strings.Split(line, ",")
	ranges := make([]Range, len(ur))
	for i, r := range ur {
		parts := strings.Split(r, "-")
		ranges[i] = Range{
			firstID: transform.StrToInt(parts[0]),
			lastID:  transform.StrToInt(parts[1]),
		}
	}

	return ranges
}

func process(ranges []Range) (sum uint64) {
	for _, rng := range ranges {
		for id := rng.firstID; id <= rng.lastID; id++ {
			if !validID(id) {
				sum += uint64(id)
			}
		}
	}

	return sum
}

func validID(id int) bool {
	digits := math.CountDigits(id)
nextSize:
	for i := 1; i <= digits/2; i++ {
		if digits%i != 0 {
			continue
		}
		exp := math.Pow(10, uint(i))
		sample := id % exp
		rest := id / exp
		for rest != 0 {
			p := rest % exp
			rest /= exp
			if sample != p {
				continue nextSize
			}
		}

		return false
	}

	return true
}

type Range struct {
	firstID, lastID int
}
