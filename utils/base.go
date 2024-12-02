package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadInput(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func WriteOutput(filePath string, part1Result interface{}) error {
	err := os.MkdirAll("outputs", os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%v\n", part1Result))
	return err
}

func ParseInts(lines []string) ([]int, error) {
	var numbers []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func SplitAndConvert(input, delimiter string) ([]int, error) {
	parts := strings.Split(input, delimiter)
	return ParseInts(parts)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ExtractIntegers(input string) []int {
	var numbers []int
	var currentNum strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			currentNum.WriteRune(char)
		} else if currentNum.Len() > 0 {
			// Convert the accumulated digits to an integer and append to the result
			num, err := strconv.Atoi(currentNum.String())
			if err == nil {
				numbers = append(numbers, num)
			}
			currentNum.Reset()
		}
	}

	// Check for any remaining number after the loop
	if currentNum.Len() > 0 {
		num, err := strconv.Atoi(currentNum.String())
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

// ExtractInt64s takes a string and returns a slice of int64
func ExtractInt64s(input string) []int64 {
	var numbers []int64
	var currentNum strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			currentNum.WriteRune(char)
		} else if currentNum.Len() > 0 {
			num, err := strconv.ParseInt(currentNum.String(), 10, 64)
			if err == nil {
				numbers = append(numbers, num)
			}
			currentNum.Reset()
		}
	}

	// Check for any remaining number after the loop
	if currentNum.Len() > 0 {
		num, err := strconv.ParseInt(currentNum.String(), 10, 64)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}
