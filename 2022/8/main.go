package main

import (
	"io/ioutil"
	"strings"
)

func reverse(s []int) []int {
	r := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return r
}

func getScenicScore(treeMap [][]int, x, y int) int {
	if x == 0 || x == len(treeMap)-1 || y == 0 || y == len(treeMap[x])-1 {
		return 0
	}
	// viewing distance in each of the four directions
	getScore := func(treeLine []int, height int) int {
		for i, v := range treeLine {
			if v >= height {
				return i + 1
			}
		}
		return len(treeLine)
	}

	score := 1
	height := treeMap[x][y]

	rightTreeLine, leftTreeLine, upTreeLine, downTreeLine := getTreeLine(treeMap, x, y)

	score *= getScore(rightTreeLine, height)
	score *= getScore(leftTreeLine, height)
	score *= getScore(upTreeLine, height)
	score *= getScore(downTreeLine, height)
	return score
}

func getTreeLine(treeMap [][]int, x int, y int) (
	rightTreeLine []int,
	leftTreeLine []int,
	upTreeLine []int,
	downTreeLine []int,
) {
	rightTreeLine = treeMap[x][y+1:]

	leftTreeLine = reverse(treeMap[x][:y])

	row := make([]int, len(treeMap))
	for i := 0; i < len(treeMap); i++ {
		row[i] = treeMap[i][y]
	}
	upTreeLine = reverse(row[:x])
	downTreeLine = row[x+1:]
	return rightTreeLine, leftTreeLine, upTreeLine, downTreeLine
}

func part2(treeMap [][]int) int {
	var highestScore int
	for i := 0; i < len(treeMap); i++ {
		for j := 0; j < len(treeMap[i]); j++ {
			if treeMap[i][j] > 0 {
				score := getScenicScore(treeMap, i, j)
				if score > highestScore {
					highestScore = score
				}
			}
		}
	}
	return highestScore
}

func isTreeVisible(treeMap [][]int, x, y int) bool {
	if x == 0 || x == len(treeMap)-1 || y == 0 || y == len(treeMap[x])-1 {
		return true
	}
	height := treeMap[x][y]

	blockedByTree := func(treeLine []int, height int) bool {
		for _, v := range treeLine {
			if v >= height {
				return true
			}
		}
		return false
	}

	leftTreeLine, rightTreeLine, upTreeLine, downTreeLine := getTreeLine(treeMap, x, y)
	if blockedByTree(leftTreeLine, height) && blockedByTree(rightTreeLine, height) && blockedByTree(upTreeLine, height) && blockedByTree(downTreeLine, height) {
		return false
	}
	return true
}

func part1(treeMap [][]int) int {
	var visibleTrees int
	for i := 0; i < len(treeMap); i++ {
		for j := 0; j < len(treeMap[i]); j++ {
			if isTreeVisible(treeMap, i, j) {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

func main() {
	input, err := ioutil.ReadFile("./2022/8/input.txt")
	if err != nil {
		panic(err)
	}

	inputSlice := strings.Split(string(input), "\r\n")
	treeMap := make([][]int, len(inputSlice))
	for i, s := range inputSlice {
		treeMap[i] = make([]int, len(s))
		for j, c := range s {
			treeMap[i][j] = int(c - '0')
		}
	}
	println("Part 1:", part1(treeMap))
	println("Part 2:", part2(treeMap))
}
