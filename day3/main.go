package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X     int
	Y     int
	Steps int
}

func (p *Point) Eq(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Point) DistFromOrigin() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

func main() {
	part1()
	part2()
}

func part2() {
	contents := util.ReadAllLines("input/day3.input")
	coords1 := buildCoordinates(contents[0], true)
	coords2 := buildCoordinates(contents[1], true)
	intersection := findIntersection(coords1, coords2)

	minPoint := intersection[0]
	for i := 0; i < len(intersection); i++ {
		if intersection[i].Steps < minPoint.Steps {
			minPoint = intersection[i]
		}
	}
	fmt.Printf("\nMin Intersection %d,%d. Steps %d\n", minPoint.X, minPoint.Y, minPoint.Steps)
}

func part1() {
	contents := util.ReadAllLines("input/day3.input")
	coords1 := buildCoordinates(contents[0], true)
	coords2 := buildCoordinates(contents[1], true)
	min := findMinIntersection(coords1, coords2)
	fmt.Printf("Closest point: %d,%d. Distance: %d", min.X, min.Y, min.DistFromOrigin())
}

func findMinIntersection(coords1 []Point, coords2 []Point) Point {
	for i := 0; i < len(coords1); i++ {
		for j := 0; j < len(coords2); j++ {
			if coords1[i].Eq(coords2[j]) {
				return coords1[i]
			} else if coords2[j].DistFromOrigin() > coords1[i].DistFromOrigin() {
				break
			}
		}
	}
	return Point{0, 0, 0}
}

func buildCoordinates(inputString string, sortList bool) []Point {
	points := []Point{{X: 0, Y: 0, Steps: 0}}
	tokens := strings.Split(inputString, ",")
	for i := 0; i < len(tokens); i++ {
		points = append(points, getPoints(points[len(points)-1], tokens[i])...)
	}
	if sortList {
		sort.Slice(points, func(i, j int) bool {
			return points[i].DistFromOrigin() < points[j].DistFromOrigin()
		})
	}
	// don't include 0,0
	return points[1:]
}

func getPoints(start Point, directive string) []Point {
	deltaX := 0
	deltaY := 0
	var points []Point
	switch directive[0] {
	case 'U':
		deltaX = 1
	case 'D':
		deltaX = -1
	case 'L':
		deltaY = -1
	case 'R':
		deltaY = 1
	}
	amt, _ := strconv.Atoi(directive[1:])
	for i := 1; i <= amt; i++ {
		points = append(points, Point{X: start.X + (i * deltaX), Y: start.Y + (i * deltaY), Steps: start.Steps + i})
	}
	return points
}

func findIntersection(coords1 []Point, coords2 []Point) []Point {
	var matches []Point
	for i := 0; i < len(coords1); i++ {
		for j := 0; j < len(coords2); j++ {
			if coords1[i].Eq(coords2[j]) {
				matches = append(matches, Point{coords1[i].X, coords1[i].Y, coords1[i].Steps + coords2[j].Steps})
				break
			} else if coords2[j].DistFromOrigin() > coords1[i].DistFromOrigin() {
				break
			}
		}
	}
	return matches
}
