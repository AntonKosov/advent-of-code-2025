package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "3", `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`)
}
