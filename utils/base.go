package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
