package main

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var inputRegex = regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([\w ,]+)`)

type Valve struct {
	Name     string
	Flow     int
	Next     []string
	Distance map[string]int
}

func part(part int, input []byte) {
	valves := make(map[string]*Valve, 0)
	for _, line := range inputRegex.FindAllStringSubmatch(string(input), -1) {
		name := line[1]
		flow, _ := strconv.Atoi(line[2])
		valves[name] = &Valve{
			Name: name,
			Flow: flow,
			Next: strings.Split(line[3], ", "),
		}
	}

	// calculate distances
	for _, v := range valves {
		v.Distance = make(map[string]int, 0)
		for _, n := range v.Next {
			v.Distance[n] = 1
		}
		queue := make([]string, 0)
		for k := range v.Distance {
			queue = append(queue, k)
		}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			for _, n := range valves[current].Next {
				if _, ok := v.Distance[n]; !ok {
					v.Distance[n] = v.Distance[current] + 1
					queue = append(queue, n)
				}
			}
		}
	}

	nonZero := make([]*Valve, 0)
	for _, v := range valves {
		if v.Flow != 0 {
			nonZero = append(nonZero, v)
		}
	}
	sort.SliceStable(nonZero, func(i, j int) bool {
		return nonZero[i].Flow > nonZero[j].Flow
	})

	if part == 1 {
		println("Part 1:", maxPressureUnder(0, 0, 0, valves["AA"], nonZero, 30))
	} else {
		max := 0
		for i := 0; i < len(nonZero); i++ {
			for _, c := range combinations(nonZero, i) {
				score := maxPressureUnder(0, 0, 0, valves["AA"], c, 26)
				score2P := maxPressureUnder(0, 0, 0, valves["AA"], minus(nonZero, c), 26)
				if score+score2P > max {
					max = score + score2P
				}
			}
		}
		println("Part 2:", max)
	}
}

func minus(fullSet, part []*Valve) []*Valve {
	// return a copy of the list without the item
	newList := make([]*Valve, 0)
	for _, v := range fullSet {
		if !contains(part, v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func contains(list []*Valve, item *Valve) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func combinations(n []*Valve, k int) [][]*Valve {
	if k == 0 {
		return [][]*Valve{{}}
	}
	if len(n) == 0 {
		return nil
	}
	var out [][]*Valve
	for i, v := range n {
		for _, c := range combinations(n[i+1:], k-1) {
			out = append(out, append([]*Valve{v}, c...))
		}
	}
	return out
}

func maxPressureUnder(currentTime int, currentPressure int, currentFlow int, currentTunnel *Valve, remaining []*Valve, timeLimit int) int {
	nScore := currentPressure + (timeLimit-currentTime)*currentFlow
	max := nScore

	for _, v := range remaining {
		distanceAndOpen := currentTunnel.Distance[v.Name] + 1
		if currentTime+distanceAndOpen < timeLimit {
			newTime := currentTime + distanceAndOpen
			newPressure := currentPressure + distanceAndOpen*currentFlow
			newFlow := currentFlow + v.Flow

			possibleScore := maxPressureUnder(newTime, newPressure, newFlow, v, removeFromList(remaining, v), timeLimit)
			if possibleScore > max {
				max = possibleScore
			}
		}
	}

	return max
}

func removeFromList(list []*Valve, item *Valve) []*Valve {
	// return a copy of the list without the item
	newList := make([]*Valve, 0)
	for _, v := range list {
		if v != item {
			newList = append(newList, v)
		}
	}
	return newList
}

func nameToIndex(name string) int {
	return int(name[0]-'A')*26 + int(name[1]-'A')
}

func main() {
	input, err := ioutil.ReadFile("./2022/16/input.txt")
	if err != nil {
		panic(err)
	}
	part(1, input)
	part(2, input)
}
