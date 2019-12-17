package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/intcomputer"
	"github.com/cfagiani/aoc2019/util"
)

type Tile struct {
	X, Y, ID int
}

type TileOutputCollector struct {
	Tiles   []Tile
	Pending [3]int
	Idx     int
}

type JoysticInputSupplier struct {
	Output    *TileOutputCollector
	LastIndex int
	CurVal    int
}

func (j *JoysticInputSupplier) GetNextInput() int {
	paddlePosition := 0
	ballPostion := 0

	for i := j.LastIndex; i < len(j.Output.Tiles); i++ {
		if j.Output.Tiles[i].ID == 3 {
			paddlePosition = j.Output.Tiles[i].X
			j.LastIndex = i
		} else if j.Output.Tiles[i].ID == 4 {
			ballPostion = j.Output.Tiles[i].X
			j.LastIndex = i
		}
	}
	if paddlePosition == 0 && ballPostion == 0 {
		//not found, return last value
		return j.CurVal
	} else {
		if paddlePosition < ballPostion {
			j.CurVal = 1

		} else if paddlePosition > ballPostion {
			j.CurVal = -1
		} else {
			j.CurVal = 0
		}
		return j.CurVal
	}
}

func (o *TileOutputCollector) WriteOutput(val int) {
	o.Pending[o.Idx] = val
	o.Idx++
	if o.Idx == 3 {
		if o.Pending[0] == -1 && o.Pending[1] == 0 {
			fmt.Printf("SCORE: %d\n", o.Pending[2])
		} else {
			o.Tiles = append(o.Tiles, Tile{o.Pending[0], o.Pending[1], o.Pending[2]})
		}
		o.Idx = 0
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	content := util.ReadFileAsString("input/day13.input")
	output := &TileOutputCollector{}
	computer := intcomputer.NewComputer(content, 0, 0, false, intcomputer.NewInputSupplier([]int{}), output)
	computer.RunToCompletion(nil)
	count := 0
	for i := 0; i < len(output.Tiles); i++ {
		if output.Tiles[i].ID == 2 {
			count++
		}
	}
	fmt.Printf("There are %d block tiles\n", count)

}

func part2() {
	content := util.ReadFileAsString("input/day13.input")
	content = "2" + content[1:]
	output := &TileOutputCollector{}
	input := &JoysticInputSupplier{Output: output}
	computer := intcomputer.NewComputer(content, 0, 0, false, input, output)
	computer.RunToCompletion(nil)
	count := 0
	for i := 0; i < len(output.Tiles); i++ {
		if output.Tiles[i].ID == 2 {
			count++
		}
	}
	fmt.Printf("There are %d block tiles\n", count)

}
