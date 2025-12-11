package main

import (
	"fmt"
	"io"
	"os"
	"slices"

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

func read(reader io.Reader) []math.Vector3[int] {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	positions := make([]math.Vector3[int], len(lines))
	for i, line := range lines {
		nums := transform.StrToInts(line)
		positions[i] = math.NewVector3(nums[0], nums[1], nums[2])
	}

	return positions
}

func process(positions []math.Vector3[int]) int {
	edges := buildEdges(positions)
	connections := make([]int, len(positions))
	circuitSizes := make([]int, len(positions))
	for i := range len(positions) {
		connections[i] = i
		circuitSizes[i] = 1
	}

	slices.SortFunc(edges, func(e1, e2 Edge) int { return e1.Dst2(positions) - e2.Dst2(positions) })
	for _, edge := range edges {
		root0Idx := circuitRoot(connections, edge[0])
		root1Idx := circuitRoot(connections, edge[1])
		if root0Idx != root1Idx {
			connections[root1Idx] = root0Idx
			circuitSizes[root0Idx] += circuitSizes[root1Idx]
			if circuitSizes[root0Idx] == len(positions) {
				pos0, pos1 := positions[edge[0]], positions[edge[1]]
				return pos0.X * pos1.X
			}
		}
	}

	panic("no solution")
}

func circuitRoot(connections []int, idx int) int {
	rootIdx := connections[idx]
	if rootIdx != idx {
		rootIdx = circuitRoot(connections, rootIdx)
		connections[idx] = rootIdx
	}

	return rootIdx
}

type Edge [2]int

func (e Edge) Dst2(positions []math.Vector3[int]) int {
	return positions[e[0]].Dst2(positions[e[1]])
}

func buildEdges(positions []math.Vector3[int]) []Edge {
	var edges []Edge
	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			edges = append(edges, Edge{i, j})
		}
	}

	return edges
}
