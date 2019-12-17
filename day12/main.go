package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"math"
	"strconv"
	"strings"
)

type Vector struct {
	v [3]int
}

type Moon struct {
	Velocity Vector
	Position Vector
}

func (v *Vector) String() string {
	return fmt.Sprintf("<x=%d,y=%d,z=%d>", v.v[0], v.v[1], v.v[2])
}

func (v *Vector) Add(v2 Vector) {
	for i := 0; i < len(v.v); i++ {
		v.v[i] += v2.v[i]
	}
}

func (v *Vector) GetEnergy() int {
	sum := 0.0
	for i := 0; i < len(v.v); i++ {
		sum += math.Abs(float64(v.v[i]))
	}
	return int(sum)
}

func (m *Moon) GetEnergy() int {
	return m.Velocity.GetEnergy() * m.Position.GetEnergy()
}

func (m *Moon) ApplyGravity(m2 *Moon) {
	for i := 0; i < len(m2.Velocity.v); i++ {
		if m.Position.v[i] < m2.Position.v[i] {
			m.Velocity.v[i]++
			m2.Velocity.v[i]--
		} else if m.Position.v[i] > m2.Position.v[i] {
			m.Velocity.v[i]--
			m2.Velocity.v[i]++
		}
	}
}

func (m *Moon) Move() {
	m.Position.Add(m.Velocity)
}

func (m *Moon) String() string {
	return fmt.Sprintf("pos=%s, vel=%s", m.Position.String(), m.Velocity.String())
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := util.ReadAllLines("input/day12.input")
	var moons []*Moon
	for _, line := range lines {
		moons = append(moons, NewMoon(line))
	}
	for i := 0; i < 1000; i++ {
		performStep(moons)
	}
	totalEnergy := 0
	for i := 0; i < len(moons); i++ {
		totalEnergy += moons[i].GetEnergy()
	}
	fmt.Printf("Total energy: %d\n", totalEnergy)
}

func part2() {
	lines := util.ReadAllLines("input/day12.input")
	var moons []*Moon
	period := make([]int, 3)
	initialPositions := [4][3]int{}
	for i, line := range lines {
		moon := NewMoon(line)
		moons = append(moons, moon)
		for j := 0; j < 3; j++ {
			initialPositions[i][j] = moon.Position.v[j]
		}
	}

	foundCount := 0

	for count := 1; foundCount < 3; count++ {
		performStep(moons)
		for i := 0; i < 3; i++ {
			allSame := true
			for j := 0; j < len(moons); j++ {
				if moons[j].Position.v[i] != initialPositions[j][i] || moons[j].Velocity.v[i] != 0 {
					allSame = false
					break
				}
			}
			if allSame {
				if period[i] == 0 {
					period[i] = count
					foundCount++
				}
			}
		}

	}
	fmt.Printf("Steps until repeat: %d\n", lcmArr(period))
}

func performStep(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		for j := i + 1; j < len(moons); j++ {
			moons[i].ApplyGravity(moons[j])
		}
	}
	for i := 0; i < len(moons); i++ {
		moons[i].Move()
	}
}

func NewMoon(input string) *Moon {
	input = strings.ReplaceAll(input, "<", "")
	input = strings.ReplaceAll(input, ">", "")
	tokens := strings.Split(input, ",")
	var pos [3]int
	for i := 0; i < len(tokens); i++ {
		parts := strings.Split(tokens[i], "=")
		val, _ := strconv.Atoi(parts[1])
		pos[i] = val
	}
	return &Moon{
		Velocity: Vector{v: [3]int{0, 0, 0}},
		Position: Vector{v: pos}}
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcmArr(a []int) int {
	multiple := lcm(a[0], a[1])
	for i := 2; i < len(a); i++ {
		multiple = lcm(multiple, a[i])
	}
	return multiple
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}
