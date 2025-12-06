package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "6295830249262", "input.txt")
}

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "4277556", `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +
`)
}
