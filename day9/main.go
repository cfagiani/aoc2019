package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/intcomputer"
	"github.com/cfagiani/aoc2019/util"
)

func main() {
	part1()
	part2()
}

func part2(){
	content := util.ReadFileAsString("input/day9.input")
	output := intcomputer.NewOutputSupplier()
	computer := intcomputer.NewComputer(content, 0, 0, false, intcomputer.NewInputSupplier([]int{2}), output)
	computer.RunToCompletion(nil)
	for i := 0; i < len(output.Output); i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Printf("%d", output.Output[i])
	}
}

func part1() {
	content := util.ReadFileAsString("input/day9.input")
	output := intcomputer.NewOutputSupplier()
	computer := intcomputer.NewComputer(content, 0, 0, false, intcomputer.NewInputSupplier([]int{1}), output)
	computer.RunToCompletion(nil)
	for i := 0; i < len(output.Output); i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Printf("%d\n", output.Output[i])
	}
}
