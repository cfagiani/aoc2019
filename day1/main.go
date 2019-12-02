package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"math"
	"strconv"
)

func main() {
	getModuleFuelReqs(false)
	getModuleFuelReqs(true)
}

func getFuelForFuel(mass float64) float64 {
	fuelForFuel := calcFuel(mass)
	if fuelForFuel > 0 {
		fuelForFuel += getFuelForFuel(fuelForFuel)
	} else {
		return 0
	}
	return fuelForFuel
}

func getModuleFuelReqs(includeFuel bool) int {
	contents := util.ReadAllLines("input/day1.input")
	sum := 0
	for i := 0; i < len(contents); i++ {
		moduleFuel := getFuelReq(contents[i])
		sum += moduleFuel
		if includeFuel {
			sum += int(getFuelForFuel(float64(moduleFuel)))
		}

	}
	fmt.Printf("Need %d fuel\n", sum)
	return sum
}

func getFuelReq(massString string) int {
	mass, err := strconv.ParseFloat(massString, 64)
	util.CheckError(err, "Can't parse float", true)
	return int(calcFuel(mass))
}

func calcFuel(mass float64) float64 {
	return math.Floor(mass/float64(3)) - 2
}
