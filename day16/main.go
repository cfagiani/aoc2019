package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"strconv"
)

var basePattern = []int{0, 1, 0, -1}

func main() {
	part1()
	part2()
}

func part1() {
	data := util.ReadFileAsString("input/day16.input")
	for i := 0; i < 100; i++ {
		data = applyPhase(data)
	}
	fmt.Printf("%s\n",data)
}

func part2() {
	data := util.GetStringAsIntArray(util.ReadFileAsString("input/day16.input"))
	newData := make([]int, len(data)*10000)
	for i := 0; i < len(newData); i += len(data) {
		copy(newData[i:], data)
	}
	offset := data[0]*1000000 + data[1]*100000 + data[2]*10000 + data[3]*1000 + data[4]*100 + data[5]*10 + data[6]

	data = newData
	newlist := make([]int, len(data))
	for i := 0; i < 100; i++ {
		newlist[len(data)-1] = data[len(data)-1]
		for j := len(data) - 2; j >= 0; j-- {
			newlist[j] = (data[j] + newlist[j+1]) % 10
		}
		data = newlist
	}

	for _, i := range data[offset : offset+8] {
		fmt.Print(i)
	}
	fmt.Println()
}

func applyPhase(val string) string {
	output := ""
	for i := 1; i <= len(val); i++ {
		pattern := getPattern(i)
		sum := 0
		for j := 0; j < len(val); j++ {
			patternOffset := (j+1) % len(pattern)
			if j < len(pattern)-1 {
				patternOffset = j + 1
			}
			digitVal, _ := strconv.Atoi(string(val[j]))
			sum += digitVal * pattern[patternOffset]
		}

		digString := strconv.Itoa(sum)
		output += string(digString[len(digString)-1])
	}
	return output
}

func getPattern(repeatCount int) []int {
	pattern := make([]int, repeatCount*4)
	for i := 0; i < repeatCount*4; i++ {
		pattern[i] = basePattern[i/repeatCount]
	}
	return pattern
}
