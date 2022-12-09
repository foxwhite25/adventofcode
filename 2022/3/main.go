package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./2022/3/input.txt")
	if err != nil {
		panic(err)
	}

	rucksacks := strings.Split(string(input), "\r\n")
	totalPriority := 0
	secondTotalPriority := 0
	var groups []string
	for _, rucksack := range rucksacks {
		groups = append(groups, rucksack)
		if len(groups) == 3 {
			//Find the character that appear in all three item
			for _, char := range groups[0] {
				if strings.Contains(groups[1], string(char)) && strings.Contains(groups[2], string(char)) {
					priority := char - 96
					if priority < 0 {
						priority += 58
					}
					secondTotalPriority += int(priority)
					break
				}
			}
			groups = []string{}
		}

		length := len(rucksack)
		firstHalf := rucksack[:length/2]
		secondHalf := rucksack[length/2:]
		// Find the character that appear in both half
		// If there is more than one, take the first one
		for _, char := range firstHalf {
			if strings.Contains(secondHalf, string(char)) {
				priority := char - 96
				if priority < 0 {
					priority += 58
				}
				totalPriority += int(priority)
				break
			}
		}
	}
	println(totalPriority)
	println(secondTotalPriority)
}
