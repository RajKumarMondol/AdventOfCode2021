package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func printArray(values []int) {
	for _, value := range values {
		fmt.Println(value)
	}
}
func printInstructions(values []instruction) {
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

type instruction struct {
	command string
	unit    int
}

func readInstruction(inputFileName string) []instruction {
	instructions := []instruction{}
	inputFile, err := os.OpenFile(inputFileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to open file error: %v", err)
		return instructions
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		value, err := strconv.Atoi(words[1])
		if err == nil {
			instruction := instruction{command: words[0], unit: value}
			instructions = append(instructions, instruction)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return instructions
	}
	return instructions
}
