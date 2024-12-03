package code

import (
	"fmt"
	"log"
	"strings"

	"advent/utils"
)

func Day3() {
	// Read input from file
	lines, err := utils.ReadInput("inputs/day3.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Part 1
	result1 := solveDay3Part1(lines)

	// Part 2
	result2 := solveDay3Part2(lines)

	// Print results to console
	fmt.Printf("Day 3 - Part 1: \n%v\n", result1)
	fmt.Printf("Day 3 - Part 2: \n%v\n", result2)

	// Write results to output files
	err = utils.WriteOutput("outputs/day3_1.txt", result1)
	if err != nil {
		log.Fatalf("Failed to write output for Part 1: %v", err)
	}
	err = utils.WriteOutput("outputs/day3_2.txt", result2)
	if err != nil {
		log.Fatalf("Failed to write output for Part 2: %v", err)
	}
}

func solveDay3Part1(input []string) interface{} {
	t := 0
	for _, in := range input {
		p := utils.ExtractPattern(in)
		for _, s := range p {
			m := utils.ExtractIntegers(s)
			t += m[0] * m[1]
		}
	}
	return t
}

func solveDay3Part2(input []string) interface{} {
	t := 0

	ci := strings.Join(input, "")
	ep := utils.ExtractEnabled(ci)

	for _, e := range ep {
		ps := utils.ExtractPattern(e)
		for _, p := range ps {
			is := utils.ExtractIntegers(p)
			t += is[0] * is[1]
		}
	}
	return t
}
