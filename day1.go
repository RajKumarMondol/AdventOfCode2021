package main

import "fmt"

func day1(inputFileName string) {
	fmt.Println("Day 01 : Problem 1 , Input : %s", inputFileName)
	values := readIntArray(inputFileName)
	fmt.Println("Input :")
	printArray(values)
	result := sonarSweepDepthIncrease(values)
	fmt.Println("Day 01 : Result : ", result)
}
func sonarSweepDepthIncrease(values []int) int {
	previousValue := 0
	depthIncreaseCount := 0
	for _, value := range values {
		if value > previousValue {
			depthIncreaseCount++
		}
		previousValue = value
		// fmt.Println(" now: ", value, " prev: ", previousValue, " dept Count: ", depthIncreaseCount)
	}
	return depthIncreaseCount - 1
}
