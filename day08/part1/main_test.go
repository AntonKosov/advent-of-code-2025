package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "66640", "input.txt")
}
