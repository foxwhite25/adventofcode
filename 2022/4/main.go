package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./2022/4/input.txt")
	if err != nil {
		panic(err)
	}

	pairs := strings.Split(string(input), "\r\n")
	counter := 0
	overlapCounter := 0
	for _, pair := range pairs {
		ranges := strings.Split(pair, ",")
		firstRangeStart, _ := strconv.Atoi(strings.Split(ranges[0], "-")[0])
		firstRangeEnd, _ := strconv.Atoi(strings.Split(ranges[0], "-")[1])
		secondRangeStart, _ := strconv.Atoi(strings.Split(ranges[1], "-")[0])
		secondRangeEnd, _ := strconv.Atoi(strings.Split(ranges[1], "-")[1])

		//Check if the range completely cover the other range
		if (firstRangeStart <= secondRangeStart && firstRangeEnd >= secondRangeEnd) || (secondRangeStart <= firstRangeStart && secondRangeEnd >= firstRangeEnd) {
			counter++
		}

		//Check if the range overlap
		if (firstRangeStart <= secondRangeStart && firstRangeEnd >= secondRangeStart) || (secondRangeStart <= firstRangeStart && secondRangeEnd >= firstRangeStart) {
			overlapCounter++
		}
	}
	println(counter)
	println(overlapCounter)
}
