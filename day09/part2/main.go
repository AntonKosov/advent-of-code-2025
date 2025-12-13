package main

import (
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
	"github.com/AntonKosov/advent-of-code-2025/aoc/math"
	"github.com/AntonKosov/advent-of-code-2025/aoc/slice"
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
	positions := make([]math.Vector2[int64], len(lines))
	for i, line := range lines {
		nums := transform.StrToInt64s(line)
		positions[i] = math.NewVector2(nums[0], nums[1])
	}

	return positions
}

func process(positions []math.Vector2[int64]) int64 {
	playground := NewCompressedPlayground(positions)
	maxArea := int64(0)
	for i, pos1 := range positions {
		for j := i + 1; j < len(positions); j++ {
			pos2 := positions[j]
			width := math.Abs(pos1.X-pos2.X) + 1
			height := math.Abs(pos1.Y-pos2.Y) + 1
			area := width * height
			if area > maxArea && playground.Inside(pos1, pos2) {
				maxArea = area
			}
		}
	}

	return maxArea
}

type Status byte

const (
	statusUnknown  Status = iota
	statusEdge     Status = iota
	statusInternal Status = iota
	statusEmpty    Status = iota
)

func (s Status) Tiled() bool {
	return s == statusEdge || s == statusInternal
}

type CompressedPlayground struct {
	playground          [][]Status
	compressedPositions map[math.Vector2[int64]]math.Vector2[int]
}

func NewCompressedPlayground(tiles []math.Vector2[int64]) *CompressedPlayground {
	playground, compressedPositions := compressPlayground(tiles)

	return &CompressedPlayground{
		playground:          playground,
		compressedPositions: compressedPositions,
	}
}

func (p *CompressedPlayground) Inside(pos1, pos2 math.Vector2[int64]) bool {
	cPos1, cPos2 := p.compressedPositions[pos1], p.compressedPositions[pos2]
	startX, endX := min(cPos1.X, cPos2.X), max(cPos1.X, cPos2.X)
	startY, endY := min(cPos1.Y, cPos2.Y), max(cPos1.Y, cPos2.Y)
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			if !p.playground[y][x].Tiled() {
				return false
			}
		}
	}

	return true
}

func compressPlayground(positions []math.Vector2[int64]) (
	playground [][]Status,
	compressedPositions map[math.Vector2[int64]]math.Vector2[int],
) {
	compressedPositions, width, height := compressPositions(positions)

	playground = make([][]Status, height+2)
	for i := range len(playground) {
		playground[i] = make([]Status, width+2)
	}

	positions = append(positions, positions[0])
	for i := 0; i < len(positions)-1; i++ {
		from, to := compressedPositions[positions[i]], compressedPositions[positions[i+1]]
		dir := to.Sub(from).Norm()
		for {
			playground[from.Y][from.X] = statusEdge
			if from == to {
				break
			}
			from = from.Add(dir)
		}
	}

	fillEmpty(playground, 0, 0)
	fillInternal(playground)

	return playground, compressedPositions
}

func fillInternal(playground [][]Status) {
	for _, row := range playground {
		for x, v := range row {
			if v == statusUnknown {
				row[x] = statusInternal
			}
		}
	}
}

func fillEmpty(playground [][]Status, x, y int) {
	if y < 0 || y >= len(playground) || x < 0 || x >= len(playground[y]) {
		return
	}

	if playground[y][x] != statusUnknown {
		return
	}

	playground[y][x] = statusEmpty

	fillEmpty(playground, x-1, y)
	fillEmpty(playground, x+1, y)
	fillEmpty(playground, x, y-1)
	fillEmpty(playground, x, y+1)
}

func compressPositions(positions []math.Vector2[int64]) (
	compressed map[math.Vector2[int64]]math.Vector2[int],
	width, height int,
) {
	realToCompressedX := compressAxis(positions, func(pos math.Vector2[int64]) int64 { return pos.X })
	realToCompressedY := compressAxis(positions, func(pos math.Vector2[int64]) int64 { return pos.Y })

	compressed = make(map[math.Vector2[int64]]math.Vector2[int], len(positions))
	for _, pos := range positions {
		compressed[pos] = math.NewVector2(realToCompressedX[pos.X], realToCompressedY[pos.Y])
	}

	return compressed, len(realToCompressedY), len(realToCompressedY)
}

func compressAxis(positions []math.Vector2[int64], transform func(math.Vector2[int64]) int64) map[int64]int {
	realValues := slice.UniqueValues(positions, transform)
	slices.Sort(realValues)
	realToCompressed := make(map[int64]int, len(realValues))
	for i, realValue := range realValues {
		realToCompressed[realValue] = i + 1
	}

	return realToCompressed
}
