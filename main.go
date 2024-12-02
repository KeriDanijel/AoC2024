package main

import (
	//"fmt"
	"log"
	"os"

	"advent/code"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go [day]")
	}

	day := os.Args[1]
	switch day {
	case "1":
		code.Day1()
	//case "2":
	//code.Day2()
	default:
		log.Fatalf("Day %s is not implemented yet", day)
	}
}
