package code

import (
	"fmt"
	"log"

	"advent/utils"
)

func Day2() {
	// Read input from file
	lines, err := utils.ReadInput("inputs/day2.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Part 1
	result1 := solveDay2Part1(lines)

	// Part 2
	result2 := solveDay2Part2(lines)

	// Print results to console
	fmt.Printf("Day 2 - Part 1: \n%v\n", result1)
	fmt.Printf("Day 2 - Part 2: \n%v\n", result2)

	// Write results to output files
	err = utils.WriteOutput("outputs/day2_1.txt", result1)
	if err != nil {
		log.Fatalf("Failed to write output for Part 1: %v", err)
	}
	err = utils.WriteOutput("outputs/day2_2.txt", result2)
	if err != nil {
		log.Fatalf("Failed to write output for Part 2: %v", err)
	}
}

// safety rules
// the levels per line are all increasong or all decreasing
// adjacent levels difffer by at leas one and at most three
func solveDay2Part1(input []string) interface{} {
	t := 0
	for _, in := range input {
		ns := utils.ExtractIntegers(in)
		tmp := ns[0] - ns[1]

		if utils.Abs(tmp) > 3 || tmp == 0 {
			continue
		}

		asc := tmp < 0

		for i := 2; i < (len(ns)); i++ {
			mov := ns[i-1] - ns[i]
			if asc {
				if mov >= 0 || mov < -3 {
					break
				}
			} else {
				if mov <= 0 || mov > 3 {
					break
				}
			}
			if (i + 1) == (len(ns)) {
				t++
			}
		}
	}
	return t
}

func solveDay2Part2(input []string) interface{} {
	t := 0

	//this probs needs to be recursive
	isSafe := func(ns []int) bool {
		if len(ns) < 2 {
			return true
		}

		tmp := ns[0] - ns[1]
		if utils.Abs(tmp) > 3 || tmp == 0 {
			return false
		}

		asc := tmp < 0

		for i := 2; i < len(ns); i++ {
			mov := ns[i-1] - ns[i]
			if asc {
				if mov >= 0 || mov < -3 {
					return false
				}
			} else {
				if mov <= 0 || mov > 3 {
					return false
				}
			}
		}
		return true
	}

	for _, in := range input {
		ns := utils.ExtractIntegers(in)

		if isSafe(ns) {
			t++
			continue
		}

		//could look at the issue area which is around i-1 or i but im in too deep
		//ill just tank the O(n^2)
		safeWithDampener := false
		for i := 0; i < len(ns); i++ {
			modified := append([]int{}, ns[:i]...)
			modified = append(modified, ns[i+1:]...)
			if isSafe(modified) {
				safeWithDampener = true
				break
			}
		}

		if safeWithDampener {
			t++
		}
	}

	return t
}
