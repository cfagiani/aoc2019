package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"strings"
)

type OrbitalObject struct {
	Name           string
	ObjectsInOrbit []*OrbitalObject
	Orbiting       *OrbitalObject
	Orbits         int
}

type OrbitalMap struct {
	ObjectMap map[string]*OrbitalObject
}

func (m *OrbitalMap) GetOrAddObject(name string) *OrbitalObject {
	value, ok := m.ObjectMap[name]
	if !ok {
		value = &OrbitalObject{Name: name}
		m.ObjectMap[name] = value
	}
	return value
}

func main() {
	orbitalMap := part1()
	part2(orbitalMap)
}

func part1() OrbitalMap {
	contents := util.ReadAllLines("input/day6.input")
	objectMap := OrbitalMap{ObjectMap: map[string]*OrbitalObject{}}
	for _, line := range contents {
		parts := strings.Split(line, ")")
		center := objectMap.GetOrAddObject(parts[0])
		objInOrbit := objectMap.GetOrAddObject(parts[1])
		center.ObjectsInOrbit = append(center.ObjectsInOrbit, objInOrbit)
		objInOrbit.Orbiting = center
	}
	centerOfMass := objectMap.GetOrAddObject("COM")
	centerOfMass.Orbits = 0
	traverseMap(centerOfMass)
	sum := 0
	for _, obj := range objectMap.ObjectMap {
		sum += obj.Orbits
	}
	fmt.Printf("Total orbits %d\n", sum)
	return objectMap
}

func part2(orbitalMap OrbitalMap) {
	startAncestors := getAncestors(orbitalMap.GetOrAddObject("YOU"))
	endAncestors := getAncestors(orbitalMap.GetOrAddObject("SAN"))
	commonAncestorLengths := []int{}
	for i := 0; i < len(startAncestors); i++ {
		for j := 0; j < len(endAncestors); j++ {
			if startAncestors[i].Name == endAncestors[j].Name {
				commonAncestorLengths = append(commonAncestorLengths, i+j)
			}
		}
	}
	min := commonAncestorLengths[0]
	for _, v := range commonAncestorLengths {
		if v < min {
			min = v
		}
	}
	fmt.Printf("Fewest number of orbital transfers %d\n", min)
}

func getAncestors(start *OrbitalObject) []*OrbitalObject {
	var ancestors []*OrbitalObject
	cur := start
	for ; cur.Orbiting != nil; {
		ancestors = append(ancestors, cur.Orbiting)
		cur = cur.Orbiting
	}
	return ancestors
}

func traverseMap(start *OrbitalObject) {
	for _, obj := range start.ObjectsInOrbit {
		obj.Orbits = start.Orbits + 1
		traverseMap(obj)
	}
}
