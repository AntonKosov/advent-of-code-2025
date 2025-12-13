package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "3", `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`)
}
