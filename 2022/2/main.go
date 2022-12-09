package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./2022/2/input.txt")
	if err != nil {
		panic(err)
	}

	rounds := strings.Split(string(input), "\r\n")
	totalScore := 0
	SecondtotalScore := 0

	scoreMap := map[string]map[string]int{
		"A": {
			"X": 1 + 3,
			"Y": 2 + 6,
			"Z": 3 + 0,
		},
		"B": {
			"X": 1 + 0,
			"Y": 2 + 3,
			"Z": 3 + 6,
		},
		"C": {
			"X": 1 + 6,
			"Y": 2 + 0,
			"Z": 3 + 3,
		},
	}
	SecondscoreMap := map[string]map[string]int{
		"A": {
			"X": 3 + 0,
			"Y": 1 + 3,
			"Z": 2 + 6,
		},
		"B": {
			"X": 1 + 0,
			"Y": 2 + 3,
			"Z": 3 + 6,
		},
		"C": {
			"X": 2 + 0,
			"Y": 3 + 3,
			"Z": 1 + 6,
		},
	}

	for _, round := range rounds {
		myInput := string(round[0])
		otherInput := string(round[2])
		totalScore += scoreMap[myInput][otherInput]
		SecondtotalScore += SecondscoreMap[myInput][otherInput]
	}
	println(totalScore)
	println(SecondtotalScore)
}
