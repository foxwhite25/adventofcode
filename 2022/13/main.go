package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	nodeType string
	rep      string
	value    int
	children []*Node
	special  int
}

func compareNodes(n1 *Node, n2 *Node) int {
	if n1.nodeType == "number" && n2.nodeType == "number" {
		if n1.value < n2.value {
			return 1
		}
		if n1.value > n2.value {
			return -1
		}
		return 0
	}
	if n1.nodeType == "number" {
		n1.nodeType = "list"
		n1.children = []*Node{{nodeType: "number", value: n1.value}}
	}
	if n2.nodeType == "number" {
		n2.nodeType = "list"
		n2.children = []*Node{{nodeType: "number", value: n2.value}}
	}
	smaller := len(n1.children)
	if len(n2.children) < smaller {
		smaller = len(n2.children)
	}

	for i := 0; i < smaller; i++ {
		check := compareNodes(n1.children[i], n2.children[i])
		if check == 0 {
			continue
		}
		return check
	}
	if len(n1.children) < len(n2.children) {
		return 1
	}
	if len(n1.children) > len(n2.children) {
		return -1
	}
	return 0
}

func parseList(input string) *Node {
	n := &Node{nodeType: "list", children: []*Node{}, rep: input}
	for i := 1; i < len(input); i++ {
		if input[i] == '[' {
			open := 1
			for j := i + 1; j < len(input); j++ {
				if input[j] == '[' {
					open++
				}
				if input[j] == ']' {
					open--
				}
				if open == 0 {
					n.children = append(n.children, parseList(input[i:j+1]))
					i = j
					break
				}
			}
		} else if input[i] == ',' {
			continue
		} else {
			for j := i + 1; j < len(input); j++ {
				if input[j] == ',' || input[j] == ']' {
					n.children = append(n.children, parseNumber(input[i:j]))
					i = j
					break
				}
			}
		}
	}
	return n
}

func parseNumber(input string) *Node {
	n := &Node{nodeType: "number", rep: input}
	n.value, _ = strconv.Atoi(input)
	return n
}

func part2(input []byte) {
	var nodes []*Node
	for _, s := range strings.Split(string(input), "\n") {
		if s == "" {
			continue
		}
		nodes = append(nodes, parseList(strings.TrimSpace(s)))
	}
	nodes = append(nodes, &Node{nodeType: "list", children: []*Node{
		{nodeType: "list", children: []*Node{
			{nodeType: "number", value: 2},
		}},
	}, special: 2})
	nodes = append(nodes, &Node{nodeType: "list", children: []*Node{
		{nodeType: "list", children: []*Node{
			{nodeType: "number", value: 6},
		}},
	}, special: 6})

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			check := compareNodes(nodes[i], nodes[j])
			if check == -1 {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}

	sixIndex := 0
	twoIndex := 0
	for i, n := range nodes {
		if n.special == 6 {
			sixIndex = i + 1
		}
		if n.special == 2 {
			twoIndex = i + 1
		}
	}

	println("Part 2:", sixIndex*twoIndex)
}

func part1(input []byte) {
	sum := 0

	for j, s := range strings.Split(string(input), "\n\n") {
		tmp := strings.Split(s, "\n")
		left := parseList(tmp[0])
		right := parseList(tmp[1])

		check := compareNodes(left, right)
		if check == 1 {
			sum += j + 1
			continue
		}
	}
	println("Part 1:", sum)
}

func main() {
	input, err := ioutil.ReadFile("./2022/13/input.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}
