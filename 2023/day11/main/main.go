package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var lines []string
var countI []uint64
var countJ []uint64

func init() {
	data, _ := os.ReadFile("day11/main/sample.txt")
	lines = strings.Split(string(bytes.TrimSpace(data)), "\n")

	countI = make([]uint64, len(lines))
	countJ = make([]uint64, len(lines[0]))

	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				countI[i]++
				countJ[j]++
			}
		}
	}
}

func sumDimention(counts []uint64, expansion uint64) uint64 {
	var sum, ongoing, counting uint64
	for _, count := range counts {
		if count == 0 {
			ongoing += counting * expansion
		} else {
			ongoing += counting
			sum += ongoing * count
			counting += count
		}
	}
	return sum
}

func solve1() uint64 {
	return sumDimention(countI, 2) + sumDimention(countJ, 2)
}

func solve2() uint64 {
	return sumDimention(countI, 1_000_000) + sumDimention(countJ, 1_000_000)
}

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}
