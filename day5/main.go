package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/intcomputer"
	"github.com/cfagiani/aoc2019/util"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	content := util.ReadFileAsString("input/day5.input")
	output := intcomputer.NewOutputSupplier()
	computer := intcomputer.NewComputer(content, 0, 0, false, intcomputer.NewInputSupplier([]int{1}), output)
	for ; !computer.ExecuteNextInstruction(); {
		// do nothing
	}
	return output.Output[len(output.Output)-1]
}


func part2() int {
	content := util.ReadFileAsString("input/day5.input")
	output := intcomputer.NewOutputSupplier()
	computer := intcomputer.NewComputer(content, 0, 0, false, intcomputer.NewInputSupplier([]int{5}), output)
	for ; !computer.ExecuteNextInstruction(); {
		// do nothing
	}
	return output.Output[len(output.Output)-1]
}