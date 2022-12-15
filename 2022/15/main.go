package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
)

type SensorBeaconPair struct {
	Sensor   Coordinate
	Beacon   Coordinate
	Distance int
}

type Coordinate struct {
	X int
	Y int
}

func (sbp SensorBeaconPair) canContainUnseenPoint(min Coordinate, max Coordinate) bool {
	corners := []Coordinate{
		{X: min.X, Y: min.Y},
		{X: min.X, Y: max.Y},
		{X: max.X, Y: min.Y},
		{X: max.X, Y: max.Y},
	}
	for _, corner := range corners {
		distance := abs(sbp.Sensor.X-corner.X) + abs(sbp.Sensor.Y-corner.Y)
		if distance > sbp.Distance {
			return true
		}
	}
	return false
}

var inputRegex = regexp.MustCompile(`Sensor at x=([-\d]+), y=([-\d]+): closest beacon is at x=([-\d]+), y=([-\d]+)`)

func part2(input []byte) {
	match := inputRegex.FindAllStringSubmatch(string(input), -1)
	pairs := make([]SensorBeaconPair, len(match))
	min := Coordinate{X: 0, Y: 0}
	max := Coordinate{X: 4000000, Y: 4000000}
	for i, m := range match {
		sensorX, _ := strconv.Atoi(m[1])
		sensorY, _ := strconv.Atoi(m[2])
		beaconX, _ := strconv.Atoi(m[3])
		beaconY, _ := strconv.Atoi(m[4])

		pairs[i] = SensorBeaconPair{
			Sensor:   Coordinate{X: sensorX, Y: sensorY},
			Beacon:   Coordinate{X: beaconX, Y: beaconY},
			Distance: abs(sensorX-beaconX) + abs(sensorY-beaconY),
		}
	}
	unseenPoint := findUnseenPoints(pairs, min, max)
	println("Part 2:", unseenPoint.X*4000000+unseenPoint.Y)
}

func findUnseenPoints(pairs []SensorBeaconPair, min Coordinate, max Coordinate) Coordinate {
	if min == max {
		return min
	}
	mid := Coordinate{
		X: (min.X + max.X) / 2,
		Y: (min.Y + max.Y) / 2,
	}

	quadrants := make([][]Coordinate, 4)
	quadrants[0] = []Coordinate{min, mid}
	quadrants[1] = []Coordinate{Coordinate{X: mid.X + 1, Y: min.Y}, Coordinate{X: max.X, Y: mid.Y}}
	quadrants[2] = []Coordinate{Coordinate{X: min.X, Y: mid.Y + 1}, Coordinate{X: mid.X, Y: max.Y}}
	quadrants[3] = []Coordinate{Coordinate{X: mid.X + 1, Y: mid.Y + 1}, max}

	for _, quadrant := range quadrants {
		if quadrant[0].X > quadrant[1].X || quadrant[0].Y > quadrant[1].Y {
			continue
		}

		allPairsCanContain := true
		for _, pair := range pairs {
			if !pair.canContainUnseenPoint(quadrant[0], quadrant[1]) {
				allPairsCanContain = false
				break
			}
		}
		if allPairsCanContain {
			k := findUnseenPoints(pairs, quadrant[0], quadrant[1])
			if k.X != -1 || k.Y != -1 {
				return k
			}
		}
	}
	return Coordinate{X: -1, Y: -1}
}

func part1(input []byte) {
	testY := 2000000
	matches := inputRegex.FindAllStringSubmatch(string(input), -1)
	block := make(map[int]bool)
	structures := make(map[int]bool)

	for _, match := range matches {
		if len(match) != 5 {
			panic("Invalid match")
		}
		sensorX, _ := strconv.Atoi(match[1])
		sensorY, _ := strconv.Atoi(match[2])
		beaconX, _ := strconv.Atoi(match[3])
		beaconY, _ := strconv.Atoi(match[4])

		if sensorY == testY {
			structures[sensorX] = true
		}
		if beaconY == testY {
			structures[beaconX] = true
		}

		radius := abs(sensorX-beaconX) + abs(sensorY-beaconY)
		distanceTo200k := abs(testY - sensorY)
		radiusAt200k := radius - distanceTo200k

		blockStart := sensorX - radiusAt200k
		blockEnd := sensorX + radiusAt200k

		for i := blockStart; i <= blockEnd; i++ {
			block[i] = true
		}
	}

	for i, _ := range structures {
		if _, ok := block[i]; ok {
			delete(block, i)
		}
	}

	println("Part 1:", len(block))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	input, err := ioutil.ReadFile("./2022/15/input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}
