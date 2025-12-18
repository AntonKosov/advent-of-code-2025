package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "7", `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`)
}

func TestFromInputL6I7(t *testing.T) {
	test.AssertStringInput(t, run, "3", `[...###] (1,3,4) (1,2) (4) (1,3,4,5) (0,1,2,5) (2,3,4) (0,1,2,3) (1,2,5) {31,60,59,57,54,26}
`)
}
