package code

import (
	"fmt"
	"log"

	"advent/utils"
)

func DayX() {
	// Read input from file
	lines, err := utils.ReadInput("inputs/dayX.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Part 1
	result1 := solveDayXPart1(lines)

	// Part 2
	result2 := solveDayXPart2(lines)

	// Print results to console
	fmt.Printf("Day X - Part 1: \n%v\n", result1)
	fmt.Printf("Day X - Part 2: \n%v\n", result2)

	// Write results to output files
	err = utils.WriteOutput("outputs/dayX_1.txt", result1)
	if err != nil {
		log.Fatalf("Failed to write output for Part 1: %v", err)
	}
	err = utils.WriteOutput("outputs/dayX_2.txt", result2)
	if err != nil {
		log.Fatalf("Failed to write output for Part 2: %v", err)
	}
}

func solveDayXPart1(input []string) interface{} {
	for _, in := range input {
		fmt.Println(in) // Replace with your logic for Part 1
	}
	return "Result for Part 1" // Replace with actual result
}

func solveDayXPart2(input []string) interface{} {
	// Implement your solution for Part 2 here
	return "Result for Part 2" // Replace with actual result
}
