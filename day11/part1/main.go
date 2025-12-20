package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2025/aoc/input"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

const (
	inputDevice  = "you"
	outputDevice = "out"
)

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) map[string][]string {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]

	connections := make(map[string][]string, len(lines))
	for _, line := range lines {
		connections[line[:3]] = strings.Split(line[5:], " ")
	}

	return connections
}

func process(connections map[string][]string) int {
	return countPaths(connections, inputDevice)
}

func countPaths(connections map[string][]string, currentDevice string) int {
	if currentDevice == outputDevice {
		return 1
	}

	count := 0
	for _, outputDevice := range connections[currentDevice] {
		count += countPaths(connections, outputDevice)
	}

	return count
}
