package main

import "fmt"

func day1(inputFileName string) {
	fmt.Println("Day 01 : Problem 1 : Input : ", inputFileName)
	sonarData := readIntArray(inputFileName)
	// fmt.Println("Input :")
	// printArray(sonarData)
	part1Result := sonarSweepDepthIncrease(sonarData)
	fmt.Println("Day 01 : Problem 1 : Result : ", part1Result)
	sonarDataSlidingWindow := windowSum(sonarData)
	// fmt.Println("Sliding window depth:")
	// printArray(sonarDataSlidingWindow)
	part2Result := sonarSweepDepthIncrease(sonarDataSlidingWindow)
	fmt.Println("Day 01 : Problem 2 : Result : ", part2Result)
}
func windowSum(values []int) []int {
	windowData := []int{}
	for index, _ := range values {
		if index >= 2 {
			windowData = append(windowData, (values[index] + values[index-1] + values[index-2]))
		}
	}
	return windowData
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
