package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
)

var inputRegex = regexp.MustCompile(`(?m)^Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.$`)

type Blueprint struct {
	oreRobotCost           int
	clayRobotCost          int
	obsidianRobotOreCost   int
	obsidianRobotClayCost  int
	geodeRobotOreCost      int
	geodeRobotObsidianCost int
}

func part2(input []byte) {
	println("Part 2:")
}

func part1(input []byte) {
	blueprints := processBlueprint(input)
	qualitySum := 0
	resultChan := make(chan int)
	for i, blueprint := range blueprints {
		i := i
		blueprint := blueprint
		go func() {
			finalState := simulateBlueprint(blueprint, State{
				oreRobotCount: 1,
			})
			quality := finalState.geodeCount * (i + 1)
			resultChan <- quality
		}()
	}
	for range blueprints {
		qualitySum += <-resultChan
	}
	println("Part 1:", qualitySum)
}

type State struct {
	oreRobotCount      int
	clayRobotCount     int
	obsidianRobotCount int
	geodeRobotCount    int

	oreCount      int
	clayCount     int
	obsidianCount int
	geodeCount    int

	timeTicked int
}

// Use recursive function to find all possible combinations of robots
func simulateBlueprint(blueprint Blueprint, state State) (finalState State) {
	state.timeTicked++
	state.oreCount += state.oreRobotCount
	state.clayCount += state.clayRobotCount
	state.obsidianCount += state.obsidianRobotCount
	state.geodeCount += state.geodeRobotCount
	if state.timeTicked == 24 {
		return state
	}

	possibleNextStates := []State{state}
	// Check if we can afford to buy a new robot
	if state.oreCount >= blueprint.oreRobotCost {
		possibleNextStates = append(possibleNextStates, State{
			oreRobotCount:      state.oreRobotCount + 1,
			clayRobotCount:     state.clayRobotCount,
			obsidianRobotCount: state.obsidianRobotCount,
			geodeRobotCount:    state.geodeRobotCount,
			oreCount:           state.oreCount - blueprint.oreRobotCost,
			clayCount:          state.clayCount,
			obsidianCount:      state.obsidianCount,
			geodeCount:         state.geodeCount,
			timeTicked:         state.timeTicked,
		})
	}
	if state.oreCount >= blueprint.clayRobotCost {
		possibleNextStates = append(possibleNextStates, State{
			oreRobotCount:      state.oreRobotCount,
			clayRobotCount:     state.clayRobotCount + 1,
			obsidianRobotCount: state.obsidianRobotCount,
			geodeRobotCount:    state.geodeRobotCount,
			oreCount:           state.oreCount - blueprint.clayRobotCost,
			clayCount:          state.clayCount,
			obsidianCount:      state.obsidianCount,
			geodeCount:         state.geodeCount,
			timeTicked:         state.timeTicked,
		})
	}
	if state.oreCount >= blueprint.obsidianRobotOreCost && state.clayCount >= blueprint.obsidianRobotClayCost {
		possibleNextStates = []State{{
			oreRobotCount:      state.oreRobotCount,
			clayRobotCount:     state.clayRobotCount,
			obsidianRobotCount: state.obsidianRobotCount + 1,
			geodeRobotCount:    state.geodeRobotCount,
			oreCount:           state.oreCount - blueprint.obsidianRobotOreCost,
			clayCount:          state.clayCount - blueprint.obsidianRobotClayCost,
			obsidianCount:      state.obsidianCount,
			geodeCount:         state.geodeCount,
			timeTicked:         state.timeTicked,
		}}
	}
	if state.oreCount >= blueprint.geodeRobotOreCost && state.obsidianCount >= blueprint.geodeRobotObsidianCost {
		possibleNextStates = []State{{
			oreRobotCount:      state.oreRobotCount,
			clayRobotCount:     state.clayRobotCount,
			obsidianRobotCount: state.obsidianRobotCount,
			geodeRobotCount:    state.geodeRobotCount + 1,
			oreCount:           state.oreCount - blueprint.geodeRobotOreCost,
			clayCount:          state.clayCount,
			obsidianCount:      state.obsidianCount - blueprint.geodeRobotObsidianCost,
			geodeCount:         state.geodeCount,
			timeTicked:         state.timeTicked,
		}}
	}
	var maxGeodeState State
	for _, nextState := range possibleNextStates {
		finalState = simulateBlueprint(blueprint, nextState)
		if finalState.geodeCount > maxGeodeState.geodeCount {
			maxGeodeState = finalState
		}
	}
	return maxGeodeState
}

func processBlueprint(input []byte) []Blueprint {
	var blueprints []Blueprint
	for _, match := range inputRegex.FindAllSubmatch(input, -1) {
		blueprints = append(blueprints, Blueprint{
			oreRobotCost:           parseInt(match[2]),
			clayRobotCost:          parseInt(match[3]),
			obsidianRobotOreCost:   parseInt(match[4]),
			obsidianRobotClayCost:  parseInt(match[5]),
			geodeRobotOreCost:      parseInt(match[6]),
			geodeRobotObsidianCost: parseInt(match[7]),
		})
	}
	return blueprints
}

func parseInt(b []byte) int {
	n, _ := strconv.Atoi(string(b))
	return n
}

func main() {
	input, err := ioutil.ReadFile("./2022/19/input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}
