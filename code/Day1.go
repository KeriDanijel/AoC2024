package code

import (
	"advent/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	// Read input from file
	lines, err := utils.ReadInput("inputs/day1.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Part 1
	result1 := solveDay1Part1(lines)

	// Part 2
	result2 := solveDay1Part2(lines)

	// Print results to console
	fmt.Printf("Day 1 - Part 1: \n%v\n", result1)
	fmt.Printf("Day 1 - Part 2: \n%v\n", result2)

	// Write results to an output file
	err = utils.WriteOutput("outputs/day1_1.txt", result1)
	if err != nil {
		log.Fatalf("Failed to write output: %v", err)
	}
	err = utils.WriteOutput("outputs/day1_2.txt", result2)
	if err != nil {
		log.Fatalf("Failed to write output: %v", err)
	}
}

func solveDay1Part1(input []string) interface{} {
	t := 0
	var ll []int
	var rl []int
	for _, in := range input {
		// Split the string by three spaces
		s := strings.Split(in, "   ")
		if len(s) != 2 {
			return "Invalid input format. Each line must have two space-separated values."
		}

		// Convert the left and right parts to integers
		left, err1 := strconv.Atoi(s[0])
		right, err2 := strconv.Atoi(s[1])
		if err1 != nil || err2 != nil {
			return "Error converting string to integers"
		}

		// Append to respective slices
		ll = append(ll, left)
		rl = append(rl, right)
	}

	// Sort both slices
	sort.Ints(ll)
	sort.Ints(rl)
	// sort ll and rl
	for i := 0; i < len(ll); i++ {
		t += int(utils.Abs(ll[i] - rl[i]))
	}

	return t
}

func solveDay1Part2(input []string) interface{} {

	var t int64 = 0
	m := make(map[int64]int64)

	var ll []int64

	for _, ln := range input {
		temp := utils.ExtractInt64s(ln)
		v, ex := m[temp[1]]
		if ex {
			m[temp[1]] = v + temp[1]
		} else {
			m[temp[1]] = temp[1]
		}
		ll = append(ll, temp[0])
	}

	for i := 0; i < len(ll); i++ {
		v, ex := m[ll[i]]
		if ex {
			t += v
		}
	}

	return t
}