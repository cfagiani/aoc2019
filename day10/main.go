package main

import (
	"fmt"
	"github.com/cfagiani/aoc2019/util"
	"math"
	"sort"
)

type Asteroid struct {
	X, Y   int
	CanSee map[float64][2]bool
}

func (a *Asteroid) eq(b Asteroid) bool {
	return a.X == b.X && a.Y == b.Y
}

func main() {
	allAsteroids, station := part1()
	part2(allAsteroids, station)
}

func part1() ([]Asteroid, Asteroid) {

	asteroids := computeAllAngles(buildPoints(util.ReadAllLines("input/day10.input")))

	var max int
	var maxAsteroid Asteroid
	for _, a := range asteroids {
		var seenCount int
		for _, v := range a.CanSee {
			if v[0] {
				seenCount++
			}
			if v[1] {
				seenCount++
			}
		}
		if seenCount > max {
			max = seenCount
			maxAsteroid = a
		}
	}
	fmt.Printf("Max asteroids monitored %d", max)
	return asteroids, maxAsteroid
}

func part2(asteroids []Asteroid, station Asteroid) {
	fmt.Printf("\nStation: %d,%d\n", station.X, station.Y)
	size := 4
	asteroidsBySlope := make([]map[float64][]Asteroid, size)
	for i := 0; i < size; i++ {
		asteroidsBySlope[i] = make(map[float64][]Asteroid)
	}
	for _, asteroid := range asteroids {
		if asteroid.eq(station) {
			continue
		}
		q, s := getQuadrantAndSlope(asteroid, station)
		temp := asteroidsBySlope[q][s]
		temp = append(temp, asteroid)
		asteroidsBySlope[q][s] = temp
	}
	// sort each of the slope arrays
	for i := 0; i < size; i++ {
		for j, _ := range asteroidsBySlope[i] {
			sort.Slice(asteroidsBySlope[i][j], func(a int, b int) bool {
				return getDistance(station, asteroidsBySlope[i][j][a]) < getDistance(station, asteroidsBySlope[i][j][b])
			})
		}
	}
	// now go over the sorted list and count the vaporizations
	count := 0
	for i := 0; i < size; i++ {
		var slopes []float64
		for k := range asteroidsBySlope[i] {
			slopes = append(slopes, k)
		}
		sort.Float64s(slopes)
		//sort.Sort(sort.Reverse(sort.Float64Slice(slopes)))

		for _, s := range slopes {
			if len(asteroidsBySlope[i][s]) > 0 {
				count++
				if count == 200 {
					fmt.Printf("\n200th: %d\n", asteroidsBySlope[i][s][0].X*100+asteroidsBySlope[i][s][0].Y)
					return
				} else {
					fmt.Printf("%d,%d\n", asteroidsBySlope[i][s][0].X, asteroidsBySlope[i][s][0].Y)
				}
				asteroidsBySlope[i][s] = asteroidsBySlope[i][s][1:]
			}
		}
	}
}

func buildPoints(lines []string) []Asteroid {
	var points []Asteroid
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				points = append(points, Asteroid{X: x, Y: y, CanSee: make(map[float64][2]bool)})
			}
		}
	}
	return points
}

func computeAllAngles(asteroids []Asteroid) []Asteroid {
	for i, asteroid := range asteroids {
		for j, asteroid2 := range asteroids {
			if i == j {
				continue
			}
			var direction int
			if asteroid.Y < asteroid2.Y || (asteroid2.Y == asteroid.Y && asteroid.X < asteroid2.X) {
				direction = 1
			} else {
				direction = 0
			}

			//horizontal is undefined so set the angle to an arbitrary value
			angle := 9999999999.0
			if (asteroid2.X - asteroid.X) != 0 {
				angle = float64(asteroid2.Y-asteroid.Y) / float64(asteroid2.X-asteroid.X)
			}
			seenArray := asteroids[i].CanSee[angle]
			seenArray[direction] = true
			asteroids[i].CanSee[angle] = seenArray
		}
	}
	return asteroids
}

func getDistance(base Asteroid, asteroid Asteroid) float64 {
	return math.Sqrt(math.Pow(float64(asteroid.X-base.X), 2) + math.Pow(float64(asteroid.Y-base.Y), 2))
}

func getQuadrantAndSlope(a Asteroid, b Asteroid) (int, float64) {
	slope := -10000.0
	if a.X-b.X != 0 {
		slope = float64(a.Y-b.Y) / float64(a.X-b.X)
	}
	q := 0
	if a.X == b.X && a.Y > b.Y {
		q = 1
	} else if a.X == b.X && a.Y <= b.Y {
		q = 0
	} else if a.X > b.X && a.Y < b.Y {
		q = 0
	} else if a.X > b.X && a.Y >= b.Y {
		q = 1
	} else if a.X < b.X && a.Y >= b.Y {
		q = 2
	} else if a.X < b.X && a.Y < b.Y {
		q = 3
	}
	return q, slope
}
