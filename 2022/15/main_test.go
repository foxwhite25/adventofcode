package main

import (
	"io/ioutil"
	"strconv"
	"testing"
)

func BenchmarkFindUnseenPoint(b *testing.B) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		b.Fatal(err)
	}
	match := inputRegex.FindAllStringSubmatch(string(input), -1)
	pairs := make([]SensorBeaconPair, len(match))
	min := Coordinate{X: 0, Y: 0}
	max := Coordinate{X: 4000000, Y: 4000000}
	for i, m := range match {
		sensorX, _ := strconv.Atoi(m[1])
		sensorY, _ := strconv.Atoi(m[2])
		beaconX, _ := strconv.Atoi(m[3])
		beaconY, _ := strconv.Atoi(m[4])

		pairs[i] = SensorBeaconPair{
			Sensor:   Coordinate{X: sensorX, Y: sensorY},
			Beacon:   Coordinate{X: beaconX, Y: beaconY},
			Distance: abs(sensorX-beaconX) + abs(sensorY-beaconY),
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findUnseenPoints(pairs, min, max)
	}
}
