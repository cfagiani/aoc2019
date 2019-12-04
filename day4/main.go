package main

import (
	"fmt"
	"strconv"
)

func main() {
	possiblePasswords := part1()
	//763 too low; 1662 too high
	part2(possiblePasswords)
}

func part2(possible []int) int {
	count := 0
	for i := 0; i < len(possible); i++ {
		digits := strconv.Itoa(possible[i])
		prevVal := -1

		repeats := map[int]int{}
		for j := 0; j < len(digits); j++ {
			cur, _ := strconv.Atoi(string(digits[j]))
			if cur == prevVal {
				if repeats[cur] == 0 {
					repeats[cur] = 2
				} else {
					repeats[cur] += 1
				}
			}
			prevVal = cur
		}
		for _, v := range repeats {
			if v == 2 {
				count++
				break
			}
		}
	}
	fmt.Printf("There are %d possibilities after filtering\n", count)
	return count
}

func part1() []int {
	min := 158126
	max := 624574
	possiblePasswords := getPossiblePasswords(min, max)
	fmt.Printf("There are %d possibilities\n", len(possiblePasswords))
	return possiblePasswords
}

func getPossiblePasswords(min int, max int) []int {
	var possible []int
	for i := min; i <= max; i++ {
		if isValid(i) {
			possible = append(possible, i)
		}
	}
	return possible
}

func isValid(val int) bool {
	digits := strconv.Itoa(val)
	if len(digits) != 6 {
		return false
	}
	hasDouble := false
	prevVal := -1
	for i := 0; i < len(digits); i++ {
		cur, _ := strconv.Atoi(string(digits[i]))
		if cur == prevVal {
			hasDouble = true
		} else if cur < prevVal {
			return false
		} else {
			prevVal = cur
		}

	}
	return hasDouble
}
