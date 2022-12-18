package main

import (
	"bytes"
	"io/ioutil"
)

type RotatingSlice[T any] struct {
	inner []T
	index int
}

func (r *RotatingSlice[T]) PushBack(v T) {
	r.inner = append(r.inner, v)
}

func (r *RotatingSlice[T]) Next() T {
	r.index++
	if r.index >= len(r.inner) {
		r.index = 0
	}
	return r.inner[r.index]
}

type TetrisPieces struct {
	shape  [][]bool
	width  int
	height int
}

type Coordinate struct {
	x int
	y int
}

var k = []TetrisPieces{
	{
		shape: [][]bool{
			{true, true, true, true},
		},
		width:  4,
		height: 1,
	},
	{
		shape: [][]bool{
			{false, true, false},
			{true, true, true},
			{false, true, false},
		},
		width:  3,
		height: 3,
	},
	{
		shape: [][]bool{
			{true, true, true},
			{false, false, true},
			{false, false, true},
		},
		width:  3,
		height: 3,
	},
	{
		shape: [][]bool{
			{true},
			{true},
			{true},
			{true},
		},
		width:  1,
		height: 4,
	},
	{
		shape: [][]bool{
			{true, true},
			{true, true},
		},
		width:  2,
		height: 2,
	},
}
var Pieces = RotatingSlice[TetrisPieces]{inner: k, index: -1}

func part2(input []byte) {
	var board [][7]bool
	wind := parseWind(input)
	Pieces.index = -1
	record := make(map[int]map[int][][20][7]bool)
	for i := 0; i < 100000; i++ {
		board = dropPieces(board, &wind)
		if record[wind.index] == nil {
			record[wind.index] = make(map[int][][20][7]bool)
		}
		record[wind.index][Pieces.index] = append(record[wind.index][Pieces.index], [20][7]bool{})
		copy(record[wind.index][Pieces.index][len(record[wind.index][Pieces.index])-1][:], board)
	}

	recordCounter := make(map[int]int)
	for _, m := range record {
		for _, i := range m {
			if len(i) > 1 {
				recordCounter[len(i)]++
			}
		}
	}
	// find the largest number in counter
	var cycle int
	for k := range recordCounter {
		if k > cycle {
			cycle = k
		}
	}

	writeBoardToFile(board)
}

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func (p TetrisPieces) TryMove(board [][7]bool, pieceBottomLeft Coordinate, direction string) (bool, Coordinate) {
	nextPieceBottomLeft := Coordinate{x: pieceBottomLeft.x, y: pieceBottomLeft.y}
	switch direction {
	case "left":
		nextPieceBottomLeft.x--
	case "right":
		nextPieceBottomLeft.x++
	case "down":
		nextPieceBottomLeft.y--
	}
	if nextPieceBottomLeft.x < 0 || nextPieceBottomLeft.x+p.width > 7 {
		return false, pieceBottomLeft
	}
	if nextPieceBottomLeft.y < 0 {
		return false, pieceBottomLeft
	}

	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			if p.shape[y][x] {
				if len(board) <= nextPieceBottomLeft.y+y {
					continue
				}
				if nextPieceBottomLeft.x+x < 0 || nextPieceBottomLeft.x+x >= 7 {
					return false, pieceBottomLeft
				}
				if board[nextPieceBottomLeft.y+y][nextPieceBottomLeft.x+x] {
					return false, pieceBottomLeft
				}
			}
		}
	}
	return true, nextPieceBottomLeft
}

func part1(input []byte) {
	var board [][7]bool
	wind := parseWind(input)
	for i := 0; i < 2022; i++ {
		board = dropPieces(board, &wind)
	}
	println("Part 1:", len(board))
}

func dropPieces(board [][7]bool, wind *RotatingSlice[string]) [][7]bool {
	piece := Pieces.Next()
	pieceBottomLeft := Coordinate{x: 2, y: len(board) + 3}
	for {
		wind := wind.Next()
		ok, newCoordinate := piece.TryMove(board, pieceBottomLeft, wind)
		if ok {
			pieceBottomLeft = newCoordinate
		}

		ok, newCoordinate = piece.TryMove(board, pieceBottomLeft, "down")
		if ok {
			pieceBottomLeft = newCoordinate
		} else {
			break
		}
	}
	if pieceBottomLeft.y+piece.height > len(board) {
		board = appendUntilLength(board, pieceBottomLeft.y+piece.height)
	}
	for y := 0; y < piece.height; y++ {
		for x := 0; x < piece.width; x++ {
			if piece.shape[y][x] {
				board[pieceBottomLeft.y+y][pieceBottomLeft.x+x] = true
			}
		}
	}
	return board
}

func writeBoardToFile(board [][7]bool) {
	var buffer bytes.Buffer
	for y := 0; y < len(board); y++ {
		for x := 0; x < 7; x++ {
			if board[y][x] {
				buffer.WriteString("X")
			} else {
				buffer.WriteString(".")
			}
		}
		buffer.WriteString("\n")
	}
	ioutil.WriteFile("./2022/17/board.txt", buffer.Bytes(), 0644)
}

func appendUntilLength(s [][7]bool, length int) [][7]bool {
	for len(s) < length {
		s = append(s, [7]bool{})
	}
	return s
}

func parseWind(input []byte) RotatingSlice[string] {
	var winds RotatingSlice[string]
	winds.index = -1
	for _, c := range bytes.TrimSpace(input) {
		if c == '<' {
			winds.PushBack("left")
		} else {
			winds.PushBack("right")
		}
	}
	return winds
}

func main() {
	input, err := ioutil.ReadFile("./2022/17/input.txt")
	if err != nil {
		panic(err)
	}
	//part1(input)
	part2(input)
}
