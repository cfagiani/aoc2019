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
	return Point{p.X + other.X, p.Y + other.Y}
}

type RepairDroid struct {
	Grid         map[Point]int
	Backtrack    MoveStack
	Position     Point
	NextPosition Point
}

type MoveStack struct {
	len int
	top *stackNode
}

type stackNode struct {
	Move int
	prev *stackNode
}

func (s *MoveStack) Push(move int) {

	s.len++
	if s.top == nil {
		s.top = &stackNode{Move: move, prev: nil}
	} else {
		s.top = &stackNode{Move: move, prev: s.top}
	}
}

func (s *MoveStack) Pop() int {
	if s.len > 0 {
		s.len--
		move := s.top.Move
		s.top = s.top.prev
		return move
	}
	return 0
}

func (s *MoveStack) Depth() int {
	return s.len
}

func (d *RepairDroid) Print() {
	var minX, minY, maxX, maxY int
	for pos := range d.Grid {
		if pos.X < minX {
			minX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		}
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			val, ok := d.Grid[Point{X: j, Y: i}]
			if ok {
				if i == 0 && j == 0 {
					fmt.Printf("S")
				} else if val == -1 {
					fmt.Printf("X")
				} else if val == 1 {
					fmt.Printf(".")
				} else if val == 2 {
					fmt.Printf("G")
				}
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func (d *RepairDroid) GetNextInput() int {
	// check each neighbor to see if we've visited, if so, move there and pus the opposite direction so we can easily
	// undo the move when backtracking
	if !d.HasVisited(d.Position.X, d.Position.Y-1) {
		d.NextPosition = Point{d.Position.X, d.Position.Y - 1}
		d.Backtrack.Push(2)
		return 1
	}
	if !d.HasVisited(d.Position.X, d.Position.Y+1) {
		d.NextPosition = Point{d.Position.X, d.Position.Y + 1}
		d.Backtrack.Push(1)
		return 2
	}
	if !d.HasVisited(d.Position.X-1, d.Position.Y) {
		d.NextPosition = Point{d.Position.X - 1, d.Position.Y}
		d.Backtrack.Push(4)
		return 3
	}
	if !d.HasVisited(d.Position.X+1, d.Position.Y) {
		d.NextPosition = Point{d.Position.X + 1, d.Position.Y}
		d.Backtrack.Push(3)
		return 4
	}

	if d.Backtrack.Depth() == 0 {
		return 0 // nowhere left to explore
	} else {
		move := d.Backtrack.Pop()
		switch move {
		case 1:
			d.NextPosition = Point{d.Position.X, d.Position.Y - 1}
		case 2:
			d.NextPosition = Point{d.Position.X, d.Position.Y + 1}
		case 3:
			d.NextPosition = Point{d.Position.X - 1, d.Position.Y}
		case 4:
			d.NextPosition = Point{d.Position.X + 1, d.Position.Y}
		}
		return move
	}
}

func (d *RepairDroid) HasVisited(x int, y int) bool {
	_, ok := d.Grid[Point{X: x, Y: y}]
	return ok
}

func (d *RepairDroid) WriteOutput(val int) {
	switch val {
	case 0:
		//wall
		d.Grid[d.NextPosition] = -1
		// need to pop last backtrack since we didn't actually move
		d.Backtrack.Pop()
	case 1:
		d.Grid[d.NextPosition] = 1
		d.Position = d.NextPosition
	case 2:
		// goal
		d.Grid[d.NextPosition] = 2
		d.Position = d.NextPosition
		fmt.Printf("Found Oxygen System at %d,%d\n", d.Position.X, d.Position.Y)
	}
}

func main() {
	robot, goal := part1()
	part2(robot, goal)
}

func part1() (*RepairDroid, Point) {
	content := util.ReadFileAsString("input/day15.input")
	robot := MakeRobot()
	computer := intcomputer.NewComputer(content, 0, 0, false, robot, robot)
	computer.RunToCompletion(nil)
	robot.Print()
	// now we have the maze discovered, find shortest path via Dijkstra's algorithm
	_, goal := shortestPath(robot)
	return robot, goal
}

func part2(robot *RepairDroid, goalPoint Point) {
	spaceToFill := 0
	offsets := []Point{Point{0, -1}, Point{0, 1}, Point{-1, 0}, Point{1, 0}}
	for _, val := range robot.Grid {
		if val != -1 {
			spaceToFill++
		}
	}
	seen := make(map[Point]bool)
	minutes := -1 // start at -1 since the goal space is filled at minute 0
	frontier := []Point{goalPoint}
	for len(seen) < spaceToFill {
		minutes++
		var nextFrontier []Point
		for _, loc := range frontier {
			seen[loc] = true
			for _, offset := range offsets {
				if isVacant(robot.Grid, seen, loc.Add(offset)) {
					nextFrontier = append(nextFrontier, loc.Add(offset))
				}
			}
		}
		frontier = nextFrontier
	}
	fmt.Printf("Took %d minutes to fill with Oxygen", minutes)
}

func isVacant(grid map[Point]int, seen map[Point]bool, loc Point) bool {
	val, ok := seen[loc]
	if ok && val {
		return false
	}
	gridVal, ok := grid[loc]
	if !ok {
		return false
	}
	return gridVal != -1
}

func shortestPath(droid *RepairDroid) (int, Point) {
	defaultVal := 10000000
	offsets := []Point{Point{0, -1}, Point{0, 1}, Point{-1, 0}, Point{1, 0}}
	distances := make(map[Point]int)
	unvisited := make(map[Point]int)
	var goalLoc Point
	for point, val := range droid.Grid {
		if val != -1 {
			unvisited[point] = defaultVal
		}
		if val == 2 {
			goalLoc = point
		}
	}
	current := Point{0, 0}
	distances[current] = 0
	delete(unvisited, current)

	for droid.Grid[current] != 2 {
		for _, offset := range offsets {
			neighbor := current.Add(offset)
			oldDist, ok := unvisited[neighbor]
			if !ok {
				// already visited this neighbor
				continue
			}
			// if the distance to current + 1 is less than the tentative distance to the neighbor, update it
			if distances[current]+1 < oldDist {
				unvisited[neighbor] = distances[current] + 1
			}
		}
		// pick new current node
		min := defaultVal
		for point, dist := range unvisited {
			if dist < min {
				current = point
				min = dist
			}
		}
		// update official distance to new current node
		distances[current] = unvisited[current]
		//remove new current node from unvisited set
		delete(unvisited, current)
	}
	fmt.Printf("Min Path is %d\n", distances[goalLoc])
	return distances[goalLoc], goalLoc
}

func MakeRobot() *RepairDroid {
	droid := RepairDroid{Grid: make(map[Point]int), Position: Point{0, 0}, Backtrack: MoveStack{}}
	droid.Grid[Point{0, 0}] = 1
	return &droid
}
