package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "24", `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`)
}
