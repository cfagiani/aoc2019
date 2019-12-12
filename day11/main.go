package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/intcomputer"
	"github.com/cfagiani/aoc2019/util"
)

type Point struct {
	X, Y int
}

func (p *Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

var UP = Point{0, 1}
var DOWN = Point{0, -1}
var LEFT = Point{-1, 0}
var RIGHT = Point{1, 0}

var DIRECTIONS = []Point{UP, LEFT, DOWN, RIGHT}

type Robot struct {
	Painted         map[Point]string
	DirectionIndex  int
	CurrentPosition Point
	OutputCount     int
}

func main() {
	part1()
	part2()
}

func (r *Robot) GetNextInput() int {
	val, ok := r.Painted[r.CurrentPosition]
	if ok {
		if val == "B" {
			return 0
		} else {
			return 1
		}
	}
	return 0
}

func (r *Robot) WriteOutput(val int) {
	if r.OutputCount%2 == 0 {
		// even outputs are paint instructions
		if val == 0 {
			r.Painted[r.CurrentPosition] = "B"
		} else if val == 1 {
			r.Painted[r.CurrentPosition] = "W"
		} else {
			fmt.Printf("Unexpected paint color %d\n", val)
		}
	} else {
		// odd outputs are TURN commands
		if val == 0 {
			r.DirectionIndex = (r.DirectionIndex + 1) % 4
		} else if val == 1 {
			r.DirectionIndex = r.DirectionIndex - 1
			if r.DirectionIndex < 0 {
				r.DirectionIndex = 4 + r.DirectionIndex
			}
		}
		r.CurrentPosition = r.CurrentPosition.Add(DIRECTIONS[r.DirectionIndex])
	}
	r.OutputCount++
}

func part1() {
	content := util.ReadFileAsString("input/day11.input")
	robot := NewRobot()
	computer := intcomputer.NewComputer(content, 0, 0, false, robot, robot)
	computer.RunToCompletion(nil)
	fmt.Printf("Painted %d panels at least once\n", len(robot.Painted))

}

func part2() {
	content := util.ReadFileAsString("input/day11.input")
	robot := NewRobot()
	robot.Painted[robot.CurrentPosition] = "W"
	computer := intcomputer.NewComputer(content, 0, 0, false, robot, robot)
	computer.RunToCompletion(nil)
	fmt.Printf("Painted %d panels at least once\n", len(robot.Painted))

	var minX, minY, maxX, maxY int
	for key, _ := range robot.Painted {
		if key.X < minX {
			minX = key.X
		} else if key.X > maxX {
			maxX = key.X
		}
		if key.Y < minY {
			minY = key.Y
		} else if key.Y > maxY {
			maxY = key.Y
		}
	}

	for i := maxY; i >= minY; i-- {
		fmt.Printf("\n")
		for j := minX; j <= maxX; j++ {
			val, ok := robot.Painted[Point{j, i}]
			if ok {
				if val == "W" {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			} else {
				fmt.Printf(" ")
			}
		}
	}

}

func NewRobot() *Robot {
	return &Robot{CurrentPosition: Point{0, 0}, DirectionIndex: 0, Painted: make(map[Point]string)}
}
