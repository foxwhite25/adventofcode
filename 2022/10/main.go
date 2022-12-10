package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Computer struct {
	cycles      int
	crtPosition int
	x           int
	totalSignal int
	crtOutput   string
}

func (c *Computer) TickCycle(count int) {
	for i := 0; i < count; i++ {
		c.cycles++
		if c.crtPosition >= c.x-1 && c.crtPosition <= c.x+1 {
			c.crtOutput += "█"
		} else {
			c.crtOutput += " "
		}

		if c.crtPosition%40 == 0 && c.crtPosition != 0 {
			c.crtOutput += "\n█"
			c.crtPosition = 0
		}
		c.crtPosition++

		if (c.cycles-20)%40 == 0 {
			c.totalSignal += c.cycles * c.x
		}
	}
}

func (c *Computer) ProcessInstruction(instruction string, argument int) {
	switch instruction {
	case "noop":
		c.TickCycle(1)
	case "addx":
		c.TickCycle(2)
		c.x += argument
	}
}

func part2(input []byte) {
	computer := processInput(input)
	println("Part 2:")
	println(computer.crtOutput)
}

func part1(input []byte) {
	computer := processInput(input)
	println("Part 1:", computer.totalSignal)
}

func processInput(input []byte) Computer {
	computer := Computer{}
	computer.x = 1
	for _, s := range strings.Split(string(input), "\n") {
		args := strings.Split(s, " ")
		instruction := args[0]
		var argument int
		if len(args) > 1 {
			argument, _ = strconv.Atoi(args[1])
		}
		computer.ProcessInstruction(instruction, argument)
	}
	return computer
}

func main() {
	input, err := ioutil.ReadFile("./2022/10/input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}
