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
	inputDevice  = "svr"
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

func process(connections map[string][]string) uint64 {
	type node struct {
		name                     string
		passedProblematicDevices int
	}

	problematicDevices := map[string]bool{
		"dac": true,
		"fft": true,
	}

	nodes := map[node]uint64{
		{name: inputDevice}: 1,
	}
	var count uint64
	for len(nodes) > 0 {
		nextNodes := make(map[node]uint64, len(nodes))
		for currentNode, currentCount := range nodes {
			for _, deviceName := range connections[currentNode.name] {
				nextNode := currentNode
				nextNode.name = deviceName

				if deviceName == outputDevice {
					if nextNode.passedProblematicDevices == len(problematicDevices) {
						count += currentCount
					}
					continue
				}

				if problematicDevices[deviceName] {
					nextNode.passedProblematicDevices++
				}

				nextNodes[nextNode] += currentCount
			}
		}
		nodes = nextNodes
	}

	return count
}
