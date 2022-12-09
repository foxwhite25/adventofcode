package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./2022/1/input.txt")
	if err != nil {
		panic(err)
	}

	grouped := strings.Split(string(input), "\r\n\r\n")
	var cals []int
	for _, group := range grouped {
		foods := strings.Split(group, "\r\n")
		cal := 0
		for _, food := range foods {
			tmp, err := strconv.Atoi(strings.TrimSpace(food))
			if err != nil {
				continue
			}
			cal += tmp
		}
		cals = append(cals, cal)
	}
	// Get the sum of top three in cals
	var sum int
	for i := 0; i < 3; i++ {
		var max int
		var index int
		for i, cal := range cals {
			if cal > max {
				max = cal
				index = i
			}
		}
		sum += max
		cals[index] = 0
	}
	println(sum)
}
