package main

import (
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	stack := make([]float64, 0)

	cars := make([][]int, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = make([]int, 2)
		cars[i][0] = position[i]
		cars[i][1] = speed[i]
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[j][0] < cars[i][0]
	})

	var timeTaken float64
	for _, car := range cars {
		timeTaken = float64(target-car[0]) / float64(car[1])
		if len(stack) == 0 || timeTaken > stack[len(stack)-1] {
			stack = append(stack, timeTaken)
		}
	}

	return len(stack)
}
