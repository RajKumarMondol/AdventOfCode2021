package main

import "fmt"

func day3(inputFileName string) {
	fmt.Println("Day 03 : Problem 1 : Input : ", inputFileName)
	diagnosticReport := readBinaryCode(inputFileName)
	// fmt.Println("Input :")
	// printBinaryCodes(diagnosticReport)
	gamma, epsilon := findGamaEpsilonRate(diagnosticReport)
	fmt.Println("Day 03 : Problem 1 : Result : ", gamma*epsilon)
	// submarinPosition = SubmarinPosition{depth: 0, location: 0, aim: 0}
	// moveAsPerInstructionWithAim(instructions, &submarinPosition)
	oxyzenRating := findOxyzenRating(diagnosticReport)
	carbonDioxideRating := findCarbonDioxideRating(diagnosticReport)
	fmt.Println("Day 03 : Problem 2 : Result : ", binaryToDecimal(oxyzenRating)*binaryToDecimal(carbonDioxideRating))
}
func binaryToDecimal(byteCode string) int {
	result := 0
	for index := 0; index < len(byteCode); index++ {
		if byteCode[index] == '1' {
			result = result*2 + 1
		} else {
			result = result*2 + 0
		}
	}
	return result
}
func findOxyzenRating(diagnosticReport []string) string {
	reportLineLength := len(diagnosticReport[0])
	for index := 0; index < reportLineLength; index++ {
		mostCommon := findMostCommon(diagnosticReport, index)
		diagnosticReport = filter(diagnosticReport, index, mostCommon)
		if len(diagnosticReport) == 1 {
			return diagnosticReport[0]
		}
	}
	return ""
}
func findCarbonDioxideRating(diagnosticReport []string) string {
	reportLineLength := len(diagnosticReport[0])
	for index := 0; index < reportLineLength; index++ {
		leastCommon := findLeastCommon(diagnosticReport, index)
		diagnosticReport = filter(diagnosticReport, index, leastCommon)
		if len(diagnosticReport) == 1 {
			return diagnosticReport[0]
		}
	}
	return ""
}
func filter(diagnosticReportLines []string, index int, bit byte) []string {
	result := []string{}
	for _, diagnosticReportLine := range diagnosticReportLines {
		if diagnosticReportLine[index] == bit {
			result = append(result, diagnosticReportLine)
		}
	}
	return result
}
func findMostCommon(diagnosticReportLines []string, index int) byte {
	numberOfOnes := 0
	numberOfZeros := 0
	for _, diagnosticReportLine := range diagnosticReportLines {
		if diagnosticReportLine[index] == '1' {
			numberOfOnes++
		} else {
			numberOfZeros++
		}
	}
	if numberOfOnes < numberOfZeros {
		return '0'
	}
	return '1'
}

func findLeastCommon(diagnosticReportLines []string, index int) byte {
	numberOfOnes := 0
	numberOfZeros := 0
	for _, diagnosticReportLine := range diagnosticReportLines {
		if diagnosticReportLine[index] == '1' {
			numberOfOnes++
		} else {
			numberOfZeros++
		}
	}
	if numberOfOnes < numberOfZeros {
		return '1'
	}
	return '0'
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
