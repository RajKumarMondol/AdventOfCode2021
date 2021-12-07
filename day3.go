package main

import "fmt"

func day3(inputFileName string) {
	fmt.Println("Day 03 : Problem 1 : Input : ", inputFileName)
	diagnosticReport := readBinaryCode(inputFileName)
	// fmt.Println("Input :")
	// printBinaryCodes(diagnosticReport)
	gamma, epsilon := findGamaEpsilonRate(diagnosticReport)
	fmt.Println("Day 03 : Problem 1 : Result : ", gamma*epsilon)
}

func findGamaEpsilonRate(diagnosticReport []string) (int, int) {
	reportLineLength := len(diagnosticReport[0])
	gamma := 0
	epsilon := 0
	for index := 0; index < reportLineLength; index++ {
		numberOfOnes := 0
		numberOfZeros := 0
		for _, diagnosticReportLine := range diagnosticReport {
			if diagnosticReportLine[index] == '1' {
				numberOfOnes++
			} else {
				numberOfZeros++
			}
		}
		if numberOfOnes > numberOfZeros {
			gamma = gamma*2 + 1
			epsilon = epsilon*2 + 0
		} else {
			gamma = gamma*2 + 0
			epsilon = epsilon*2 + 1
		}
	}
	return gamma, epsilon
}
