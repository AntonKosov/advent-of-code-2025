package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2025/aoc/test"
)

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "2", `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`)
}
