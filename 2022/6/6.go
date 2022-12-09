package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./2022/6/input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := strings.TrimSpace(string(input))

	k := 14
	//Find the first index where 14 letters are all different
	//Put the 14 letter in a set and see if the set has 4 elements
	for i := 0; i < len(inputStr)-k+1; i++ {
		set := make(map[byte]bool)
		for j := 0; j < k; j++ {
			set[inputStr[i+j]] = true
		}
		if len(set) == k {
			println(i + k)
			break
		}
	}
}
