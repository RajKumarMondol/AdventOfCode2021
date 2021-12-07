package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func printArray(values []int) {
	for _, value := range values {
		fmt.Println(value)
	}
}
func readIntArray(inputFileName string) []int {
	values := []int{}
	inputFile, err := os.OpenFile(inputFileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to open file error: %v", err)
		return values
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err == nil {
			values = append(values, value)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return values
	}
	return values
}
