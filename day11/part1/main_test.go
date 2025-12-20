package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "5", `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`)
}
