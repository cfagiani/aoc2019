package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Pos 0: %s\n", part1(12, 2))
	n, v := part2()
	fmt.Printf("Answer: %d\n", (100 * n + v))
}

func part2() (int, int) {
	target := "19690720"
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if part1(i, j) == target {
				fmt.Printf("N=%d, V=%d\n", i, j)
				return i, j
			}
		}
	}
	return 0,0
}

func part1(noun int, verb int) string {
	content := util.ReadFileAsString("input/day2.input")
	tokens := strings.Split(content, ",")
	tokens[1] = fmt.Sprintf("%d", noun)
	tokens[2] = fmt.Sprintf("%d", verb)
	for i := 0; i < len(tokens)-4; i += 4 {
		if tokens[i] == "99" {
			return tokens[0]
		} else {
			pos1, _ := strconv.Atoi(tokens[i+1])
			pos2, _ := strconv.Atoi(tokens[i+2])
			val1, _ := strconv.Atoi(tokens[pos1])
			val2, _ := strconv.Atoi(tokens[pos2])
			dest, _ := strconv.Atoi(tokens[i+3])
			if tokens[i] == "1" {
				tokens[dest] = fmt.Sprintf("%d", (val1 + val2))
			} else if tokens[i] == "2" {
				tokens[dest] = fmt.Sprintf("%d", (val1 * val2))
			}
		}
	}
	return tokens[0]
}
