package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func part2(input []byte) {
	numberOfSand := 0
	sandPosition := Coordinates{500, 0}
	board, height := parseInput(input)
	for i := 0; i < 1000; i++ {
		board[i][height+2] = true
	}
	defaultPosition := Coordinates{500, 0}

L:
	for {
		sandPosition = Coordinates{500, 0}
		for {
			newSandPosition, _ := tickSand(board, sandPosition)
			if newSandPosition == defaultPosition {
				numberOfSand++
				break L
			}
			if newSandPosition == sandPosition {
				board[sandPosition.X][sandPosition.Y] = true
				numberOfSand++
				break
			}
			sandPosition = newSandPosition
		}
	}
	println("Part 2:", numberOfSand)
}

func part1(input []byte) {
	numberOfSand := 0
	sandPosition := Coordinates{500, 0}
	board, _ := parseInput(input)
L:
	for {
		sandPosition = Coordinates{500, 0}
		for {
			newSandPosition, oob := tickSand(board, sandPosition)
			if oob {
				break L
			}
			if newSandPosition == sandPosition {
				board[sandPosition.X][sandPosition.Y] = true
				numberOfSand++
				break
			}
			sandPosition = newSandPosition
		}
	}
	println("Part 1:", numberOfSand)
}

func tickSand(blocked [][]bool, current Coordinates) (next Coordinates, outOfBounds bool) {
	if current.Y+1 >= len(blocked[0]) {
		return current, true
	}
	if !blocked[current.X][current.Y+1] {
		return Coordinates{current.X, current.Y + 1}, false
	}
	if !blocked[current.X-1][current.Y+1] {
		return Coordinates{current.X - 1, current.Y + 1}, false
	}
	if !blocked[current.X+1][current.Y+1] {
		return Coordinates{current.X + 1, current.Y + 1}, false
	}
	return current, false
}

func parseInput(input []byte) ([][]bool, int) {
	board := make([][]bool, 1000)
	height := 0
	for i := range board {
		board[i] = make([]bool, 1000)
	}
	for _, s := range strings.Split(string(input), "\n") {
		if s == "" {
			continue
		}
		var lineNodes []Coordinates
		for _, s2 := range strings.Split(s, "->") {
			if s2 == "" {
				continue
			}
			tmp := strings.Split(s2, ",")
			if len(tmp) != 2 {
				panic("Invalid input")
			}
			x, err := strconv.Atoi(strings.TrimSpace(tmp[0]))
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(strings.TrimSpace(tmp[1]))
			if err != nil {
				panic(err)
			}
			if y > height {
				height = y
			}
			lineNodes = append(lineNodes, Coordinates{x, y})
		}
		for i := 0; i < len(lineNodes)-1; i++ {
			if lineNodes[i].X == lineNodes[i+1].X {
				for _, y := range between(lineNodes[i].Y, lineNodes[i+1].Y) {
					board[lineNodes[i].X][y] = true
				}
			} else {
				for _, x := range between(lineNodes[i].X, lineNodes[i+1].X) {
					board[x][lineNodes[i].Y] = true
				}
			}
		}
	}
	return board, height
}

func between(a, b int) []int {
	result := make([]int, 0)
	if a > b {
		b, a = a, b
	}
	for i := a; i <= b; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	input, err := ioutil.ReadFile("./2022/14/input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}
