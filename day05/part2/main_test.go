package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "353716783056994", "input.txt")
}

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "14", `3-5
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
