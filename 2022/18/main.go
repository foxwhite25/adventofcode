package main

import (
	"bytes"
	"io/ioutil"
)

type Cube struct {
	x int
	y int
	z int
}

func part2(input []byte) {
	cubes := processCube(input)
	cubeSize := Cube{}
	for _, cube := range cubes {
		if cube.x > cubeSize.x {
			cubeSize.x = cube.x
		}
		if cube.y > cubeSize.y {
			cubeSize.y = cube.y
		}
		if cube.z > cubeSize.z {
			cubeSize.z = cube.z
		}
	}
	cubeSpace := make(map[Cube]bool)
	for _, cube := range cubes {
		cubeSpace[cube] = true
	}
	exterior := make(map[Cube]bool)
	markExterior(Cube{x: 0, y: 0, z: 0}, cubeSpace, exterior, cubeSize)
	// Find the total exterior surface area of the cubes
	var totalSurfaceArea int
	for _, cube := range cubes {
		var surfaceArea int
		for _, neighbor := range neighbors(cube) {
			if !cubeSpace[neighbor] && exterior[neighbor] {
				surfaceArea++
			}
		}
		totalSurfaceArea += surfaceArea
	}
	println("Part 2:", totalSurfaceArea)
}

func markExterior(cube Cube, cubeSpace map[Cube]bool, exterior map[Cube]bool, cubeSize Cube) {
	if exterior[cube] {
		return
	}
	if cube.x < -1 || cube.x > cubeSize.x+1 || cube.y < -1 || cube.y > cubeSize.y+1 || cube.z < -1 || cube.z > cubeSize.z+1 {
		return
	}
	if cubeSpace[cube] {
		return
	}
	exterior[cube] = true
	for _, neighbor := range neighbors(cube) {
		markExterior(neighbor, cubeSpace, exterior, cubeSize)
	}
}

func neighbors(cube Cube) []Cube {
	return []Cube{
		{x: cube.x - 1, y: cube.y, z: cube.z},
		{x: cube.x + 1, y: cube.y, z: cube.z},
		{x: cube.x, y: cube.y - 1, z: cube.z},
		{x: cube.x, y: cube.y + 1, z: cube.z},
		{x: cube.x, y: cube.y, z: cube.z - 1},
		{x: cube.x, y: cube.y, z: cube.z + 1},
	}
}

func part1(input []byte) {
	cubes := processCube(input)
	cubeSpace := make(map[Cube]bool)
	for _, cube := range cubes {
		cubeSpace[cube] = true
	}
	// Find the total surface area of the cubes
	// Each cube is 1x1x1
	// Each cube has 6 sides
	var totalSurfaceArea int
	for _, cube := range cubes {
		var surfaceArea int
		for _, neighbor := range neighbors(cube) {
			if !cubeSpace[neighbor] {
				surfaceArea++
			}
		}
		totalSurfaceArea += surfaceArea
	}
	println("Part 1:", totalSurfaceArea)
}

func processCube(input []byte) []Cube {
	var cubes []Cube
	for _, line := range bytes.Split(bytes.TrimSpace(input), []byte{'\n'}) {
		coordinates := bytes.Split(bytes.TrimSpace(line), []byte{','})
		if len(coordinates) != 3 {
			panic("invalid input")
		}
		cubes = append(cubes, Cube{
			x: atoi(coordinates[0]),
			y: atoi(coordinates[1]),
			z: atoi(coordinates[2]),
		})
	}
	return cubes
}

func atoi(b []byte) int {
	var n int
	for _, c := range b {
		n = n*10 + int(c-'0')
	}
	return n
}

func main() {
	input, err := ioutil.ReadFile("./2022/18/input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}
