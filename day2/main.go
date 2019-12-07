package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/intcomputer"
	"github.com/cfagiani/aoc2019/util"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(12, 2))
	n, v := part2()
	fmt.Printf("Part 2: %d\n", 100*n+v)
}

func part2() (int, int) {
	target := 19690720
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if part1(i, j) == target {
				return i, j
			}
		}
	}
	return 0, 0
}

func part1(noun int, verb int) int {
	content := util.ReadFileAsString("input/day2.input")
	computer := intcomputer.NewComputer(content, noun, verb, true, intcomputer.NewInputSupplier([]int{}), intcomputer.NewOutputSupplier())
	computer.RunToCompletion(nil)
	return computer.Memory[0]
}
