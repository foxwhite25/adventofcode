package main

import (
	"github.com/RyanCarrier/dijkstra"
	"io/ioutil"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func part2(input []byte) {
	heightMap, start, end := processInput(input)
	lowest := 9999999
	for y, line := range heightMap {
		for x, height := range line {
			if height == 'a'-96 {
				start = Coordinate{X: x, Y: y}
				length := findShortestPathToHighest(heightMap, start, end)
				if length < lowest && length > 0 {
					lowest = length
				}
			}
		}
	}
	println("Part 2:", lowest)
}

func part1(input []byte) {
	heightMap, start, end := processInput(input)
	println("Part 1:", findShortestPathToHighest(heightMap, start, end))
}

func findShortestPathToHighest(heightMap [][]int, start Coordinate, end Coordinate) int {
	//you should do it in as few steps as possible
	//the elevation of the destination square can be at most one higher than the elevation of your current square
	//you can move exactly one square up, down, left, or right
	coordinateHash := func(c Coordinate) int {
		return c.X*1000 + c.Y
	}
	graph := dijkstra.NewGraph()
	for y, line := range heightMap {
		for x, _ := range line {
			graph.AddVertex(coordinateHash(Coordinate{X: x, Y: y}))
		}
	}

	for y, line := range heightMap {
		for x, height := range line {
			neighbors := getNeighbors(heightMap, x, y)
			for _, neighbor := range neighbors {
				if heightMap[neighbor.Y][neighbor.X] <= height+1 {
					err := graph.AddArc(coordinateHash(Coordinate{X: x, Y: y}), coordinateHash(neighbor), 1)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
	path, _ := graph.Shortest(coordinateHash(start), coordinateHash(end))
	return int(path.Distance)
}

func getNeighbors(heightMap [][]int, x int, y int) []Coordinate {
	var neighbors []Coordinate
	if y > 0 {
		neighbors = append(neighbors, Coordinate{X: x, Y: y - 1})
	}
	if y < len(heightMap)-1 {
		neighbors = append(neighbors, Coordinate{X: x, Y: y + 1})
	}
	if x > 0 {
		neighbors = append(neighbors, Coordinate{X: x - 1, Y: y})
	}
	if x < len(heightMap[y])-1 {
		neighbors = append(neighbors, Coordinate{X: x + 1, Y: y})
	}
	return neighbors
}

func processInput(input []byte) (heightMap [][]int, start Coordinate, end Coordinate) {
	for y, line := range strings.Split(string(input), "\n") {
		heightMap = append(heightMap, []int{})
		for x, char := range line {
			if char == 'S' {
				start = Coordinate{X: x, Y: y}
				heightMap[y] = append(heightMap[y], 'a'-96)
			} else if char == 'E' {
				end = Coordinate{X: x, Y: y}
				heightMap[y] = append(heightMap[y], 'z'-96)
			} else {
				heightMap[y] = append(heightMap[y], int(char)-96)
			}
		}
	}
	return heightMap, start, end
}

func main() {
	input, err := ioutil.ReadFile("./2022/12/input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}
