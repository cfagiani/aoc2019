package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/intcomputer"
	"github.com/cfagiani/aoc2019/util"
	"sync"
)

func main() {
	part1()
	part2()
}

func part1() {

	content := util.ReadFileAsString("input/day7.input")

	possibleInputs := util.AllPermutations([]int{0, 1, 2, 3, 4})
	maxOut := 0
	for _, inputs := range possibleInputs {
		nextInput := 0
		for i := 0; i < len(inputs); i++ {
			output := intcomputer.NewOutputSupplier()
			computer := intcomputer.NewComputer(content, 0, 0, false, intcomputer.NewInputSupplier([]int{inputs[i], nextInput}), output)
			computer.RunToCompletion(nil)
			nextInput = output.Output[0]
		}
		if nextInput > maxOut {
			maxOut = nextInput
		}
	}

	fmt.Printf("Largest output value is %d\n", maxOut)
}

func part2() {
	content := util.ReadFileAsString("input/day7.input")

	possibleInputs := util.AllPermutations([]int{5, 6, 7, 8, 9})
	maxOut := 0
	for _, inputs := range possibleInputs {
		firstIOSupplier := intcomputer.NewChannelIOSupplier([]int{})
		previousIOSupplier := firstIOSupplier
		var wg sync.WaitGroup
		for i := 0; i < len(inputs); i++ {
			wg.Add(1)
			previousIOSupplier.WriteOutput(inputs[i])
			if i == 0 {
				firstIOSupplier.WriteOutput(0)
			}

			var outputSupplier *intcomputer.ChannelIOSupplier
			if i < len(inputs)-1 {
				outputSupplier = intcomputer.NewChannelIOSupplier([]int{})
			} else {
				outputSupplier = firstIOSupplier
			}
			computer := intcomputer.NewComputer(content, 0, 0, false, previousIOSupplier, outputSupplier)
			previousIOSupplier = outputSupplier
			go computer.RunToCompletion(&wg)
		}
		wg.Wait()
		outputValue := <-firstIOSupplier.IOChannel
		if outputValue > maxOut {
			maxOut = outputValue
		}
	}
	fmt.Printf("Largest output value is %d\n", maxOut)
}
