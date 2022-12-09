package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

func followKnots(head Coordinate, tail Coordinate) Coordinate {
	deltaX := head.X - tail.X
	deltaY := head.Y - tail.Y

	if abs(deltaX) > 1 || abs(deltaY) > 1 {
		return Coordinate{tail.X + sign(deltaX), tail.Y + sign(deltaY)}
	}
	return tail
}

type Data struct {
	coords []Coordinate
	grid   []Coordinate
}

func part2(input []byte) []Data {
	regex := regexp.MustCompile(`(\w) (\d+)`)
	knotsXY := make([]Coordinate, 10)
	grid := make(map[Coordinate]bool)
	grid[Coordinate{0, 0}] = true

	var result []Data

	for _, s := range strings.Split(string(input), "\n") {
		match := regex.FindStringSubmatch(s)
		direction := match[1]
		distance, _ := strconv.Atoi(match[2])
		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				knotsXY[0].Y++
			case "D":
				knotsXY[0].Y--
			case "L":
				knotsXY[0].X--
			case "R":
				knotsXY[0].X++
			}

			for i := 1; i < 10; i++ {
				knotsXY[i] = followKnots(knotsXY[i-1], knotsXY[i])
			}
			grid[knotsXY[9]] = true
		}
		temp := make([]Coordinate, 10)
		copy(temp, knotsXY)
		temp2 := make([]Coordinate, len(grid))
		for coordinate, b := range grid {
			if b {
				temp2 = append(temp2, coordinate)
			}
		}

		result = append(result, Data{
			coords: temp,
			grid:   temp2,
		})
	}

	println("Part 2:", len(grid))
	return result
}

func part1(input []byte) {
	regex := regexp.MustCompile(`(\w) (\d+)`)
	head := Coordinate{0, 0}
	tail := Coordinate{0, 0}
	grid := make(map[Coordinate]bool)
	grid[head] = true

	for _, s := range strings.Split(string(input), "\n") {
		match := regex.FindStringSubmatch(s)
		direction := match[1]
		distance, _ := strconv.Atoi(match[2])
		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "L":
				head.X--
			case "R":
				head.X++
			}
			tail = followKnots(head, tail)
			grid[tail] = true
		}
	}
	println("Part 1:", len(grid))
}

func main() {
	input, err := ioutil.ReadFile("./2022/9/input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	result := part2(input)
	saveToImages(result)
}
