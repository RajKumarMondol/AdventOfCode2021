package main

import "fmt"

type SubmarinPosition struct {
	depth    int
	location int
	aim      int
}

func day2(inputFileName string) {
	fmt.Println("Day 02 : Problem 1 : Input : ", inputFileName)
	instructions := readInstruction(inputFileName)
	// fmt.Println("Input :")
	// printInstructions(instructions)
	submarinPosition := SubmarinPosition{depth: 0, location: 0, aim: 0}
	moveAsPerInstruction(instructions, &submarinPosition)
	part1Result := submarinPosition.depth * submarinPosition.location
	fmt.Println("Day 02 : Problem 1 : Result : ", part1Result)
	submarinPosition = SubmarinPosition{depth: 0, location: 0, aim: 0}
	moveAsPerInstructionWithAim(instructions, &submarinPosition)
	part2Result := submarinPosition.depth * submarinPosition.location
	fmt.Println("Day 02 : Problem 2 : Result : ", part2Result)
}
func moveAsPerInstruction(instructions []instruction, submarinPosition *SubmarinPosition) {
	for _, instruction := range instructions {
		switch {
		case instruction.command == "forward":
			submarinPosition.location += instruction.unit
		case instruction.command == "up":
			submarinPosition.depth -= instruction.unit
		case instruction.command == "down":
			submarinPosition.depth += instruction.unit
		}
		// fmt.Println(" Instruction : ", instruction, "\n - Position : ", submarinPosition)
	}
}
func moveAsPerInstructionWithAim(instructions []instruction, submarinPosition *SubmarinPosition) {
	for _, instruction := range instructions {
		switch {
		case instruction.command == "forward":
			submarinPosition.location += instruction.unit
			submarinPosition.depth += instruction.unit * submarinPosition.aim
		case instruction.command == "up":
			submarinPosition.aim -= instruction.unit
		case instruction.command == "down":
			submarinPosition.aim += instruction.unit
		}
		// fmt.Println(" Instruction : ", instruction, "\n - Position : ", submarinPosition)
	}
}
