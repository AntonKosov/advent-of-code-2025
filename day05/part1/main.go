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

func run(reader io.Reader, writer io.Writer) {
	ranges, ids := read(reader)
	answer := process(ranges, ids)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) ([]Range, []ID) {
	lines := input.Lines(reader)
	var ranges []Range
	for lines[0] != "" {
		parts := strings.Split(lines[0], "-")
		lines = lines[1:]
		ranges = append(ranges, Range{
			start: ID(transform.StrToUInt64(parts[0])),
			end:   ID(transform.StrToUInt64(parts[1])),
		})
	}

	lines = lines[1 : len(lines)-1]
	ids := make([]ID, len(lines))
	for i, line := range lines {
		ids[i] = ID(transform.StrToUInt64(line))
	}

	return ranges, ids
}

func process(ranges []Range, ids []ID) int {
	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if r.Contains(id) {
				count++
				break
			}
		}
	}

	return count
}

type ID uint64

type Range struct {
	start ID
	end   ID
}

func (r Range) Contains(id ID) bool {
	return id >= r.start && id <= r.end
}
