package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"strconv"
)

func main() {
	image := part1()
	part2(image)
}

func part1() [][][]int {
	content := util.ReadFileAsString("input/day8.input")
	return readImage(content, 6, 25)
}

func part2(layers [][][]int) {
	image := make([][]string, len(layers[0]))
	for i := range layers[0] {
		image[i] = make([]string, len(layers[0][i]))
	}
	for l := len(layers) - 1; l >= 0; l-- {
		for i := 0; i < len(layers[l]); i++ {
			fmt.Printf("\n")
			for j := 0; j < len(layers[l][i]); j++ {
				if layers[l][i][j] == 0 {
					image[i][j] = " "
				} else if layers[l][i][j] == 1 {
					image[i][j] = "*"
				}
			}
		}
	}
	for i := 0; i < len(image); i++ {
		fmt.Printf("\n")
		for j := 0; j < len(image[i]); j++ {
			fmt.Printf("%s", image[i][j])
		}
	}
}

func readImage(content string, height int, width int) [][][]int {
	digitCount := make([][]int, 0)
	var layers [][][]int
	minIdx := 0
	layerCount := 0
	for idx := 0; idx < len(content); {
		digits := make([][]int, height)
		for i := range digits {
			digits[i] = make([]int, width)
		}
		digitCount = append(digitCount, make([]int, 10))
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				val := content[idx]
				digit, _ := strconv.Atoi(string(val))
				digitCount[layerCount][digit]++

				digits[i][j] = digit + digits[i][j]

				idx++
			}
		}
		if digitCount[layerCount][0] < digitCount[minIdx][0] {
			minIdx = layerCount
		}
		layers = append(layers, digits)
		layerCount++
	}
	fmt.Printf("Product of min layer is: %d\n", digitCount[minIdx][1]*digitCount[minIdx][2])
	return layers
}
