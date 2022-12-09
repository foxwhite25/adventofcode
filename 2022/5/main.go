package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Reverse() {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	part2 := true
	input, err := ioutil.ReadFile("./2022/5/input.txt")
	if err != nil {
		panic(err)
	}

	tmpSplit := strings.Split(string(input), "\r\n\r\n")
	if len(tmpSplit) != 2 {
		panic("Invalid input")
	}
	initialState := strings.ReplaceAll(tmpSplit[0], "    ", " [0]")
	var stacks []Stack
	for i := 0; i < 9; i++ {
		stacks = append(stacks, Stack{})
	}

	for _, s := range strings.Split(initialState, "\r\n") {
		for i, s2 := range strings.Split(strings.TrimSpace(s), " ") {
			if !strings.Contains(s2, "[") {
				continue
			}
			val := strings.Trim(s2, "[]")
			if val == "0" {
				continue
			}
			stacks[i].Push(val)
		}
	}
	for _, stack := range stacks {
		stack.Reverse()
	}

	instructions := strings.Split(tmpSplit[1], "\r\n")
	regex := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	for _, instruction := range instructions {
		matches := regex.FindStringSubmatch(instruction)
		if len(matches) != 4 {
			panic("Invalid instruction" + instruction)
		}
		count, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])
		from--
		to--
		if !part2 {
			for i := 0; i < count; i++ {
				val, ok := stacks[from].Pop()
				if !ok {
					continue
				}
				stacks[to].Push(val)
			}
		} else {
			var tmp Stack
			for i := 0; i < count; i++ {
				val, ok := stacks[from].Pop()
				if !ok {
					continue
				}
				tmp.Push(val)
			}
			tmp.Reverse()
			for _, s := range tmp {
				stacks[to].Push(s)
			}
		}
	}
	result := ""
	for _, stack := range stacks {
		if tmp, ok := stack.Pop(); ok {
			result += tmp
		}
	}
	println(result)
}
