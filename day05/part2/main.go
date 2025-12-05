package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	"github.com/AntonKosov/advent-of-code-2025/aoc/transform"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	ranges := read(reader)
	answer := process(ranges)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []Range {
	lines := input.Lines(reader)
	var ranges []Range
	for lines[0] != "" {
		parts := strings.Split(lines[0], "-")
		lines = lines[1:]
		ranges = append(ranges, Range{
			start: transform.StrToUInt64(parts[0]),
			end:   transform.StrToUInt64(parts[1]),
		})
	}

	return ranges
}

func process(ranges []Range) uint64 {
	var count uint64
	for _, r := range merge(ranges) {
		count += r.end - r.start + 1
	}

	return count
}

func merge(ranges []Range) []Range {
	slices.SortFunc(ranges, func(a, b Range) int {
		if a.start < b.start {
			return -1
		}
		return 1
	})

	mergedRanges := []Range{ranges[0]}
	for len(ranges) > 0 {
		r := ranges[0]
		ranges = ranges[1:]
		lastRange := &mergedRanges[len(mergedRanges)-1]
		if lastRange.end < r.start-1 {
			mergedRanges = append(mergedRanges, r)
			continue
		}
		lastRange.end = max(lastRange.end, r.end)
	}

	return mergedRanges
}

type Range struct {
	start uint64
	end   uint64
}
