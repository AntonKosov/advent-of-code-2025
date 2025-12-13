package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "3121910778619", `987654321111111
811111111111119
234234234234278
818181911112111
`)
}
