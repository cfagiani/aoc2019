package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"math"
	"strconv"
	"strings"
)

type Chemical struct {
	Name     string
	Quantity int
}

type Reaction struct {
	Product  Chemical
	Reagents []Chemical
}

func main() {
	part1()
	part2()
}

func part1() {
	reactions := buildReactions()
	quantitiesOnHand := make(map[string]*Chemical)
	for chem := range reactions {
		quantitiesOnHand[chem] = &Chemical{Name: chem, Quantity: 0}
	}

	fmt.Printf("Need %d ORE to make 1 FUEL\n", getOreNeededForFuel(1, quantitiesOnHand, reactions))
}

func part2() {
	initialStock := 1000000000000
	step := 100000
	var fuel, lastStock int
	reactions := buildReactions()
	quantitiesOnHand := make(map[string]*Chemical)
	for chem := range reactions {
		quantitiesOnHand[chem] = &Chemical{Name: chem, Quantity: 0}
	}

	// compute the amount of fuel we can produce in big steps at first
	for oreUsed := 0; oreUsed < initialStock; fuel++ {
		lastStock = initialStock - oreUsed
		oreUsed += getOreNeededForFuel(step, quantitiesOnHand, reactions)
	}
	fuel = (fuel - 1) * step
	for oreUsed := 0; oreUsed < lastStock; fuel++ {
		oreUsed += getOreNeededForFuel(1, quantitiesOnHand, reactions)
	}
	fmt.Printf("Can make %d FUEL with %d ORE\n", fuel-1, initialStock)

}

func getOreNeededForFuel(qty int, quantitiesOnHand map[string]*Chemical, reactions map[string]Reaction) int {
	reqs := getRequirementsFor(Chemical{Name: "FUEL", Quantity: qty}, quantitiesOnHand, reactions)
	oreCount := 0
	for len(reqs) > 0 {
		var nextReqs []Chemical
		for _, req := range reqs {
			if req.Name == "ORE" {
				oreCount += req.Quantity
			} else {
				nextReqs = append(nextReqs, getRequirementsFor(req, quantitiesOnHand, reactions)...)
			}
		}
		reqs = nextReqs
	}
	return oreCount
}

func getRequirementsFor(chem Chemical, quantitiesOnHand map[string]*Chemical, reactions map[string]Reaction) []Chemical {
	reaction := reactions[chem.Name]
	results := make([]Chemical, len(reaction.Reagents))
	onHand := 0

	onHand = quantitiesOnHand[chem.Name].Quantity

	timesNeeded := int(math.Ceil(float64(chem.Quantity-onHand) / float64(reaction.Product.Quantity)))

	for i := 0; i < len(reaction.Reagents); i++ {
		results[i] = Chemical{Name: reaction.Reagents[i].Name, Quantity: reaction.Reagents[i].Quantity * timesNeeded}

	}
	quantitiesOnHand[chem.Name].Quantity = int(math.Max(float64(onHand+reaction.Product.Quantity*timesNeeded-chem.Quantity), 0.0))
	return results
}

func buildReactions() map[string]Reaction {
	reactionStrings := util.ReadAllLines("input/day14.input")
	reactions := make(map[string]Reaction)
	for _, line := range reactionStrings {
		inputOutput := strings.Split(line, "=>")
		var reagents []Chemical
		for _, part := range strings.Split(inputOutput[0], ",") {
			reagents = append(reagents, getChemical(part))
		}
		product := getChemical(inputOutput[1])
		reactions[product.Name] = Reaction{Product: product, Reagents: reagents}
	}
	return reactions
}

func getChemical(input string) Chemical {
	parts := strings.Split(strings.TrimSpace(input), " ")
	amt, _ := strconv.Atoi(parts[0])
	return Chemical{strings.TrimSpace(parts[1]), amt}
}
